## Methods in Java

Methods are blocks of code that perform specific tasks and are defined within a class. They help in code reusability and modularity.

### Types of Methods (Based on Return Types)
- **Void Methods:** Do not return any value.  
    ```java
    void display() { ... }
    ```
- **Non-Void Methods:** Return a value of a specified type.  
    ```java
    int add(int a, int b) { return a + b; }
    ```
    - **Return Type:**
                - The return type specifies what type of value a method returns.
                - `void` indicates that the method does not return any value.

    - **Roles of Methods:**
                - **Encapsulation:** Group related code into reusable blocks.
                - **Abstraction:** Hide implementation details and expose only functionality.
                - **Code Reusability:** Avoid code duplication by calling methods multiple times.
                - **Modularity:** Break down complex problems into smaller, manageable pieces.

### Initializing Instance Variables

Instance variables in Java can be initialized in several ways:

- **By Object Reference:** Assign values directly using the object.
    ```java
    Example obj = new Example();
    obj.value = 10;
    ```

- **Instance Initialization Block (IIB):** A block inside the class that runs every time an object is created, before the constructor.
    ```java
    class Example {
        int value;
        {
            value = 20; // IIB
        }
    }
    ```

- **Constructors:** Assign values when the object is created.
    ```java
    class Example {
        int value;
        Example(int v) {
            value = v;
        }
    }
    ```

- **Setter Methods:** Use methods to set values after object creation.
    ```java
    class Example {
        int value;
        void setValue(int v) {
            value = v;
        }
    }
    ```

## Constructors in Java

Constructors are special methods used to initialize objects. They have the same name as the class and no return type.

### Types of Constructors
- **Default Constructor:** Provided by Java if no constructor is defined.  
    ```java
    class Example {
            Example() { }
    }
    ```
- **No-Argument Constructor:** Defined by the programmer without parameters.
- **Parameterized Constructor:** Takes arguments to initialize fields.
- **Copy Constructor:** Used to create a copy of an object (not built-in, but can be defined).

## Important Keywords

- **this:** Refers to the current object instance.
    ```java
    this.value = value;
    ```
- **super:** Refers to the parent class object.
    ```java
    super.methodName();
    ```
- **this():** Calls another constructor in the same class.
    ```java
    this(arg1, arg2);
    ```
- **super():** Calls the parent class constructor.
    ```java
    super();
    ```

## Class Loading

Class loading is the process by which the Java Virtual Machine (JVM) loads class files into memory. It involves:
- **Loading:** Finding and importing the class.
- **Linking:** Verifying and preparing the class.
- **Initialization:** Executing static initializers and assigning static variables.
