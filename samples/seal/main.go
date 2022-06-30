package main

import (
    "fmt"
    "github.com/syndtr/goleveldb/leveldb"
    "log"
)

func main() {
    fmt.Println("In main")
    Put("key1", "value1")
    val := Get("key1")
    fmt.Println(val)
}

func Put(key string, value string)  {
    db, err := leveldb.OpenFile("/tmp/sgx.db", nil)
    if err != nil {
        log.Fatalf("Open faled %v", err.Error())
    }
    defer db.Close()

    err = db.Put([]byte(key), []byte(value), nil)
    if err != nil {
        log.Fatalf("Put faled %v", err.Error())
    }
}

func Get(key string) string  {
    db, err := leveldb.OpenFile("/tmp/sgx.db", nil)
    if err != nil {
        log.Fatalf("Open faled %v", err.Error())
    }
    defer db.Close()
    val, err := db.Get([]byte(key), nil)
    if err != nil {
        log.Fatalf("Get faled %v", err.Error())
    }
    return string(val)
}