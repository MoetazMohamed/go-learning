# MultiThreading

## Multi threading in C#
### Basic OS concept
Definition 
* MutlThreading : simultaneous execution of tasks over a certain period of time.
* Process : Program in execution (part of operating system that is repsonsible for executing an application)
* Thread is lightweight process (unit that execute code under program)
* Every program has logic (code) and a thread is repsonsible for executig the lgoic
* Every Program has a Main thread that execute the logic of the program
* C# program ==> Main method is the main thread that execute the logic of the program.
* By default every program is single threaded 
### Single threaded model
Single thread runs all processes present in the program in synchronizing manner (one after the other)

``` c#


    public void greet(){
        Console.WriteLine("Hello World");
    }

    public void score(){
         Console.WriteLine("scored a goal");
    }

    public static void main(){
        // Main thread that is responsible for executing the logic of the program
        // Single threaded model where the main method is a single thread that execute all the program processes in synchronizing manner one after the other
        score();
        greet();
    }

    // Single threaded model example
using System;
using System.Threading;

public class Geek
{
    public static void method1()
    {
        // It prints numbers from 0 to 4
        for (int i = 1; i < 5; i++)
        {

            Console.WriteLine("Method1 is : {0}", i);

            if (i == 2)
            {
                Thread.Sleep(100);
                // What will sleep is the main thread that is executed when we execute the program
            }
        }
    }
    public static void method2()
    {
        // It prints numbers from 0 to 4
        for (int j = 1; j < 5; j++)
        {

            Console.WriteLine("Method2 is : {0}", j);
        }
    }
}

public class Geeks
{
    static public void Main()
    {

        // Calling static methods
        Geek.method1();
        // There is only one method that is executing right now. Main thread is the only one running and if u ran different methods in the main method they all execute under the same thread which is the main thread 
        
        Geek.method2();
    }
}

```
* We want to execute multiple programs or tasks simultaneously. For example, I want to execute chrome while listening to songs on spotify. How can I do this ?
* Definition
* Race condition : when timing or order of execution affects the correctnes of piece of code
* Data race : occurs when processes have to acces the same resource concurrently