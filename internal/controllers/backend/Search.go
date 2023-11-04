package backend

import (
	u "dashboard/internal/utils"
	"embed"
	"html/template"
	"net/http"

	"github.com/containers/podman/v4/pkg/domain/entities"
)

type Data struct {
	PodsInfo        []entities.ListContainer
	ContainersExist bool
}

var FSResult embed.FS

func Search(w http.ResponseWriter, r *http.Request) {
	q := r.PostFormValue("containerSearch")
	var qresults []entities.ListContainer
	containerArr := u.GetPods()
	var f bool

	for _, v := range containerArr {
		if containsSubstring(v.Names[0], q) {
			qresults = append(qresults, v)
		}
	}
	if len(qresults) == 0 {
		f = false
		w.Write([]byte("<h1>Make more Containers!</h1>"))
	} else {
		f = true
	}

	tmpl := template.Must(template.ParseFS(FSResult, "web/layouts/base.html", "web/components/navbar.html", "web/components/body/podinfo.html"))
	tmpl.ExecuteTemplate(w, "containers", Data{
		PodsInfo:        qresults,
		ContainersExist: f,
	})
}

func containsSubstring(str, substr string) bool {
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
