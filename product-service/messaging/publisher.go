package messaging

import "github.com/nats-io/nats.go"

func Publish() {

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// todo pass actual publish event and conigure dockerfile
}
