package routes

import (
	"majoo/constants"
	"majoo/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e:= echo.New()
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	eAuth.POST("/produk", controllers.CreateProdukController)
	eAuth.GET("/produk", controllers.GetProdukController)
	eAuth.PUT("/produk/:id", controllers.UpdateProdukController)
	eAuth.DELETE("/produk/:id", controllers.DeleteProdukController)

	eAuth.GET("/user", controllers.GetUserController)
	eAuth.PUT("/user/:id", controllers.UpdateUserController)
	eAuth.DELETE("/user/:id", controllers.DeleteUserController)

	return e
}