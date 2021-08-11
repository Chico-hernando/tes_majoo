package controllers

import (
	"majoo/configs"
	"majoo/lib/database"
	"majoo/middlewares"
	"majoo/models/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateUserController(c echo.Context) error{
	var userCreate user.UserCreate
	c.Bind(&userCreate)
	hashedPassword, err := HashPassword(userCreate.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Register User",
			err.Error(),
		))
	}

	var userDB user.User
	userDB.Email = userCreate.Email
	userDB.Nama = userCreate.Nama
	userDB.Password = string(hashedPassword)

	e := configs.DB.Create(&userDB).Error
	if e != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Register Customer",
			e.Error(),
		))
	}

	var userResponse = user.UserResponse{
		Id: userDB.Id,
		Nama: userDB.Nama,
		Email: userDB.Email,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success Register Customer",
		userResponse,
	))
}

func LoginController(c echo.Context) error {
	var userLogin user.UserLogin
	c.Bind(&userLogin)

	var userDB user.User
	configs.DB.First(&userDB,"email",userLogin.Email)
	hashedPassword := userDB.Password

	err := VerifyPassword(hashedPassword, userLogin.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Login",
			"Wrong Email or Password",
		))
	}

	token, _ := middlewares.GenerateJWT(userDB.Id)

	var userResponse = user.LoginResponse{
		Email: userDB.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Login",
		userResponse,
	))
}

func GetUserController(c echo.Context) error{
	var userData []user.User
	var err error

	userData, err = database.GetUser()

	if err != nil {
		return c.JSON(http.StatusOK, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			userData,
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get Data",
		userData,
	))
}

func UpdateUserController(c echo.Context) error{
	var userUpdate user.UserUpdate
	userId,_ := strconv.Atoi(c.Param("id"))
	c.Bind(userUpdate)

	hashedPassword, _ := HashPassword(userUpdate.Password)

	var userDB user.User
	configs.DB.First(&userDB,"id",userId)
	userDB.Password = string(hashedPassword)

	err := configs.DB.Save(&userDB).Error

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
		userDB,
	))
}

func DeleteUserController(c echo.Context) error{
	userId, _ := strconv.Atoi(c.Param("id"))

	var userDB user.User
	err := configs.DB.Where("id", userId).Delete(&userDB).Error

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