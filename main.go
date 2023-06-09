package main

import (
	"Validacao-e-contagem-de-numeros-de-telefone/validation"
	"bufio"
	"fmt"
	"os"
)

func main() {
	phones := []string{}
	count := 0
	file, err := os.Open("../data/phones.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phones[count] = scanner.Text()
		count++
	}
	resp := validation.Validation(phones)
	fmt.Println(resp)
}
