# Golang  

+ Created by the engineers at Google.  

### Go was created to combine:  

- The ease of programming of an interpreted, dynamically typed language (such as Python).  
- The efficiency and safety of a statically typed, compiled language (such as C++).  
- Modern capabilities with support for networked and multicore computing. 

~~~
package main             // for package
import ("fmt")           // fmt == format
// this is for comment
func main() {            // main --> main function for execution
  fmt.Println("Hello World!")
}
~~~

### comments 
single line commands : //
multiline comment : 
~~~
/* this
 is the multiline command
 */
~~~

### why are data types needed ?
- categorize a set of related values
- describe the operations that can be done on them
- define the way the data is stored

### dynamically types : python, javascript
- no need to declare data type
### dynamically types : c, c++, java, go
- need to declare data type( def add(int a, int b){retrun a+b;})

### Static typed advantages:
 - Better performance.
- Bugs can often be caught by a compiler.
- Better data integrity.

### Dynamic typed advantages:
- Faster to write code.
- Generally, less rigid.
- Shorter learning curve.

## Golang

- Go has a concept of types that is either explicitly declared by a programmer or is inferred by the compiler.
- It's a fast, statically typed,compiled language that feels like a dynamically typed, interpreted language.

~~~
package main             
import ("fmt")           

func main() {            
    name:= "Purandhar"
  fmt.Println(name)
}
~~~
# Golang Data Types

Go provides a variety of built-in data types to handle different kinds of data. These types can be categorized into basic types, composite types, reference types, and interface types.

## 1. Basic Data Types

### 1.1 Boolean Type
- **bool**: Represents a boolean value (true or false).
- Default value: `false`.

```go
var isActive bool = true
```

### 1.2 Numeric Types
#### Integer Types
| Type       | Size (bits) | Range |
|------------|------------|--------------------------------|
| int8       | 8          | -128 to 127                    |
| int16      | 16         | -32,768 to 32,767              |
| int32      | 32         | -2,147,483,648 to 2,147,483,647|
| int64      | 64         | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| int        | Depends on architecture (32 or 64 bits) |

```go
var x int = 100
var y int64 = 100000
```

#### Unsigned Integer Types
| Type       | Size (bits) | Range |
|------------|------------|--------------------------------|
| uint8      | 8          | 0 to 255                      |
| uint16     | 16         | 0 to 65,535                   |
| uint32     | 32         | 0 to 4,294,967,295           |
| uint64     | 64         | 0 to 18,446,744,073,709,551,615 |
| uint       | Depends on architecture (32 or 64 bits) |

```go
var a uint8 = 255
var b uint32 = 40000
```

#### Floating-Point Types
| Type      | Size (bits) | Precision |
|-----------|------------|------------|
| float32   | 32         | ~6-7 decimal digits |
| float64   | 64         | ~15 decimal digits |

```go
var pi float64 = 3.1415926535
```

#### Complex Types
| Type      | Size (bits) |
|-----------|------------|
| complex64 | 64         |
| complex128| 128        |

```go
var c complex128 = complex(5, 7) // 5 + 7i
```

### 1.3 String Type
- Immutable sequence of characters.
- UTF-8 encoded.

```go
var name string = "Golang"
```

## 2. Composite Types

