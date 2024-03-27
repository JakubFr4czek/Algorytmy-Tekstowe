package main

import (
	"fmt"
	"regexp"
)

var Word = regexp.MustCompile(`\pL+`) //kompiluje regexpa

var IP = regexp.MustCompile(
	"^(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])[.]" +
		"(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])[.]" +
		"(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])[.]" +
		"(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$")

var Animal = regexp.MustCompile("([Kk]ot|[Pp]ies)")

var Comma = regexp.MustCompile(`, +\pL+[^iy]( |$)`)

func main() {

	text := "Wydział Wiertnictwa, Nafty i Gazu"
	fmt.Printf("%#v", Word.FindAllString(text, -1)) //zwraca wszystkie stringi pasujace do Word dla -1 lub n dla n > 0

	fmt.Printf("\n")

	fmt.Printf("%#v", IP.FindStringSubmatch("127.0.0.1"))

	fmt.Printf("\n")

	text2 := "Kot i pies leżą na kanapie"
	fmt.Printf("%#v", Animal.ReplaceAllString(text2, "${1}ek")) //podmienia pasujace fragmenty stringa na podany tekst, lub ${1}ek, to znaczy ze za ${1} to bedzie ten tekst

	fmt.Printf("\n")

	text3 := "Wydział Informatyki, Elektroniki " + "i Telekomunikacji, Biuro Dziekana"
	fmt.Printf("%#v", Comma.Split(text3, -1)[0]) // Split dzieli łańcuch s na części. Te części są rozdzielone takimi podłańcuchami,które pasują do regexpa

	fmt.Printf("\n")

}
