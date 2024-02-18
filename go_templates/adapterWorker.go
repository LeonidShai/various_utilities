package main

import (
	"fmt"
	adapter "go_templates/adapter"
)

func adapterWorker() {
	fmt.Println("---------- Adapter -----------")

	transport := adapter.NewBusAdapter(&adapter.Bus{})
	transport.Travel()

	transport = adapter.NewPlaneAdapter(&adapter.Plane{})
	transport.Travel()

	transport = adapter.NewTrainAdapter(&adapter.Train{})
	transport.Travel()

	fmt.Println("-------------------")
}
