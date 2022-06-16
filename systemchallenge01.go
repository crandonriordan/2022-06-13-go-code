package main

import (
    "fmt"
    "os/exec"
    "io/ioutil"
    "log"
)

func main() {
    cmd := exec.Command("nslookup", "google.com")

    stdout, err := cmd.StdoutPipe()

    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
        
    data, err := ioutil.ReadAll(stdout)

    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", string(data))
}
