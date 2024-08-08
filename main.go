package main

import (
    "fmt"
    "math/rand"
    "time"
)

var acakan = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
    var isiAcak int

    isiAcak = randomWithRange(2, 10)
    fmt.Println("Angka acak:", isiAcak)

    isiAcak = randomWithRange(2, 10)
    fmt.Println("Angka acak:", isiAcak)

    isiAcak = randomWithRange(2, 10)
    fmt.Println("Angka acak:", isiAcak)
}

func randomWithRange(min, max int) int {
    var value = acakan.Int()%(max-min+1) + min
    return value
}