### 2.1 Array
- Fixed-length collection of elements of the same type.

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5}
```

### 2.2 Slice
- Dynamic-sized, more flexible than arrays.

```go
var slice []int = []int{1, 2, 3, 4, 5}
```

### 2.3 Map
- Key-value pair collection.

```go
var m map[string]int = map[string]int{"one": 1, "two": 2}
```

### 2.4 Struct
- User-defined type that groups fields.

```go
type Person struct {
    Name string
    Age  int
}
var p = Person{"Alice", 25}
```

## 3. Reference Types

### 3.1 Pointer
- Stores memory address.

```go
var ptr *int
```

### 3.2 Function
- Functions can be assigned to variables.

```go
func add(a int, b int) int {
    return a + b
}
var fn func(int, int) int = add
```

### 3.3 Channel
- Used for goroutine communication.

```go
var ch chan int = make(chan int)
```

## 4. Interface Type
- Defines behavior without implementation.

```go
type Shape interface {
    Area() float64
}
```

# Golang Variables

## Declaring Variables
Go provides multiple ways to declare variables.

### Using `var` Keyword
```go
var name string = "Golang"
var age int = 10
var isAvailable bool = true
```

### Type Inference
```go
var language = "Go" // Type inferred as string
var number = 42     // Type inferred as int
```

### Short Variable Declaration (`:=`)
```go
message := "Hello, Go!" // Implicitly inferred as string
count := 5             // Implicitly inferred as int
```

## Multiple Variable Declaration
```go
var (
    firstName string = "John"
    lastName  string = "Doe"
    age       int    = 30
)
```

### Constants
Constants are declared using the `const` keyword.
```go
const Pi float64 = 3.14159
const Greeting = "Hello, World!"
```

## Zero Values
If a variable is declared without an initial value, it gets a default "zero value".
```go
var defaultInt int       // 0
var defaultString string // ""
var defaultBool bool     // false
```

## Variable Scope
- **Local variables**: Declared inside a function and accessible only within it.
- **Global variables**: Declared outside functions and accessible throughout the package.

```go
var globalVar = "I am global"

func main() {
    localVar := "I am local"
    fmt.Println(globalVar)
    fmt.Println(localVar)
}

# Golang Format Specifiers

Go provides various format specifiers for formatting input and output using the `fmt` package.

## General Format Specifiers
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%v` | Default format for the value | `fmt.Printf("Value: %v\n", 42)` → `Value: 42` |
| `%+v` | Struct with field names | `fmt.Printf("Struct: %+v\n", person)` → `Struct: {Name:John Age:30}` |
| `%#v` | Go syntax representation | `fmt.Printf("Go Syntax: %#v\n", person)` → `Go Syntax: main.Person{Name:"John", Age:30}` |
| `%T` | Type of the value | `fmt.Printf("Type: %T\n", 42)` → `Type: int` |
| `%%` | Literal percent sign | `fmt.Printf("Percentage: %%\n")` → `Percentage: %` |

## Integer Format Specifiers
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%b` | Binary representation | `fmt.Printf("Binary: %b\n", 10)` → `Binary: 1010` |
| `%o` | Octal representation | `fmt.Printf("Octal: %o\n", 10)` → `Octal: 12` |
| `%d` | Decimal representation | `fmt.Printf("Decimal: %d\n", 10)` → `Decimal: 10` |
| `%x` | Hexadecimal (lowercase) | `fmt.Printf("Hexadecimal: %x\n", 255)` → `Hexadecimal: ff` |
| `%X` | Hexadecimal (uppercase) | `fmt.Printf("Hexadecimal: %X\n", 255)` → `Hexadecimal: FF` |
| `%c` | Character representation | `fmt.Printf("Character: %c\n", 65)` → `Character: A` |

## Floating-Point Format Specifiers
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%e` | Scientific notation (lowercase) | `fmt.Printf("Scientific: %e\n", 123.456)` → `Scientific: 1.234560e+02` |
| `%E` | Scientific notation (uppercase) | `fmt.Printf("Scientific: %E\n", 123.456)` → `Scientific: 1.234560E+02` |
| `%f` | Decimal notation | `fmt.Printf("Decimal: %f\n", 123.456)` → `Decimal: 123.456000` |
| `%g` | Compact representation | `fmt.Printf("Compact: %g\n", 123.456)` → `Compact: 123.456` |
| `%G` | Compact representation (uppercase) | `fmt.Printf("Compact: %G\n", 123.456)` → `Compact: 123.456` |

## String Format Specifiers
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%s` | String representation | `fmt.Printf("String: %s\n", "Hello")` → `String: Hello` |
| `%q` | Quoted string | `fmt.Printf("Quoted: %q\n", "Hello")` → `Quoted: "Hello"` |
| `%x` | Hex representation of bytes | `fmt.Printf("Hex Bytes: %x\n", "Hello")` → `Hex Bytes: 48656c6c6f` |

## Boolean Format Specifier
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%t` | Boolean value | `fmt.Printf("Boolean: %t\n", true)` → `Boolean: true` |

