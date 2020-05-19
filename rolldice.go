package main

import (
	"fmt"
	"math/rand"
	"project/binar/module"
	"sync"
	"time"
)

var onlyOnce sync.Once

// prepare the dice
var dice = []int{1, 2, 3, 4, 5, 6}

func rollDice() int {

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func main() {
	var s string
	fmt.Print("Input nama anda :")
	fmt.Scanln(&s)
	b := &module.Board{}
	c1 := make(chan map[string]*module.Board)
	status := make(chan string)
	for {
		go b.Login(c1, s, status)
		check := <-status
		if check != "Play" {
			fmt.Println(check)
			break
		}

	}
	data := <-c1
	fmt.Println(data)

}
