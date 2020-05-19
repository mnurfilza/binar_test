package module

import (
	"fmt"
	"math/rand"
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

func getDice() int {
	wait := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)

		wait <- true
	}()

	out := <-wait
	for {
		dice1 := rollDice()
		fmt.Println("Dice 1: ", dice1)
		if out == true {
			return dice1
		}
	}
}

func PlaYGame(member string) {
	b := &Board{}
	c1 := make(chan map[string]*Board)
	status := make(chan string)
	go b.Login(c1, member, status)

}

type Board struct {
	Score int
}

func (b *Board) Login(c1 chan map[string]*Board, name string, status chan string) {
	// fmt.Println(name)
	m := make(map[string]*Board)
	if b != nil {
		m[name] = &Board{Score: 0}
	}
	if len(m) == 5 {
		status <- "Play"
	}
	c1 <- m
}
