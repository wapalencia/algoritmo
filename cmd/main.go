package main

import (
	//"fmt"
	//"strings"

	"fmt"

	"github.com/wapalencia/algoritmo/internal"
)

func main() {

	var values = []string{"a", "a", "c", "c", "s", "x", "l", "x", "L"}
	var values2 = []string{" ", " ", " ", "c", "s", "", "", "x", "L"}
	var values3 = []string{"1", "1", "1", "c", "s", "", "1", "x", "L"}

	/*var LeerArray = []string{"1", "m", "u", "y", "-", "f", "a", "c", "i", "l", "9", "0"}
	internal.ValidarPar(values)
	fmt.Println(internal.RecorrerPares(-5, 6, 2))
	fmt.Println(internal.EncontrarPalabra(LeerArray))
	fmt.Println(internal.IsPalindrome("anar"))

	ss := "hola"
	fmt.Println(strings.Contains(ss, "o"))
	fmt.Println(strings.Count(ss, "o"))*/
	
	fmt.Println(internal.IsPalindrome("ana"))
	fmt.Println(internal.ReadArry(values))
	fmt.Println(internal.ReadArry(values2))
	fmt.Println(internal.ReadArry(values3))

}
