package interpreter

// Шаблон Интерпретатор
// определяет грамматику простого языка,
// представляет предложения на этом языке
// и интерпретирует их.
// Состоит из парсера и объектов, каждый
// из которых отвечает конкретной интерпретации

type Interpreter interface {
	Interpret(int, int) int
}

type AddValue struct {
}

func (v *AddValue) Interpret(left, right int) int {
	return left + right
}

type SubValue struct {
}

func (v *SubValue) Interpret(left, right int) int {
	return left - right
}
