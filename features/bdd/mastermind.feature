Feature: Mastermind Game
  As a player
  I want to play Mastermind via the HTTP API
  So that I can test my guessing skills

  Background:
    Given the API is running on local test server
    And the request headers are:
      | Content-Type | application/json |

  Scenario: Start a new game
    Given I send a POST request to "/games"
    Then the response status should be 201
    And the response JSON should have a field "gameId"
    And I save the JSON field "gameId" as gameId

  Scenario Outline: Make a guess and receive feedback
    Given the game answer is "<answer>"
    And the gameId is "<realGameId>"
    And the request payload is:
      """
      {
        "guess": "<guess>"
      }
      """
    When I send a POST request to "/games/<requestGameId>/guesses"
    Then the response status should be 200
    And the response JSON should have a field "feedback"
    And the response JSON field "feedback" should equal "<expectedFeedback>"

    Examples:
      | guess | answer | realGameId | requestGameId | expectedFeedback |
      | 1234  | 1234   | 12345      | 12345         | 4A0B             |
      | 1234  | 1235   | 12345      | 12345         | 3A0B             |