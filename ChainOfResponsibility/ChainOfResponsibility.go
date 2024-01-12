package main

import "fmt"

// Handler interface
type Handler interface {
	handleRequest()
}

// ConcreteHandler1 struct
type ConcreteHandler1 struct {
	successor Handler
}

// NewConcreteHandler1 constructor
func NewConcreteHandler1(successor Handler) *ConcreteHandler1 {
	return &ConcreteHandler1{successor: successor}
}

// handleRequest method for ConcreteHandler1
func (ch1 *ConcreteHandler1) handleRequest() {
	if true { // Can handle request.
		fmt.Println("Finally handled by ConcreteHandler1")
	} else if ch1.successor != nil {
		fmt.Println("Message passed to next in chain by ConcreteHandler1")
		ch1.successor.handleRequest()
	}
}

// ConcreteHandler2 struct
type ConcreteHandler2 struct {
	successor Handler
}

// NewConcreteHandler2 constructor
func NewConcreteHandler2(successor Handler) *ConcreteHandler2 {
	return &ConcreteHandler2{successor: successor}
}

// handleRequest method for ConcreteHandler2
func (ch2 *ConcreteHandler2) handleRequest() {
	if false { // Can't handle request.
		fmt.Println("Finally handled by ConcreteHandler2")
	} else if ch2.successor != nil {
		fmt.Println("Message passed to next in chain by ConcreteHandler2")
		ch2.successor.handleRequest()
	}
}

func main() {
	// Client code.
	ch1 := NewConcreteHandler1(nil)
	ch2 := NewConcreteHandler2(ch1)
	ch2.handleRequest()
}

/*
Message passed to next in chain by ConcreteHandler2
Finally handled by ConcreteHandler1
*/