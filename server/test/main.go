package main

import (
	"fmt"
	"github.com/didintri196/mikorm"
	"github.com/didintri196/mikrotik-restfull/domain/models"
	"github.com/didintri196/mikrotik-restfull/repositories"
)

func main() {
	config := mikorm.Configs{
		Ip:        "127.0.0.1",
		Port:      "28728",
		Username:  "admin",
		Password:  "",
		ModeDebug: true,
	}
	connRouteOS := mikorm.New(config)

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
		Name: "xx@gmail.com",
	}

	data, err := repo.Read(filter)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	// delete secret
	err = repo.Disable(data.ID)
	if err != nil {
		fmt.Println(err)
	}
}
