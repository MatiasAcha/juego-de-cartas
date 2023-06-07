package juego

import (
	"fmt"
	"math/rand"
	"time"
)

type Mazo struct {
	Cartas []*Carta
	Total  int
}

func NewMazo() *Mazo {
	return &Mazo{}
}

func (m *Mazo) Mezclar() {
	rand.Seed(time.Now().UnixNano())

	for i := m.Total - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		m.Cartas[i], m.Cartas[j] = m.Cartas[j], m.Cartas[i]
	}
}

func (m *Mazo) AgregarCarta(carta *Carta) {
	if m.Total >= 80 {
		fmt.Println("El mazo está completo. No se puede agregar más cartas.")
		return
	}

	m.Cartas = append(m.Cartas, carta)
	m.Total++
}

func (m *Mazo) Robar() *Carta {
	if m.Total == 0 {
		fmt.Println("El mazo está vacío. No se puede robar ninguna carta.")
		return nil
	}

	carta := m.Cartas[m.Total-1]
	m.Cartas = m.Cartas[:m.Total-1]
	m.Total--

	return carta
}

func (m *Mazo) RobarCarta(carta *Carta) *Carta {
	var result *Carta
	var indice int

	for i := 0; i < m.Total; i++ {
		if m.Cartas[i] == carta {
			result = m.Cartas[i]
			indice = i
			break
		}
	}

	if carta != result {
		fmt.Println("La carta no se encuentra en el mazo.")
		return nil
	}

	for i := indice; i < m.Total-1; i++ {
		m.Cartas[i] = m.Cartas[i+1]
	}

	m.Total--
	m.Mezclar()
	return carta
}

func (m *Mazo) GetMazo() []*Carta {
	return m.Cartas
}

func (m *Mazo) String() string {
	result := ""
	for _, v := range m.Cartas {
		result += "("
		result += v.String()
		result += ")\n"

	}
	return result
}
