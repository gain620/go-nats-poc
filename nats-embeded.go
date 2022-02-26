package main

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	opts := &server.Options{}

	// Initialize new server with options
	ns, err := server.NewServer(opts)

	if err != nil {
		log.Panicln(err)
	}

	// Start the server via goroutine
	go ns.Start()

	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		log.Panicln("not ready for connection")
	}

	// Connect to server
	nc, err := nats.Connect(ns.ClientURL())

	if err != nil {
		log.Panicln(err)
	}

	subject := "my-subject"

	// Subscribe to the subject
	nc.Subscribe(subject, func(msg *nats.Msg) {
		// Print message data
		data := string(msg.Data)
		log.Infoln(data)

		// Shutdown the server (optional)
		ns.Shutdown()
	})

	// Publish data to the subject
	nc.Publish(subject, []byte("Hello embedded NATS!"))

	// Wait for server shutdown
	ns.WaitForShutdown()
}
