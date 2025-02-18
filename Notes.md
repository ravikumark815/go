# GO Notes

Main features:
- Efficient Compilation
- Efficient Execution
- Ease of Programming

> [Official Documentation](https://go.dev/doc/)

> [Go Packages](https://go.dev/pkg)


### Hello World:
```go
package main // Packages used to organize and reuse code

import "fmt" // Package of formatted I/O functions

func main() { // execution starts here
    fmt.Println("Hello World!") // Print Line
}
```
```sh
terminal@linux$ go run helloworld.go
Hello World!
terminal@linux$ 
```

### Go CLI
|   |   |
|---|---|
|`go build` | Compiles go source code |
|`go run`   | Compiles and executes go source code |
|`go fmt`   | Formats all the go code in each file in current directory|
|`go install`| Compiles and "installs" a package|
|`go get`   | Downloads the raw source code of someone else's package|
|`go test`  | Runs any tests associated with current project|
|`go mod init testproj`| Creates a new project named testproj|

### Variables
- Start with `[a-z][A-Z][_]`
- Contains `[a-z][A-Z][0-9][_]`
- No Keywords
- Identifier with only `_` is a blank identifier and used as a placeholder
- Exported identiers start with **Upper** Case Letter and are accessed across files

> Declaring a variable:
- Using `var` keyword: Can be used anywhere
- Using `:=` short declaration operator: Can be used inside functions only.

```go
var variable_name type = expression
variable_name := expression
```

```go
// Examples
var vName string = "Ravi"
var vName2 = "Ravi"
var v1, v2 = 1, 2.3
v3 := 5.78
var pl = fmt.Println //Alias for function
```

### Keywords
|   |   |   |   |   |
|---|---|---|---|---|
| break | case | chan | const | continue | 
default | defer | else | fallthrough | for | 
func | go | goto | if | import | 
interface | map | package | range | return | 
select | struct | switch |  type | var

### Data Types
1. Basic Types: Integers, Float, Strings, Boolean
2. Aggregate Types: Arrays, Structs
3. Reference Types: Pointers, Slices, Maps, Functions, Channels
4. Interface Types

| Data Type | Description |
|---|---|
int8 | 8-bit signed integer
int16 | 16-bit signed integer
int32 | 32-bit signed integer
int64 | 64-bit signed integer
uint8 | 8-bit unsigned integer
uint16 | 16-bit unsigned integer
uint32 | 32-bit unsigned integer
uint64 | 64-bit unsigned integer
int | Both int and uint contain same size, either 32 or 64 bit.
uint | Both int and uint contain same size, either 32 or 64 bit.
rune | It is a synonym of int32 and also represent Unicode code points.
byte | It is a synonym of uint8.
uintptr | It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
float32 | 32-bit IEEE 754 floating-point number
float64	| 64-bit IEEE 754 floating-point number
complex64 | Complex numbers which contain float32 as a real and imaginary component.
complex128 | Complex numbers which contain float64 as a real and imaginary component

```
Runes and Strings:
- Text is represented using the rune type (char in other languages)
- Rune is an alias for int32.
- A rune can represent any symbol - letters, numbers, language characters etc.
- String is the data type for storing multiple runes
- Strings are arrays of bytes and a string length
```

```go
// Printing Data type of a variable
fmt.Printf("Type: %T\n", "hello")   // string
fmt.Println(reflect.TypeOf(25))     // int
fmt.Println(reflect.TypeOf(3.14))   // float64
fmt.Println(reflect.TypeOf(true))   // bool
fmt.Println(reflect.TypeOf(🦍))     // int32
```

> Typecasting:
```go
package main
import "fmt"
func main() {
	// To cast type the type to convert to with the variable to
	// convert in parentheses
	// Doesn't work with bools or strings
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println(cV2)

	// Convert string to int (ASCII to Integer)
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	fmt.Println(cV4, err, reflect.TypeOf(cV4))

	// Convert int to string (Integer to ASCII)
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	fmt.Println(cV6)

	// Convert string to float
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println(cV8)
	}

    // Convert Byte arr into string
    byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	fmt.Println("Byte String :", bStr) // abc

	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	fmt.Println(cV9)
}
```

### Operators
| Category | Operators |
|---|---|
Arithmetic | `+ - * / %`
Relational | `== != > < >= <=`
Logical | `&&` &#124;&#124; `!`
Bitwise | `&` &#124; `^ << >> &^`
Assignment | `= += -= *= /= %= &= ^= <<= >>=` &#124;= 
Unary | `+ - ! & * ^ ++ --`
Misc | `&:`(Address) `*:`(Pointer) `<-`(Receive) `:=` (Declaration)

### Comments:
```go
// Single Line Comment

/*
Multi 
Line
Comment
*/
```

### Control and Loop Statements
**if-else-if:**
```go
if v1 == 100 {
    fmt.Printf("100\n")
} else if v1 == 200 {
    fmt.Printf("200\n")
} else if v1 == 300 {
    fmt.Printf("300\n")
} else {
    fmt.Printf("No Matches\n")
}
```
**for:**
```go
for i := 0; i < 5; i++{ 
    fmt.Printf("%d",i)   
}
```
```go
// Infinite loop 
for { 
    fmt.Printf("GeeksforGeeks\n")   
}
```
```go
// for loop as while
i:= 0 
for i < 3 { 
    i += 2 
}
```
```go
// Range Loop
arr := []string{"a", "b"}
for i, j := range arr { 
    fmt.Println(i, j)  
}
/* Output:
0 a
1 b
*/
```
```go
// Strings
for i,j := range "abc" {
    fmt.Printf("%d: %c\n", i, j)
}
```
```go
// Maps
mmap := map[int]string {
    1:"a",
    2:"b",
    3:"c",
}
for k, v := range mmap {
    fmt.Println(k, v)
}
```
**Switch:**
```go
func main() {
    day := 1
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    default:
        fmt.Println("Invalid")
    }
}
```
```go
// Type Switch
func main() {
    var day interface{} = 2
    switch v := day.(type) {
    case int:
        switch v {
        case 1:
            fmt.Println("Monday")
        case 2:
            fmt.Println("Tuesday")
        default:
            fmt.Println("Invalid day")
        }
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

**Loop Control:**
```go
// break
for i:=0; i<5; i++ { 
    fmt.Println(i) 
    if i == 3{ 
         break; 
    } 
}
```
```go
// continue
for i < 8 {
    if i == 5 {
        i = i + 2;
        continue;
    }
    fmt.Printf("value is: %d\n", i);
    i++;
} 
```
```go
// goto
label: for x < 8 {
    if x == 5 {
        x = x + 1; 
        goto label
    }
    x++;
}
```
    
### I/O functions

|Placeholder|Purpose|
|---|---|
%d | Integer
%c | Character
%f | Float
%t | Boolean
%s | String
%o | Base 8
%x | Base 16
%v | Guesses based on data type
%T | Type of supplied value


```go
// Print
import fmt
fmt.Println("Hello Go")

fmt.Printf("%s %d %c %f %t %o %x\n", "String", 1, 'A', 3.14, true, 1, 1)
fmt.Printf("%9f\n", 3.14)      // Width 9
fmt.Printf("%.2f\n", 3.141592) // Decimal precision 3.14

sp := fmt.Sprintf("%9.f\n", 3.141592) // Returns a formatted string
fmt.Println(sp) // Width 9 no precision  3

// Scan
import (
    "bufio"
    "os"
)
reader := bufio.NewReader(os.Stdin)
name, err := reader.ReadString('\n')
```

### Strings
```go
import "strings"

sV1 := "A word"

// Replacer that can be used on multiple strings
// to replace one string with another
replacer := strings.NewReplacer("A", "Another")
sV2 := replacer.Replace(sV1)
fmt.Println(sV2)

// Get length
fmt.Println("Length : ", len(sV2)) // 12

// Contains string
fmt.Println("Contains Another :", strings.Contains(sV2, "Another")) // true

// Get first index match
fmt.Println("o index :", strings.Index(sV2, "o")) // 2

// Replace all matches with 0
// If -1 was 2 it would replace the 1st 2 matches
fmt.Println("Replace :", strings.Replace(sV2, "o", "0", -1)) // An0ther w0rd

// Remove whitespace characters from beginning and end of string
sV3 := "\nSome words\n"
sV3 = strings.TrimSpace(sV3) //Some words

// Split at delimiter
fmt.Println("Split :", strings.Split("a-b-c-d", "-")) //[a b c d]

// Upper and lowercase string
fmt.Println("Lower :", strings.ToLower(sV2)) //another word
fmt.Println("Upper :", strings.ToUpper(sV2)) //ANOTHER WORD

// Prefix or suffix
fmt.Println("Prefix :", strings.HasPrefix("tacocat", "taco")) // true
fmt.Println("Suffix :", strings.HasSuffix("tacocat", "cat")) // true

rStr := "abcdefg"

// Runes in string
fmt.Println("Rune Count :", utf8.RuneCountInString(rStr))

// Print runes in string
for i, runeVal := range rStr {
    // Get index, Rune unicode and character
    fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
}
```

### Time
```go
import "time"
now := time.Now()
fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()) // 2024 December 25 19 4 25

```

### Math
```go
import "math"
randNum := rand.Intn(50) // Generate a rand int between 0 and 50
fmt.Println(randNum)
fmt.Println("Abs(-10) =", math.Abs(-10)) // 10
fmt.Println("Pow(4, 2) =", math.Pow(4, 2)) // 16
fmt.Println("Sqrt(16) =", math.Sqrt(16)) // 4
fmt.Println("Cbrt(8) =", math.Cbrt(8)) // 2
fmt.Println("Ceil(4.4) =", math.Ceil(4.4)) // 5
fmt.Println("Floor(4.4) =", math.Floor(4.4)) // 4
fmt.Println("Round(4.4) =", math.Round(4.4)) // 4
fmt.Println("Log2(8) =", math.Log2(8)) // 3
fmt.Println("Log10(100) =", math.Log10(100)) // 2
// Get the log of e to the power of 2
fmt.Println("Log(7.389) =", math.Log(math.Exp(2))) // 2
fmt.Println("Max(5,4) =", math.Max(5, 4)) // 5
fmt.Println("Min(5,4) =", math.Min(5, 4)) // 4
// Convert 90 degrees to radians
r90 := 90 * math.Pi / 180
// Convert 1.5708 radians to degrees
d90 := r90 * (180 / math.Pi)
fmt.Printf("%f radians = %f degrees\n", r90, d90) // 1.57096 radians = 90 degrees
fmt.Println("Sin(90): ", math.Sin(90 * math.Pi / 180)) // 1
fmt.Println("Cos(0): ", math.Cos(0 * math.Pi / 180)) // 1
fmt.Println("Tan(45): ", math.Tan(45 * math.Pi / 180)) // 1
fmt.Println("Hypot: ", math.Hypot(3,4)) // 5
```

### Arrays
```go
var arr1 [5] int //Declare
arr2 := [5] int {1, 2, 3, 4, 5} // Declare and Initialize
arr1[0] = 10 // Assign
fmt.Println (arr1[0]) // Access
fmt.Println (len(arr2)) // Length
arr3 := [][]int {} // Multidimension array
fmt.Println(arr1[:3]) // Slice from beginning
fmt.Println(arr1[2:]) // Slice to the end

//  Iterate: 
for i:=0; i<len(arr2); i++ {
    fmt.Println(arr2[i])
}

for k,v := range arr2 {
    fmt.Printf("%d : %d\n", k, v)
}

// Slices of Array
str1 := "abcdef"
rArr := [] rune(str1) // Convert to array of runes
for _,v := range rArr {
    fmt.Printf("%d", v) // 97 98 99 100 101
}
```

### Slices
Slices are arrays that can grow.
```go
// Create
sl1 := make([]string, 3) // size 3
sl1[0] = "Ravi"
sl1[1] = "Kumar"
sl1[2] = "Reddy"

// Length
fmt.Println(len(sl1)) // 3

// Iterate
for i:=0; i<len(sl1); i++ {
    fmt.Println(sl1[i])
}
for _,v := range sl1 {
    fmt.Println(sl1[v])
}

// Creating slice from an arry. Slice changes when the array it points to, changes and vice versa. 
sArr := [5] int {10,20,30,40,50}
sl2 := sArr[0:2] // Creating Slice for range 0-2
fmt.Println(s)
for _,v := range sl2 {
    fmt.Println(sl1[v]) 
}
sl2.append(sl2, 10)
```

### Functions
*Syntax:*
```go
func funcName(parameters) returnTypes {
    // BODY
}
```
```go
// Example:
func divide(x float64, y float64) (ans float64, err error) {
    if y==0 {
        return 0, fmt.Errorf("Division by 0")
    }
    return x/y, nil
}
func main() {
    divide(5/4)
}

// Voradic Functions: receive unknown number of values:
func getSum(nums ...int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

// Passing Arrays:
func getArraySum(arr[] int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

// Passing Pointer:
func changeVal(numPtr *int) int {
    *numPtr = 12
}
num := 10
changeVal(&num)

// Passing Pointer of an Array
func arrayPtrFunc(arr*[]int) int {
    return 1
}

func main() {
    arr := []int {1,2,3,4}
    fmt.Println(arrayPtrFunc(&arr))
}

// Passing a slice
func sliceFunc(nums ...int) int {
    return len(nums)
}

func main() {
    sSlice := []int{10,20,30}
    fmt.Println(sliceFunc(sSlice...)) // Notice ... in func call
}
```

### Pointers
```go
// Declaration
num := 10
var numPtr *int = &num

// Access
*numPtr = 12

// Address
fmt.Println("Addr:", numPtr)
```

### File I/O
```go
// Create a file
f,err := os.Create("data.txt")
defer f.Close() // Closes the file as soon as we are out of scope

// Write to file
arr := []int{2,3,5,7,11}
for _,num := range arr {
    _,err := f.WriteString(strcov.Itoa(num) + "\n")
}

// Check if file exists
_, err = os.Stat("data.txt")
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File Doesn't Exist")
}

// Open a file
f,err = os.Open("data.txt")

// Read from file
scan1 := bufio.NewScanner(f)
for scan1.Scan() {
    fmt.Println(scan1.Text())
}

// Append to file
/*
        Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

    O_RDONLY : open the file read-only
    O_WRONLY : open the file write-only
    O_RDWR   : open the file read-write

    These can be or'ed

    O_APPEND : append data to the file when writing
    O_CREATE : create a new file if none exists
    O_EXCL   : used with O_CREATE, file must not exist
    O_SYNC   : open for synchronous I/O
    O_TRUNC  : truncate regular writable file when opened
*/
_, err = os.Stat("data.txt")
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File Doesn't Exist")
} else {
    f, err = os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    if _, err := f.WriteString("13\n"); err != nil {
        log.Fatal(err)
    }
}
```

### Command Line Args
```go
func main() {
    fmt.Println(os.Args)
}
/*
go run test.go 24 abc
Output: [test.go, 24, abc]
*/
```

### Packages/Modules
- Packages allow you to keep related together and go looks for packages codes in specific directories. 
- Variables/Functions starting with uppercase letters are accessible outside the package
- To retrieve values or variables for valid authenticated access, we can write getter functions that start with uppercase letters and simply return the variable value.
- `package main` must have `main` function in it. Package names are mostly lowercase
- Example:
```go
// testproj/mypackage/mypackage.go
package test
var GlobalVar := 10 // G uppercase for external access

// main.go
import (
    test "testproj/mypackage"
)
func main() {
    fmt.Println("Var:", test.GlobalVar) // Var: 10
}
```

### Maps
- Maps are collections of k-v pairs.
- Key can be anything that can be compared using `==`

Syntax:
```go
var mapName map [keyType]valueType
mapName := make(map[keyType]valueType)
mapName := map[int]string {1: "a",
    2: "b",
    3: "c"
}
```
Example:
```go
var heroes map[string]string
villains := make(map[string]string)
mapName := map[int]string {1: "a",
    2: "b",
    3: "c"
}

// Access and Assignment
heroes["Ironman"] = "Tony Stark"
heroes["Batman"] = "Bruce Wayne"
heroes["Superman"] = "Clark Kent"

fmt.Println(mapName) // O/p: map[1:a 2:b 3:c]

// Iterate
for k,v := range mapName {
    fmt.Printf("%d:%s", k, v) // 1:a2:b3:c
}

// Delete
delete(heroes, "Superman")
```

### Generics
- With generics we can specify the data type to be used at a later time. 
- Mainly used when we want to use functions with moldable data types
- Generic type parameter is capitalized and in square brackets

```go
type MyConstraint interface {
    int | float64
}

// T is the generic and can be int/Float64 only
func getSumGen[T MyConstraint] (x T, y T) T {
    return x + y
}

func main() {
    fmt.Println("5 + 4 =", getSumGen(5, 4)) // 9
    fmt.Println("5.6 + 4.7 =", getSumGen(5.6, 4.7)) // 10.3
}
```

### Structs
- Allow you to store values with many different types in a structured manner 
```go
// Define
type customer struct {
    name string
    address string
    balance float64
    id int64
}

// Composition: Struct inside struct
type business struct {
    businessName string
    workaddr string
    contact
}

// Struct Args
func getCustInfo(c customer) {
    fmt.Println("%s owes us %2.f", c.name, c.bal)
}

// Instance and Access
func main() {
    var test customer
    test.name = "John Smith"
    test.address = "9 Main St"
    test.balance = 234.56
    getCustInfo(test) // John Smith owes us 234.56
    
    var bTest business
    bTest.name = "ABC Corp"
    bTest.contact.name = "Jane Doe"
}
```

### Defined Types
- Defined type used previously with structs
- You can use them also to enhance the quality of other data types

```go
	type Tsp float64
    type TBs float64
    type ML float64

    func tspToML (tsp Tsp) ML {
        return ML(tsp * 4.92)
    }

    func TBToML (tsp TBs) ML {
        return ML(tsp * 14.79)
    }

    func (tsp Tsp) ToMLs() ML {
        return ML(tsp * 4.92)
    }
    
    // Convert from tsp to mL
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f mL\n", ml1)

	// Convert from TBs to mL
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f mL\n", ml2)

	// You can use arithmetic and comparison operators
	fmt.Println("2 tsp + 4 tsp =", Tsp(2) + Tsp(4))
	fmt.Println("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// We can convert with functions
	// Bad Way
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 TBs = %.2f mL\n", TBToML(3))

	// We can solve this by using methods which
	// are functions associated with a type
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
```

### Interfaces
- Interfaces basically group types together based on their methods
- They allow you to create contracts that say if anything inherits it that they will implement defined methods
- If we had animals and wanted to define that they all perform certain actions, but in their specific way we could use an interface
- With Go you don't have to say a type uses an interface. When your type implements the required methods it is automatic

```go
// Creating interface
type shape interface {
    area() float64
    perimeter() float64
}

type square struct {
    length float64
}

type circle struct {
    radius float64
}

func (s square) area() float64{
    return s.length * s.length
}

func (s square) perimeter() float64{
    return s.length * 4
}

func (c circle) area() string {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() string {
    return 2 * math.Pi * c.radius
}

/* 
Originally square and circle have their own structs and methods.
From here on, the methods are common to both of them and we don't want to write redundant code. Interface groups them together.
So s belongs to a struct that has both area and perimeter functions.
*/

func printShapeInfo(s shape) {
    fmt.Printf("Area: %0.2f", s.area())
    fmtPrintf("Perimeter: %0.2f", s.perimeter())
}

func main() {
    printShapeInfo(square_a{length: 15.2})
    printShapeInfo(circle_a{radius: 7.5})
}
```

### Concurrency
- Concurrency is the concept of executing multiple tasks simultaneously at the same time
- Concurrency allows us to have multiple blocks of code share execution time by pausing their execution. 
- We can also run blocks of codes in parallel at the same time. (Concurrent tasks in Go are called goroutines) 
- To execute multiple functions in new goroutines add the word go in front of the function calls (Those functions can't have return values)

```go
func printTo15() {
	for i := 1; i <= 15; i++ {
		fmt.Println("Func 1 :", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Func 2 :", i)
	}
}

func main() {
    go printTo15() // Adding go infront would make them run as go routines (threads)
    go printTo10()
    time.Sleep(2*time.Second)
}

/* O/p
Fun 1: 1
Fun 2: 1
Fun 1 : 3
....
*/
```

### Channels
- We can have go routines communicate using channels. One routine acts a sender another acts as a receiver
- For example, if you want the above `printTo15` and `printTo10` functions to print in order, but execute concurrently

```go
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
}
func nums2(channel chan int) {
	channel <- 3
	channel <- 4
}

func main() {
    // Creating channels
    channel1 := make(chan int)
    channel2 := make(chan int)
    // Running nums1 and nums2 concurrently
    go nums1(channel1)
    go nums2(channel2)
    // Printing them in order using channels
    fmt.Println(<-channel1)
    fmt.Println(<-channel1)
    fmt.Println(<-channel2)
    fmt.Println(<-channel2)
}
```

### Closures
- Closures are functions that don't have to be associated with an identifier (Anonymous)
```go
// Pass a function to a function
func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
    intSum := func(x, y, )
	
    // Create a closure that sums values
	intSum := func(x, y int) int { 
        return x + y 
    }
	fmt.Println("5 + 4 =", intSum(5, 4)) // 9

	// Closures can change values outside the function
	samp1 := 1
	changeVar := func() { 
        samp1 += 1 
    }
	changeVar() // Changing value of a var outside of the function
	fmt.Println("samp1 =", samp1) // sampl = 2

	// Pass a function to a function
	useFunc(sumVals, 5, 8)
}
```

### Recursion
```go
import regexp
func main() {
    reStr := "The ape was at the apex"
    match,_ := regexp.MatchString("ape[^ ]", reStr) // ape not followed by a space
    fmt.Println(match) // true

    reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

    // Did you find any matches?
	fmt.Println("MatchString :", r.MatchString(reStr2)) // true

	// Return first match
	fmt.Println("FindString :", r.FindString(reStr2)) // rat

	// Starting and ending index for 1st match
	fmt.Println("Index :", r.FindStringIndex(reStr2)) // [4 7]

	// Return all matches
	fmt.Println("All Matches :", r.FindAllString(reStr2, -1)) // [rat mat fat pat]

	// Get 1st 2 matches
	fmt.Println("First 2 Matches :", r.FindAllString(reStr2, 2)) // [rat mat]

	// Get indexes for all matches
	fmt.Println("Indexes for all matches :", r.FindAllStringSubmatchIndex(reStr2, -1)) //[[4 7 4 7][8 11 8 11][12 15 12 15][16 19 16 19]]

	// Replace all matches with Dog
	fmt.Println(r.ReplaceAllString(reStr2, "Dog")) // Cat Dog Dog Dog Dog
}
```