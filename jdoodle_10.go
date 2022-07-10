package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    file, err := os.Open("/uploads/text.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println("start sleep")
        comp, _ := time.ParseDuration(scanner.Text())
        time.Sleep(comp)
        fmt.Println("finish sleep")
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}