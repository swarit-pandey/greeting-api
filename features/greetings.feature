Feature: Greeting message
  Scenario: Client makes a GET request to /greeting
    When the client requests GET /greeting
    Then the response should contain "Hello from server!"

