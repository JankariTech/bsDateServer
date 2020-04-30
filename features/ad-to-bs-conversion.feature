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
      | 2020-02-29 | 2076-11-17 |
      | 2020-01-31 | 2076-10-17 |

  Scenario Outline: converting an invalid AD date
    When a "GET" request is sent to the endpoint "/bs-from-ad/<ad-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "cannot convert date, invalid or missing data"
    Examples:
      | ad-date       |
      | 97-04         |
      | 97-04-08      |
      | aeiou         |
      | 1984-04-31    |
      | 1987-04-05-01 |
      | 2022-02-29    |

  Scenario Outline: unhandled request types
    When a "<type>" request is sent to the endpoint "/bs-from-ad/<ad-date>"
    Then the HTTP-response code should be "400"
    And the response content should be "<response>"
    Examples:
      | type | ad-date    | response                      |
      | POST | 97-04      | Could not create POST request |
      | PUT  | 97-04      | Could not create PUT request  |
      | POST | 2020-02-29 | Could not create POST request |
      | PUT  | 2020-02-29 | Could not create PUT request  |
