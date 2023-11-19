package controller

import (
	"context"
	"net/http"

	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	einar "github.com/Ignaciojeria/einar/app/shared/archetype/echo_server"

	"github.com/labstack/echo/v4"
)

func init() {
	container.InjectInboundAdapter(func() error {
		einar.Echo.POST("/api/shutdown", shutdown)
		return nil
	})
}

func shutdown(c echo.Context) error {
	if err := einar.Echo.Shutdown(context.Background()); err != nil {
		c.Logger().Fatal(err)
		return c.JSON(http.StatusOK, "err")
	}
	return c.JSON(http.StatusOK, "Server is shutting down")
}
