# Channels

## What is this topic?

Channels are Go's way to communicate between goroutines. This project covers channel creation, send/receive, buffered channels, and channel directions.

## Channel Creation

```go
ch := make(chan int)        // unbuffered
ch := make(chan int, 10)    // buffered (capacity 10)
```

## Sending and Receiving

```go
ch <- 42        // send
value := <-ch   // receive
```

## Unbuffered Channels

- Synchronous communication
- Sender blocks until receiver is ready
- Receiver blocks until sender is ready

## Buffered Channels

- Asynchronous up to capacity
- Sender blocks only when buffer is full
- Receiver blocks only when buffer is empty

## Channel Directions

```go
func send(ch chan<- int)    // send-only
func receive(ch <-chan int) // receive-only
```

## Closing Channels

```go
close(ch)
value, ok := <-ch  // ok is false when closed
```

## Key Concepts

- **Unbuffered**: Synchronous communication
- **Buffered**: Asynchronous up to capacity
- **Send/receive**: `<-` operator
- **Close**: Signal no more values
- **Direction**: Restrict channel usage
