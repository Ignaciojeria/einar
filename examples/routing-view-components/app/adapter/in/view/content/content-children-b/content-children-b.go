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

const component = "content-children-b"

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
		einar.Echo.GET("/content/"+component, render)
		einar.Echo.GET(component+constants.DOT_CSS, echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

// Ver la posibilidad de trasladar esto a un middleware
func render(c echo.Context) error {
	data := map[string]interface{}{
		"layoutComponentDefault":  "content",
		"contentComponentDefault": "content/content-children-b",
	}
	standalone := c.Request().Header.Get("standalone")
	if standalone == "true" {
		return c.Render(http.StatusOK, "content-children-b.html", data)
	}
	return c.Render(http.StatusOK, "layout.html", data)
}
