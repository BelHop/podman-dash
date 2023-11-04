package internal

import (
	"dashboard/internal/controllers"
	"dashboard/internal/controllers/backend"
	"dashboard/internal/controllers/commands"
	"embed"
	_ "embed"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PageData struct {
	ModalHTML string
}

//go:embed web/*
var embedFS embed.FS

func Routes() {
	controllers.FSIndex, controllers.FSContainer, commands.FSReplaceI, backend.FSResult = embedFS, embedFS, embedFS, embedFS

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.IndexController)
	r.Route("/containers", func(r chi.Router) {
		r.Get("/", controllers.ContainerController)
		r.Post("/new", commands.NewContainer)
		r.Post("/search", backend.Search)

		r.Route("/{container}", func(r chi.Router) {
			r.Post("/run", commands.RunContainer)
			r.Post("/stop", commands.StopContainer)
			r.Post("/delete", commands.DeleteContainer)
			r.Get("/inspect", commands.InspectPod)
		})
	})
	r.Route("/sys", func(r chi.Router) {
		r.Get("/mem", commands.MemPercentage)
	})
	http.ListenAndServe(":8080", r)
}