## Pointer Format Specifier
| Specifier | Description | Example |
|-----------|-------------|---------|
| `%p` | Pointer address | `var x = 42; fmt.Printf("Pointer: %p\n", &x)` → `Pointer: 0xc0000140a0` |

These format specifiers help in printing data in various formats effectively in Go!

## Input Formatting
Golang uses `fmt.Scan`, `fmt.Scanf`, and `fmt.Scanln` for taking user input.

### Using `fmt.Scan`
Reads space-separated values into variables.
```go
var name string
var age int
fmt.Print("Enter name and age: ")
fmt.Scan(&name, &age)
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

### Using `fmt.Scanf`
Reads formatted input.
```go
var name string
var age int
fmt.Print("Enter name and age: ")
fmt.Scanf("%s %d", &name, &age)
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

### Using `fmt.Scanln`
Reads input until a newline is encountered.
```go
var name string
fmt.Print("Enter your name: ")
fmt.Scanln(&name)
fmt.Printf("Hello, %s!\n", name)
```

### Input with Multiple Data Types
Reads a string, an integer, a float, and a boolean value.
```go
package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var height float64
	var isStudent bool

	fmt.Print("Enter your name, age, height, and student status (true/false): ")
	fmt.Scanf("%s %d %f %t", &name, &age, &height, &isStudent)

	fmt.Printf("Name: %s\nAge: %d\nHeight: %.2f\nStudent: %t\n", name, age, height, isStudent)
}
```

This program takes multiple data types as input and prints them in a formatted way.
## In Go, fmt.Scanf returns two values:

- **count** → The number of successfully scanned items.
- **error** → An error value if the input does not match the expected format.
In the given code:
~~~
package main

import (
	"fmt"
)

func main() {
	var a string
	var b int

	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a string followed by an integer.")
		return
	}

	fmt.Println("count:", count)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

~~~
# print the of data type name using %T or "reflect.TypeOf()" 

~~~
package main
import "fmt"

func main() {
var grades int = 42
    var message string = "hello world"
    var isCheck bool = true
    var amount float32 = 5466.54

    fmt.Printf("variable grades = &v is of type %T \n", grades, grades)
    fmt.Printf ("variable message = '%v' is of type %T \n", message, message)
    fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
    fmt.Printf ("variable amount = %v is of type %T \n", amount, amount)
}
~~~

### reflect.TypeOf()

~~~
package main
import (
"fmt"
"reflect"
)

func main() {
fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
fmt.Printf("Type: %v \n", reflect.TypeOf("priyanka"))
fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
fmt.Printf("Type: %v \n", reflect.TypeOf(true))

}

>>> go run main.go

Type: int
Type: string
Type: float64
Type: bool
~~~

~~~
package main
import (
"fmt"
"reflect")

func main() {
var grades int = 42
var message string = "hello world"

fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
fmt.Printf ("variable message='%v' is of type %v \n", message, reflect.TypeOf(message))

}
~~~

>>> go run main.go
variable grades = 42 is of type int
variable message ='hello world' is of type string

# Go Type Casting

## Introduction
Type casting in Go (Golang) refers to converting a variable from one type to another. Unlike some other languages, Go does not support implicit type conversions, so all type conversions must be explicit.

## Syntax for Type Casting
```go
var newTypeVariable = TypeName(variable)
```

## Example: Integer to Float Conversion
```go
package main
import "fmt"

func main() {
    var intVal int = 42
    var floatVal float64 = float64(intVal) // Explicit conversion
    fmt.Println("Converted Value:", floatVal)
}
```

## Example: Float to Integer Conversion
```go
package main
import "fmt"

func main() {
    var floatVal float64 = 42.7
    var intVal int = int(floatVal) // Truncates decimal part
    fmt.Println("Converted Value:", intVal)
}
```

## Example: String to Integer Conversion
Using the `strconv` package:
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    intVal, err := strconv.Atoi(str) // Converts string to integer
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }
    fmt.Println("Converted Value:", intVal)
}
```

## Example: Integer to String Conversion
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    intVal := 123
    str := strconv.Itoa(intVal) // Converts integer to string
    fmt.Println("Converted Value:", str)
}
```

