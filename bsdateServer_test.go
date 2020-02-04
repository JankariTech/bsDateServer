package main

import (
    "fmt"
    "github.com/DATA-DOG/godog"
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

func theHTTPresponseCodeShouldBe(arg1 string) error {
    return godog.ErrPending
}

func theResponseContentShouldBe(arg1 string) error {
    return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
    s.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
    s.Step(`^the HTTP-response code should be "([^"]*)"$`, theHTTPresponseCodeShouldBe)
    s.Step(`^the response content should be "([^"]*)"$`, theResponseContentShouldBe)
}
