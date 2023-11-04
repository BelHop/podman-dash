package controllers

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/unix"

	u "dashboard/internal/utils"
)

type Data struct {
	MemoryPercent string
	UsedMemory    string
	Storage       string
	PodsNum       int
	PodsInfo      []entities.ListContainer
	Display       bool
}

var FSIndex embed.FS

func IndexController(w http.ResponseWriter, r *http.Request) {

	d := Data{
		MemoryPercent: fmt.Sprintf("%.2f", MemPercentage()),
		UsedMemory:    fmt.Sprintf("%.2f", 100-MemPercentage()),
		Storage:       GetStorage(),
		PodsNum:       u.GetPodNums(),
		PodsInfo:      u.GetPods(),
	}
	tmpl := template.Must(template.ParseFS(FSIndex, "web/layouts/base.html", "web/components/navbar.html", "web/components/statsCards.html"))
	tmpl.Execute(w, d)
}

func MemPercentage() float64 {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal("Could not initialize module")
	}
	return m.UsedPercent
}

func GetStorage() string {
	bf := float64(getStorageBytes())
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.2fYiB", bf)
}

func getStorageBytes() uint64 {
	var stat unix.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	unix.Statfs(wd, &stat)
	// Available blocks * size per block = available space in bytes
	return stat.Bavail * uint64(stat.Bsize)
}
