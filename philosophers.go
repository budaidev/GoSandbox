package main

import (
	"fmt"
	"strconv"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	name            string
}

func (p Philo) eat() {
	fmt.Println("starting to eat", p.name)
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println(p.name, "eating", i+1, "times")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	fmt.Println("finishing eating", p.name)
}

func master(queue chan int, p []*Philo, done chan bool) {
	for true {
		select {
		case k := <-queue:
			p[k].eat()
			done <- true
		}
	}
}

func main() {

	maxPhilosphers := 2
	numberOfPhilosophers := 5

	Csticks := make([]*ChopS, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5], strconv.Itoa(i + 1)}
	}

	q := make(chan int)
	done := make(chan bool)

	//add a master queue to limit the number of workers
	for i := 0; i < maxPhilosphers; i++ {
		go master(q, philos, done)
	}

	//start philosopher eats
	for j := 0; j < numberOfPhilosophers; j++ {
		go func(j int) {
			q <- j
		}(j)
	}

	// wait for all jobs to finish
	for c := 0; c < numberOfPhilosophers; c++ {
		<-done
	}
}
