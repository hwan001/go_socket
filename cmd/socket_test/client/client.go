package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

type Config struct {
    ServerAddress string `json:"serverAddress"`
}

func loadConfig(path string) (*Config, error) {
    var config Config
    configFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer configFile.Close()

    bytes, _ := ioutil.ReadAll(configFile)
    json.Unmarshal(bytes, &config)

    return &config, nil
}

func main() {
	config, err := loadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config file: %s", err)
    }

    conn, err := net.Dial("tcp", config.ServerAddress)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

	for {
		_, err = conn.Write([]byte("Hello from client"))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}
