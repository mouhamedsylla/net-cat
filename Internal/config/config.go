package config


var (
	IdClient = 0
	RefuseMessage = "Connection refused\n"
	WelcomeMessage = `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    ` + "`.       | `" + `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     ` + "`-'       `" + `--'
[ENTER YOUR NAME]:`
)

// The type Options is a struct that contains two fields, MAX_CONNECTIONS and PORT.
// @property {int} MAX_CONNECTIONS - The MAX_CONNECTIONS property represents the maximum number of
// connections that can be established at the same time.
// @property {string} PORT - The PORT property is a string that represents the port number for a
// network connection.
type Options struct {
	MAX_CONNECTIONS int
	PORT string
}

type OptionsFunc func(*Options)

// The function SetMaxConnect returns an OptionsFunc that sets the maximum number of connections in the
// Options struct.
func SetMaxConnect(max int) OptionsFunc {
	return func(o *Options) {
		o.MAX_CONNECTIONS = max
	}
}

// The function SetPort returns an OptionsFunc that sets the PORT field of an Options struct to the
// provided port string.
func SetPort(port string) OptionsFunc {
	return func(o *Options) {
		o.PORT = port
	}
}

// The function returns a default set of options for a server, including a maximum number of
// connections and a port number.
func DefaultOptions() Options{
	return Options{
		MAX_CONNECTIONS: 10,
		PORT: ":8989",
	}
}

