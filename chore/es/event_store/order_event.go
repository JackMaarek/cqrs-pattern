package estore

import (
	"encoding/json"
)

type OrderEvent struct {
	*Event
	UserID         uint64
	ItemCount      uint64
	TotalPrice 	   uint64
	OrderStatus    string
}

func (o *OrderEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *OrderEvent) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	return nil
}