package main

import (
	"fmt"
	"github.com/didintri196/mikrotik-restfull/domain/models"
	"github.com/didintri196/mikrotik-restfull/repositories"
	"github.com/didintri196/mikrotik-restfull/utils/routeros"
)

func main() {
	config := routeros.Configs{
		Auth: models.Auth{
			Ip:       "192.168.4.1",
			Port:     "8728",
			Username: "admin",
			Password: "segopecel12s",
		},
		Debug: true,
	}
	connRouteOS := routeros.New(config)

	repo := repositories.NewResourceRepository(&connRouteOS)
	var data models.Resource

	err := repo.Read(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
