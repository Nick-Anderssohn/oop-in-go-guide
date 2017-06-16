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

func main() {
	// Can have normally instantiated structs...in this example, unknownAnimal is of type *animal
	unknownAnimal := &animal{
		name: "Bob the unknown",
	}
	unknownAnimal.printName()
	unknownAnimal.makeCoolAnimalSound()

	// Unlike normal oop languages, you CANNOT do this:
	//     var polymorphicDog *animal
	//     polymorphicDog = &dog{}
	// It would give you an error like this:
	// cannot use dog literal (type *dog) as type *animal in assignment

	// Instead, you utilize the interface to accomplish polymorphism
	// First, note that because type animal has methods printName() and makeCoolAnimalSound(),
	// it automatically implements iAnimal without you doing anything else
	// This means you could do something like this:

	var polymorphicAnimal iAnimal
	polymorphicAnimal = unknownAnimal
	polymorphicAnimal.printName()
	polymorphicAnimal.makeCoolAnimalSound()

	// Since dog has animal as an anonymous field, it "inherited" both printName() and
	// makeCoolAnimalSound() from animal (note that it also overrode makeCoolAnimalSound())
	// This means that iAnimal can also be instantiated as a dog:
	polymorphicAnimal = &dog{
		animal: animal{
			name: "Cozmo the dog",
		},
		dogCollarBrand: "Barky",
	}
	polymorphicAnimal.printName()
	polymorphicAnimal.makeCoolAnimalSound()

	// So how is this different than normal oop? Notice that with normal oop, we could access
	// genericAnimal's dogCollarBrand field by casting it to a dog. You CANNOT do this in go! Interfaces
	// are simply a list of functions and nothing else!
	// This brings us to the PROPER (in my opinion) way to do inheritance and polymorphism in go.
	// Separate the polymorphic behavior from  the rest of your code utilizing functions that accept an
	// interface as the parameter.
	properWay()
}

func properWay() {
	fmt.Println("\n***Proper way***")
	// create a variable of type *animal
	unknownAnimal := &animal{
		name: "Bob the unknown",
	}
	// run polymorphic functions since it implements iAnimal
	runPolymorphicAnimalFuncs(unknownAnimal)

	// create a variable of type *dog
	polymorphicDog := &dog{
		animal: animal{
			name: "Cozmo the dog",
		},
		dogCollarBrand: "Barky",
	}
	// run polymorphic functions since it implements iAnimal
	runPolymorphicAnimalFuncs(polymorphicDog)
	// and now, since polymorphicDog is of type *dog instead of type iAnimal,
	// we can access dogCollarBrand
	fmt.Println(polymorphicDog.dogCollarBrand)
}

func runPolymorphicAnimalFuncs(polymorphicAnimal iAnimal) {
	polymorphicAnimal.printName()
	polymorphicAnimal.makeCoolAnimalSound()
}
