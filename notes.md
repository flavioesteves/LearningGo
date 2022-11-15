# Section 1: Getting Started
01. How to Get Help
02. Course Resources
03. Join Our Community
04. Environment Setup
05. VSCode Installation
06. Go Support in VSCode

# Section 2: A Simple Start
07. Boring Ol' Hello World
08. Five Important Questions
  * How do we run the code in our project? --> go CLI 
    - go build: Compiles a bunch of go source code files
    - go run: Compiles and executes one or two files
    - go fmt: Formats all the code in each file in the current directory
    - go install: Compiles and "installs" a package
    - go get: Downloads the raw source code of someone else's package
    - go test: Runs any tests associated with the current project
09. Go Packages
  * What does 'package main' mean?
    Package == Project == Workspace
  * 2 types of packages:
    - Executable: Generates a file that we can run
    - Reusable: Code used as 'helpers'. Good place to put reusable logic
  * package "main" is a executable type package, is mandatory to have a func main()
10. Import Statements
  * What does 'import "fmt"' mean?
  - import other go packages
  - [https://pkg.go.dev/std](https://pkg.go.dev/std)
11. File Organization
  * What's that 'func' thing?
  * How is the main.go file organized?
    - Package declaration
    - Import other packages that we need
    - Declare functions, tell Go to do things
12. How to Access Course Diagrams

# Section 3: Deeper Into Go
13. Project Overview
 * Cards project: newDeck, print, shuffle, deal, saveToFile, newDeckFromFile
14. New Project Folder
15. Variable Declarations
  * Some Basic Go types:
  - bool: true/false
  - string: "string"
  - int: 1, -1 
  - float64: 1.0, -1.0
16. Functions and Return Types
17. Slices nad For Loops
  * Array: Fixed length list of things
  * Slice: An Array that can grow or shrink
  * TO add a element to a slice -> "var" = append("var", newVar) 
  * Loop: 
    - for index, "varOfSlice" := range "slice" { fmt.Printls(index, "varOfSlice") }
18. OO Aproach vs GO Approach
19. Custom types Declarations
  * Custom types "extend" a base type and add some extra functionality to it.
  * A function with a receiver is like a "method" - a function that belongs to an "instance"
  *  func (d deck) print(){}
20. Receiver Functions
  * func (d deck) print (){} -> Any variable of the type "deck" now gets access to the "print" method
  * d --> The actual copy of the deck we're working with is available in the function as a variable called 'd'
  * deck --> Every variable of type 'deck' can call this function on itself
  * receiver variable is like "this" in JS & "self" in Rust
  * by th econvention is the first letter of the type. ex:.. (d deck)
21. Creating a new deck
  * "_" replaces a a variable that is not used, and tell the compiler to ignore it. Example iterators on the for loop
22. Slice Range Syntax
  * Slices are zero-indexed like arrays in other languages
  *  A range of selection on the slice: Slice[starIndexIncluding : upToNotIncluding]
23. Multiple Returns Values
  * return "value1", "value2"
  * Syntax to capture multiple returns: result1, result2 := <value1>, <value2>
24. Byte Slices
  *  WriteFile in 1.16 change form std package from "ioutils" to "os" --> [https://pkg.go.dev/os#WriteFile](https://pkg.go.dev/os#WriteFile)
25. Deck to String
  * Type Conversion. Example: Type we want -> Value we have - []byte("Hi there!")
  * deck --> []string --> string --> []byte
26. Joining a Slice of Strings
  * [https://pkg.go.dev/strings](https://pkg.go.dev/strings)
27. Saving Data to the Hard Drive
  * OS package instead of ioutils
28. Reading From the Hard Drive
