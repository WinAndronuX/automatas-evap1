package main

import (
	"bufio"
	"evap1/internal/aut"
	"fmt"
	"log"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {

	fmt.Print(prompt)

	line, err := stdin.ReadString('\n')
	if err != nil && line == "" {
		log.Fatal(err)
	}

	return strings.TrimSpace(line)
}

func main() {

	name := readLine("Ingrese su nombre completo: ")
	id := readLine("Ingrese su matricula: ")

	var automata = aut.New(name, id)

	for {
		text := readLine("\nIngrese la cadena a validar: ")

		if aut.Eval(automata, text) {
			fmt.Println("\n-> Valid Expression.")
		} else {
			fmt.Println("\n-> Invalid Expression.")
		}

		if strings.EqualFold(readLine("\nContinue? y/N: "), "n") {
			break
		}
	}
}
