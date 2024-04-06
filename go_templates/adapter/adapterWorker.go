package adapter

import (
	"fmt"
)

func AdapterWorker() {
	fmt.Println("---------- Adapter -----------")

	transport := NewBusAdapter(&Bus{})
	transport.Travel()

	transport = NewPlaneAdapter(&Plane{})
	transport.Travel()

	transport = NewTrainAdapter(&Train{})
	transport.Travel()

	fmt.Println("-------------------")
}
