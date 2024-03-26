package main

import (
	"fmt"
	"go_templates/strategy"
)

func strategyWorker() {
	fmt.Println("---------- Strategy -----------")
	// создадим наш класс-клиент-исполнитель
	ourStrategy := strategy.WhatDo{}

	// можем создать сначала одну стратегию
	sitting := strategy.ChairSitting{}
	// и использовать ее
	ourStrategy.SetNewStrategy(&sitting)
	ourStrategy.ExecuteOurStrategy()

	// а теперь можем создать (тут и/или) другую стратегию
	eatting := strategy.EatSoup{}

	// и использовать уже ее
	ourStrategy.SetNewStrategy(&eatting)
	ourStrategy.ExecuteOurStrategy()

	fmt.Println("-------------------")
}
