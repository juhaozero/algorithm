package main

import (
	"algorithm/search"
	"math/rand"

	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano() / 1e6)
	search.FindPath()
	//tree.CreatTree()
	//skip.CreatSkipList()
}
