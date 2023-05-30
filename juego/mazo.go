package juego

type Mazo struct {
	mazo *Stack
}

func NewMazo() *Mazo {
	m := NewStack()
	return &Mazo{mazo: m}
}

func (m *Mazo) robar() *Carta {
	v, _ := m.mazo.Top()
	return v
}
