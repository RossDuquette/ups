package main

import (
	"flag"
	"os"

	"github.com/RossDuquette/ups/pkg/pubsub"
)

func main() {
	topic := flag.String("topic", "", "Topic to publish to.")
	message := flag.String("msg", "", "Message to publish.")
	flag.Parse()

	conn := pubsub.Connect(os.Args[0], "127.0.0.1")
	conn.Publish(*topic, []byte(*message))
}
