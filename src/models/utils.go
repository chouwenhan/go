package models

import (
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    CouchDB_server string
    CouchDB_port int
    CouchDB_username string
    CouchDB_password string
}

func ReadConfig() Configuration{
    file, _ := os.Open("config.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    return configuration
}
