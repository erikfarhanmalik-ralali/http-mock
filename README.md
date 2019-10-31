# http-mock

## Prerequisite
- go (https://golang.org/doc/install).

## How to run
- Go to source directory.
- Execute `go run main.go`.
- Application will be run on localhost:8888.
- End point can be accessed like this: `http://localhost:8888/v1/mock/level`

## Add mock end point
- Open main.go file
- Create a function with router as parameter (use addLevelMockAPI function as sample)
- Set the desired API end point URL
- Set the response data