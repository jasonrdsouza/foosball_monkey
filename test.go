package main

import (
    "fmt"
    //"io/ioutil"
    //"log"
    "path/filepath"
)

func test() {
    fmt.Println("Hello, 世界")
    // dir_contents, err := ioutil.ReadDir("/Users/jason/Downloads")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // for _, f := range dir_contents {
    //     fmt.Println(f.Name())
    //     fmt.Println(uint32(f.Mode()))
    // }
    // file_contents, err := ioutil.ReadFile("/Users/jason/Downloads/test.txt")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // err = ioutil.WriteFile("/Users/jason/Desktop/backup_test.txt", file_contents, 0644)
    // if err != nil {
    //     log.Fatal(err)
    // }
    a := "test1"
    b := "/"
    c := "file"
    d := fmt.Sprint(a,b,c)
    fmt.Println(d)
    fmt.Println(filepath.Join(a,c))
}