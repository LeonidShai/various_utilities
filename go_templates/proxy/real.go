package proxy

// Структура реального объекта (какое-нибудь хранилище)
type warehouse struct {
	shelfs map[int]string
}

func NewWarehouse() *warehouse {
	return &warehouse{
		shelfs: map[int]string{
			0: "c++",
			1: "golang",
			2: "python",
		},
	}
}

func (w *warehouse) GetData(indx int) (string, bool) {
	result, ok := w.shelfs[indx]

	return result, ok
}
