package main

import "fmt"

type Animal struct {
	name string
}

type Bird struct {
	Animal
}

func (b *Bird) fly() {
	// To be overridden by subclasses
}

type Dodo struct {
	Bird
}

func (d *Dodo) fly() {
	fmt.Println("The dodo is extinct and cannot fly.")
}

type Penguin struct {
	Bird
}

func (p *Penguin) fly() {
	fmt.Println("The penguin cannot fly.")
}

func (p *Penguin) slide() {
	fmt.Println("The penguin is sliding on its belly!")
}

func (p *Penguin) swim() {
	fmt.Println("The penguin is swimming in the water!")
}

type Eagle struct {
	Bird
}

func (e *Eagle) fly() {
	fmt.Println("The eagle is soaring through the sky!")
}

type Sparrow struct {
	Bird
}

func (s *Sparrow) fly() {
	fmt.Println("The sparrow is fluttering its wings!")
}

// Client code.
func main() {
	bird1 := &Eagle{Bird{Animal{"Eagle"}}}
	bird1.fly()

	bird2 := &Dodo{Bird{Animal{"Dodo"}}}
	bird2.fly()

	/*
		The eagle is soaring through the sky!
		The dodo is extinct and cannot fly.
	*/

	bird3 := &Penguin{Bird{Animal{"Pigeon"}}}
	bird3.fly()

	/*
		The pigeon is fluttering its wings!
	*/
}

/*
The eagle is soaring through the sky!
The dodo is extinct and cannot fly.
The penguin cannot fly.
*/