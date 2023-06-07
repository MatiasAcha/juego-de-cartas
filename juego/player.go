package juego

type Player struct {
	nombre     string
	deck       *Mazo
	cementerio *Mazo
	vida       int
	mano       []*Carta
}

func NewPlayer(name string, mazo *Mazo) *Player {
	c := NewMazo()
	return &Player{nombre: name, deck: mazo, cementerio: c, vida: 20}
}

func (p *Player) Robar() bool {
	carta := p.deck.Robar()
	if carta == nil {
		return true
	}
	p.mano = append(p.mano, carta)
	return false
}

func (p *Player) IniciarTurno() {

}

func (p *Player) GetMano() []*Carta {
	return p.mano
}

func (p *Player) GetDeck() *Mazo {
	return p.deck
}
