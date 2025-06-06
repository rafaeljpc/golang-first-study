# Learning with Golang: Consuming Kafka or HTTP Endpoint and Saving to Elasticsearch

This project demonstrates how to consume messages from a Kafka server or HTTP endpoint and save the data into Elasticsearch. The project includes an E2E test package that verifies the functionality of both Kafka and HTTP interfaces.

## Project Structure

The project is structured as follows:

- `cmd`: Contains the main executable.
- `internal`: Contains internal packages.
- `internal\di`: Dependency injection container.
- `internal\domain\service`: Domain service layer
- `internal\adapter`: Adapter layer
- `test`: Contains end-to-end tests.

## Prerequisites

Ensure you have the following tools installed:

- Go 1.24 or higher
- Docker (for E2E tests using TestContainer)

## E2E Tests

Run the E2E tests with the following command:

```sh
go test ./test
```

## Contributing

Contributions are welcome! Open a Pull Request or contact us to discuss improvements.

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.
