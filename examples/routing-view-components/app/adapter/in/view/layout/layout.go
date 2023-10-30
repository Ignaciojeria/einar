package controller

import (
	"embed"
	"my-project-name/app/shared/archetype/container"
	einar "my-project-name/app/shared/archetype/echo_server"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:embed *.html
var html embed.FS

//go:embed *.css
var css embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: "layout.html",
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: "layout.css",
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/", render)
		einar.Echo.GET("/layout.css", echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	componentName := c.QueryParam("component-name")

	if componentName == "" {
		componentName = "content"
	}

	data := map[string]interface{}{
		"componentName": componentName,
	}
	return c.Render(http.StatusOK, "layout.html", data)
}