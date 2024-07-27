# Go Concurrency Example

## Project Setup

### Initiate Go Modules

```sh
go mod init go_concurrency_example
go mod tidy
```

### Launch the Web API

```sh
go run ./webapi
```

## Performance Testing

### Single Client Test

```sh
go run ./simple_client http://localhost:3000 20
```

This should typically take about 20 seconds, depending on the response time of the web API and the number of requests (20 in this example).

### Concurrent Client Test with Goroutines

```sh
go run ./concurent_client http://localhost:3000 20
```

This is expected to significantly reduce the time, potentially down to about 1 second, illustrating the power of concurrency in Go.
