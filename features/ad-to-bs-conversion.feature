Feature: convert dates from AD to BS using an API
  As an app-developer in Nepal
  I want to be able to send AD dates to an API endpoint and receive the corresponding BS dates
  So that I have a simple way to convert AD to BS dates, that can be used in other apps

  Scenario Outline: converting a valid AD date
    When a "GET" request is sent to the endpoint "/bs-from-ad/<ad-date>"
    Then the HTTP-response code should be "200"
    And the response content should be "<bs-date>"
    Examples:
      | ad-date    | bs-date    |
      | 2003-07-17 | 2060-04-01 |
      | 1983-04-14 | 2040-01-01 |
      | 1984-04-12 | 2040-12-30 |

  Scenario: converting an invalid AD date
    When a "GET" request is sent to the endpoint "/bs-from-ad/97-13-01"
    Then the HTTP-response code should be "400"
    And the response content should be "not a valid date"

  Scenario Outline: unhandled request types
    When a "<type>" request is sent to the endpoint "/bs-from-ad/60-13-01"
    Then the HTTP-response code should be "400"
    Examples:
      | type |
      | POST |
      | PUT  |
