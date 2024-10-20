package chanplus

import (
	"sync"
)

// ChannelState represents a tracked channel with introspection capabilities.
type ChannelState[T any] struct {
	ch     chan T
	closed bool
	mu     sync.Mutex
}

// New creates and returns a new Chan instance with the specified capacity.
// If capacity is 0, an unbuffered channel is created.
//
// Example:
//
//	ch := chanplus.New[int](3) // Creates a buffered channel of integers with capacity 3
//	unbufferedCh := chanplus.New[string](0) // Creates an unbuffered channel of strings
func New[T any](capacity int) *ChannelState[T] {
	return &ChannelState[T]{
		ch: make(chan T, capacity),
	}
}

// Send sends a value to the channel. It returns an error if the channel is closed.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](1)
//	err := ch.Send(42)
//	if err != nil {
//	    // Handle error (e.g., channel closed)
//	}
func (cs *ChannelState[T]) Send(value T) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if cs.closed {
		return false
	}

	cs.ch <- value
	return true
}

// Close closes the channel. It returns an error if the channel is already closed.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](1)
//	err := ch.Close()
//	if err != nil {
//	    // Handle error (e.g., channel already closed)
//	}
func (cs *ChannelState[T]) Close() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if !cs.closed {
		close(cs.ch)
		cs.closed = true
	}
}

// IsClosed returns true if the channel is closed, false otherwise.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](1)
//	fmt.Println("Is closed:", ch.IsClosed())
//	ch.Close()
//	fmt.Println("Is closed:", ch.IsClosed())
func (cs *ChannelState[T]) IsClosed() bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.closed
}

// Receive receives a value from the channel. It returns the received value
// and a boolean indicating whether the channel is closed.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](1)
//	ch.Send(42)
//	value, ok := ch.Receive()
//	if !ok {
//	    // Channel is closed
//	} else {
//	    fmt.Println("Received:", value)
//	}
func (cs *ChannelState[T]) Receive() (T, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if cs.closed && len(cs.ch) == 0 {
		var zero T
		return zero, false
	}

	select {
	case value, ok := <-cs.ch:
		return value, ok
	default:
		var zero T
		return zero, false
	}
}

func (cs *ChannelState[T]) GetChannel() chan T {
	return cs.ch
}

// Len returns the number of items currently in the channel.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](5)
//	fmt.Println("Length:", ch.Len()) // Output: 0
func (cs *ChannelState[T]) Len() int {
	return len(cs.ch)
}

// Cap returns the capacity of the channel.
// This method is thread-safe.
//
// Example:
//
//	ch := chanplus.New[int](5)
//	fmt.Println("Capacity:", ch.Cap()) // Output: 5
func (cs *ChannelState[T]) Cap() int {
	return cap(cs.ch)
}
