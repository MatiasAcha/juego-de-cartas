package juego

type Carta struct {
	nombre   string
	atributo string
	poder    *Puntos
}

func NewCarta(poder *Puntos, nombre string, atributo string) *Carta {
	return &Carta{poder: poder, nombre: nombre, atributo: atributo}
}

func (c *Carta) Pelear(b *Carta) {

}

func (c *Carta) String() string {
	result := c.nombre + ", Puntos: " + c.poder.String()
	return result
}
