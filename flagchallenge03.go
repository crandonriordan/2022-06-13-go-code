package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
)

var num *int
var fileNameFlag *string

func init() {
    num = flag.Int("n", 0, "Number of lines to from the top to parse")
    fileNameFlag = flag.String("fileName", "/tmp/data.txt", "File to read from")
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    flag.Parse()
    n := *num
    fileName := *fileNameFlag
    file, err := os.Open(fileName)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    tmp := 0
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        tmp++
        if(tmp >= n) {
            break
        }
    }
}
