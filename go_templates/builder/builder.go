package builder

type ChairBuilder interface {
	Sitting(string) ChairBuilder
	ChairLegs(int) ChairBuilder
	ChairBack(string) ChairBuilder

	Build() chair
}

type chair struct {
	chairLegs int
	chairBack string
	sitting   string
}

type chairBuilder struct {
	chairLegs int
	chairBack string
	sitting   string
}

func NewChairBuilder() ChairBuilder {
	return &chairBuilder{
		chairLegs: 4,
		chairBack: "red",
		sitting:   "wood",
	}
}

func (c *chairBuilder) Sitting(material string) ChairBuilder {
	c.sitting = material
	return c
}

func (c *chairBuilder) ChairLegs(num int) ChairBuilder {
	c.chairLegs = num
	return c
}

func (c *chairBuilder) ChairBack(backColor string) ChairBuilder {
	c.chairBack = backColor

	return c
}

func (c *chairBuilder) Build() chair {
	return chair{
		chairLegs: c.chairLegs,
		chairBack: c.chairBack,
		sitting:   c.sitting,
	}
}

// разные виды билдеров (в зависимости от типа)
type officeChairBuilder struct {
	chairBuilder
}

func NewOfficeChairBuilder() ChairBuilder {
	return &officeChairBuilder{
		chairBuilder: chairBuilder{
			chairLegs: 5,
			chairBack: "black",
			sitting:   "plastic",
		},
	}
}

func (o *officeChairBuilder) Build() chair {
	return chair{
		chairLegs: o.chairLegs,
		chairBack: o.chairBack,
		sitting:   o.sitting,
	}
}

type homeChairBuilder struct {
	chairBuilder
}

func NewHomeChairBuilder() ChairBuilder {
	return &homeChairBuilder{
		chairBuilder: chairBuilder{
			chairLegs: 4,
			chairBack: "green",
			sitting:   "wood",
		},
	}
}

func (o *homeChairBuilder) Build() chair {
	return chair{
		chairLegs: o.chairLegs,
		chairBack: o.chairBack,
		sitting:   o.sitting,
	}
}
