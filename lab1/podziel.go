package main

import(
	
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

)

func main(){

	wordList := Read("slowa.txt")
	wordSet := map[string]bool{}
	for _, w := range wordList {

		wordSet[w] = true

	}
	wordCounter := map[string]int{}
	for _, w := range wordList{
		for _, parts := range Split(w){
			AddIfIn(parts, wordSet, &wordCounter)
		}

	}

	pairs := Sort(wordCounter)
	for _, p := range pairs{
		fmt.Printf("%d %s\n", p.number, p.word)
	}		
}

// Read wczytuje wyrazy z pliku o nazwie 'filename' i zwraca je w
// wycinku tablicy z łańcuchów
func Read(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// Split dzieli wyraz 'word' na dwa niepuste łańcuchy na wszystkie
// możliwe sposoby i zwraca te dwa łańcuchy w wycinku tablicy wycinków
// tablic łańcuchów
func Split(word string) [][]string{

	r := [][]string{}
	for i := 1; i < len(word); i++ {

		r = append(r, []string{word[:i], word[i:]})

	}
	return r

}

// AddIfIn zwiększa licznik 'counter' przy tych łańcuchach z wycinka
// 'parts', które występują w zbiorze łańcuchów 'set'.
func AddIfIn(parts []string, set map[string]bool, counter *map[string]int) {

	if set[parts[0]] && set[parts[1]] {
		(*counter)[parts[0]]++
		(*counter)[parts[1]]++
	}

}

type Pair struct{
	word string
	number int
}

// Sort sortuje pary w malejącej kolejnośći pół 'number'. Takie pary,
// które mają jednakową wartość pól 'counter', sortuje w kolejności
// leksykograficznej pól 'word'.
func Sort(counter map[string]int) []Pair{

	r:= []Pair{}
	for w, c := range counter{
		r = append(r, Pair{w, c})
	}
	slices.SortFunc(r, func(a, b Pair) int {
		if n := b.number - a.number; n != 0 {
			return n
		}
		return cmp.Compare(a.word, b.word)
	})
	return r

}