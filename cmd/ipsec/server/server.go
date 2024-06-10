package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen for incoming connections on both IPv4 and IPv6
    listener4, err := net.Listen("tcp4", ":8080")
    if err != nil {
        fmt.Println("Error starting IPv4 listener:", err)
        return
    }
    listener6, err := net.Listen("tcp6", ":8080")
    if err != nil {
        fmt.Println("Error starting IPv6 listener:", err)
        return
    }

    go handleConnections(listener4)
    go handleConnections(listener6)

    select {} // Keep the main function running
}

func handleConnections(listener net.Listener) {
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            return
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading from connection:", err)
        return
    }
    fmt.Printf("Received message: %s\n", string(buffer[:n]))
    conn.Write([]byte("Message received"))
    conn.Close()
}
