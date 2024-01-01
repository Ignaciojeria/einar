package echo_server

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type EmbeddedPattern struct {
	Pattern string
	Content embed.FS
}

var EmbeddedPatterns []EmbeddedPattern

type TemplateRegistry struct {
	templates *template.Template
}

func NewTemplateRegistry(embeddedPatterns []EmbeddedPattern) *TemplateRegistry {
	t := template.New("")
	for _, embeddedPattern := range embeddedPatterns {
		_, err := t.ParseFS(embeddedPattern.Content, embeddedPattern.Pattern)
		if err != nil {
			panic(err)
		}
	}
	return &TemplateRegistry{
		templates: t,
	}
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func setUpRenderer(globPatterns ...EmbeddedPattern) {
	renderer := NewTemplateRegistry(globPatterns)
	Echo.Renderer = renderer
}
