package main

import "fmt"

// DesiredInterface interface
type DesiredInterface interface {
	Operation()
}

// Adapter struct
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Operation() {
	a.adaptee.SomeOperation()
}

// Adaptee struct
type Adaptee struct{}

func (a *Adaptee) SomeOperation() {
	fmt.Println("Adaptee some_operation() function called.")
}

// Client Code
func main() {
	adapter := Adapter{adaptee: &Adaptee{}}
	adapter.Operation()
}

/*
Adaptee some_operation() function called.
*/