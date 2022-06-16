package main

import (
    "fmt"
    "os"
)

func main() {
    numOfArgs := len(os.Args[1:])
    fmt.Println("Number of arguments: ", numOfArgs)
}
