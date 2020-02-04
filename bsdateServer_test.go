package main

import (
    "github.com/DATA-DOG/godog"
)

func aRequestIsSentToTheEndpoint(arg1, arg2 string) error {
    return godog.ErrPending
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
