package internal

import (

	//"fmt"
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
			if !isCharcad(c, x, xletrep) {
				cont++

				c2 = append(c2, GetChar(nList, i-1))

			}
		}

	}

	return c2
}
func isCharcad(xr []string, srt string, letrep string) bool {
	numList := xr

	for _, x := range numList {

		if x == srt {
			return true
		}

	}

	return false
}
func IntArrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
