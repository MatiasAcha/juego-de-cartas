package juego

type Juego struct {
	tablero *Tablero
	turno   int
	ganador string
	p1      *Player
	p2      *Player
}

func NewJuego(tab *Tablero, player1 *Player, player2 *Player) *Juego {
	return &Juego{tablero: tab, turno: 1, ganador: "Empate", p1: player1, p2: player2}
}

func (j *Juego) IniciarJuego() {
	j.IniciarTurno()
}

func (j *Juego) IniciarTurno() {
	if j.turno == 1 {
		for i := 0; i < 6; i++ {
			j.p1.robar()
			j.p2.robar()
			j.p1.IniciarTurno()
		}
	} else if j.turno%2 != 0 {
		j.p1.robar()
	} else {
		return
	}
}
