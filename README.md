[![Build Status](https://github.com/ioriver/ioriver-go/actions/workflows/master.yml/badge.svg)](https://github.com/ioriver/ioriver-go/actions/workflows/master.yml)

# ioriver-go

A Go client library for interacting with the IORiver API.

## Features

- Manage services, domains, certificates, origins, traffic policies, and more
- Simple, idiomatic Go API
- Designed for integration with Go applications and tools

## Installation

```
go get github.com/ioriver/ioriver-go
```

## Usage

```go
import "github.com/ioriver/ioriver-go"

// Example usage:
client := ioriver.NewClient("<api-key>")

// List services
services, err := client.ListServices()
if err != nil {
	// handle error
}
for _, svc := range services {
	fmt.Println(svc.Name)
}
```

## Documentation

See source code and GoDoc comments for API details.

- [IORiver Guides](https://www.ioriver.io/docs/category/guides/)
- [IO River REST API Documentation](https://www.ioriver.io/docs/api/io-river-api)

## License

MIT License
