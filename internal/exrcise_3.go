package internal

import (

	//"fmt"
	"fmt"
	"strings"
)

//"fmt"

func EncontrarPalabra(data []string) string {

	response := ""
	for _, item := range data {
		if !isNumber(item) {
			response += item
		}
	}
	return response
}

func isNumber(num string) bool {
	numList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for _, x := range numList {
		if x == num {
			return true
		}
	}
	return false
}

func IsPalindrome(w string) bool {

	i := 0
	for _, item := range w {
		i++
		rev := item
		ini := w[len(w)-i]
		if string(rev) != string(ini) {
			println("el for", string(rev), "vs cadena", string(ini))
			return false
		}
	}
	return true
}

func GetChar(str []string, iterado int) []string {
	var xcad string
	var result []string

	c := make([]string, 0)
	numList := str
	//l:=len(numList[len(str)])

	for _, x := range numList {

		if x == numList[iterado] {
			xcad += x
		}
	}
	//fmt.Println(len(xcad))
	if len(xcad) > 1 {
		c = append(c, xcad)
	}

	result = c

	return result

}

func ReadArry(r []string) [][]string {
	var xletrep string
	cont := 0
	nList := r
	c := make([]string, 0)
	c2 := make([][]string, 0)
	//fmt.Println(c)
	i := 0
	for _, x := range nList {

		i++
		if !strings.Contains(xletrep, x) {
			xletrep += x
			if !isCharcad(c, x) {
				cont++

				c2 = append(c2, GetChar(nList, i-1))

			}
		}

	}

	return c2
}
func isCharcad(xr []string, srt string) bool {
	numList := xr

	for _, x := range numList {

		if x == srt {
			return true
		}

	}

	return false
}
func CreateSeed(seed []string, peso []string) int {

	var resultPeso string
	var resultSeed string
	var posi int
	var countSeed int
	var totalSeed int

	for _, s := range seed {

		if s == "a" {
			resultSeed += s
			fmt.Println("Semillas ...", s)
		}
	}

	fmt.Println("total Semillas ...En SEMILLAS", len(resultSeed))
	countSeed += len(resultSeed)
	i := -1
	for _, p := range peso {
		i++
		fmt.Println("peso ...", i, "posicion", p)
		if p == "a" {
			resultPeso += p

			posi = i

		}
	}
	///SABER QUE METODO USAR
	if len(resultPeso) != 0 {

		switch resultSeed {
		case "":
			fmt.Println("resultado semilla ", resultSeed, "  is my favorite!")

			countSeed = 0
			totalSeed = countSeed

		case "a":
			fmt.Println("resultado semilla ", resultSeed, "  is my favorite!")

			countSeed = 1
			totalSeed = (countSeed) + (countSeed)
		case "aa":

			fmt.Println("resultado semilla ", resultSeed, "is great!")

			countSeed = len(resultPeso) + len(resultSeed)
			totalSeed = (posi + countSeed + posi) + (posi * countSeed * posi)

		default:
			fmt.Println("I've never tried", resultSeed, "before")
			countSeed = len(resultPeso) + len(resultSeed)
			totalSeed = (posi + countSeed + posi) + (posi * countSeed * posi)
		}

	} else {
		fmt.Println("total   resultPeso ", resultPeso)
		countSeed = len(resultPeso) + len(resultSeed)
		totalSeed = countSeed
	}

	fmt.Println("total Semillas ...En pESO", len(resultPeso))
	fmt.Println("Posicion PESO ...", posi)
	fmt.Println("toal semilla y peso", countSeed)

	fmt.Println("(", posi, "+", countSeed, "+", posi, ") + (", posi, "*", countSeed, "*", posi, ")")

	return totalSeed
}
