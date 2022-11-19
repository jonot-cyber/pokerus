package pokerus

import (
	"fmt"
	"math/rand"
    "time"
)

func init() {
    source := rand.NewSource(time.Now().UnixNano())
    if rand.New(source).Int() % 65_536 < 3 {
        fmt.Println("Your gopher has pokérus!")
    }
}
