package controllers

import (
	"embed"
	"html/template"
	"net/http"

	u "dashboard/internal/utils"

	"github.com/containers/podman/v4/pkg/domain/entities"
)

type ContainerData struct {
	PodsInfo        []entities.ListContainer
	ContainersExist bool
}

var FSContainer embed.FS

func ContainerController(w http.ResponseWriter, r *http.Request) {
	var f bool
	if len(u.GetPods()) == 0 {
		f = false
	} else {
		f = true
	}

	d := ContainerData{
		PodsInfo:        u.GetPods(),
		ContainersExist: f,
	}
	tmpl := template.Must(template.ParseFS(FSContainer, "web/layouts/base.html", "web/components/navbar.html", "web/components/body/podinfo.html"))
	tmpl.Execute(w, d)
}
