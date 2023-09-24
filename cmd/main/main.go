package main

import (
    "fmt"
    "github.com/ivanGuerreschi/whatversion/internal/utility"
    "github.com/ivanGuerreschi/whatversion/internal/version"
)

func main() {
    fmt.Println(utility.GetHome())
    fmt.Println(version.GetLocal())
}
