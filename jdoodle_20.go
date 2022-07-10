package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "time"
)
func f(t string) {
    fmt.Println("start sleep " + t)
    comp, _ := time.ParseDuration(t)
    time.Sleep(comp)
    fmt.Println("finish sleep " + t)
}
func main() {
    file, err := os.Open("/uploads/text.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        go f(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}