package core

func SampleLoginHandlerCode() string {
	return `package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) login(c echo.Context) error {
	const op = "httpserver.login"

	var req param.GetUserReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.svc.GetByID(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, resp)
}`
}

func SampleHandlerCode() string {
	return `package httpserver

type Handler struct {
	svc            service.Service
}

func New(svc  service.Service) Handler {
	return Handler{
		svc:               svc,
	}
}`
}

func SampleRouteCode() string {
	return `package httpserver

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	v1 := e.Group("/v1/users")

	v1.GET("/login", h.login)
}`
}
