package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	words := strings.Fields(text)
	wordCount := len(words)
	charCount := len(text)

	fmt.Println("--- Анализатор текста ---")
	fmt.Printf("Исходный текст: %s\n", text)
	fmt.Printf("Количество слов: %d\n", wordCount)
	fmt.Printf("Количество символов: %d\n", charCount)
}
