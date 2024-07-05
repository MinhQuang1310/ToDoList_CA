package main

import "fmt"

type Animal interface {
	Move(move string)
	Eat(eat string)
	Name(name string)
	getMove() string
	getName() string
	getEat() string
}

type Dog struct {
	name string
	move string
	eat string
}

func newDog () Animal {
	return &Dog{}
}

func (d *Dog) Move(move string) {
	d.move = move
}

func (d *Dog) Eat(eat string)  {
	d.eat = eat
}

func (d *Dog) Name(name string)  {
	d.name = name
}

func (d Dog) getMove() string {
    return d.move
}

func (d Dog) getName() string {
    return d.name
}

func (d Dog) getEat() string {
    return d.eat
}

func main() {
	animal := newDog()
	animal.Move("barks")
	animal.Eat("eats bones")
	animal.Name("tình")
	fmt.Printf("Name: %s\nMove: %s\nEat: %s\n", animal.getName(), animal.getMove(), animal.getEat())
}
