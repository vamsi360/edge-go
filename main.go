package main

import (
	"fmt"

	"github.com/edge-go/core"
	"github.com/edge-go/service"
)

func main() {
	fmt.Printf("Starting app..\n")

	serviceRepo := service.NewServiceRepo()

	serviceDef1 := core.NewServiceDef("http", "localhost", 18080)
	serviceRepo.Register("/hello", serviceDef1)

	repo1 := serviceRepo.Get("/hello")
	fmt.Printf("Repo1: %v\n", repo1)
}
