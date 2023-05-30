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

func (p *Player) robar() {
	p.mano = append(p.mano, p.deck.robar())
}

func (p *Player) IniciarTurno() {

}
