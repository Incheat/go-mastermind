package bdd_test

import "github.com/cucumber/godog"

func anExistingTodoWithTitleSavedAs(arg1, arg2 string) error {
        return godog.ErrPending
}

func iSaveTheJSONFieldAs(arg1, arg2 string) error {
        return godog.ErrPending
}

func iSendADELETERequestTo(arg1 string) error {
        return godog.ErrPending
}

func iSendAGETRequestTo(arg1 string) error {
        return godog.ErrPending
}

func iSendAPOSTRequestTo(arg1 string) error {
        return godog.ErrPending
}

func iSendAPUTRequestTo(arg1 string) error {
        return godog.ErrPending
}

func theAPIBaseURLIs(arg1 string) error {
        return godog.ErrPending
}

func theJSONArrayLengthShouldBe(arg1 int) error {
        return godog.ErrPending
}

func theJSONArrayShouldContainAnItemWithFieldEqual(arg1, arg2 string) error {
        return godog.ErrPending
}

func theRawRequestBodyIs(arg1 *godog.DocString) error {
        return godog.ErrPending
}

func theRequestHeadersAre(arg1 *godog.Table) error {
        return godog.ErrPending
}

func theRequestPayloadIs(arg1 *godog.DocString) error {
        return godog.ErrPending
}

func theResponseBodyShouldContain(arg1 string) error {
        return godog.ErrPending
}

func theResponseJSONFieldShouldEqual(arg1, arg2 string) error {
        return godog.ErrPending
}

func theResponseJSONShouldHaveAField(arg1 string) error {
        return godog.ErrPending
}

func theResponseShouldBeAJSONArray() error {
        return godog.ErrPending
}

func theResponseStatusShouldBe(arg1 int) error {
        return godog.ErrPending
}

func initTodoSteps(ctx *godog.ScenarioContext) {
        ctx.Step(`^an existing todo with title "([^"]*)" saved as "([^"]*)"$`, anExistingTodoWithTitleSavedAs)
        ctx.Step(`^I save the JSON field "([^"]*)" as "([^"]*)"$`, iSaveTheJSONFieldAs)
        ctx.Step(`^I send a DELETE request to "([^"]*)"$`, iSendADELETERequestTo)
        ctx.Step(`^I send a GET request to "([^"]*)"$`, iSendAGETRequestTo)
        ctx.Step(`^I send a POST request to "([^"]*)"$`, iSendAPOSTRequestTo)
        ctx.Step(`^I send a PUT request to "([^"]*)"$`, iSendAPUTRequestTo)
        ctx.Step(`^the API base URL is "([^"]*)"$`, theAPIBaseURLIs)
        ctx.Step(`^the JSON array length should be (\d+)$`, theJSONArrayLengthShouldBe)
        ctx.Step(`^the JSON array should contain an item with field "([^"]*)" equal "([^"]*)"$`, theJSONArrayShouldContainAnItemWithFieldEqual)
        ctx.Step(`^the raw request body is:$`, theRawRequestBodyIs)
        ctx.Step(`^the request headers are:$`, theRequestHeadersAre)
        ctx.Step(`^the request payload is:$`, theRequestPayloadIs)
        ctx.Step(`^the response body should contain "([^"]*)"$`, theResponseBodyShouldContain)
        ctx.Step(`^the response JSON field "([^"]*)" should equal "([^"]*)"$`, theResponseJSONFieldShouldEqual)
        ctx.Step(`^the response JSON should have a field "([^"]*)"$`, theResponseJSONShouldHaveAField)
        ctx.Step(`^the response should be a JSON array$`, theResponseShouldBeAJSONArray)
        ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
}