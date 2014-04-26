package main

import (
    "flag"
    "fmt"
)

const (
    version = "0.0.1"
)

var (
    versionFlag = flag.Bool("version", false, "version flag display version")
)

func main() {
    flag.Parse()

    if *versionFlag {
        fmt.Println(version)
    }
}
