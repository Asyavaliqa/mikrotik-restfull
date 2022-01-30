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
			Ip:       "127.0.0.1",
			Port:     "28728",
			Username: "admin",
			Password: "",
		},
		Debug: true,
	}
	connRouteOS := routeros.New(config)

	//repo := repositories.NewResourceRepository(&connRouteOS)
	//var data models.Resource
	//repo.Read(&data)
	//fmt.Println(data)

	// ========================== //

	//repo := repositories.NewSecretRepository(&connRouteOS)
	//data := models.Secret{
	//	Name:     faker.Email(),
	//	Password: "passwordX",
	//	Service:  "any",
	//	Profile:  "default",
	//	Comment:  faker.Name(),
	//}
	//
	//err := repo.Add(data)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//======================//

	//repo := repositories.NewSecretRepository(&connRouteOS)
	//data, err := repo.Browse()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(data)

	//=====================//

	repo := repositories.NewSecretRepository(&connRouteOS)
	filter := models.Secret{
		Comment: "asd",
	}
	data, err := repo.Read(filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

}
