package documents
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Ошибка: передайте ровно один аргумент (путь к файлу)")
		os.Exit(1)
	}

	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}

	text := string(data)

	charCount := utf8.RuneCountInString(text)

	lineCount := strings.Count(text, "\n") + 1

	fmt.Printf("--- Анализ файла: %s ---\n", filePath)
	fmt.Printf("Количество символов: %d\n", charCount)
	fmt.Printf("Количество строк: %d\n", lineCount)
}
