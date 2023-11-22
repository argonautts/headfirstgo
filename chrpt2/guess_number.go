// guess_number - игра, в которой игрок должен угадать случайное число.
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

func inputNumber(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	Handle(err)

	return input
}

func inputCorrect(reader *bufio.Reader) int {
	input := inputNumber(reader)
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input) // convert to int
	Handle(err)

	if number > 100 || number < 1 {
		log.Fatal("The entered number does not match the range")
	}

	return number
}

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var success bool
	target := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= 10; i++ {
		fmt.Print(i, ". Enter number (1-100): ")

		number := inputCorrect(reader)

		if target > number {
			fmt.Println("Oops. Your guess was LOW»")
		} else if target < number {
			fmt.Println("«Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Congratulations! It was: ", target)
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
