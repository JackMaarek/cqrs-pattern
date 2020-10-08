package models

type Order struct {
	ID    	   uint64 `gorm:"primary_key" json:"id"`
	TotalPrice uint64 `json:"total_price"`
	ItemCount  uint64 `json:"item_count"`
	UserId uint64   `json:"-"`
}