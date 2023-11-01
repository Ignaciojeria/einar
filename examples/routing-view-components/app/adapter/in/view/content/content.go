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
		Pattern: "content.html",
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: "content.css",
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/content", render)
		einar.Echo.GET("/content.css", echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		"layoutComponentDefault":  "content",
		"contentComponentDefault": "empty",
	})
	if c.Request().Header.Get("FlatContext") != "" {
		return c.Render(http.StatusOK, "content.html", routerState)
	}
	return c.Render(http.StatusOK, "layout.html", routerState)
}
