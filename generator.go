package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())

    var output strings.Builder
    colorGreen := "\033[32m"
    timestamp := time.Now()

    //Lowercase and Uppercase Both
    charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP{}[];',./?$#@!*()&^%"
    var length int
    fmt.Print("[!] Enter length of password: ")
    fmt.Scan(&length)
    fmt.Println("[+] Generated password at", timestamp.Format(time.RFC822))
    for i := 0; i < length; i++ {
        random := rand.Intn(len(charSet))
        randomChar := charSet[random]
        output.WriteString(string(randomChar))
    }
    fmt.Println("    *", (colorGreen), output.String())
}