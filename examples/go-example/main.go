package main

import "fmt"

// iAnimal will be used to achieve polymorphic behavior between different animal structs
type iAnimal interface {
	printName()
	makeCoolAnimalSound()
}

// animal is what I will refer to as the "parent" struct
type animal struct {
	name string
}

func (a *animal) printName() {
	fmt.Println(a.name)
}

func (a *animal) makeCoolAnimalSound() {
	fmt.Println("I do not know what type of animal I am. :(")
}

// dog will be the example "child" struct
type dog struct {
	animal         // anonymous field
	dogCollarBrand string
}

func (d *dog) makeCoolAnimalSound() {
	fmt.Println("Woof!")
}

func (d *dog) peeOnFireHydrant() {
	fmt.Println(d.name, "marks his territory!")
}

func main() {
	// instantiate animal and dog as themselves rather than a parent or interface
	unknownAnimal := &animal{"Bob the unknown"}
	cozmo := &dog{
		animal: animal{
			name: "Cozmo the dog",
		},
		dogCollarBrand: "Barky",
	}

	// now pass them to something that utilizes an interface in order to achieve polymorphic behavior
	runPolymorphicAnimalFuncs(unknownAnimal)
	runPolymorphicAnimalFuncs(cozmo)

	// do things specific to the structs rather than the interfaces
	fmt.Println() // newline
	fmt.Println(unknownAnimal.name)
	fmt.Println(cozmo.dogCollarBrand)
	cozmo.peeOnFireHydrant()
}

func runPolymorphicAnimalFuncs(polymorphicAnimal iAnimal) {
	polymorphicAnimal.printName()
	polymorphicAnimal.makeCoolAnimalSound()
}
