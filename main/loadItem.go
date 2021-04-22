package main

import (
	"fmt"
	"shopping/db"
	"os"
)

func main() {
	itemId := os.Args[1]
	fmt.Println(db.LoadItem(itemId))
}
