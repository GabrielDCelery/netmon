package netstat

// Connection represents a single network connection from ss output.
type Connection struct {
	Protocol string
	State    string
	RecvQ    string
	SendQ    string
	Local    string
	Peer     string
	Process  string
}
