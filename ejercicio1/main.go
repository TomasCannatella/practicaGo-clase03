/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados
al momento de depositar el sueldo, para cumplir el objetivo es necesario
crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo
y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

package main

import "fmt"

func main() {
	salary := 150000.1
	fmt.Println("El impuesto para el empleado con un sueldo de $", salary, " es: $", calculateTax(salary))
}

func calculateTax(salary float64) (tax float64) {
	switch {
	case salary > 150000:
		tax += (salary * 10) / 100
		fallthrough
	case salary > 50000:
		tax += (salary * 17) / 100
	default:
		tax = 0
	}
	return
}
