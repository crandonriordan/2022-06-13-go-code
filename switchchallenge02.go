package main

import (
    "fmt"
    "regexp"
)

var zipcode = regexp.MustCompile(`^\d{5}$`)

func main() {
    zips := []string { 
        "60525",
        "42765",
        "42701",
        "abcdef",
    }

    for _, zip := range zips {
        var matches bool = zipcode.MatchString(zip)
        fmt.Printf("%s is a zipcode? %v\n", zip, matches)
    }
}
