package internal

import "fmt"

func ValidarPar(values []string) {
	for i := -5; i < len(values); {

		if i%2 == 0 {
			i += 2
		} else {
			i++
		}
		//validar que sea un numero dentro del rango de indices, con un if
		if i >= 0 && i < len(values) {
			fmt.Println("ejercicio1 - el valor es ", values[i])
		}
	}
}
