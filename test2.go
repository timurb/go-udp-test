package main

import "os"
import "fmt"
import "net"
import "strconv"

const MAX = 1000000

func main() {
/*  Do the initalization of UDP socket to reduce difference between 2 tests */
  conn, err := net.Dial("udp", "localhost:1234")
  if err != nil {
    fmt.Printf("Error while connecting to socket\n")
    os.Exit(1)
  }

  defer conn.Close()

/*  Iterate MAX times */
  for i := 0; i < MAX; i++ {
    x :=  strconv.Itoa(i)
    buf := []byte(x)
    conn.Write(buf)
  }
}
