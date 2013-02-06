// +build !appengine
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "os"
    "encoding/json"
    "time"
)

func files_dirs() {
    dir_contents, err := ioutil.ReadDir("/Users/jason/Downloads")
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range dir_contents {
        fmt.Println(f.Name())
        fmt.Println(uint32(f.Mode()))
    }
    file_contents, err := ioutil.ReadFile("/Users/jason/Downloads/test.txt")
    if err != nil {
        log.Fatal(err)
    }
    err = ioutil.WriteFile("/Users/jason/Desktop/backup_test.txt", file_contents, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    a := "test1"
    b := "/"
    c := "file"
    d := fmt.Sprint(a,b,c)
    fmt.Println(d)
    fmt.Println(filepath.Join(a,c))
}

func jsoning() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func not_main() {
    fmt.Println("Hello, 世界")

    var c chan string = make(chan string, 4)
    
    go pinger(c)
    go ponger(c)
    go printer(c)
    
    var input string
    fmt.Scanln(&input)
}