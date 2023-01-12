package handler

import (
	"CleanArch/features/user"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv user.UserService
}

func New(usrv user.UserService) user.UserHandler {
	return &userController{
		srv: usrv,
	}
}

// Register dari kontrak interface UserHandler di entity
func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := uc.srv.Register(*RequestToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already exist") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			}
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", res))
	}
}

// Login dari kontrak interface UserHandler di entity
func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(&input)
		if err != nil {
			log.Println("input error", err.Error())
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}
		useToken, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Println("", err.Error())
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "login successful", res, useToken))

	}
}

// Profile dari kontrak interface UserHandler di entity
func (uc *userController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		// log.Println(token)
		res, err := uc.srv.Profile(token)
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", res))
	}
}

// Update implements user.UserHandler
func (uc *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := UpdateRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := uc.srv.Update(c.Get("user"), *RequestToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}
		return c.JSON(PrintSuccessReponse(http.StatusCreated, "update successdul", res))

	}
}

// Delete implements user.UserHandler
func (uc *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			log.Println("fail to delete")
			if strings.Contains(err.Error(), "fail") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "fail to delete, account id not found",
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "server error",
				})
			}

		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "deleting account successful",
		})
	}
}
