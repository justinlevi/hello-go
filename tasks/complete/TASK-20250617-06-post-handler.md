# TASK-20250617-06 - Implement POST Handler
Status: COMPLETE
Priority: MEDIUM
Assignee: TBD
Created: 2025-06-17
Updated: 2025-06-17
Completed: 2025-06-17

## User Story
As an API consumer
I want to send JSON data via POST requests
So that I can create personalized greetings with structured input

## Acceptance Criteria
- New POST endpoint at /hello accepts JSON body
- Request body structure: {"name": "string"}
- Validate Content-Type is application/json
- Return 400 for invalid JSON
- Return 415 for wrong Content-Type
- Limit request body size to prevent abuse
- Return personalized greeting based on input

## Gherkin Scenarios
```gherkin
Feature: POST Handler

Scenario: Successful POST with valid JSON
  Given the server is running
  When I POST {"name": "Alice"} to /hello
  Then I should receive 200 status
  And response contains {"message": "Hello, Alice!"}

Scenario: Invalid JSON body
  Given the server is running
  When I POST malformed JSON to /hello
  Then I should receive 400 status
  And error message about invalid JSON

Scenario: Missing Content-Type
  Given the server is running
  When I POST without Content-Type header
  Then I should receive 415 status
```

## Testing & Validation Steps
- Unit tests for JSON parsing
- Test request body size limits
- Validate Content-Type checking
- Test malformed JSON handling
- Verify proper error responses
- Performance test with concurrent POSTs
- Security test for JSON injection

## Parent/Sub Tasks
- Parent Task: None
- Sub Tasks: None
- Dependencies: TASK-20250617-03