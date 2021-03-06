# Test Microservice

## Overview
* Receives an IP address and slice of whitedlisted countries
* Returns boolean value indicating if country is whitelisted

## Usage

### HTTP

* To use as HTTP server, run `go run ./cmd/http/`
* Create a POST request to `/api/checkip` with the following:
    ```json
    {
      "ip":"<IP Address>",
      "whitelist":["<Array of Countries>"]
    }
    ```


### gRPC

* To use as gRPC server, run `go run ./cmd/grpc/`
* In your gRPC client, structure your IP address and list of countries to satisfy the `Input` type:
    ```go
    type Input struct {
      IP        string   
      WhiteList []string
    }
    ```
## Testing

Run tests with `go test ./...`

## Next Steps

* ~~Confirm that API response is useful and sufficient for other service~~
* ~~Determine whether country codes would be preferable to English country names~~
* ~~Use Environemt Variables to make it easier to swap mmdbs for `FindCountryByIP`~~
* Add more unit tests
* Look into [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to provide API in both gRPC and HTTP at the same time.
* Determine better naming for Input and Output in `whitelist.proto`
* Determine what error to return is `json.Marshall` fails
* Determine if I should change name of `microservices.CheckIPAddress` to differentiate it from `app.CheckIPAddress`?