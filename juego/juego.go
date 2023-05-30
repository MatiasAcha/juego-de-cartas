package juego

type juego struct {
	tablero Tablero
	turno   int
	ganador string
}

func NewJuego(tab Tablero) *juego {
	return &juego{tablero: tab, turno: 1, ganador: "Empate"}
}
