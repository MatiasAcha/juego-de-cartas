package juego

type Puntos struct {
	arriba    int
	abajo     int
	izquierda int
	derecha   int
}

func NewPuntos(a, b, c, d int) *Puntos {
	return &Puntos{arriba: a, abajo: b, izquierda: c, derecha: d}
}
