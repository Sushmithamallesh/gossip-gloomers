package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	var counter uint64

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"

		nodeID := n.ID()
		timeStamp := time.Now().UnixNano() / 1000000

		count := atomic.AddUint64(&counter, 1)

		body["id"] = fmt.Sprintf("%s-%d-%d", nodeID, timeStamp, count)
		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
