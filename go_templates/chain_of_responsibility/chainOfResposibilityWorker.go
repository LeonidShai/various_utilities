package chainofresponsibility

import (
	"fmt"
)

func ChainOfResponsibilityWorker() {
	fmt.Println("---------- Chain of Responsibility -----------")
	// будем варить суп =))
	soup := &Soup{}

	// заведем процессор для варки супа =))
	cookSoup := CookSoupProcessor{}

	// заведем объекты-действия для варки супа
	pourWater := &PourWaterPanProcess{}
	peelVegetables := &PeelVegetablesProcess{}
	putVegetablesPan := &PutVegetablesPanProcess{}
	putSeasoning := &PutSeasoningProcess{}
	boil := &BoilProcess{}
	eat := &EatProcess{}

	// расставим последовательность действий
	cookSoup.SetNextHandler(pourWater)
	pourWater.SetNextHandler(peelVegetables)
	peelVegetables.SetNextHandler(putVegetablesPan)
	putVegetablesPan.SetNextHandler(putSeasoning)
	putSeasoning.SetNextHandler(boil)
	boil.SetNextHandler(eat)

	// запуск последовательности
	cookSoup.Handle(soup)
	if soup.Eat {
		fmt.Println("Soup was ready and eaten")
	} else {
		fmt.Println("Something wrong !!!")
	}

	fmt.Println("-------------------")
}
