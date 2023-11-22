package main

import "fmt"

// Interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

// Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) Says() string {
	return d.Sound
}
func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}
func main() {
	// ask a riddle (задати загадку)
	dog := Dog{
		Name:         "Dog",
		Sound:        "Gaf",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "Cat"
	cat.Sound = "meow"
	cat.NumberOfLegs = 4
	cat.HasTail = true

	Riddle(&cat)

}
func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says %s and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs)
	fmt.Println(riddle)
}
