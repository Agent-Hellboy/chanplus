package main

import (
	"chanplus/chanplus"
	"fmt"
)

func main() {
	// Create a new channel with buffer size 2
	ch := chanplus.New[int](2)

	// Send values
	ch.Send(10)
	ch.Send(20)

	// Receive values
	val1, ok1 := ch.Receive()
	val2, ok2 := ch.Receive()

	fmt.Printf("Received: %d (ok: %v), %d (ok: %v)\n", val1, ok1, val2, ok2)

	fmt.Println(ch.Len())
	fmt.Println(ch.Cap())

	// Try to receive from an empty channel
	val3, ok3 := ch.Receive()
	fmt.Printf("Received from empty channel: %d (ok: %v)\n", val3, ok3)

	// Close the channel
	ch.Close()

	// Try to send to a closed channel
	sendOk := ch.Send(30)
	fmt.Printf("Send to closed channel: %v\n", sendOk)
}
