package forms

type OrderForm struct {
	TotalPrice uint64 `json:"total_price" validate:"required"`
	ItemCount  uint64 `json:"item_count" validate:"required"`
}
