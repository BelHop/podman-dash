package commands

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v3/mem"
)

var FSReplaceI embed.FS

type MemReplace struct {
	MemoryPercent string
	UsedMemory    string
}

func MemPercentage(w http.ResponseWriter, r *http.Request) {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal("Could not initialize module")
	}
	s := fmt.Sprintf("%.2f", m.UsedPercent)
	su := fmt.Sprintf("%.2f", 100-m.UsedPercent)
	tmpl := template.Must(template.ParseFS(FSReplaceI, "web/layouts/base.html", "web/components/navbar.html", "web/components/statsCards.html"))

	d := MemReplace{
		MemoryPercent: s,
		UsedMemory:    su,
	}

	tmpl.ExecuteTemplate(w, "mem", d)

}
