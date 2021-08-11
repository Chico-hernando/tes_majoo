package produk

import "majoo/models/base"

type ProductResponse struct {
	base.BaseResponse
	Data []Produk `json:"data"`
}