# Go QUIC Proxy

Go QUIC Proxy is a simple application that proxies TCP data over QUIC connections. This project utilizes the `quic-go` package to facilitate the conversion of TCP traffic into QUIC streams, providing a modern and efficient way to handle network communications.

## Features

- Proxy TCP connections over QUIC
- Easy to configure server settings
- Lightweight and efficient

## Prerequisites

- Go 1.16 or later
- `quic-go` package

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/go-quic-proxy.git
   cd quic-proxy
   ```

2. Install the dependencies:

   ```
   go mod tidy
   ```

## Configuration

The application requires a configuration file to specify the server address and port. You can modify the `internal/config.go` file to set your desired settings.

## Running the Application

To run the application, execute the following command:

```
go run cmd/main.go
```

This will start the QUIC server and begin listening for incoming connections.

## Usage

Once the server is running, you can connect to it using a QUIC-compatible client. The server will accept TCP connections and proxy the data over QUIC.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.