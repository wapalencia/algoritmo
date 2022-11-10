package main

import (
	"fmt"

	"github.com/wapalencia/algoritmo/internal"
	
)

func main() {

	var values = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var LeerArray = []string{"1", "m", "u", "y", "-", "f", "a", "c", "i", "l", "9", "0"}
	internal.ValidarPar(values)
	fmt.Println(internal.RecorrerPares(-5, 6, 2))
	fmt.Println(internal.EncontrarPalabra(LeerArray))
	fmt.Println(internal.EncontrarPalabra1(LeerArray))


}
