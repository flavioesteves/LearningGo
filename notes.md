# Section 1: Getting Started
01. How to Get Help
02. Course Resources
03. Join Our Community
04. Environment Setup
05. VSCode Installation
06. Go Support in VSCode

# Section 2: A Simple Start
08. Boring Ol' Hello World
09. Five Important Questions
  * How do we run the code in our project? --> go CLI 
    - go build: Compiles a bunch of go source code files
    - go run: Compiles and executes one or two files
    - go fmt: Formats all the code in each file in the current directory
    - go install: Compiles and "installs" a package
    - go get: Downloads the raw source code of someone else's package
    - go test: Runs any tests associated with the current project
10. Go Packages
  * What does 'package main' mean?
    Package == Project == Workspace
  * 2 types of packages:
    - Executable: Generates a file that we can run
    - Reusable: Code used as 'helpers'. Good place to put reusable logic
  * package "main" is a executable type package, is mandatory to have a func main()
11. Import Statements
  * What does 'import "fmt"' mean?
  - import other go packages
  - [https://pkg.go.dev/std](https://pkg.go.dev/std)
12. File Organization
  * What's that 'func' thing?
  * How is the main.go file organized?
    - Package declaration
    - Import other packages that we need
    - Declare functions, tell Go to do things
13. How to Access Course Diagrams

# Section 3: Deeper Into Go
14. Project Overview
 * Cards project: newDeck, print, shuffle, deal, saveToFile, newDeckFromFile
15. New Project Folder
16. Variable Declarations
  * Some Basic Go types:
  - bool: true/false
  - string: "string"
  - int: 1, -1 
  - float64: 1.0, -1.0
17. Functions and Return Types
18. Slices nad For Loops
  * Array: Fixed length list of things
  * Slice: An Array that can grow or shrink
  * TO add a element to a slice -> "var" = append("var", newVar) 
  * Loop: 
    - for index, "varOfSlice" := range "slice" { fmt.Printls(index, "varOfSlice") }
19. OO Aproach vs GO Approach
20. Custom types Declarations
  * Custom types "extend" a base type and add some extra functionality to it.
  * A function with a receiver is like a "method" - a function that belongs to an "instance"
  *  func (d deck) print(){}
21. Receiver Functions
  * func (d deck) print (){} -> Any variable of the type "deck" now gets access to the "print" method
  * d --> The actual copy of the deck we're working with is available in the function as a variable called 'd'
  * deck --> Every variable of type 'deck' can call this function on itself
  * receiver variable is like "this" in JS & "self" in Rust
  * by th econvention is the first letter of the type. ex:.. (d deck)
22. Creating a new deck
  * "_" replaces a a variable that is not used, and tell the compiler to ignore it. Example iterators on the for loop
23. Slice Range Syntax
  * Slices are zero-indexed like arrays in other languages
  *  A range of selection on the slice: Slice[starIndexIncluding : upToNotIncluding]
24. Multiple Returns Values
  * return "value1", "value2"
  * Syntax to capture multiple returns: result1, result2 := <value1>, <value2>
25. Byte Slices
  *  WriteFile in 1.16 change form std package from "ioutils" to "os" --> [https://pkg.go.dev/os#WriteFile](https://pkg.go.dev/os#WriteFile)
26. Deck to String
  * Type Conversion. Example: Type we want -> Value we have - []byte("Hi there!")
  * deck --> []string --> string --> []byte
27. Joining a Slice of Strings
  * [https://pkg.go.dev/strings](https://pkg.go.dev/strings)    
28. Saving Data to the Hard Drive
  * OS package instead of ioutils
29. Reading From the Hard Drive
  * os.ReadFile
  * [ byteSlice ] [ err ] := os.ReadFile(filename)
  * err -> Value of the type 'error' if nothing went wrong, it will have a value of 'nil'
  * Quit application os.Exit()
30. Error Handling
31. Shuffling a Deck
  * Package math/rand (https://pkg.go.dev/math/rand#Rand.Intn)
  * Swap positions multiple elements in the Slice. Example:. d[i], d[newPosition] = d[newPosition], d[i] // swap elements
32. Random Number Generation
33. Creating a go.mod File
34. Testing With Go
  * To make a test creat a new file ending in _test.go
  * To run all tests in a package, run the command: go test
35. Errorf call has arguments but no formatting directives
36. Writting Useful Tests
  * *testing.T --> It is the testing handler is required to be an argument in Go test functions
37. Asserting Elements in a Slice 
38. Testing File IO
39. Project Review
40. Assignment Even or Odd?
41. Even or Odd Solution

# Section 4: Organizing Data With Structs
42. Structs in Go
43. Defining Structs
44. Declaring Structs
45. Updating Struct Values
46. Embedding Structs
47. Structs with Receiver Functions
48. Pass By Value
49. Structs with Pointers
50. Pointer Operations
    * "&variable" this operator, give the "memory" address of the value this variable is pointing at
    * "*pointer" this operator, give the "value" this memory address is pointing at
51. Pointer Shortcut
52. Gotchas With Pointers(check the lesson)
53. References vs Value Types
    * Arrays:
        - Primitive data structure
        - Can't be resized
        - Rarely used directly
    * Slices:
        - Can grow and shrink
        - Used 99% of the time for lists of elements
    * Value types: int, float, string, bool, structs 
        - use pointers to change these things in a function
    * Reference types: slices, maps, channels, pointers, functions
        - Is not necessary pointers in these

# Section 5: Maps
54. What's a Map?
    * In other languages is kind like:
        - Go - Map
        - Ruby - Hash
        - JS - Object
        - Python - Dict
55. Manipulating Maps
56. Iterating Over Maps
57. Differences Between Maps and Structs
    * Map:
        - All keys must be the same type
        - All values must be the same type
        - Keys are indexed - we can iterate over them
        - Use to represent a collection of related properties
        - Don't need to know all the keys at compile time
        - Reference Type!
    * Struct
        - Values can be of different type
        - Keys don't support indexing
        - You need to know all the different fields at compile time
        - Use to represent a "thing" with a lot of different properties
        - Value Type!

# Section 6: Interfaces
58. Purpose of Interfaces
    * We know that
        - Every value has a type
        - Every function has to specify the type of its arguments
    * So doed that mean
        - Every function we every write has to be rewritten to accomdodate different types even if the logic in it is identical?

59. Problems Without Interfaces
60. Interfaces in Practice
    * 2 types of bots structs and respective functions
    * type bot interface: If you are a type in this program with a function called "getGreeting" and you return a string then you are now an honorary member og type "bot"
61. Rules of Interfaces
    * Concrete Type: map, struc, int, string, englishBot
    * Inteface Type: bot
62. Extra Interface Notes
    * Interface are not generic types
    * Interfaces are 'implicit'
    * Interfaces are a contract to help us manage types
    * Interfaces are tough. Step#1 is understanding how to read them
63. The HTTP Package
    * Package net/http
64. Reading the Docs
65. More Interface Syntax
