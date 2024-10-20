package chanplus_test

import (
	"testing"

	"chanplus/chanplus"
)

// TestChannelSendAndReceive tests basic Send and Receive operations.
func TestChannelSendAndReceive(t *testing.T) {
	ch := chanplus.New[int](1)
	if !ch.Send(10) {
		t.Errorf("Send failed, expected true")
	}

	value, ok := ch.Receive()
	if !ok || value != 10 {
		t.Errorf("Expected to receive 10, got %v", value)
	}
}

// TestChannelClose tests closing a channel and verifying it doesn't accept further sends.
func TestChannelClose(t *testing.T) {
	ch := chanplus.New[int](1)

	// Close the channel
	ch.Close()

	if !ch.IsClosed() {
		t.Errorf("Expected channel to be closed")
	}

	if ch.Send(20) {
		t.Errorf("Expected Send to fail after channel is closed")
	}

	_, ok := ch.Receive()
	if ok {
		t.Errorf("Expected Receive to return false on closed channel")
	}
}

// TestLenAndCap tests Len and Cap methods for introspection.
func TestLenAndCap(t *testing.T) {
	ch := chanplus.New[int](3)

	if ch.Cap() != 3 {
		t.Errorf("Expected capacity to be 3, got %d", ch.Cap())
	}

	if ch.Len() != 0 {
		t.Errorf("Expected length to be 0, got %d", ch.Len())
	}

	ch.Send(1)
	if ch.Len() != 1 {
		t.Errorf("Expected length to be 1 after one send, got %d", ch.Len())
	}
}

// TestStringChannel tests the channel with string type.
func TestStringChannel(t *testing.T) {
	ch := chanplus.New[string](2)

	if !ch.Send("hello") {
		t.Errorf("Send failed, expected true")
	}

	if ch.Len() != 1 {
		t.Errorf("Expected length to be 1, got %d", ch.Len())
	}

	value, ok := ch.Receive()
	if !ok || value != "hello" {
		t.Errorf("Expected to receive 'hello', got %v", value)
	}

	if ch.Len() != 0 {
		t.Errorf("Expected length to be 0 after receive, got %d", ch.Len())
	}
}

// TestStructChannel tests the channel with a custom struct type.
func TestStructChannel(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	ch := chanplus.New[Person](1)

	person := Person{Name: "Alice", Age: 30}

	if !ch.Send(person) {
		t.Errorf("Send failed, expected true")
	}

	if ch.Len() != 1 {
		t.Errorf("Expected length to be 1, got %d", ch.Len())
	}

	receivedPerson, ok := ch.Receive()
	if !ok || receivedPerson != person {
		t.Errorf("Expected to receive %v, got %v", person, receivedPerson)
	}

	if ch.Len() != 0 {
		t.Errorf("Expected length to be 0 after receive, got %d", ch.Len())
	}
}
