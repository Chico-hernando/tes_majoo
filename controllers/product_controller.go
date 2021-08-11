package controllers

import (
	"majoo/configs"
	"majoo/lib/database"
	"majoo/models/produk"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProdukController(c echo.Context) error {

	var produkCreate produk.ProdukCreate
	c.Bind(&produkCreate)

	var produkDB produk.Produk
	produkDB.NamaProduk = produkCreate.NamaProduk
	produkDB.SKU = produkCreate.SKU
	produkDB.Harga = produkCreate.Harga

	err := configs.DB.Create(&produkDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Create Data",
			err.Error(),
		))
	}


	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Create Produk",
		produkDB,
	))
}

func GetProdukController(c echo.Context) error {
	var produkData []produk.Produk
	var err error

	produkData, err = database.GetProduk()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			produkData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		produkData,
	))
}

func UpdateProdukController(c echo.Context) error {
	var produkUpdate produk.ProdukUpdate
	c.Bind(&produkUpdate)

	produkId,_ := strconv.Atoi(c.Param("id"))

	var produkDB produk.Produk
	configs.DB.First(&produkDB,"id",produkId)
	produkDB.Harga = produkUpdate.Harga

	err := configs.DB.Save(&produkDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data",
		produkDB,
	))
}

func DeleteProdukController(c echo.Context) error {
	produkId, _ := strconv.Atoi(c.Param("id"))

	var produkDB produk.Produk
	err := configs.DB.Where("id", produkId).Delete(&produkDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Delete Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success delete data",
		"",
	))
}