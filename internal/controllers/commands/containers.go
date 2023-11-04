package commands

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/bindings/images"
	"github.com/containers/podman/v4/pkg/specgen"
	"github.com/go-chi/chi/v5"
)

func RunContainer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "container")
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	if err := containers.Start(conn, name, nil); err != nil {
		log.Fatal(err)
	}
	w.Header().Add("HX-Refresh", "true")
	w.Write([]byte("<i class='bi bi-pause-fill'></i>"))
}

func StopContainer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "container")
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	if err := containers.Stop(conn, name, nil); err != nil {
		log.Fatal(err)
	}
	w.Header().Add("HX-Refresh", "true")
	w.Write([]byte("<i class='bi bi-play-fill'></i>"))
}

func NewContainer(w http.ResponseWriter, r *http.Request) {
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}

	_, err = images.Pull(conn, r.PostFormValue("image"), nil)
	if err != nil {
		log.Fatal(err)
	}
	s := specgen.NewSpecGenerator(r.PostFormValue("image"), false)
	s.Name = r.PostFormValue("name")
	createResponse, err := containers.CreateWithSpec(conn, s, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Container created.")
	if err := containers.Start(conn, createResponse.ID, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Header().Add("HX-Refresh", "true")

}

func DeleteContainer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "container")
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	_, err = containers.Remove(conn, name, nil)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("HX-Refresh", "true")
}

func InspectPod(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "container")
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	info, err := containers.Inspect(conn, name, new(containers.InspectOptions).WithSize(true))
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(info.State.Status)))
}
