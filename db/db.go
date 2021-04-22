package db

import (
	"fmt"
	"shopping/models"
	"encoding/json"
	"github.com/go-redis/redis"
)

func LoadItem(itemId string) (string) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	res, err := client.Get(itemId).Result()
	if err != nil {
			fmt.Println(err)
	}

	// item := models.Item{}
	// json.Unmarshal([]byte(res), &item)

	return res
}

func AddItem(itemId, itemName string, price float64, qty int) (string){
	client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		})

	item, err1 := json.Marshal(models.Item{Desc: itemName, Price: price, Qty: qty})
	if err1 != nil {
			fmt.Println(err1)
	}

	err2 := client.Set(itemId, item, 0).Err()
	if err2 != nil {
		fmt.Println(err2)
	}

	return itemId + " - " + string(item) + " added to shop"
}