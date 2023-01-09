package main

import (
    "fmt"
    "github.com/cucumber/godog"
    "io/ioutil"
    "net/http"
    "strings"
)

var host = "http://localhost:10000"

var res *http.Response

func aRequestIsSentToTheEndpoint(method, endpoint string) error {
    var reader = strings.NewReader("")
    var request, err = http.NewRequest(method, host+endpoint, reader)
    if err != nil {
        return fmt.Errorf("could not create request %s", err.Error())
    }

    res, err = http.DefaultClient.Do(request)
    if err != nil {
        return fmt.Errorf("could not send request %s", err.Error())
    }
    return nil
}

func theHTTPresponseCodeShouldBe(expectedCode int) error {
    if expectedCode != res.StatusCode {
        return fmt.Errorf("status code not as expected! Expected '%d', got '%d'", expectedCode, res.StatusCode)
    }
    return nil
}

func theResponseContentShouldBe(expectedContent string) error {
    body, _ := ioutil.ReadAll(res.Body)
    if expectedContent != strings.TrimSpace(string(body)) {
        return fmt.Errorf("response content not as expected! Expected '%s', got '%s'", expectedContent, string(body))
    }
    return nil
}

func FeatureContext(ctx *godog.ScenarioContext) {
    ctx.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
    ctx.Step(`^the HTTP-response code should be "(\d+)"$`, theHTTPresponseCodeShouldBe)
    ctx.Step(`^the response content should be "([^"]*)"$`, theResponseContentShouldBe)
}
