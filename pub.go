package main

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

const NatsTopic = "my_app"

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish(NatsTopic, []byte("Hello from publisher"))

	log.Infoln("Sent...")

	nc.Drain()
}
