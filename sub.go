package main

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

const NatsTopic = "my_app"

func main() {
	//wait := make(chan bool)

	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Panic(err)
	}

	nc.Subscribe(NatsTopic, func(m *nats.Msg) {
		log.Infof("Received: %s\n", string(m.Data))
		nc.Publish(m.Reply, []byte("Hello from subscriber"))
	})

	log.Infoln("Subscribed to", NatsTopic)

	//<-wait
	select {}
}
