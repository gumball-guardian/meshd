package meshd

// Works for websocket client and server
type Transport interface {
	Open()         // Client will connect, server will listen
	Close()        // Client will disconnect, server will stop listening and disconnect all clients
	OnConnect()    // When we connect or re-connect
	OnDisconnect() // Handle disconnection
	OnMessage()    // When we receive a new message from the other endpoint
	OnError()      // When we have an error

	SendMessage() // Send a new message to the other endpoint
	SendReply()   // Send a response to a message we sent
}

type WebsocketClient struct {
	port int16
}

type WebsocketServer struct {
	port int16
}

func (ws WebsocketServer) Open() {

}

func (ws WebsocketServer) Close() {

}

func (ws WebsocketServer) OnConnect() {

}

func (ws WebsocketServer) SendMessage() {

}

func (ws WebsocketServer) OnMessage() {

}

func (ws WebsocketServer) SendReply() {

}

func (ws WebsocketServer) OnError() {

}

func (wc WebsocketClient) Open() {

}

func (wc WebsocketClient) Close() {

}

func (wc WebsocketClient) OnConnect() {

}

func (wc WebsocketClient) SendMessage() {

}

func (wc WebsocketClient) OnMessage() {

}

func (wc WebsocketClient) SendReply() {

}

func (wc WebsocketClient) OnError() {

}

type EmailAgent interface {
	// TODO: Some interface for polling via IMAP or POP3 or ???
	SendPlainTextEmail(targetEmail string, subject string, body string)
}
