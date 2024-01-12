package main

import (
	"fmt"
	"sync"
)

// MethodRequest encapsulates a method call along with its arguments
type MethodRequest struct {
	method func(args ...interface{})
	args   []interface{}
}

// ActiveObject encapsulates its own thread of control and executes methods asynchronously
type ActiveObject struct {
	queue      chan MethodRequest
	stopSignal sync.WaitGroup
	isRunning  bool
}

// NewActiveObject creates a new ActiveObject
func NewActiveObject() *ActiveObject {
	return &ActiveObject{
		queue:      make(chan MethodRequest),
		isRunning:  true,
	}
}

// Run starts the ActiveObject's goroutine
func (a *ActiveObject) Run() {
	for a.isRunning || len(a.queue) > 0 {
		methodRequest := <-a.queue
		methodRequest.method(methodRequest.args...)
		a.stopSignal.Done() // Decrement the WaitGroup counter
	}
}

// ScheduleMethod schedules a method to be executed
func (a *ActiveObject) ScheduleMethod(method func(args ...interface{}), args ...interface{}) {
	a.stopSignal.Add(1) // Increment the WaitGroup counter
	methodRequest := MethodRequest{
		method: method,
		args:   args,
	}
	a.queue <- methodRequest
}

// Stop stops the ActiveObject's goroutine
func (a *ActiveObject) Stop() {
	a.isRunning = false
	go func() {
		a.stopSignal.Wait() // Wait for the goroutine to finish
		close(a.queue)
	}()
}

func main() {
	// Create an instance of ActiveObject
	activeObject := NewActiveObject()

    var wg sync.WaitGroup
	wg.Add(1)

	// Start the ActiveObject goroutine
	go func() {
		activeObject.Run()
		wg.Done()
	}()

	// Invoke methods on the ActiveObject
	activeObject.ScheduleMethod(printMessage, "Hello")
	activeObject.ScheduleMethod(printMessage, "World")

	// Stop the ActiveObject goroutine
	activeObject.Stop()
    wg.Wait()

}

func printMessage(args ...interface{}) {
	message := args[0].(string)
	fmt.Println(message)
}
