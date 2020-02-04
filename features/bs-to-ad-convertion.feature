Feature: convert dates from BS to AD using an API
  As an app-developer in Nepal
  I want to be able to send BS dates to an API endpoint and receive the corresponding AD dates
  So that I have a simple way to convert BS to AD dates, that can be used in other apps

  Scenario: converting a valid BS date
  When a "GET" request is sent to the endpoint "/ad-from-bs/2060-04-01"
  Then the HTTP-response code should be "200"
  And the response content should be "2003-07-17"

  Scenario: converting an invalid BS date
  When a "GET" request is sent to the endpoint "/ad-from-bs/60-13-01"
  Then the HTTP-response code should be "404"
  And the response content should be "not a valid date"
