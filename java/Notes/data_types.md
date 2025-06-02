# Java Data Types

Java data types specify the size and type of values that can be stored in a variable. They are mainly categorized into **primitive** and **non-primitive** types.

## 1. Primitive Data Types

Primitive data types are the most basic data types built into the Java language. There are 8 primitive types:

| Type     | Size    | Description                  | Example         |
|----------|---------|-----------------------------|-----------------|
| `byte`   | 1 byte  | 8-bit signed integer        | `byte b = 10;`  |
| `short`  | 2 bytes | 16-bit signed integer       | `short s = 20;` |
| `int`    | 4 bytes | 32-bit signed integer       | `int i = 100;`  |
| `long`   | 8 bytes | 64-bit signed integer       | `long l = 100L;`|
| `float`  | 4 bytes | 32-bit floating point       | `float f = 10.5f;`|
| `double` | 8 bytes | 64-bit floating point       | `double d = 20.99;`|
| `char`   | 2 bytes | Single 16-bit Unicode char  | `char c = 'A';` |
| `boolean`| 1 bit   | true or false value         | `boolean b = true;`|

## 2. Non-Primitive Data Types

Non-primitive data types are also called reference types because they refer to objects. Examples include:

- **String**: Represents a sequence of characters.  
    Example: `String str = "Hello";`
- **Arrays**: Collection of similar data types.  
    Example: `int[] arr = {1, 2, 3};`
- **Classes**: User-defined blueprints for objects.
- **Interfaces**: Abstract types used to specify behavior.

## 3. Class Data Types in Java

In Java, **class data types** are non-primitive types that refer to objects created from user-defined or built-in classes. When you declare a variable of a class type, it holds a reference to an object of that class.
    - Object is data then class is a data type. this prove the class in non-primitive types ( ex: Car myCar = new Car(); car = class_name, car()--> non-primitive datatype )

For example, consider a `Car` class:

```java
class Car {
    String color;
    int year;

    void drive() {
        System.out.println("Driving...");
    }
}
```

You can create an object (instance) of the `Car` class and assign it to a variable:

```java
Car myCar = new Car();
myCar.color = "Red";
myCar.drive();
```

Here, `myCar` is a variable of the class data type `Car`. It stores a reference to a `Car` object in memory.
Class data types allow you to model real-world entities and their behaviors in your Java programs.