## Type Assertions (for Interface{})
Type assertions are used to extract the underlying value of an interface type.
```go
package main
import "fmt"

func main() {
    var i interface{} = "hello"
    str, ok := i.(string) // Asserting that i is of type string
    if ok {
        fmt.Println("String Value:", str)
    } else {
        fmt.Println("Type assertion failed")
    }
}
```

## Summary
- **Explicit type conversion** is required in Go.
- **Use strconv package** for string to numeric conversions.
- **Use type assertions** to extract values from `interface{}`.

Understanding type casting helps prevent type-related errors and ensures type safety in Go programs.
# Go Constants

## Introduction
Constants in Go are immutable values that are set at compile-time and cannot be changed during program execution. They are declared using the `const` keyword.

## Declaring Constants
```go
const Pi = 3.14159
const Greeting = "Hello, Go!"
```

## Multiple Constants Declaration
```go
const (
    Pi      = 3.14159
    Gravity = 9.8
    Language = "Go"
)
```

## Typed and Untyped Constants
- **Typed Constants**: Have a specific type assigned.
- **Untyped Constants**: Do not have a specific type until they are used in a context that requires one.

```go
const typedInt int = 42    // Typed constant
const untypedInt = 42      // Untyped constant
```

## Constant Expressions
Constants can be defined using expressions, but only if the result can be evaluated at compile-time.

```go
const (
    A = 2
    B = A * 3   // Valid
    C = B + 4   // Valid
)
```

## Enumerated Constants with `iota`
The `iota` identifier is used to create a sequence of related constants.

```go
const (
    Red = iota   // 0
    Green        // 1
    Blue         // 2
)
```

### Using `iota` with Bit Shifting
```go
const (
    Read   = 1 << iota // 1
    Write              // 2
    Execute            // 4
)
```

## Summary
- Constants are immutable and declared using `const`.
- They can be **typed** or **untyped**.
- Use `iota` for sequential values.
- Constants must be **evaluatable at compile-time**.

Understanding constants is crucial for writing efficient and safe Go programs.
# Go Operators

## Introduction
Operators in Go are symbols that perform operations on operands. Go supports various types of operators, including arithmetic, comparison, logical, bitwise, and assignment operators.

## Arithmetic Operators
Arithmetic operators perform mathematical operations.

| Operator | Description | Example |
|----------|-------------|---------|
| `+`  | Addition  | `a + b`  |
| `-`  | Subtraction  | `a - b`  |
| `*`  | Multiplication  | `a * b`  |
| `/`  | Division  | `a / b`  |
| `%`  | Modulus (remainder)  | `a % b`  |

Example:
```go
package main
import "fmt"

func main() {
    a, b := 10, 5
    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Division:", a/b)
    fmt.Println("Modulus:", a%b)
}
```

## Comparison Operators
Comparison operators compare two values and return a boolean result.

| Operator | Description | Example |
|----------|-------------|---------|
| `==` | Equal to | `a == b` |
| `!=` | Not equal to | `a != b` |
| `>` | Greater than | `a > b` |
| `<` | Less than | `a < b` |
| `>=` | Greater than or equal to | `a >= b` |
| `<=` | Less than or equal to | `a <= b` |

Example:
```go
package main
import "fmt"

func main() {
    a, b := 10, 20
    fmt.Println("a == b:", a == b)
    fmt.Println("a != b:", a != b)
    fmt.Println("a > b:", a > b)
    fmt.Println("a < b:", a < b)
    fmt.Println("a >= b:", a >= b)
    fmt.Println("a <= b:", a <= b)
}
```

## Logical Operators
Logical operators are used to combine boolean expressions.

| Operator | Description | Example |
|----------|-------------|---------|
| `&&` | Logical AND | `a && b` |
| `||` | Logical OR | `a || b` |
| `!` | Logical NOT | `!a` |

Example:
```go
package main
import "fmt"

func main() {
    a, b := true, false
    fmt.Println("a && b:", a && b)
    fmt.Println("a || b:", a || b)
    fmt.Println("!a:", !a)
}
```

## Bitwise Operators
Bitwise operators perform bitwise computations on integers.

