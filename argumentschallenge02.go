package main

import (
    "fmt"
    "os"
)

func main() {
    for index, arg := range os.Args[1:] {
        fmt.Printf("Position %d has value %s\n", index, arg)
    }
}

