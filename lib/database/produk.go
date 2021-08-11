package database

import (
	"majoo/configs"
	"majoo/models/produk"
)

func GetProduk() (dataresult []produk.Produk, err error) {
	err = configs.DB.Find(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}