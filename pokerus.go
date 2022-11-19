package pokerus

import (
    "fmt"
    "math/rand"
)

func init() {
    if rand.Int() % 65_536 < 3 {
        fmt.Println("Your gopher has pokérus!")
    }
}
