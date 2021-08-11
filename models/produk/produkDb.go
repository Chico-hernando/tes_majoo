package produk

import "time"

type Produk struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	NamaProduk string    `json:"namaProduk"`
	SKU        string    `json:"sku"`
	Harga      int       `json:"harga"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
