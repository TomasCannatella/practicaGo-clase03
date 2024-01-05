/*
Ejercicio 5 - Calcular cantidad de alimento


Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos,
pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

Perro     010  kg de alimento.
Gato      005  kg de alimento.
Hamster   250   g de alimento.
Tarántula 150   g de alimento.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:
*/

package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalDog, err := animal(dog)
	if err != "" {
		fmt.Println("Error:", err)
		return
	}
	animalCat, err := animal(cat)
	if err != "" {
		fmt.Println("Error:", err)
	}

	var amount float64

	amount += animalDog(10)

	amount += animalCat(10)
	fmt.Println("La cantidad de alimento que se necesita son: ", amount, "kg")
}

func animal(animalType string) (function func(quantity int) float64, err string) {
	switch animalType {
	case dog:
		function = animalDog
	case cat:
		function = animalCat
	case hamster:
		function = animalHamster
	case tarantula:
		function = animalTarantula
	default:
		err = "This animal not exist"
	}
	return
}

func animalDog(quantity int) float64 {
	return float64(quantity * 10)
}

func animalCat(quantity int) float64 {
	return float64(quantity * 5)
}

func animalHamster(quantity int) float64 {
	return float64(quantity) * 0.250
}

func animalTarantula(quantity int) float64 {
	return float64(quantity) * 0.150
}
