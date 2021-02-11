package main

import (
	"fmt"
	"math/rand"
)

func main() {
	v := rand.Int31n(300)
	fmt.Println(v)
}
