package models

import (
	"context"
	"log"
	"net/http"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/go-chi/chi/v5"
)

type RunModel struct {
	name string
}

func HandleContainer(w http.ResponseWriter, r *http.Request) string {
	name := chi.URLParam(r, "container")
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	inspect, err := containers.Inspect(conn, name, new(containers.InspectOptions).WithSize(true))
	if inspect.State.Status == "running" {
		containers.Stop(conn, name, nil)
		return "<li><button hx-Post=''><i class='bi bi-pause-fill'></i></button></li>"
	}

	containers.Start(conn, name, nil)
	return "<li><button hx-Post=''><i class='bi bi-play-fill'></i></button></li>"
}

func (rm *RunModel) buttonChange() string {
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	inspect, err := containers.Inspect(conn, rm.name, new(containers.InspectOptions).WithSize(true))
	if inspect.State.Status == "running" {
		return "<li><button hx-Post=''><i class='bi bi-pause-fill'></i></button></li>"
	}

	return "<li><button hx-Post=''><i class='bi bi-play-fill'></i></button></li>"
}
