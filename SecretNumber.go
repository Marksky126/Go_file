package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret Number is", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil { //err为nil 的时候才是没有任何错误发生
			fmt.Println("An error occured while reading input.Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input,Please enter an integer value")
			continue
		}
		fmt.Println("Your guess is ", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is bigger than secretNumber")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than secretNumber")
		} else {
			fmt.Println("Congratulations! It's")
			break
		}
	}
}
