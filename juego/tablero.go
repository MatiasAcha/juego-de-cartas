package juego

type Tablero struct {
	p1    *Player
	p2    *Player
	campo [3][3]*Carta
}

func NewTablero(player1 *Player, player2 *Player) *Tablero {
	return &Tablero{p1: player1, p2: player2}
}
