# Core-java
- **Software**: it is collection of program using which human problem can be sloved with less effect
- **Program**: it is set of instruction given to hardware machine in order to solved the problem
- **Programming Language**: Any software language which used in order to build a program is called programming language like c, c++, java, python etc
- **Platform**: A Platform is a Place or envronment given in order to run a software application or program
   - there are three Platforms
    1) linux
    2) MS windows
    3) MAC OS

## Object orientation: 
  - pursuing and seeing everything in the prospective of object is called Object orientation (facts and touch).
  ### java: example for Object orientation
   - java is an Object oriented programming langauge also is a platform independed programming language. Jave see's everything around itself as a prospective of object **(java --> computer, car, application, remote etc)**

## Association of Objects (HAS-A Relationship)

In object-oriented programming, association represents the relationship between two separate classes, established through their objects.

### HAS-A Relationship
When an object of one class contains a reference to an object of another class, it is called a HAS-A relationship. This helps in code reusability and models real-world relationships between objects.

There are two main types of HAS-A relationships:

### 1. Composition
- **Definition**: Composition is a strong form of association where one class owns another class. The contained object cannot exist independently of the container object.
- **Example**: A `House` class has `Room` objects. If the `House` is destroyed, its `Room` objects are also destroyed.
- **Java Example**:
    ```java
    class Room {
        // Room properties and methods
    }

    class House {
        private Room room; // House HAS-A Room

        public House() {
            room = new Room();
        }
    }
    ```

### 2. Aggregation
- **Definition**: Aggregation is a weaker form of association where one class contains references to other classes, but the contained objects can exist independently.
- **Example**: A `Library` class has `Book` objects. If the `Library` is destroyed, the `Book` objects can still exist elsewhere.
- **Java Example**:
    ```java
    class Book {
        // Book properties and methods
    }

    class Library {
        private List<Book> books; // Library HAS-A collection of Books

        public Library(List<Book> books) {
            this.books = books;
        }
    }
    ```

**Summary**:  
- **Composition**: Strong ownership, lifecycle is tied.
- **Aggregation**: Weak ownership, lifecycle is independent.

## Diagram: HAS-A Relationship

Below is a simple diagram to illustrate Composition and Aggregation:

```
Composition (filled diamond):
[House] ◆───[Room]

Aggregation (empty diamond):
[Library] ◇───[Book]
```

- **Composition**: The filled diamond (◆) shows a strong relationship; if the container is destroyed, so are its parts.
- **Aggregation**: The empty diamond (◇) shows a weaker relationship; the parts can exist independently of the container.

---
## JDK, JRE, JVM

### JVM (Java Virtual Machine)
- **Definition**: JVM is an abstract machine that enables your computer to run Java programs. It provides a runtime environment in which Java bytecode can be executed.
- **Role**: Converts bytecode into machine-specific code so it can be executed on any platform.
- **Key Point**: JVM is platform-dependent (different JVMs for Windows, Linux, etc.), but Java bytecode is platform-independent.

### JRE (Java Runtime Environment)
- **Definition**: JRE is a package of software that provides the libraries, JVM, and other components to run applications written in Java.
- **Role**: Contains everything needed to run Java programs, but not to develop them.
- **Includes**: JVM + core libraries + supporting files.

### JDK (Java Development Kit)
- **Definition**: JDK is a full-featured software development kit required to develop Java applications and applets.
- **Role**: Provides tools for developing, debugging, and monitoring Java applications.
- **Includes**: JRE + development tools (compiler, debugger, etc.).

---
### Diagram: JDK, JRE, JVM Relationship (Detailed)

```
+---------------------------------------------------+
|                     JDK                           |
|  (Java Development Kit)                           |
|  - Development tools (javac, javadoc, debugger)   |
|  - Java source code (.java files)                 |
|                                                   |
|   +-------------------------------------------+   |
|   |                  JRE                      |   |
|   |   (Java Runtime Environment)              |   |
|   |   - Java class libraries                  |   |
|   |   - Supporting files                      |   |
|   |                                           |   |
|   |    +----------------------------------+   |   |
|   |    |              JVM                 |   |   |
|   |    |   (Java Virtual Machine)         |   |   |
|   |    |   - Class loader                 |   |   |
|   |    |   - Bytecode verifier            |   |   |
|   |    |   - Interpreter/JIT compiler     |   |   |
|   |    |   - Garbage collector            |   |   |
|   |    +----------------------------------+   |   |
|   +-------------------------------------------+   |
+---------------------------------------------------+
```

