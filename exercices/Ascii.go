package student

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ascii(str string) string {

	file, _ := os.Open("standard.txt")
	fileVal := scanFile(file)

	narg := strings.Split(str, "\\n")

	for _, v := range narg {

		printLetter(v, fileVal)

	}
	//fmt.Println("tes")
	//	return printLetter(str, fileVal)
	return "teést"
}

func printLetter(s string, fileVal []string) string {
	rep := ""
	i := 1
	err := 0
	verif := ""
	for i <= 8 {
		for _, arg := range s {

			indexBase := int(rune(arg)-32) * 9

			if indexBase <= 855 {
				verif += (fileVal[indexBase+i])
			} else {
				i = 8
				err = 1
				break
			}

		}
		i++
		if err == 0 {
			fmt.Println(verif)
			rep += verif
			verif = ""
			//return verif
		} else {
			fmt.Println("Error")
		}

	}
	//fmt.Println(rep)
	//	rep = ""
	return "éééé"
}

func scanFile(arg *os.File) []string {

	var fileValue []string

	scanner := bufio.NewScanner(arg)

	for scanner.Scan() {

		lines := scanner.Text()

		fileValue = append(fileValue, lines)

	}

	return fileValue
}
