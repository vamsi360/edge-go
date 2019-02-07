package main

import (
	"fmt"

	"github.com/edge-go/core"
	"github.com/edge-go/service"
)

func main() {
	fmt.Printf("Starting app..\n")

	serviceRepo := service.NewServiceRepo()

	tstoreSD := core.NewServiceDef("http", "localhost", 18080)
	serviceRepo.Register(tstoreSD)
	serviceRepo.RegisterEdge("/changes", tstoreSD)

	sdef1 := serviceRepo.GetDef("/changes")
	fmt.Printf("Repo1: %v\n", sdef1)

	sdefs := serviceRepo.GetDefs()
	fmt.Printf("Registered ServiceDefs: %v\n", sdefs)
}
