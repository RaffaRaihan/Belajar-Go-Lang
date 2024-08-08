package main

import "fmt"
import "strings"

func main() {
	var nama = []string{"Raffa", "Raihan"}
    printMessage("hallo", nama)
}
func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}