package api

import (
	utils "echo-template/utils"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func SendMessageToSlack(c echo.Context) error {
	if err := utils.SendMessageToSlack(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func Login(c echo.Context) error {
	jwt, err := utils.CreateJWT(1, 10)
	if err != nil {
		log.Printf("%v\n", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	log.Printf("%v\n", jwt)
	c.Response().Header().Set("Authorization", "Bearer"+jwt)
	return c.NoContent(http.StatusOK)
}

func GetUser(c echo.Context) error {
	jwt := c.Get("jwt").(*jwt.Token)
	claims := jwt.Claims.(*utils.JwtCustomClaims)

	log.Printf("%v\n", claims.UserID)
	return c.NoContent(http.StatusOK)
}

// func Healthcheck(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "health check.")
// }

// type (
// 	User struct {
// 		Name  string `json:"name" form:"name"`
// 		Email string `json:"email" form:"email"`
// 	}
// 	handler struct {
// 		db map[string]*User
// 	}
// )

// func (h *handler) CreateUser(c echo.Context) error {
// 	s := "[CreateUser]"

// 	u := new(User)
// 	if err := c.Bind(u); err != nil {
// 		log.Printf("%s %v\n", s, err)
// 		return err
// 	}
// 	return c.JSON(http.StatusCreated, u)
// }

// func (h *handler) GetUser(c echo.Context) error {
// 	email := c.Param("email")
// 	user := h.db[email]
// 	if user == nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "user not found")
// 	}
// 	return c.JSON(http.StatusOK, user)
// }
