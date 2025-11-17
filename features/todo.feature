Feature: Todo management
  As an API client
  I want to manage todos via the HTTP API
  So that I can list, create, retrieve, update, and delete todos

  Background:
    Given the API base URL is "http://localhost:8080/api/v1"
    And the request headers are:
      | Content-Type | application/json |

  Scenario: List todos when none exist
    When I send a GET request to "/todos"
    Then the response status should be 200
    And the response should be a JSON array
    And the JSON array length should be 0

  Scenario: Create a new todo
    Given the request payload is:
      """
      {
        "title": "Buy milk",
        "description": "2 liters of whole milk"
      }
      """
    When I send a POST request to "/todos"
    Then the response status should be 201
    And the response JSON should have a field "id"
    And the response JSON field "title" should equal "Buy milk"
    And I save the JSON field "id" as "buyMilkId"

  Scenario: List todos when one exists
    Given an existing todo with title "Walk dog" saved as "walkDogId"
    When I send a GET request to "/todos"
    Then the response status should be 200
    And the response should be a JSON array
    And the JSON array should contain an item with field "title" equal "Walk dog"

  Scenario: Get an existing todo
    Given an existing todo with title "Read book" saved as "readBookId"
    When I send a GET request to "/todos/{readBookId}"
    Then the response status should be 200
    And the response JSON field "title" should equal "Read book"

  Scenario: Get a non-existent todo returns 404
    When I send a GET request to "/todos/non-existent-id-123"
    Then the response status should be 404

  Scenario: Create a todo without title should fail
    Given the request payload is:
      """
      {
        "description": "no title here"
      }
      """
    When I send a POST request to "/todos"
    Then the response status should be 400
    And the response body should contain "title"

  Scenario: Create a todo with empty title should fail
    Given the request payload is:
      """
      {
        "title": "",
        "description": "empty title"
      }
      """
    When I send a POST request to "/todos"
    Then the response status should be 400
    And the response body should contain "title"

  Scenario: Create a todo with invalid JSON should fail
    Given the raw request body is:
      """
      { "title": "Bad JSON"
      """
    When I send a POST request to "/todos"
    Then the response status should be 400

  Scenario: Update an existing todo
    Given an existing todo with title "Original title" saved as "updateId"
    And the request payload is:
      """
      {
        "title": "Updated title"
      }
      """
    When I send a PUT request to "/todos/{updateId}"
    Then the response status should be 200
    And the response JSON field "title" should equal "Updated title"

  Scenario: Update a non-existent todo returns 404
    Given the request payload is:
      """
      {
        "title": "Whatever"
      }
      """
    When I send a PUT request to "/todos/non-existent-id-456"
    Then the response status should be 404

  Scenario: Delete an existing todo
    Given an existing todo with title "To be deleted" saved as "deleteId"
    When I send a DELETE request to "/todos/{deleteId}"
    Then the response status should be 200

  Scenario: Delete a non-existent todo returns 404
    When I send a DELETE request to "/todos/non-existent-id-789"
    Then the response status should be 404
