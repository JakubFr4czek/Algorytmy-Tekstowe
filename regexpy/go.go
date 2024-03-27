package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Przykładowy tekst zawierający daty w formacie DD-MM-YYYY
    text := "Data zakupu: 15-07-2023, Data dostawy: 20-07-2023"

    // Wzorzec dla daty w formacie DD-MM-YYYY
    pattern := regexp.MustCompile(`(\d{2})-(\d{2})-(\d{4})`)

    // Użycie FindAllString
    matches := pattern.FindAllString(text, -1)
    fmt.Println("FindAllString:")
    for _, match := range matches {
        fmt.Println(match)
    }

    // Użycie FindAllStringSubmatch
    submatches := pattern.FindAllStringSubmatch(text, -1)
    fmt.Println("\nFindAllStringSubmatch:")
    for _, submatch := range submatches {
        for _, group := range submatch {
            fmt.Println(group)
        }
    }
}
