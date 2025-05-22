# Java Solutions

### 1) Calculate Simple Interest
**Question:** Write a Java program to calculate the simple interest given the principal, rate of interest, and time period.
```java
import java.util.Scanner;
public class SimpleInterest {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter Principal: ");
        double principal = sc.nextDouble();
        System.out.print("Enter Rate: ");
        double rate = sc.nextDouble();
        System.out.print("Enter Time: ");
        double time = sc.nextDouble();
        double interest = (principal * rate * time) / 100;
        System.out.println("Simple Interest: " + interest);
    }
}
```

### 2) Convert Fahrenheit to Celsius
**Question:** Write a Java program to convert temperature from Fahrenheit to Celsius.
```java
import java.util.Scanner;
public class TemperatureConversion {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter Fahrenheit: ");
        double fahrenheit = sc.nextDouble();
        double celsius = (fahrenheit - 32) * 5/9;
        System.out.println("Celsius: " + celsius);
    }
}
```

### 3) Calculate Area and Circumference of Circle
**Question:** Write a Java program to calculate the area and circumference of a circle given its radius.
```java
import java.util.Scanner;
public class Circle {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter Radius: ");
        double radius = sc.nextDouble();
        double area = Math.PI * radius * radius;
        double circumference = 2 * Math.PI * radius;
        System.out.println("Area: " + area);
        System.out.println("Circumference: " + circumference);
    }
}
```

### 4) Arithmetic Operations
**Question:** Write a Java program to perform all basic arithmetic operations (addition, subtraction, multiplication, division, and modulus) on two numbers.
```java
import java.util.Scanner;
public class ArithmeticOperations {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter two numbers: ");
        double a = sc.nextDouble();
        double b = sc.nextDouble();
        System.out.println("Addition: " + (a + b));
        System.out.println("Subtraction: " + (a - b));
        System.out.println("Multiplication: " + (a * b));
        System.out.println("Division: " + (a / b));
        System.out.println("Modulus: " + (a % b));
    }
}
```

### 5) Calculate BMI
**Question:** Write a Java program to calculate the Body Mass Index (BMI) given weight (kg) and height (m).
```java
import java.util.Scanner;
public class BMI {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter weight (kg): ");
        double weight = sc.nextDouble();
        System.out.print("Enter height (m): ");
        double height = sc.nextDouble();
        double bmi = weight / (height * height);
        System.out.println("BMI: " + bmi);
    }
}
```

### 6) Area and Perimeter of Rectangle
**Question:** Write a Java program to calculate the area and perimeter of a rectangle given its length and breadth.
```java
import java.util.Scanner;
public class Rectangle {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter Length: ");
        double length = sc.nextDouble();
        System.out.print("Enter Breadth: ");
        double breadth = sc.nextDouble();
        System.out.println("Area: " + (length * breadth));
        System.out.println("Perimeter: " + (2 * (length + breadth)));
    }
}
```

### 7) Special Digits
**Question:** Write a Java program to check if a two-digit number is a special number (e.g., 89 → 8+9+8*9 = 89).
```java
import java.util.Scanner;
public class SpecialDigit {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter two-digit number: ");
        int num = sc.nextInt();
        int d1 = num / 10;
        int d2 = num % 10;
        if ((d1 + d2 + d1 * d2) == num) {
            System.out.println("Special Number");
        } else {
            System.out.println("Not a Special Number");
        }
    }
}
```

### 8) Even or Odd
**Question:** Write a Java program to check whether a number is even or odd.
```java
import java.util.Scanner;
public class EvenOdd {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter a number: ");
        int num = sc.nextInt();
        if (num % 2 == 0)
            System.out.println("Even");
        else
            System.out.println("Odd");
    }
}
```

### 9) Voting Eligibility
**Question:** Write a Java program to check if a person is eligible to vote (age should be 18 or above).
```java
import java.util.Scanner;
public class Voting {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter age: ");
        int age = sc.nextInt();
        if (age >= 18)
            System.out.println("Eligible to vote");
        else
            System.out.println("Not eligible to vote");
    }
}
```

