package chi_server

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func init() {
	config.Installations.EnableHTTPServer = true
	container.InjectHTTPServer(func() error {
		fmt.Println("starting server on port :" + config.PORT.Get())
		r := chi.NewRouter()
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
		err := http.ListenAndServe(":"+config.PORT.Get(), r)
		if err != nil {
			log.Error().Err(err).Msg("error initializing application server")
			return err
		}
		return nil
	}, container.InjectionProps{Paralel: false, DependencyID: uuid.NewString()})
}
