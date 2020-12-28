# Test Microservice

## Overview
* Receives an IP address and slice of whitedlisted countries
* Returns boolean value indicating if country is whitelisted

## Usage

* To use as HTTP server, run `go run ./cmd/http/`
* To use as gRPC server, run `go run ./cmd/grpc/`

## Testing

Run tests with `go test ./...`

## Next Steps

* Confirm API response is useful and sufficient for other service
* Add more unit tests
* Look into [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to provide API in both gRPC and HTTP at the same time.