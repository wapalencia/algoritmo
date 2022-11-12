package internal

import (
	"fmt"
	//"strings"
)

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

func ReadSTR(str []string) {
	numList := make([]string, 1)
	numList2 := make([][]string, 1)
	var xred string

	i := 0

	for _, x := range str {

		//fmt.Println(x)

		/* 		if i > 0 || x == numList[i] {

		   			numList[i] = xred
		   			//fmt.Println(numList[i] )
		   		}
		   		xred = "" */

		numList[i] = x
		fmt.Println(numList[i])

		for _, leninterno := range str {
			xred += x
			if x == leninterno {
				numList = append(numList, xred)
			}

		}

		numList2 = append(numList2, numList)

	}
	//fmt.Println(xred)
	fmt.Println(numList2)

}
