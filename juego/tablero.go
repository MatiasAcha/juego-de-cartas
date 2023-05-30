package juego

type Tablero struct {
	campo [3][3]*Carta
}

func NewTablero() *Tablero {
	return &Tablero{}
}