- **JDK**: Contains development tools (compiler, debugger, etc.) and the JRE.
- **JRE**: Contains the JVM and Java class libraries needed to run Java applications.
- **JVM**: Executes Java bytecode, manages memory, and provides platform independence.

- **JDK** contains **JRE**, which contains **JVM**.
- Use **JDK** for development, **JRE** for running Java programs, and **JVM** is the engine that executes Java bytecode.

## Platform Independence in Java

Java is known for its platform independence, which means that Java programs can run on any operating system without modification.

### How Platform Independence Works

1. **Source Code (.java)**: You write your Java program in a `.java` file.
2. **Compilation**: The Java compiler (`javac`) checks for syntax errors and converts the `.java` file into a `.class` file containing **bytecode**.
3. **Bytecode (.class)**: Bytecode is an intermediate, platform-independent code that is not specific to any operating system.
4. **JVM (Java Virtual Machine)**: Each operating system (Windows, Linux, macOS, etc.) has its own implementation of the JVM. The JVM reads and executes the bytecode on the respective platform. (JVM is also called an interpreter.)

### Diagram

```
YourCode.java
    |
    |  (javac)
    v
YourCode.class (bytecode)
    |
    |  (JVM for Windows)
    |  (JVM for Linux)
    |  (JVM for macOS)
    v
Runs on any OS!
```

**Summary:**  
- Write once, run anywhere: Java code is compiled to bytecode, which can be executed on any OS with a compatible JVM.
- The JVM abstracts away the underlying platform, enabling true platform independence.

### class landing

- when the .class file is created  by compiler and stored in main memory is called Class Loading.
- (**class loading**): the process of loading the .class file from the main memory to the JVM memory(heap) is called class loading.

## Class and Object in Java

### What is a Class?

A **class** in Java is a blueprint or template for creating objects. It defines the properties (fields/attributes) and behaviors (methods/functions) that the objects created from the class will have.

- **Syntax Example:**
    ```java
    class Car {
        // Fields (properties)
        String color;
        String model;

        // Method (behavior)
        void drive() {
            System.out.println("Car is driving");
        }
    }
    ```

- **Explanation:**
    - `Car` is a class.
    - It has two fields: `color` and `model`.
    - It has one method: `drive()`.

### What is an Object?

An **object** is an instance of a class. It represents a real-world entity and has its own state (**values for fields or variables**) and behavior (can call **methods**).

### Syntax for Creating an Object in Java

```java
ClassName objectName = new ClassName();
```

- **Example:**
    ```java
    Car myCar = new Car();
    ```
- `Car` is the class name.
- `myCar` is the object name (reference variable).
- `new Car()` creates a new object of the `Car` class.

- **Creating an Object:**
    ```java
    Car myCar = new Car();
    myCar.color = "Red";
    myCar.model = "Sedan";
    myCar.drive(); // Output: Car is driving
    ```

- **Explanation:**
    - `myCar` is an object of the `Car` class.
    - It has its own `color` and `model` values.
    - It can use the `drive()` method defined in the class.

### Key Points

- A class is a template; an object is a real instance.
- Multiple objects can be created from a single class, each with its own data.
- Objects interact with each other using methods.

### Diagram

```
Class (Blueprint)
   |
   |--- Object 1 (myCar): color = "Red", model = "Sedan"
   |--- Object 2 (yourCar): color = "Blue", model = "SUV"
```

**Summary:**  
- **Class**: Defines structure and behavior.
- **Object**: Actual entity with state and behavior, created from a class.

## Variables in Java

A **variable** in Java is a container that holds data which can be changed during the execution of a program. Variables are used to store information to be referenced and manipulated in a program.
- works as right to left (a=10)

### Types of Variables
- 1) Local Variables
- 2) Instance Variables (global variables)
- 3) Static Variables

1. **Local Variables**
    - Declared inside methods, constructors, or blocks.
    - local variable gets the life only when the method/constructors/blocks is called.
    - Only accessible within the method/block where they are declared.
    - Must be initialized before use.

    ```java
    void exampleMethod() {
         int localVar = 10; // local variable
         System.out.println(localVar);
    }
    ```
