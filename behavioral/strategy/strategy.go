package main

import "fmt"

type Duck struct {
	FlyBehavior
	QuackBehavior
	Name string
}

type FlyBehavior interface {
	Fly()
}

type RocketFly struct{}

func (b *RocketFly) Fly() {
	fmt.Println("rocket-fly")
}

type CannotFly struct{}

func (b *CannotFly) Fly() {
	fmt.Println("cannot-fly")
}

type Quock struct{}

func (b *Quock) Quack() {
	fmt.Println("quock")
}

type Quick struct{}

func (b *Quick) Quack() {
	fmt.Println("quick")
}

type QuackBehavior interface {
	Quack()
}

func NewSuperDuck() *Duck {
	return &Duck{
		Name:          "superduck",
		FlyBehavior:   &RocketFly{},
		QuackBehavior: &Quock{},
	}
}

func NewMiniDuck() *Duck {
	return &Duck{
		Name:          "miniduck",
		FlyBehavior:   &CannotFly{},
		QuackBehavior: &Quick{},
	}
}

func main() {
	miniduck := NewMiniDuck()
	fmt.Println(miniduck.Name)
	miniduck.Quack()
	miniduck.Fly()

	superduck := NewSuperDuck()
	fmt.Println(superduck.Name)
	superduck.Quack()
	superduck.Fly()
}
