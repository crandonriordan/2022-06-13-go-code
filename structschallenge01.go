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
    astro3 := astro{"Brian", 42, "Virgin Enterprise Galactic", true}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

}
