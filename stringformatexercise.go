package main

import (
    "fmt"
)

func main() {
    uri := "https://example.org:6001/v2/snacks?"
    r := "req=snickers"
    q := "quantity=1"
    s := "size=king"
    fullUri := fmt.Sprintf(uri + "%s%s%s", r, q, s);
    fmt.Println(fullUri)
}
