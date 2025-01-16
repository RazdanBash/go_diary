package main

import (
	"bufio"
	"fmt"
	"os"
	"strings" // Добавляем пакет strings для обрезки символов
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Читаем строку
	input = strings.TrimSpace(input)    // Убираем лишние пробелы и символы новой строки
	fmt.Println("You entered:", input)  // Выводим результат


	70
}