package builder

// Директору выдаём билдер (можно задавать разных) и директор может билдитьт объекты
type director struct {
	chairBuilder ChairBuilder
}

func NewDirector(b ChairBuilder) director {
	return director{
		chairBuilder: b,
	}
}

func (d *director) BuildChair() chair {
	return d.chairBuilder.Build()
}

func (d *director) SetBuilder(b ChairBuilder) {
	d.chairBuilder = b
}
