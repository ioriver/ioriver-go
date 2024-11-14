package main

import (
	"fmt"

	ioriver "github.com/ioriver/ioriver-go"
)

func main() {
	client := ioriver.NewClient("xgzcnygxb733310edbd2cdd2d043957b36316243e78b396e")
	client.EndpointUrl = "http://127.0.0.1:8000/api/v1/"

	service := ioriver.Service{
		Name:        "test-new",
		Description: "ffff",
		Certificate: "dab94e85-c078-4c12-be50-8c1c8f0163b4",
	}

	newSp, err := client.CreateService(service)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newSp)
	}

	// sp := ioriver.ServiceProvider{
	// 	AccountProvider: "406ce0f8-e868-44c5-b9eb-282b1635bc6b",
	// 	Service:         "2164c699-277c-48d1-a81d-d53929062d0a",
	// 	Weight:          0,
	// }
	// fmt.Println(sp)

	// newSp, err := client.CreateServiceProvider(sp)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(newSp)
	// }

	// client.DeleteServiceProvider(newSp.Service, newSp.Id, "delete")
}
