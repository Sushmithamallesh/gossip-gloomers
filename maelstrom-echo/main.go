package main

import (
	"encoding/json"
	"log"
	"os"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message body to a loosely types map
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message body type to be echo_ok
		body["type"] = "echo_ok"

		//Echo the original message back
		return n.Reply(msg, body)
	})

	// Execute the node's message loop. This will run until the STDIN is closed.
	if err := n.Run(); err != nil {
		log.Println("Error %s", err)
		os.Exit(1)
	}
}
