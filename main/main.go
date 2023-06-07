package main

import (
	"fmt"
	"modulo/juego"
)

func main() {
	puntosCarta1 := juego.NewPuntos(10, 6, 9, 4)
	puntosCarta2 := juego.NewPuntos(5, 3, 9, 10)
	carta1 := juego.NewCarta(puntosCarta1, "Squall", "")
	carta2 := juego.NewCarta(puntosCarta2, "Laguna", "")
	Mazo1 := juego.NewMazo()
	Mazo1.AgregarCarta(carta1)
	Mazo1.AgregarCarta(carta2)
	p1 := juego.NewPlayer("Pepe", Mazo1)
	fmt.Println("Mano: " + fmt.Sprint(p1.GetMano()))
	fmt.Println("Deck: " + p1.GetDeck().String())
	fmt.Println()
	fmt.Println("¡ROBA!")
	fmt.Println()
	p1.Robar()
	fmt.Println("Mano: " + fmt.Sprint(p1.GetMano()))
	fmt.Println("Deck: " + p1.GetDeck().String())

}
