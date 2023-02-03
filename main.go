package main

import (
	"github.com/kevinand11/go-cache/cache"
)

func main () {
	c := cache.New()

	save := [] string {
		"parrot", "avocado", "dragonfruit",
		"tree", "potato", "tomato", "tree",
		"dog",
	}

	for _, word := range save {
		c.Save(word)
		c.Display()
	}
}