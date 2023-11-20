package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/Ignaciojeria/einar/app/domain/ports/out"
	einar "github.com/Ignaciojeria/einar/app/shared/archetype/resty"
)

var Shutdown out.Shutdown = func(ctx context.Context) (err error) {
	einar.LoadDependency()()
	client := einar.Client
	// Realiza una solicitud POST al endpoint de apagado
	resp, err := client.R().
		SetContext(ctx).
		Post("http://localhost:5555/api/shutdown")

	if err != nil {
		return err // Manejo de errores en caso de fallar la solicitud
	}

	// Verifica si la respuesta es exitosa
	if resp.StatusCode() != http.StatusOK {
		return errors.New("unexpected status code") // O maneja el error como prefieras
	}

	return nil
}
