package main

import "os"
import "fmt"
import "net"
import "strconv"

const MAX = 1000000

func main() {
  conn, err := net.Dial("udp", "localhost:1234")
  if err != nil {
    fmt.Printf("Error while connecting to socket\n")
    os.Exit(1)
  }

  fmt.Fprintf(conn, "init")

  for i := 0; i < MAX; i++ {
    x :=  strconv.Itoa(i)
    fmt.Printf(x)
    fmt.Fprintf(conn, x)
  }
}
