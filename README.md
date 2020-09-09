# Quich Intro to Go

## The Idea

- simple and small
- easy to read
- backwards compatible
- complete standart library
- concorrency
- compiled static linked binary

## The Syntax

```go
//Main
package main

import "fmt"

func main() {
  fmt.Println("Bye, asteroid!")
}

//Vars
var first string
first = "This is first string"
  
var second = "This is second string"
  
third := "This is third string"

// conditional
if a > b {
    b = a
}

// for loop
for i := 0; i < c; i++ {
  fmt.Print(i, " ")
}

// equivalent to a `while` loop
for c > 0 {
    fmt.Print(c, " ")
    c--
}
```

Basic Types
```go
bool
string
byte
rune
int ( int8 int16 int32 int64 )
uint ( uint8 uint16 uint32 uint64 uintptr )
float32, float64
complex64 complex128

//struct
type Person struct{
  Name    string `json:"name"`
  Age     int    `json:"age"`
}
```

Funcs
```go
func echo(name string) string {
    return name
}

// equivalent to an Object Method
func (p Person) greet() {
    fmt.Printf("Hello %s\n" + p.Name)
}

func (p Person) isAuthorized() boolean {
    return p.isAdmin
}

func parseToJson(p Person) ([]byte, error) {
    json, err := json.Marshal(p)
    if err != nil {
        fmt.Printf("error: %s\n", err.Error)
        return nil, err
    }
    return json, nil
}

//goroutines
go func() {
  fmt.Println("concurrent")
}

p := Person{}
go p.greet("John")
```



## The Tooling and Standart Library

go fmt

## Exemples

- cli tool
- web server

...
## The Community

- Dave Cheney
https://dave.cheney.net/practical-go

- Gophercises
https://gophercises.com/