| Operator | Description | Example |
|----------|-------------|---------|
| `&` | Bitwise AND | `a & b` |
| `|` | Bitwise OR | `a | b` |
| `^` | Bitwise XOR | `a ^ b` |
| `<<` | Left shift | `a << 1` |
| `>>` | Right shift | `a >> 1` |

Explanation and Output:
```go
package main
import "fmt"

func main() {
    a, b := 5, 3 // 5 = 0101, 3 = 0011 in binary
    fmt.Println("a & b:", a & b)  // 0101 & 0011 = 0001 (1)   binary + binary ==> convent to decimal (use AND and OR  true tables)
    fmt.Println("a | b:", a | b)  // 0101 | 0011 = 0111 (7)
    fmt.Println("a ^ b:", a ^ b)  // 0101 ^ 0011 = 0110 (6)
    fmt.Println("a << 1:", a << 1) // 0101 << 1 = 1010 (10)
    fmt.Println("a >> 1:", a >> 1) // 0101 >> 1 = 0010 (2)
}
```

## Assignment Operators
Assignment operators assign values to variables.

| Operator | Description | Example |
|----------|-------------|---------|
| `=` | Assign | `a = b` |
| `+=` | Add and assign | `a += b` |
| `-=` | Subtract and assign | `a -= b` |
| `*=` | Multiply and assign | `a *= b` |
| `/=` | Divide and assign | `a /= b` |
| `%=` | Modulus and assign | `a %= b` |

Example:
```go
package main
import "fmt"

func main() {
    a := 10
    a += 5
    fmt.Println("a after += 5:", a)
    a -= 3
    fmt.Println("a after -= 3:", a)
    a *= 2
    fmt.Println("a after *= 2:", a)
    a /= 4
    fmt.Println("a after /= 4:", a)
    a %= 3
    fmt.Println("a after %= 3:", a)
}
```

## Summary
- **Arithmetic operators** perform mathematical operations.
- **Comparison operators** compare values.
- **Logical operators** evaluate boolean expressions.
- **Bitwise operators** manipulate bits.
- **Assignment operators** assign values and perform compound operations.

Understanding operators is essential for performing computations and logical operations in Go programming.
# Go Conditions

Go provides several types of conditional statements to control the flow of execution in your program. These include `if`, `else`, `else if`, `switch`, and `select` (for concurrency). This document outlines their usage.

## `if` Statement

The `if` statement is used to execute a block of code if a specified condition evaluates to true.

### Syntax:

```go
if condition {
    // code to be executed if the condition is true
}
~~~
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}

~~~
## else Statement
The else statement is used to execute a block of code if the condition in the if statement evaluates to false.
~~~
if condition {
    // code to be executed if the condition is true
} else {
    // code to be executed if the condition is false
}
~~~
~~~
package main

import "fmt"

func main() {
    x := 3
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
~~~
## else if Statement
The else if statement is used to check multiple conditions sequentially. It allows you to add more than one conditional branch.
~~~
if condition1 {
    // code to be executed if condition1 is true
} else if condition2 {
    // code to be executed if condition2 is true
} else {
    // code to be executed if none of the conditions are true
}
~~~
~~~
package main

import "fmt"

func main() {
    x := 7
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else if x > 5 {
        fmt.Println("x is greater than 5 but less than or equal to 10")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
~~~
## switch Statement
The switch statement allows you to select one of many code blocks to be executed. Its often used as an alternative to multiple if-else conditions.
~~~
switch expression {
case value1:
    // code to be executed if the expression matches value1
case value2:
    // code to be executed if the expression matches value2
default:
    // code to be executed if the expression doesn't match any value
}
~~~
~~~
package main

import "fmt"

func main() {
    x := 2
    switch x {
    case 1:
        fmt.Println("x is 1")
    case 2:
        fmt.Println("x is 2")
    default:
        fmt.Println("x is neither 1 nor 2")
    }
}
~~~
## fallthrough

The fallthrough keyword is used in switch-case to force the execution flow to fall through the successive case block.miain.go
~~~
func main() {
var i int = 10
switch i {
    case -5:
       fmt.Println("-5")
    case 10:
       fmt.Println("10")
       fallthrough
    case 20:
       fmt.Println("20")
       fallthrough
    default:
    fmt.Println("default")
}
~~~
## select Statement
The select statement is used in Go to handle multiple channels in concurrent programming. It allows a goroutine to wait on multiple communication operations.
~~~
select {
case ch1 := <-chan1:
    // code to be executed if data is received from chan1
case ch2 := <-chan2:
    // code to be executed if data is received from chan2
default:
    // code to be executed if no channels are ready
}

~~~
~~~
package main

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        ch1 <- "hello from ch1"
    }()

    go func() {
        ch2 <- "hello from ch2"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
~~~

# Loops in Go

## For Loop
Go has only one looping construct: the `for` loop. It can be used in different ways.

### Basic `for` Loop
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

### `for` as a While Loop
```go
package main
import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

