# A Guide to OOP in Go

How to accomplish object-oriented/polymorphic behavior in golang. This is meant for programmers who have at least a basic understanding of inheritance and polymorphism, experience with a standard OOP language (such as C#, Java, C++, etc...), know the basics of golang, and want to learn how to write better code in go and use OOP design patterns.

## A Basic OOP Example
This example will be basic and brief because you should have at least this level of understanding on inheritance and polymorphism. You might have a parent class Animal and child class Dog that look like this in C#:
```
class Animal
    {
        public string Name { get; set; }

        public void PrintName()
        {
            Console.WriteLine(Name);
        }

        public virtual void MakeCoolAnimalSound()
        {
            Console.WriteLine("I do not know what type of animal I am. :(");
        }

    }

    class Dog : Animal
    {
        public string DogCollarBrand {get; set;}
        override public void MakeCoolAnimalSound()
        {
            Console.WriteLine("Woof!");
        }

        public void PeeOnFireHydrant()
        {
          Console.WriteLine(Name + " marked his territory!");
        }
    }
```

We can create an instance of animal and have it do stuff.

```
Animal unknownAnimal = new Animal {Name = "Bob the unknown"};
unknownAnimal.PrintName();
unknownAnimal.MakeCoolAnimalSound();
```

We can also create a Dog. A Dog can be assigned to an animal because it inherits from Animal (thus you have polymorphism).

```
Animal polymorphicDog = new Dog
{
  Name = "Cozmo the dog",
  DogCollarBrand = "Barky"
};
```

You can access fields and methods from polymorphicDog that come from both class Animal and class Dog  since it is a Dog and Dog inherits from Animal.

```
polymorphicDog.PrintName();
polymorphicDog.MakeCoolAnimalSound();
Console.WriteLine(((Dog)polymorphicDog).DogCollarBrand);
```

## The Same Example in Go
This will be the same example as above, but implemented in go. It will show you how anonymous fields and interfaces can be used to achieve the same polymorphic behavior as above.
### Anonymous Fields
In go, we will have "parent" struct animal and "child" struct dog. Parent and child are in quotes because that is not the technically correct term for them, but it is okay to think of them like that for now.

```
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
```

Notice that dog has Animal as an anonymous field. You can think of this as a combination of inheritance and composition. In other words, dog has both an is-a and has-a relationship to animal. That sounds confusing, so lets clear it up.

##### An Anonymous Field's Inheritance-Like (is-a) Behavior

In the C# example, we saw that Dog inherited from Animal. This meant that Dog acquired Animal's methods and fields. In go, dog also acquired animal's methods and fields. Lets say there is a variable coolDog that has been instantiated like so:

```
coolDog := &dog{
  animal: animal{
    name: "Cozmo the dog",
  },
  dogCollarBrand: "Barky",
}
```
This instantiated coolDog to type *dog. As mentioned, dog has animal as an anonymous field and acquired both its methods and fields. They can be accessed like normal:

```
coolDog.name
coolDog.printName()
coolDog.makeCoolAnimalSound()
```

Remember that coolDog had its own implementation for makeCoolAnimalSound(). Accessing it like above would use dog's implementation of the method instead of animal's. There is a way to access animal's implementation of the method as well, but we will come to that later. dog's other methods and fields can also be accessed like normal:

```
coolDog.peeOnFireHydrant()
coolDog.dogCollarBrand
```

You can see how animal as an anonymous field acted like a parent class with virtual methods. In fact, you may think this does everything that inheriting a parent class in another language would do, but it doesn't. We will come to how and why later.

##### An Anonymous Field's Composition-Like (has-a) behavior

In general, when a class has a field of the same type, or of a parent class, it is referred to as composition. Since, dog has animal as a field (an anonymous one), this is an example of composition. In fact, animal can also be accessed like a normal field.

```
theAnimalThatCoolDogContains := coolDog.animal
```

This means you can access all of animal's fields and methods in a different way than the inheritance-like way.

```
coolDog.animal.name
```
will give you the same thing as
```
coolDog.name
```
Another example is
```
coolDog.animal.printName()
```
will give you the same thing as
```
coolDog.printName()
```

But, they only give you the same thing because animal and dog share the same definition for that field and method. Remember that dog overrode animal's definition of makeCoolAnimalSound()? To access animal's definition instead of dog's, you can do:

```
coolDog.animal.makeCoolAnimalSound()
```
instead of
```
coolDog.makeCoolAnimalSound()
```

The first one would give you an output of:
```
I do not know what type of animal I am. :(
```
whereas, the second one would give you an output of:
```
Woof!
```
### Polymorphism

Okay awesome. Now that I know what anonymous fields are, I can use the same design as in the C# application, right? Not quite, but lets try anyways. Lets assume we have created both animal and dog, and are ready to create the main logic of our application. First lets create a normal animal and make it do stuff.

```
unknownAnimal := &animal{
  name: "Bob the unknown",
}
unknownAnimal.printName()
unknownAnimal.makeCoolAnimalSound()
```

Yup, that worked perfectly! Now we can instantiate an animal as a dog, right? No we can't, and that is the first big reason why anonymous fields don't offer exactly what inheritance does. But, we need a way to achieve polymorphic behavior, so we define an interface.

```
// iAnimal will be used to achieve polymorphic behavior between different animal structs
type iAnimal interface {
	printName()
	makeCoolAnimalSound()
}
```

In go, an interface is just a list of methods. Any struct that has an implementation for all of the methods in the interface "implements" the interface. animal has a definition for all of the methods listed in the iAnimal interface, so it automatically implements the iAnimal interface. dog has animal as an anonymous field. This means it also has all of the methods listed in the iAnimal interface and thus implements it too. Therefore, if we want, both animal and dog can be assigned to iAnimal.

```
var polymorphicAnimal iAnimal
polymorphicAnimal = &animal{ name: "Bob the unknown" }
polymorphicAnimal.printName()
polymorphicAnimal.makeCoolAnimalSound()

var polymorphicDog iAnimal
polymorphicDog = &dog{ name: "Cozmo the dog" }
polymorphicDog.printName()
polymorphicDog.makeCoolAnimalSound()

```

However, there is a serious problem if we do it like this. We cannot access anything that animal or dog has that is not described by the interface. This means you cannot access any fields or dog's peeOnFireHydrant() function. You cannot cast from an interface to a struct. This brings us to how to properly use anonymous fields and interfaces.

### Proper Usage of interfaces and fields

The key is to separate polymorphic action logic from other logic. This means instead of taking a style where you create variables that are of a parent and instantiate them using a child, you should create child variables and write functions that accept interfaces, or potentially accept a list of interfaces (or some other data structure). Here is the proper way to write this example in go:

```
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
	fmt.Println(unknownAnimal.name)
	fmt.Println(cozmo.dogCollarBrand)
	cozmo.peeOnFireHydrant()
}

func runPolymorphicAnimalFuncs(polymorphicAnimal iAnimal) {
	polymorphicAnimal.printName()
	polymorphicAnimal.makeCoolAnimalSound()
}
```

## Summary

Normally in OOP you might do:
```
ParentClass instance = new ChildClass()
```
You might then do everything to instance and cast when you need to.
In go, you should separate child specific logic from parent logic and also separate out polymorphic action logic. This means:
1. Define parent struct, child struct, and the interface(s) for the actions (aka methods) that both the parent and child share.
2. If struct specific logic is needed, create the struct and perform that logic.
3. Use the interface to perform the polymorphic actions.

## Examples

See the examples folder for C# and go examples of the same program with nearly identical designs. 


Copyright (c) 2017 Nick Anderssohn