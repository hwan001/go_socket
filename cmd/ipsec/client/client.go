package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run client.go <server_hostname> <network>")
        return
    }
    serverHostname := os.Args[1]
    network := os.Args[2]
    connectToServer(serverHostname, network)
}

func connectToServer(serverHostname, network string) {
    conn, err := net.Dial(network, serverHostname+":8080")
    if err != nil {
        fmt.Printf("Error connecting to server (%s): %v\n", network, err)
        return
    }
    message := "Hello, Server!"
    conn.Write([]byte(message))
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from connection:", err)
        return
    }
    fmt.Printf("Received response: %s\n", string(buffer[:n]))
    conn.Close()
}
