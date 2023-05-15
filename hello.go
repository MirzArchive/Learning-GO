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

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	timeSeed := time.Now().Unix()
	rand.Seed(timeSeed)
	randNum := rand.Intn(50) + 1
	r := bufio.NewReader(os.Stdin)

	for {
		guess, _ := r.ReadString('\n')
		trimmedGuess := strings.TrimSpace(guess)
		convGuess, _ := strconv.Atoi(trimmedGuess)

		pl(randNum)
		if convGuess < randNum {
			pl("Guess Higher")
		} else if convGuess > randNum {
			pl("Guess Lower")
		} else {
			pl("You are correct!")
			break
		}

	}

}
