// Copyright (c) 2017 Nick Anderssohn
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

        public void PeeOnFireHydrant()
        {
          Console.WriteLine(Name + " marked his territory!");
        }
    }

    class Program
    {
        static void Main(string[] args)
        {
            Animal unknownAnimal = new Animal {Name = "Bob the unknown"};
            Animal cozmo = new Dog
            {
              Name = "Cozmo the dog",
              DogCollarBrand = "Barky"
            };

            RunPolymorphicAnimalFuncs(unknownAnimal);
            RunPolymorphicAnimalFuncs(cozmo);

            Console.WriteLine();
            Console.WriteLine(unknownAnimal.Name);
            Console.WriteLine(((Dog)cozmo).DogCollarBrand);
            ((Dog)cozmo).PeeOnFireHydrant();
        }

        static void RunPolymorphicAnimalFuncs(Animal polymorphicAnimal)
        {
            polymorphicAnimal.PrintName();
            polymorphicAnimal.MakeCoolAnimalSound();
        }
    }
}