### 10) Product Price Calculation
**Question:** Write a Java program to read the price of a product, calculate the total amount, and determine how much change needs to be returned to the customer.
```java
import java.util.Scanner;
public class ProductPrice {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter total amount: ");
        double totalAmount = sc.nextDouble();
        System.out.print("Enter amount given by customer: ");
        double givenAmount = sc.nextDouble();
        double change = givenAmount - totalAmount;
        System.out.println("Change to return: " + change);
    }
}
```
### 11. Program to Check if the Entered Input is a Digit

**Question:** Write a Java program to check if the user-entered input is a digit.

**Answer:**
```java
import java.util.Scanner;

public class CheckDigit {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter a character: ");
        char input = scanner.next().charAt(0);
        if (Character.isDigit(input)) {
            System.out.println("It is a digit.");
        } else {
            System.out.println("It is not a digit.");
        }
        scanner.close();
    }
}
```
# 2. Program to Validate Month Number
### Question: Write a Java program to read a month number and check if it is valid (1-12).

Answer:
~~~
import java.util.Scanner;

public class CheckMonth {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter month number (1-12): ");
        int month = scanner.nextInt();
        if (month >= 1 && month <= 12) {
            System.out.println("Valid month.");
        } else {
            System.out.println("Invalid month.");
        }
        scanner.close();
    }
}
~~~
# 13. Program to Check if a Number is Positive or Negative
### Question: Write a Java program to check if the user-entered number is positive or negative.

Answer:
~~~
import java.util.Scanner;

public class CheckNumber {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter a number: ");
        int number = scanner.nextInt();
        if (number > 0) {
            System.out.println("Positive number.");
        } else if (number < 0) {
            System.out.println("Negative number.");
        } else {
            System.out.println("Zero.");
        }
        scanner.close();
    }
}
~~~
# 14. Program to Read and Display Academic Details
### Question: Write a Java program to read and display a student's academic details (name, roll number, and course).

Answer:
~~~
import java.util.Scanner;

public class AcademicDetails {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter name: ");
        String name = scanner.nextLine();
        System.out.print("Enter roll number: ");
        int roll = scanner.nextInt();
        scanner.nextLine(); // Consume newline
        System.out.print("Enter course: ");
        String course = scanner.nextLine();
        System.out.println("Student Details: " + name + ", Roll No: " + roll + ", Course: " + course);
        scanner.close();
    }
}
~~~
# 15. Program to Find the Biggest Among Three Numbers
### Question: Write a Java program to read three integer values from the user and find the largest among them.

Answer:
~~~
import java.util.Scanner;

public class LargestOfThree {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter three numbers: ");
        int a = scanner.nextInt(), b = scanner.nextInt(), c = scanner.nextInt();
        int max = (a > b) ? (a > c ? a : c) : (b > c ? b : c);
        System.out.println("The largest number is: " + max);
        scanner.close();
    }
}
~~~
# 16. Program to Find the Biggest Among Four Numbers
### Question: Write a Java program to find the largest among four numbers.

Answer:
~~~
import java.util.Scanner;

public class LargestOfFour {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter four numbers: ");
        int a = scanner.nextInt(), b = scanner.nextInt(), c = scanner.nextInt(), d = scanner.nextInt();
        int max = Math.max(Math.max(a, b), Math.max(c, d));
        System.out.println("The largest number is: " + max);
        scanner.close();
    }
}
~~~
# 17. Program to Determine Student's Result Based on Marks
### Question: Write a Java program to read six subject marks and determine the result (Distinction, First Class, Second Class, Pass, or Fail).

Answer:

~~~
import java.util.Scanner;

public class StudentResult {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int[] marks = new int[6];
        System.out.println("Enter marks for six subjects:");
        for (int i = 0; i < 6; i++) {
            marks[i] = scanner.nextInt();
        }
        int total = 0;
        for (int mark : marks) {
            total += mark;
        }
        double average = total / 6.0;
        if (average >= 75) {
            System.out.println("Distinction");
        } else if (average >= 60) {
            System.out.println("First Class");
        } else if (average >= 50) {
            System.out.println("Second Class");
        } else if (average >= 40) {
            System.out.println("Pass");
        } else {
            System.out.println("Fail");
        }
        scanner.close();
    }
}
~~~
# 18. Program to Check Leap Year
### Question: Write a Java program to check if the entered year is a leap year or not.

Answer:
~~~
import java.util.Scanner;

