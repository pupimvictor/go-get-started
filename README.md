# A Introduction to Go

## Lightning talk at Staten Island Software Meetup

https://www.meetup.com/Staten-Island-Computer-Software-Meetup-Group/events/272794111/

## Why Go

- simple and small
- easy to read
- backwards compatible
- complete standart library
- concorrency
- compiled static linked binary

## Syntax

```go
//Main
package main

import "fmt"

func main() {
  fmt.Println("Bye, asteroid!")
}

//Vars
var foo string
foo = "bar"
bar := "foo"

// conditional
if a > b {
    b = a
}

// for loop
for i := 0; i < c; i++ {
  fmt.Print(i, " ")
}

//while loop
for c > 0 {
    fmt.Print(c, " ")
    c--
}

// for each loop
for i, x := range mySlice {
  fmt.Println(x)
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

// a 'Method'
func (p Person) greet() {
    fmt.Printf("Hello %s\n" + p.Name)
}

// multiple return values
func parseFromJson(payload []byte) (*Person, error) {
    var p Person
    err := json.Unmarshal(payload)
    if err != nil {
        fmt.Printf("error: %s\n", err.Error)
        return nil, err
    }
    return json, nil
}
```

webServer
```go
package main

import (
  "net/http"
  "log"
  "os"
)

func main() {
  http.HandleFunc("/greet", func(writer http.ResponseWriter, request *http.Request) {
      writer.WriteHeader(200)
      writer.Write([]byte("hello, friend"))
  })
  log.Printf("listening on port %s", os.Args[1])
  log.Fatal(http.ListenAndServe(os.Args[1], nil))
}
```

## The Tooling and Standart Library

go fmt

## Exemples

- cli tool
- web server

...

## projects
docker
k8s
malware

## The Community

- Dave Cheney
https://dave.cheney.net/practical-go

- Gophercises
https://gophercises.com/