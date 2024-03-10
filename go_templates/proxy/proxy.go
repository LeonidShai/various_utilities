package proxy

// Заместитель ведет себя точно так же как и Реальный объект (real.go)
// и при необходимости создает экземпляр реального объекта
// можно отложить создание реального объекта до его фактической необходимости
type Service interface {
	GetData(int) (string, bool)
}

type ProxyWarehouse struct {
	warechouse *warehouse
}

func (p *ProxyWarehouse) GetData(indx int) (string, bool) {
	if p.warechouse == nil {
		p.warechouse = NewWarehouse()
	}

	return p.warechouse.GetData(indx)
}
