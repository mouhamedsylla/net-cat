package server

import (
	"bufio"
	"netcat/Internal/App/client"
	"netcat/Internal/config"
	"netcat/Internal/utils"
	"fmt"
	"net"
)

// The Server type represents a server with options, a listener, a list of clients, a message channel,
// and a quit channel.
// @property option - The `option` property is of type `config.Options`. It likely contains various
// configuration options for the server, such as the server's address, port, and other settings.
// @property listener - The `listener` property is a variable of type `net.Listener`. It represents a
// network listener that listens for incoming connections on a specific network address.
// @property {[]client.Client} clients - The `clients` property is a slice of `client.Client` objects.
// It is used to store the connected clients to the server.
// @property Msgch - The `Msgch` property is a channel of type `client.Message`. It is used to send and
// receive messages between the server and the clients.
// @property quitch - The `quitch` channel is used to signal the server to stop and gracefully shut
// down. It is a channel of type `struct{}` which means it does not carry any data. When a value is
// sent on this channel, it indicates that the server should stop and exit.
type Server struct {
	option   config.Options
	listener net.Listener
	clients  []client.Client
	Msgch    chan client.Message
	quitch   chan struct{}
}

// The NewServer function creates a new server instance with customizable options.
func NewServer(options ...config.OptionsFunc) *Server {
	opt := config.DefaultOptions()
	for _, fn := range options {
		fn(&opt)
	}
	return &Server{
		option: opt,
		Msgch:  make(chan client.Message),
		quitch: make(chan struct{}),
	}
}


// The `Start()` function is a method of the `Server` struct. It is responsible for starting the server
// and handling incoming connections.
func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.option.PORT)
	if err != nil {
		return
	}
	fmt.Printf("Listening on the port %s", s.option.PORT)

	s.listener = ln
	go s.AcceptLoop()
	<-s.quitch
	close(s.Msgch)
}


// The `AcceptLoop()` function is a method of the `Server` struct. It is responsible for accepting
// incoming connections and starting the `ReadLoop()` and `Chat()` goroutines for each connection.
func (s *Server) AcceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}
		go s.ReadLoop(conn)
		go s.Chat()
	}
}


// The `ReadLoop` function is responsible for reading messages from a client connection and handling
// them. Here is a breakdown of what it does:
func (s *Server) ReadLoop(conn net.Conn) {
	defer func() {
		s.Disconnect(conn)
		conn.Close()
	}()
	if config.IdClient > 9 {
		conn.Write([]byte(config.RefuseMessage))
		conn.Close()
		return
	}
	conn.Write([]byte(config.WelcomeMessage))
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	c := client.NewClient(name[:len(name)-1], conn)
	config.IdClient++
	s.clients = append(s.clients, c)
	notif := fmt.Sprintf("%s has joined our chat...\n", c.Name)
	utils.NotifGroup(notif, s.ToSend(conn))
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			break
		}
		
		s.Msgch <- client.Message{
			From:    c,
			Payload: buffer[:n],
		}
	}
}


// The `Disconnect` function is responsible for disconnecting a client from the server. It takes a
// `net.Conn` object as a parameter, which represents the client's connection to the server.
func (s *Server) Disconnect(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	for i, client := range s.clients {
		addr2 := client.Connect.RemoteAddr().String()
		if addr2 == addr {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			notif := fmt.Sprintf("%s has left the chat...\n", client.Name)
			utils.NotifGroup(notif, s.clients)
			break
		}
	}
}


// The `Chat()` function is responsible for handling incoming messages from clients and broadcasting
// them to all other connected clients.
func (s *Server) Chat() {
	for msg := range s.Msgch {
		message := fmt.Sprintf("\n[%s]: %s", msg.From.Name, msg.Payload)
		conn := msg.From.Connect
		utils.BroadcastMessage(s.ToSend(conn), []byte(message))
	}
}

// The `ToSend` function is a method of the `Server` struct. It is responsible for creating a list of
// clients that should receive a message, excluding the client that sent the message.
func (s *Server) ToSend(conn net.Conn) (tabClient []client.Client) {
	addr := conn.RemoteAddr().String()
	for _, client := range s.clients {
		addr1 := client.Connect.RemoteAddr().String()
		if addr1 != addr {
			tabClient = append(tabClient, client)
		} else {
			remis := fmt.Sprintf("[%s]:", client.Name)
			client.Connect.Write([]byte(remis))
		}
	}
	return
}
