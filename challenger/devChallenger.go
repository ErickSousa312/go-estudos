//go:build !windows
// +build !windows

// go env GOOS GOARCH -- para verificar la plataforma e arquitectura de Go.

package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// words2 := loadWords("repo.json") // pode ser um vetor ou requi
	words := []string{
		"gato",
		"cachorro",
		"elefante",
		"leão",
		"tigre",
		"panda",
		"gorila",
		"girafa",
		"hipopótamo",
		"rinoceronte",
		"zebra",
		"macaco",
		"cobra",
		"jacaré",
		"pássaro",
		"urso",
		"leopardo",
		"lobo",
		"coelho",
		"peixe",
		"camelo",
		"rinoceronte",
		"orangotango",
		"crocodilo",
		"baleia",
		"pinguim",
		"águia",
		"coruja",
		"búfalo",
		"caranguejo",
	}
	word := getRandomWord(words)
	guessed := strings.Repeat("_ ", len(word))
	guesses := make(map[rune]bool)
	maxAttempts := 6
	attempts := 0

	fmt.Println("Bem-vindo ao Jogo da Forca!")

	for {
		fmt.Println("\nPalavra:", guessed)
		fmt.Printf("Tentativas restantes: %d\n", maxAttempts-attempts)

		if strings.Index(guessed, "_") == -1 {
			fmt.Println("Parabéns! Você ganhou!")
			break
		}

		if attempts >= maxAttempts {
			fmt.Println("Você perdeu! A palavra era:", word)

			FakeReader := bufio.NewReader(os.Stdin)
			FakeReader.Read(make([]byte, 3))
			break
		}

		fmt.Print("Digite uma letra ou a palavra inteira: ")
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			clearScreen()
			continue
		}

		if string(input) == word {
			fmt.Println("Parabéns, Você é o brabo!")
			FakeReader := bufio.NewReader(os.Stdin)
			FakeReader.Read(make([]byte, 3))
			break
		}

		if len(input) == 1 {
			guess := rune(input[0])
			if guesses[guess] {
				fmt.Println("Você já tentou essa letra, tente outra amigo")
				clearScreen()
				continue
			}

			guesses[guess] = true

			if strings.ContainsRune(word, guess) {
				for i, char := range word {
					if char == guess {
						guessed = guessed[:i*2] + string(char) + guessed[i*2+1:]
						clearScreen()
					}
				}
			} else {
				fmt.Println("Essa letra não existe nessa palavra.")
				clearScreen()
				attempts++
			}
		} else {
			fmt.Println("Palavra incorreta.")
			clearScreen()
			attempts++
		}
	}
}

//usar com json
// func loadWords(filename string) []string {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("Erro ao abrir arquivo:", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	var words []string
// 	err = json.NewDecoder(file).Decode(&words)
// 	if err != nil {
// 		fmt.Println("Erro ao ler arquivo JSON:", err)
// 		os.Exit(1)
// 	}
// 	return words
// }

func getRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "r/documents")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
