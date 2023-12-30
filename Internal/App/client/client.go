package client

import (
	"netcat/Internal/config"
	"net"
)

// The type "Client" represents a client connection with an ID, a network connection, and a name.
// @property {int} Id - The Id property is an integer that represents the unique identifier of the
// client.
// @property Connect - The `Connect` property is of type `net.Conn`, which represents a network
// connection. It is used to establish and manage network connections in Go.
// @property {string} Name - The "Name" property of the "Client" struct is a string that represents the
// name of the client.
type Client struct {
	Id      int
	Connect net.Conn
	Name    string
}

// The Message type represents a message sent from a client with a payload of bytes.
// @property {Client} From - The "From" property is of type "Client". It represents the sender of the
// message.
// @property {[]byte} Payload - The `Payload` property is a byte array that represents the data being
// sent in the message. It can be any sequence of bytes, such as text, binary data, or serialized
// objects.
type Message struct {
	From    Client
	Payload []byte
}

// The NewClient function creates a new client object with the given name and connection.
func NewClient(name string, conn net.Conn) Client {
	return Client{
		Id:      config.IdClient,
		Connect: conn,
		Name:    name,
	}
}