~~~
public class mobile {
	String name= "redmi";
	int cost=25000;
	
	void call() {
		long number=9948094438l;
		System.out.println("ringing.... "+ number);
	}

	public static void main(String[] args) {
     mobile m1= new mobile();
     m1.call();   
	}
}
~~~
- **output**
~~~
ringing.... 9948094438
~~~

2. **Instance Variables**
    - Declared inside a class but outside any method, constructor, or block.
    - Each object has its own copy.
    - Also called non-static fields.

    ```java
    class Car {
         String color; // instance variable
    }
    ```
~~~
public class mobile {
	String name= "redmi";
	int cost=25000;
	
	void call() {
		System.out.println("ringing....");
	}

	public static void main(String[] args) {
     mobile m1= new mobile();
     System.out.println(m1.name);
     System.out.println(m1.cost);
     m1.call();
     mobile m2= new mobile();
     System.out.println(m2.name);
     System.out.println(m2.cost);
     m2.call();
	}

}
~~~
- **output**
~~~
redmi
25000
ringing....
redmi
25000
ringing....
~~~
- Any variable which is decleared inside the class and outside the method is called instances variable, Instances variable get the life only when the object is created.
- if we create 'n' no.of object then 'n' no.of times the instance variable get life

3. **Static Variables**
    - Declared with the `static` keyword inside a class.
    - Shared among all instances of the class.(static variable created once and copied to all the instances)
    - Also called class variables.
    - Static variable are created only once's
    - static variables can be access directly by using class_name and dot operator without creating and objects
     

    ```java
    class Car {
         static int numberOfCars; // static variable
    }
    ```
~~~
package purandhar_proj;

public class mobile {
	static String name= "redmi";
	
	static void call() {
		long number=9948094438l;
		System.out.println("ringing.... "+ number);
	}

	public static void main(String[] args) {
     System.out.println(mobile.name);
     mobile.call();
	}
}
~~~
- **output**
~~~
redmi
ringing.... 9948094438
~~~
### Variable Declaration and Initialization

- **Declaration**: Tells the compiler about the variable's name and type.
- **Initialization**: Assigns a value to the variable.

```java
int age;         // declaration
age = 25;        // initialization
int year = 2024; // declaration and initialization
```

### Naming Rules
- Must not start with a letter, `$`, or `_`.
- Cannot start with a digit.
- Cannot use Java reserved keywords.
- Case-sensitive.

### Example

```java
class Student {
     String name;         // instance variable
     static String school = "ABC School"; // static variable

     void display() {
          int rollNumber = 101; // local variable
          System.out.println(name + " " + rollNumber + " " + school);
     }
}
```

### Summary

- Variables are used to store data.
- Java has local, instance, and static variables.
- Proper naming and initialization are important for correct program behavior.
varibles


## Java Identifier Rules

An **identifier** is the name given to elements such as classes, methods, variables, and interfaces in Java. Identifiers help uniquely identify these elements in your code.

### Rules for Identifiers in Java

1. **Allowed Characters**: (class numbers must be camelcase) 
    - Identifiers can contain letters (`A-Z`, `a-z`), digits (`0-9`), underscore (`_`), and dollar sign (`$`).
2. **Cannot Start with a Digit**:  
    - Identifiers must begin with a letter, underscore, or dollar sign.  (- but avoid _ and $ for class names as per convention)
    - Example: `int 1value;` // Invalid  
    - Example: `int value1;` // Valid
3. **No Spaces or Special Characters**:  
    - Spaces and special characters like `@`, `#`, `%`, etc., are not allowed.
4. **Case Sensitive**:  
    - Java is case sensitive, so `MyVar` and `myvar` are different identifiers.
5. **Cannot Use Reserved Keywords**:  
    - Java keywords (like `class`, `int`, `public`, etc.) cannot be used as identifiers.
6. **No Length Limit**:  
    - There is no restriction on the length of an identifier, but it is recommended to keep names meaningful and readable.

### Examples

| Valid Identifiers   | Invalid Identifiers |
|---------------------|--------------------|
| `myVar`             | `1stValue`         |
| `student123`        | `@score`           |

### Best Practices
- Use meaningful names that describe the purpose (e.g., `studentName`, `totalAmount`).
- Follow camelCase for variables and methods (`myVariable`), PascalCase for classes (`MyClass`).
- Avoid using `$` and `_` unless necessary.