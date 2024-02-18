package adapter

type Transport interface {
	Travel()
}

// адаптеры для этих структур
type BusAdapter struct {
	Bus *Bus
}

func (ba *BusAdapter) Travel() {
	ba.Bus.TravelByBus()
}

func NewBusAdapter(bus *Bus) Transport {
	return &BusAdapter{
		Bus: bus,
	}
}

type PlaneAdapter struct {
	Plane *Plane
}

func (pa *PlaneAdapter) Travel() {
	pa.Plane.TravelByPlane()
}

func NewPlaneAdapter(plane *Plane) Transport {
	return &PlaneAdapter{
		Plane: plane,
	}
}

type TrainAdapter struct {
	Train *Train
}

func (ta *TrainAdapter) Travel() {
	ta.Train.TravelByTrain()
}

func NewTrainAdapter(train *Train) Transport {
	return &TrainAdapter{
		Train: train,
	}
}
