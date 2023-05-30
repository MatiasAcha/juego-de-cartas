package juego

import (
	"errors"
)

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack struct {
	array []*Carta
}

func NewStack() *Stack {
	return &Stack{}
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack) Push(v *Carta) {
	s.array = append(s.array, v)
}

// Pop saca un elemento de la pila. O(1)
func (s *Stack) Pop() (*Carta, error) {
	if (s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return v, nil
}

// Top devuelve el elemento del tope de la pila. O(1)
func (s *Stack) Top() (*Carta, error) {
	if (s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := (s.array)[len(s.array)-1]
	return v, nil
}

// IsEmpty verifica si la pila esta vacia. O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.array) == 0
}