public class LeapYear {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter a year: ");
        int year = scanner.nextInt();
        if ((year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)) {
            System.out.println("Leap year.");
        } else {
            System.out.println("Not a leap year.");
        }
        scanner.close();
    }
}
~~~
# 19. Program to Check if a Number is a Three-Digit Number
### Question: Write a Java program to check if a given number is a three-digit number.

Answer:
~~~
import java.util.Scanner;

public class ThreeDigitNumber {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter a number: ");
        int number = scanner.nextInt();
        if (number >= 100 && number <= 999) {
            System.out.println("It is a three-digit number.");
        } else {
            System.out.println("It is not a three-digit number.");
        }
        scanner.close();
    }
}
~~~
# 20. Program to Read a Birthday Date and Print the Day of the Week
### Question: Write a Java program to read a birth date (day, month, year) and determine the day of the week.

Answer:
~~~
import java.util.Scanner;
import java.time.LocalDate;
import java.time.format.TextStyle;
import java.util.Locale;

public class BirthdayDay {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter day: ");
        int day = scanner.nextInt();
        System.out.print("Enter month: ");
        int month = scanner.nextInt();
        System.out.print("Enter year: ");
        int year = scanner.nextInt();
        LocalDate date = LocalDate.of(year, month, day);
        String dayOfWeek = date.getDayOfWeek().getDisplayName(TextStyle.FULL, Locale.ENGLISH);
        System.out.println("Day of the week: " + dayOfWeek);
        scanner.close();
    }
}
~~~
# 21. Program to Read Month Name and Print the Number of Days
### Question: Write a Java program to read the name of a month and print how many days it has.

Answer:
~~~
import java.util.Scanner;

public class MonthDays {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Enter month name: ");
        String month = scanner.nextLine().toLowerCase();
        int days;
        switch (month) {
            case "january":
            case "march":
            case "may":
            case "july":
            case "august":
            case "october":
            case "december":
                days = 31;
                break;
            case "april":
            case "june":
            case "september":
            case "november":
                days = 30;
                break;
            case "february":
                days = 28; // Assuming non-leap year
                break;
            default:
                days = -1; // Invalid month
                break;
        }
        if (days != -1) {
            System.out.println("Number of days: " + days);
        } else {
            System.out.println("Invalid month name.");
        }
        scanner.close();
    }
}
~~~


# Write a java program to print the entered charatered is ovels and consonent 
~~~
package purandhar_proj;

import java.util.Scanner;

public class OvelConcasonent {

	public static void main(String[] args) {
        System.out.println("Enter a letter to find if it's a vowel or consonant:");
        Scanner s = new Scanner(System.in);
        char a = s.next().charAt(0);
        s.close();  // Closing the Scanner to prevent resource leaks
        
        // Check if the input is a valid alphabet letter
        if ((a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z')) {
            if (a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' ||
                a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U') {
                System.out.println(a + " is a vowel");
            } else {
                System.out.println(a + " is a consonant");
            }
        } else {
            System.out.println("Invalid input! Please enter an alphabet letter.");
        }
    }
}

~~~

# write the program to print the number of vowels and consonant in the sentances
~~~
import java.util.Scanner;

public class VowelConsonantCounter {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        // Get user input
        System.out.println("Enter a sentence:");
        String sentence = scanner.nextLine().toLowerCase();  // Convert to lowercase for uniformity
        scanner.close();

        int vowels = 0, consonants = 0;

        // Define vowel characters
        String vowelChars = "aeiou";

        // Loop through each character in the sentence
        for (char c : sentence.toCharArray()) {
            if (Character.isLetter(c)) {  // Check if character is a letter
                if (vowelChars.indexOf(c) != -1) {
                    vowels++;  // Increment vowel count
                } else {
                    consonants++;  // Increment consonant count
                }
            }
        }

        // Output result
        System.out.println("Number of vowels: " + vowels);
        System.out.println("Number of consonants: " + consonants);
    }
}
~~~

##  Enhanced For Loop
~~~
for (char c : sentence.toCharArray()) {
    // Process each character
}
~~~
## Regular for Loop
~~~
char[] chars = sentence.toCharArray();
for (int i = 0; i < chars.length; i++) {
    char c = chars[i]; // Manually accessing each character
}
~~~