package bdd_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

// --- Step implementations ---

var (
	vars map[string]string
	headers http.Header
	testServer *httptest.Server
)

func startTestServer() {
	if testServer != nil {
		return
	}
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate POST /games
		if r.Method == http.MethodPost && r.URL.Path == "/games" {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"gameId":"12345"}`))
			return
		}
		// Simulate POST /games/{gameId}/guesses
		if r.Method == http.MethodPost && r.URL.Path == "/games/12345/guesses" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"feedback":"4A0B"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
}


func stopTestServer() {
	if testServer != nil {
		testServer.Close()
		testServer = nil
	}
}

func theAPIisRunningOnLocalTestServer() error {
	if testServer == nil {
		return fmt.Errorf("test server is not running")
	}
	return nil
}

func theRequestHeadersAre(table *godog.Table) error {
	if len(table.Rows) < 1 {
		return fmt.Errorf("headers table must have at least one data row")
	}

	// Expecting:
	// | Header-Name | Header-Value |
	for _, row := range table.Rows {
		if len(row.Cells) != 2 {
			return fmt.Errorf("header row must have 2 columns (header name, value)")
		}
		name := row.Cells[0].Value
		value := row.Cells[1].Value
		headers.Set(name, value)
	}

	return nil
}

type APISteps struct {
	url string
	response *http.Response
	data map[string]interface{}
	payload []byte
	gameId string
	answer string
}

func newAPISteps() *APISteps {
	return &APISteps{data: make(map[string]interface{})}
}

func (a *APISteps) iSendAPOSTRequestTo(path string) error {
	fullURL := testServer.URL + path

	req, err := http.NewRequest(http.MethodPost, fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	// Copy headers from the global headers map (if any)
	req.Header = headers.Clone()

	resp, err := testServer.Client().Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	a.response = resp
	return nil
}

func (a *APISteps) theResponseStatusShouldBe(status int) error {
	
	if a.response == nil {
		return fmt.Errorf("no response received")
	}
	if a.response.StatusCode != status {
		return fmt.Errorf("expected status %d but got %d", status, a.response.StatusCode)
	}

	return nil
}

func (a *APISteps) theResponseJSONShouldHaveField(field string) error {

	if a.response == nil {
		return fmt.Errorf("no response received")
	}

	body, err := io.ReadAll(a.response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	a.response.Body.Close()

	bodyTrimmed := bytes.TrimSpace(body)
	if len(bodyTrimmed) == 0 {
		return fmt.Errorf(
			"expected JSON body but response was empty (status=%d, path=%s)",
			a.response.StatusCode,
			a.response.Request.URL.Path,
		)
	}

	if err := json.Unmarshal(bodyTrimmed, &a.data); err != nil {
		return fmt.Errorf(
			"failed to unmarshal response JSON: %w (raw body: %q)",
			err, string(bodyTrimmed),
		)
	}

	if _, ok := a.data[field]; !ok {
		return fmt.Errorf("expected JSON field %q to be present, but it was missing", field)
	}

	return nil
}

func (a *APISteps) iSaveTheJSONFieldAsGameID(field string) error {
	if a.data == nil {
		return fmt.Errorf("no response received")
	}
	a.gameId = a.data[field].(string)
	return nil
}

func (a *APISteps) theGameAnswerIs(answer string) error {
	a.answer = answer
	return nil
}

func (a *APISteps) theRequestPayloadIs(payload string) error {
	// Optional: validate JSON so tests fail fast
	if !json.Valid([]byte(payload)) {
		return fmt.Errorf("invalid JSON payload:\n%s", payload)
	}

	a.payload = []byte(payload)
	return nil
}

func (a *APISteps) theGameIdIs(gameId string) error {
	a.gameId = gameId
	return nil
}

func (a *APISteps) iSendAPOSTRequestToGameGuesses() error {
	fullURL := testServer.URL + "/games/" + a.gameId + "/guesses"

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(a.payload))
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	req.Header = headers.Clone()

	resp, err := testServer.Client().Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	a.response = resp
	return nil
}

func (a *APISteps) theResponseJSONFieldShouldEqual(field, expected string) error {
	if a.data == nil {
		return fmt.Errorf("no response received")
	}

	if a.data[field] != expected {
		return fmt.Errorf("expected JSON field %q to be %q, but got %q", field, expected, a.data[field])
	}
	return nil
}



// ---- Godog wiring ----

func initMastermindTestSuite(suite *godog.TestSuiteContext) {
	vars = make(map[string]string)
	headers = http.Header{}
	suite.BeforeSuite(startTestServer)
	suite.AfterSuite(stopTestServer)
}

func initMastermindScenario(ctx *godog.ScenarioContext) {
	apiSteps := newAPISteps()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	ctx.Step(`^the API is running on local test server$`, theAPIisRunningOnLocalTestServer)
	ctx.Step(`^the request headers are:$`, theRequestHeadersAre)
	ctx.Step(`^I send a POST request to "([^"]+)"$`, apiSteps.iSendAPOSTRequestTo)
	ctx.Step(`^the response status should be (\d+)$`, apiSteps.theResponseStatusShouldBe)
	ctx.Step(`^the response JSON should have a field "([^"]+)"$`, apiSteps.theResponseJSONShouldHaveField)
	ctx.Step(`^I save the JSON field "([^"]+)" as gameId$`, apiSteps.iSaveTheJSONFieldAsGameID)
	ctx.Step(`^the request payload is:$`, apiSteps.theRequestPayloadIs)
	ctx.Step(`^the game answer is "([^"]+)"$`, apiSteps.theGameAnswerIs)
	ctx.Step(`^the gameId is "([^"]+)"$`, apiSteps.theGameIdIs)
	ctx.Step(`^I send a POST request to "/games/{gameId}/guesses"$`, apiSteps.iSendAPOSTRequestToGameGuesses)
	ctx.Step(`^the response JSON field "([^"]+)" should equal "([^"]+)"$`, apiSteps.theResponseJSONFieldShouldEqual)
}