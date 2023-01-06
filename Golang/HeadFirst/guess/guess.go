package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//генератор рандомных чисел
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 //загадонное число
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	//игрок вводит свое число
	reader := bufio.NewReader(os.Stdin) //ввод с клавиатуры

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // удалили энтер
		guess, err := strconv.Atoi(input) //перевели в число
		if err != nil {
			log.Fatal(err)
		}

		//проверка угадал или нет
		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
