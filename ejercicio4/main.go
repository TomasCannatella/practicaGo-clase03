/* Ejercicio 4 - Calcular estadísticas


Los profesores de una universidad de Colombia necesitan calcular
algunas estadísticas de calificaciones de los/as estudiantes de un curso.
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se
quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido)
que se le pueda pasar una cantidad N de enteros
devuelva el cálculo que se indicó en la función anterior.

Ejemplo: */

package main

import (
	"fmt"
	"slices"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operation(minimum)

	averageFunc, err := operation(average)

	maxFunc, err := operation(maximum)

	if err != "" {
		fmt.Println("Error is: ", err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Min value is: ", minValue)
	fmt.Println("Average value is:", averageValue)
	fmt.Println("Max value is: ", maxValue)

}

func operation(calculateType string) (func(values ...int) float64, string) {
	switch calculateType {
	case minimum:
		return minFunc, ""
	case average:
		return averageFunc, ""
	case maximum:
		return maxFunc, ""
	default:
		return nil, "It's function not exist"
	}
}

func minFunc(grades ...int) float64 {
	return float64(slices.Min(grades))
}

func averageFunc(grades ...int) (avarage float64) {
	var totalGrades float64
	for _, grade := range grades {
		totalGrades += float64(grade)
	}
	avarage = totalGrades / float64(len(grades))
	return
}

func maxFunc(grades ...int) float64 {
	return float64(slices.Max(grades))
}
