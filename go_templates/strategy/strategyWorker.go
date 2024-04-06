package strategy

import (
	"fmt"
)

func StrategyWorker() {
	fmt.Println("---------- Strategy -----------")
	// создадим наш класс-клиент-исполнитель
	ourStrategy := WhatDo{}

	// можем создать сначала одну стратегию
	sitting := ChairSitting{}
	// и использовать ее
	ourStrategy.SetNewStrategy(&sitting)
	ourStrategy.ExecuteOurStrategy()

	// а теперь можем создать (тут и/или) другую стратегию
	eatting := EatSoup{}

	// и использовать уже ее
	ourStrategy.SetNewStrategy(&eatting)
	ourStrategy.ExecuteOurStrategy()

	fmt.Println("-------------------")
}
