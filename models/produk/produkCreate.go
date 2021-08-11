package produk

type ProdukCreate struct {
	NamaProduk string    `json:"namaProduk"`
	SKU        string    `json:"sku"`
	Harga      int       `json:"harga"`
}