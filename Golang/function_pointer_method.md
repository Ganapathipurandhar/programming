# functions
- self contained units of code which carry out a certain job.
- help us divide a program into small manageable,repeatable and organisable chunks.
### why use functions?
• Reusability
• Abstraction

### functions syntax
func <function_name>(<params>) <return type> {
// body  of the function
}
### example
func addNumbers(a int, b int) int {
 sum := a + b
 return sum
}
### calling a function

<function_name>(<argument(s)>)    // ex: addNumbers(2, 3) or sumOfNumbers := addNumbers(2, 3)

### naming convention for functions
- must begin with a letter.
- can have any number of additional letters and symbols.
- cannot contain spaces.
- case-sensitive.

## parameters vs arguments

- **Function parameters** are the names listed in the function's definition.
~~~
func addNumbers(a int, b int) int {
 sum := a + b
 return sum
}
~~~

- **Function arguments** are the real values passed into the function.
~~~
func main() {
  sumOfNumbers := addNumbers(2, 3)
  fmt.Print(sumOfNumbers)
}
~~~




