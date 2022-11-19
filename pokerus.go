package pokerus

import (
	"fmt"
	"math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
    if rand.Int() % 65_536 < 3 {
        fmt.Println("Your gopher has pokérus!")
    }
}
