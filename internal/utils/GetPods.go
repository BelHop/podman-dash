package utils

import (
	"context"
	"log"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities"
)

func GetPods() []entities.ListContainer {
	conn, err := bindings.NewConnection(context.Background(), "unix:///run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	Pods, err := containers.List(conn, new(containers.ListOptions).WithAll(true))
	return Pods
}

func GetPodNums() int {
	conn, err := bindings.NewConnection(context.Background(), "unix://run/user/1000/podman/podman.sock")
	if err != nil {
		log.Fatal(err)
	}
	Pods, err := containers.List(conn, new(containers.ListOptions).WithAll(true))
	return len(Pods)
}
