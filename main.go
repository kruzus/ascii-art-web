package main

import (
	"bufio"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))
var templatesDir = os.Getenv("TEMPLATES_DIR")

func main() {
	fmt.Println("==>  * localhost:9000 *  <==")
	openbrowser("http://localhost:9000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/Image", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		img := r.Form["machaine"][0]
		if r.Form["machaine"][0] == "rickroll.png" {
			openbrowser("https://www.youtube.com/watch?v=xvFZjo5PgG0&ab_channel=RickRoll")
		}
		//	fmt.Fprintln("<script>alert('you have been pwned')</script>")
		//	//	fmt.Fprintf(w, DrawAscii("Capture.jpg"))

		fmt.Fprintln(w, DrawAscii(img))
		//fmt.Println(DrawAscii(img))
		//	fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/Ascii", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ascii := r.Form["machaine"][0]
		fmt.Fprintln(w, Ascii(ascii))
	})
	log.Fatal(http.ListenAndServe(":9000", nil))

}
func DrawAscii(v string) string {
	base := "0ND8OZ$7I?+=~:,.."
	//base := "$"
	f, _ := os.Open(v)
	img, _ := png.Decode(f)
	bounds := img.Bounds()
	ascii := ""
	for y := 0; y < bounds.Dy(); y += 2 {
		for x := 0; x < bounds.Dx(); x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			r = r & 0xFF
			g = g & 0xFF
			b = b & 0xFF
			gray := 0.299*float64(r) + 0.578*float64(g) + 0.114*float64(b)
			temp := fmt.Sprintf("%.0f", gray*float64(len(base)+1)/255)
			index, _ := strconv.Atoi(temp)
			if index >= len(base) {
				ascii += "."
				ascii += ""
				//fmt.Print(" ")
			} else {
				ascii += string(base[index])
				//fmt.Print(string(base[index]))
			}

		}

		ascii += "<br>\n"

	}
	f.Close()

	//	fmt.Println(ascii)
	return ascii
}
func openbrowser(zz string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", zz).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}

//Ascii pokok
func Ascii(str string) string {

	file, _ := os.Open("standard.txt")
	fileVal := scanFile(file)

	k := ""

	for i := 1; i <= 8; i++ {

		for _, arg := range str {
			indexBase := int(rune(arg)-32) * 9
			if indexBase > 856 {
				k = "index out of range"
				return k
			} else if indexBase < 856 {

				k += fileVal[indexBase+i]
				k = strings.Replace(k, " ", "&nbsp;", -1)
			}

		}
		k += "<br>"

	}

	//fmt.Print(k)
	return k

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
