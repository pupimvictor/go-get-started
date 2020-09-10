# A Introduction to Go

## Lightning talk at Staten Island Software Meetup

https://www.meetup.com/Staten-Island-Computer-Software-Meetup-Group/events/272794111/

## Why Go

- Simple and Small
- Easy to read
- Backwards compatible
- Complete standart library
- Concorrency
- Compiled static linked binary

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

go run
go test
go build
go fmt
...
https://www.alexedwards.net/blog/an-overview-of-go-tooling

## Sample App

The Sample app in this repository demonstrates in a simple way how to create your first Go web server and a CLI to interact with it.

The app is a simple Shopping List using memory storage. It supports Add Item and Get Items operations.

- app.go
  app logic and http requests handlers

- app_test.go
  app.go unit tests

- go.mod and go.sum
  Go dependency managemenet files

- cmd/server.go
  main function for the server. Register http handlers and listen on a tcp port

- cmd/cli.go
  main function for CLI tool. Reads inputs from OS stdin and makes a http request to the server

...

## projects and articles

Docker - https://github.com/docker

Go at Cloudflare - https://blog.cloudflare.com/go-at-cloudflare/

FritzFrog (BotNet) - https://www.guardicore.com/2020/08/fritzfrog-p2p-botnet-infects-ssh-servers/

## Learning Go

- A Tour of Go
https://tour.golang.org/

- justforfunc: Programming in Go
https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw

- Dave Cheney
https://dave.cheney.net/practical-go

- Gophercises
https://gophercises.com/