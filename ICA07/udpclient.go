package main
import (
    "fmt"
    "net"
    "bufio"
    //"./crypting"
)

func main() {
    p :=  make([]byte, 2048)
    //p := crypting.Encrypt("snart sommerferie")
    conn, err := net.Dial("udp", "127.0.0.1:8080")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Møte Fr 5.5 14:45 Flåklypa")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}
