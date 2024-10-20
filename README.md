# Chanplus

Chanplus is a Go package that provides a safe, introspective wrapper for channels. It allows users to easily create, manage, and inspect channels while ensuring thread-safe operations.

[![Go Reference](https://pkg.go.dev/badge/github.com/Agent-Hellboy/chanplus.svg)](https://pkg.go.dev/github.com/Agent-Hellboy/chanplus)
[![Go Report Card](https://goreportcard.com/badge/github.com/Agent-Hellboy/chanplus)](https://goreportcard.com/report/github.com/Agent-Hellboy/chanplus)
[![codecov](https://codecov.io/gh/Agent-Hellboy/chanplus/branch/main/graph/badge.svg)](https://codecov.io/gh/Agent-Hellboy/chanplus)


## Disclaimer

This package isn't meant to replace the standard `chan` type. It's more like a sidekick, throwing in some extra features and safety checks.

Okay, I'll level with you - I created this package 'cause I was bored and thought, "Hey, why not make channels a bit fancier?"

But for real this time - I made this to get my head around how Go packages are built, how to put them on GitHub, how to make them available for other devs, and all that jazz about documentation, linters, formatters, and writing tests. You know, the whole shebang of package development.

## Features
- Generic support for channels (`T any`)
- Safely send and receive values from the channel
- Introspect channel state (check if closed, view length and capacity)
- Thread-safe operations

## Installation

To install Chanplus, use the following command:

```
go get github.com/Agent-Hellboy/chanplus
```


## Usage

```go
package main

import (
	"fmt"
	"github.com/Agent-Hellboy/chanplus"
)

func main() {
	ch := chanplus.New[int](3)

	// Send values
	ch.Send(20)

	// Receive values
	val, _ := ch.Receive()
	fmt.Println("Received:", val)

	// Introspect
	fmt.Println("Length:", ch.Len())
	fmt.Println("Capacity:", ch.Cap())

	// Close the channel
	ch.Close()
}
```

## API Reference

For detailed API documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/Agent-Hellboy/chanplus) page.

**EVERYTHING BELOW IS JUST A JOKE, PLEASE IGNORE IT.**


## Why Chanplus?

Chanplus simplifies channel usage in Go by:

1. **Eliminating race conditions**: All operations are thread-safe by default.
2. **Simplifying error handling**: No need to use `select` or worry about sending on closed channels.
3. **Providing easy introspection**: Check channel state, length, and capacity with simple method calls.
4. **Offering a consistent API**: Use the same methods regardless of whether the channel is buffered or unbuffered.


## Future Scope

We have several ideas for expanding and improving Chanplus in the future:

1. **Timeout operations**: Implement Send and Receive methods with timeout options.
2. **Channel composition**: Provide utilities for combining multiple channels (e.g., merge, split, multiplex).
3. **Advanced introspection**: Add more detailed statistics and monitoring capabilities.
4. **Buffered-to-unbuffered conversion**: Allow dynamic conversion between buffered and unbuffered channels.
5. **Channel patterns**: Implement common channel patterns like worker pools and pipelines.
6. **Context integration**: Add support for cancellation and deadlines using Go's context package.
7. **Performance optimizations**: Continuously improve performance for high-throughput scenarios.
8. **Evaluate benefits for concurrency patterns**: Assess the advantages of using Chanplus in various scenarios such as:
   - Pipelines
   - Fan-in/fan-out
   - Worker pools
   - Pub/sub systems
   - Rate limiting
   - Timeouts and cancellation
   - Broadcast channels
   - Multiplexing and demultiplexing


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
