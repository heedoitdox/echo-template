package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func Healthcheck(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "health check.")
// }

type (
	User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
	handler struct {
		db map[string]*User
	}
)

func (h *handler) CreateUser(c echo.Context) error {
	s := "[CreateUser]"

	u := new(User)
	if err := c.Bind(u); err != nil {
		log.Printf("%s %v\n", s, err)
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *handler) GetUser(c echo.Context) error {
	email := c.Param("email")
	user := h.db[email]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}
