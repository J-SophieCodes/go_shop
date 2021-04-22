package main

import (
	"fmt"
	"shopping/db"
	"os"
	"strconv"
)

func main() {
	itemId := os.Args[1]
	itemName := os.Args[2]
	price, _ := strconv.ParseFloat(os.Args[3], 64)
	qty, _ := strconv.Atoi(os.Args[4])

	fmt.Println(db.AddItem(itemId, itemName, price, qty))
}
