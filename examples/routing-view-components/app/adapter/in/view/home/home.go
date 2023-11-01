package controller

import (
	"embed"
	"my-project-name/app/shared/archetype/container"
	einar "my-project-name/app/shared/archetype/echo_server"
	"my-project-name/app/shared/constants"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const component = "home"

//go:embed *.html
var html embed.FS

//go:embed *.css
var css embed.FS

func init() {
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: html,
		Pattern: component + constants.DOT_HTML,
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: component + constants.DOT_CSS,
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/"+component, render)
		einar.Echo.GET("/"+component+constants.DOT_CSS, echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

// Ver la posibilidad de trasladar esto a un middleware
func render(c echo.Context) error {
	routerState := einar.NewRoutingState(c, map[string]string{
		"layoutComponentDefault":  "home",
		"contentComponentDefault": "empty",
	})
	if c.Request().Header.Get("FlatContext") != "" {
		return c.Render(http.StatusOK, "home.html", routerState)
	}
	return c.Render(http.StatusOK, "layout.html", routerState)
}
