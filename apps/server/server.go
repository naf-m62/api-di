package server

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"

	"api-di/apps/handlers"
	"api-di/apps/middlewares"
	"api-di/services"
)

func Start(app di.Container) {
	var port string
	port = os.Getenv("SERVER_PORT")
	r := mux.NewRouter()

	// Function to apply the middlewares:
	// - recover from panic
	// - add the container in the http requests
	m := func(h http.HandlerFunc) http.HandlerFunc {
		return middlewares.PanicRecoveryMiddleware(
			di.HTTPMiddleware(h, app, func(msg string) {
				services.Logger.Error(msg)
			}),
			services.Logger,
		)
	}

	getHand := handlers.Handler{
		Tmpl: template.Must(template.ParseGlob("./apps/handlers/templates/*")),
	}

	r.HandleFunc("/", m(getHand.GetLinkListHandler)).Methods("GET")
	r.HandleFunc("/create", m(getHand.CreateLinkForm)).Methods("GET")
	r.HandleFunc("/stand/create", m(getHand.CreateStandLinkForm)).Methods("GET")
	r.HandleFunc("/links/{id}", m(getHand.UpdateLinkForm)).Methods("GET")

	r.HandleFunc("/create", m(handlers.CreateLinkHandler)).Methods("POST")
	r.HandleFunc("/stand/create", m(handlers.CreateStandLinkHandler)).Methods("POST")
	r.HandleFunc("/links/{id}", m(handlers.UpdateLinkHandler)).Methods("POST")
	r.HandleFunc("/links/{id}", m(handlers.DeleteLinkHandler)).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	services.Logger.Info("Listening on port " + port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			services.Logger.Error(err.Error())
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	services.Logger.Info("Stopping the http server")

	if err := srv.Shutdown(ctx); err != nil {
		services.Logger.Error(err.Error())
	}
}
