package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess:")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)

		for guess := 10; guess == 0; guess-- {
		}
		if guess < target {
			fmt.Println("Меньше загаданного числа")
		}
		if guess > target {
			fmt.Println("Больше загаданного числа")
		}
		if guess == target {
			fmt.Println("Вы угадали !")
		}
	}
}
