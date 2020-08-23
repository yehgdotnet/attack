package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    
)

var (
path = "sensitive.ini"
)
func access_test() {
    file, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      path:= scanner.Text()
      file := filepath.Base(path)
      fmt.Println(file)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func main(){
   access_test()
}
