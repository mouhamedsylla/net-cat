
# Netcat Chat Server

A simple chat server and client implementation in Go using TCP. The server allows multiple clients to connect, chat, and broadcast messages to each other.



## Usage

### Server
To start the server, run the following command:
```bash
./TCPChat $port
```
Replace `$port` with the desired port number.

### Client

No direct client implementation is provided. However, you can use tools like `netcat` or create a custom Go client to connect to the server on the specified port.

## Server Architecture

The server is structured with the following components:

- **Server**: Manages incoming connections, client data, and message broadcasting.
- **Client**: Represents a connected client with a unique ID, connection details, and a name.
- **Message**: Represents a message sent by a client.

## Server Features

- **Connection Limit**: The server has a maximum connection limit, set to 10 by default. Clients attempting to connect after the limit is reached will be refused.
- **Welcome Message**: Clients receive a welcome message upon successful connection.
- **Chat**: Clients can send and receive messages in the chat room.
- **User Join/Leave Notifications**: Clients receive notifications when a new user joins or an existing user leaves the chat.

## Customization

### Server Configuration

Server configuration options can be modified in the `Internal/config/config.go` file. The following options are available:

- **MAX_CONNECTIONS**: Maximum number of allowed connections.
- **PORT**: Port number for the server to listen on.

### Extending the Client

To create a custom client, import the `netcat/Internal/App/client` package and use the `NewClient` function to create a new client instance.

## Dependencies
This program does not have external dependencies beyond the standard Go library.

## Acknowledgments
This project is a basic implementation to demonstrate Go concurrency and network programming concepts. Feel free to extend and enhance it based on your requirements.








