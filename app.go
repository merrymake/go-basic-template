package main

import (
	"fmt"
	"os"
	m "github.com/merrymake/go-service-library"
	l "github.com/merrymake/go-service-library/lib"
)

func handleHello(payloadBytes []byte, envelope l.Envelope) {
	payload := string(payloadBytes)

	m.ReplyStringToOrigin(fmt.Sprintf("Hello, %s!", payload), l.GetMimeType("txt"))
}

func main() {
	args := os.Args[1:]
	new(m.Merrymake).Service(args).
		Handle("handleHello", handleHello)
}
