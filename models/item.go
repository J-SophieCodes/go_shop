package models

type Item struct {
	Desc string `json:"description"`
	Price float64 `json:"price"`
	Qty int `json:"quantity"`
}
