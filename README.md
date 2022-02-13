# calculator

## Simple calculator api

### To Run:
#### Navigate to the root directory and run the command:
`go run calculator.go`

#### Send a request to 'localhost:8080/math' that looks like:
`{
    "op": "+",
    "left": 4,
    "right": 2
}`

#### You should receive a response similar to:
`{
    "error": "",
    "result": 6
}`

#### Supported operations:
`"+", "-", "*", "/"

