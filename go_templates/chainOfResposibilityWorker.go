package main

import (
	"fmt"
	chainofresponsibility "go_templates/chain_of_responsibility"
)

func chainOfResponsibilityWorker() {
	fmt.Println("---------- Chain of Responsibility -----------")
	// будем варить суп =))
	soup := &chainofresponsibility.Soup{}

	// заведем процессор для варки супа =))
	cookSoup := chainofresponsibility.CookSoupProcessor{}

	// заведем объекты-действия для варки супа
	pourWater := &chainofresponsibility.PourWaterPanProcess{}
	peelVegetables := &chainofresponsibility.PeelVegetablesProcess{}
	putVegetablesPan := &chainofresponsibility.PutVegetablesPanProcess{}
	putSeasoning := &chainofresponsibility.PutSeasoningProcess{}
	boil := &chainofresponsibility.BoilProcess{}
	eat := &chainofresponsibility.EatProcess{}

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
