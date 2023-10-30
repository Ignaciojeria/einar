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
		Pattern: "home.html",
	})
	einar.EmbeddedPatterns = append(einar.EmbeddedPatterns, einar.EmbeddedPattern{
		Content: css,
		Pattern: "home.css",
	})
	container.InjectInboundAdapter(func() error {
		einar.Echo.GET("/home", render)
		einar.Echo.GET("/styles/home.css", echo.WrapHandler(http.FileServer(http.FS(css))))
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

// Ver la posibilidad de trasladar esto a un middleware
func render(c echo.Context) error {
	standalone := c.Request().Header.Get("standalone")
	data := map[string]interface{}{
		"componentName": "home",
	}
	if standalone == "true" {
		return c.Render(http.StatusOK, "home.html", data)
	}
	c.Request().Header.Set("component-name", "home")
	return c.Redirect(http.StatusSeeOther, "/?component-name=home")
}
