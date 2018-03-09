// HelloWorld in Go - Brian Dwyer - 2018

package main

// Dependencies
import (
  _ "context"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  _ "os/signal"
  _ "reflect"
  _ "strconv"
  _ "syscall"

  "github.com/k0kubun/pp"
)

// Constants
const (
  version = "v1.0.0"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

    // Figure out the Type of Object that R is
    // fmt.Println(reflect.TypeOf(r))

    // Attempt to Convert HTTP Request to JSON - Out to Console
    // This errors with: [] json: unsupported type: func() (io.ReadCloser, error)
    fmt.Println(json.Marshal(r))

    // Get the Hostname
    fmt.Println(os.Hostname())

    // Return Pretty-Printed HTTP Request (Still not JSON)
    pp.Fprintf(w, "%s", r)

    // Print the Remote Address
    fmt.Println(r.RemoteAddr)

    // Print the Request Method
    fmt.Println(r.Method)

    // Print the Request URL
    fmt.Println(r.URL)

    // Convert structs to JSON
    // data, err := json.Marshal(r)
    // if err != nil {
    //   log.Fatal(err)
    // }
    // fmt.Printf("%s\n", data)

    // fmt.Fprintf(w, json.Marshal(r))
    // fmt.Fprintf(w, reflect.TypeOf(r))

    // fmt.Fprintf(w, r)
}

// Code
func main() {
  log.Println("Starting helloworld application...")

  // Call a Function for a given Route
  http.HandleFunc("/", handler)

  // Launch the Web Server
  log.Fatal(http.ListenAndServe(":5000", nil))
}
