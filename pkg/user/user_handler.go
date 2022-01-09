package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	userService Service
}

func NewHandler(userService Service) *handler {
	return &handler{userService: userService}
}

func InitRESTHandler(e *echo.Echo, userService Service) {
	h := handler{
		userService: userService,
	}

	e.POST("/users/login", h.login)
	e.POST("/users/register", h.register)
}

func (h *handler) login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	fmt.Println(req)

	ctx := c.Request().Context()

	res, err := h.userService.Login(ctx, &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	res, err := h.userService.Register(ctx, &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
