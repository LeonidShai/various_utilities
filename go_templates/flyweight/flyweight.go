package flyweight

import "fmt"

// Паттерн Приспособленец
// нужен для переиспользования одних и тех же похожих объектов (храним в специальной мапе,
// из которой при необходимости их достаем).
// Есть один или несколько объектов, которые можно использовать
// одновременно в нескольких контекстах

// заведем хранилище простых объектов, из которого
// будем дергать при необходимости эти объекты
type dressClosetFactory struct {
	dresses map[string][]Dress
}

func NewDressClosetFactory(dresses ...Dress) dressClosetFactory {
	dcf := dressClosetFactory{
		dresses: make(map[string][]Dress),
	}
	for _, dress := range dresses {
		dcf.dresses[dress.DressType()] = append(dcf.dresses[dress.DressType()], dress)
	}

	return dcf
}

func (dcf *dressClosetFactory) GetDress(dressType string, color string) Dress {
	dresses, ok := dcf.dresses[dressType]
	if !ok {
		fmt.Printf("Bad dress type: %s\n", dressType)

		return nil
	}

	for _, dress := range dresses {
		if dress.Color() == color {
			return dress
		}
	}

	fmt.Printf("Maybe dress with %s color in washing =))\n", color)

	return nil
}