### Infinite Loop
```go
package main
import "fmt"

func main() {
    for {
        fmt.Println("Infinite loop")
    }
}
```

### Looping Over a Slice
```go
package main
import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### Looping Over a Map
```go
package main
import "fmt"

func main() {
    myMap := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range myMap {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

### Breaking and Continuing Loops
#### Using `break`
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            break
        }
        fmt.Println(i)
    }
}
```

#### Using `continue`
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i == 3 {
            continue
        }
        fmt.Println(i)
    }
}
```

### Nested Loops
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
```

### Looping Over Strings (Rune Iteration)
```go
package main
import "fmt"

func main() {
    str := "Hello"
    for index, char := range str {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }
}

# Arrays, Slices, and Maps in Go

## Arrays
An array in Go is a fixed-size collection of elements of the same type. with sequenal memory address

### Declaring and Initializing an Array
```go
package main
import "fmt"

func main() {
    var arr [5]int // Declaring an array of size 5
    fmt.Println(arr) // Outputs: [0 0 0 0 0]

    arr2 := [3]int{1, 2, 3} // Declaring and initializing an array
    fmt.Println(arr2) // Outputs: [1 2 3]
}
```

### Accessing Elements
```go
fmt.Println(arr2[1]) // Outputs: 2
```

### Modifying an Array
```go
arr[2] = 42
fmt.Println(arr) // Outputs: [0 0 42 0 0]
```

---

## Slices
A slice is a dynamically-sized, more powerful abstraction of an array.

### Declaring and Initializing Slices
```go
package main
import "fmt"

func main() {
    var slice []int // Declaring a slice (nil by default)
    fmt.Println(slice) // Outputs: []

    slice2 := []int{1, 2, 3} // Initializing a slice
    fmt.Println(slice2) // Outputs: [1 2 3]
}
```

### Creating a Slice from an Array
```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Slice from index 1 to 3 (excluding 4)
fmt.Println(slice) // Outputs: [2 3 4]
```

### Using `make` to Create a Slice
```go
slice := make([]int, 5) // Creates a slice of length 5
fmt.Println(slice) // Outputs: [0 0 0 0 0]
```

### Appending to a Slice
```go
slice = append(slice, 6, 7, 8)
fmt.Println(slice) // Outputs: [0 0 0 0 0 6 7 8]
```

### Iterating Over a Slice
```go
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

---

## Maps
A map is a key-value data structure.

### Declaring and Initializing Maps
```go
package main
import "fmt"

func main() {
    myMap := make(map[string]int) // Declaring a map
    myMap["one"] = 1
    myMap["two"] = 2
    fmt.Println(myMap) // Outputs: map[one:1 two:2]

    myMap2 := map[string]int{"a": 10, "b": 20} // Initializing a map
    fmt.Println(myMap2) // Outputs: map[a:10 b:20]
}
```

### Accessing Values in a Map
```go
value := myMap["one"]
fmt.Println(value) // Outputs: 1
```

### Checking if a Key Exists
```go
val, exists := myMap["three"]
if exists {
    fmt.Println("Key exists with value", val)
} else {
    fmt.Println("Key does not exist")
}
```

### Deleting a Key from a Map
```go
delete(myMap, "one")
fmt.Println(myMap) // Outputs: map[two:2]
```

### Iterating Over a Map
```go
for key, value := range myMap2 {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}


