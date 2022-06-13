package main

import "fmt"

type astro struct {
    name        string
    age         int
    lastmission string
    isneeded    bool
}

func main() {
    astro1 := astro{"Bob", 22, "Bezo's Enterprise", false}
    astro2 := astro{"Bill", 38, "Musk's Mars Terraforming", false}
    astro3 := astro{"Brian", 42, "SpaceX Crew-3", true}
   

    p := []astro{ astro1, astro2, astro3 }
    fmt.Println(p)
    fmt.Println(p[2].lastmission)

}
