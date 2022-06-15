package main

import "fmt"

func main() {
    programmingLanguages := map[string]string {
        "Python": ".py",
        "Java": ".java",
        "Ruby": ".rb",
        "JavaScript": ".js",
        "Golang": ".go",
        "C++": ".cpp",
    }

    fmt.Println("initial map ", programmingLanguages)
    delete(programmingLanguages, "C++")
    programmingLanguages["Julia"] = ".jl"
    fmt.Println("map now ", programmingLanguages)
}
