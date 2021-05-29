## Useful commands

- initialize your new project
```go
go mod init example.com/moduleName
```

- run main.go (default file that go looks to run)
```go
go run .
```

- run tests with `-v` for verbose option.
```go
go test -v
```

- run all tests in the directory
```go
go test -v ./...
```

- build your app
```go
go build
```

- edit go mod
```go
go mod edit -replace=example.com/greetings=../greetings
```