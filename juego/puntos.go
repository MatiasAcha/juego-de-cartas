package juego

import "fmt"

type Puntos struct {
	arriba    int
	abajo     int
	izquierda int
	derecha   int
}

func NewPuntos(a, b, c, d int) *Puntos {
	return &Puntos{arriba: a, abajo: b, izquierda: c, derecha: d}
}

func (p *Puntos) String() string {
	result := "Arriba: " + fmt.Sprint(p.arriba) + ", Abajo: " + fmt.Sprint(p.abajo) + ", Izquierda: " + fmt.Sprint(p.izquierda) + ", Derecha: " + fmt.Sprint(p.derecha)
	return result
}
