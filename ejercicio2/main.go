/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
y devuelva el promedio. No se pueden introducir notas negativas.
*/
package main

import "fmt"

func main() {
	grades := []float64{10, 7, 8, 8, 8, 7, 8, 9, 10}
	fmt.Println("El promedio de los alumnos es de: ", avgGrade(grades...))
}

func avgGrade(grades ...float64) (avarage float64) {
	var totalGrades float64
	for _, grade := range grades {
		totalGrades += grade
	}
	avarage = totalGrades / float64(len(grades))
	return
}
