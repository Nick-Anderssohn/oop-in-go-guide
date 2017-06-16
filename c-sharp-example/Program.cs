using System;

namespace c_sharp_example
{
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
    }

    class Program
    {
        static void Main(string[] args)
        {
          // Can have an instance of the parent class
            Animal unknownAnimal = new Animal {Name = "Bob the unknown"};
            unknownAnimal.PrintName();
            unknownAnimal.MakeCoolAnimalSound();

            Animal polymorphicDog = new Dog
            {
              Name = "Cozmo the dog",
              DogCollarBrand = "Barky"
            };
            polymorphicDog.PrintName();
            polymorphicDog.MakeCoolAnimalSound();
            Console.WriteLine(((Dog)polymorphicDog).DogCollarBrand);
        }
    }
}
