package main

import "fmt"

type Virtmach struct {
    IP string
    Hostname string
    Diskgb int
    Ram int
}

func (self Virtmach) ramInGb() {
    fmt.Printf("Size of ram in GB: ")
    fmt.Println(float64(self.Ram) / 1000.0)
}

func (self *Virtmach) addRamStick() {
    self.Ram = self.Ram + 1024;
}

func main() {
    mac := Virtmach{"192.92.21.2", "nasa", 12, 1024} 
    mac.ramInGb()
    mac.addRamStick()
    mac.ramInGb()
}
