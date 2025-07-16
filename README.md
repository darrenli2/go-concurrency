# Go Concurrency Examples

A collection of Go programs demonstrating various concurrency patterns and synchronization techniques.

## Files Overview

### Basic Concurrency Patterns

- **`concurrent.go`** - Task processing using worker goroutines with a buffered channel
- **`concurrent2.go`** - Enhanced task processing with detailed comments and unbuffered channels
- **`concurrent3.go`** - Semaphore-based concurrency control using buffered channels

### Synchronization Examples

- **`nondeadlock.go`** - Demonstrates a potential deadlock scenario with channel synchronization
- **`nondeadlock2.go`** - Fixed version using `sync.WaitGroup` to prevent deadlocks

### Special Cases

- **`neverending.go`** - Example of an infinite program using channel-based message passing

## Running the Examples

Execute any file using:
```bash
go run <filename>.go
```

For example:
```bash
go run concurrent.go
go run neverending.go
```

## Key Concepts Demonstrated

- **Worker Pool Pattern** - Limiting concurrent workers processing tasks
- **Channel Communication** - Using channels for goroutine synchronization
- **Semaphore Pattern** - Controlling concurrency with buffered channels
- **WaitGroup Usage** - Proper synchronization of multiple goroutines
- **Deadlock Prevention** - Common pitfalls and their solutions

## Learning Notes

- `concurrent.go` shows basic worker pools with error handling
- `concurrent2.go` and `concurrent3.go` provide different approaches to the same problem
- `nondeadlock.go` vs `nondeadlock2.go` illustrates proper synchronization techniques
- `neverending.go` demonstrates how to keep a program running indefinitely