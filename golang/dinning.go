package main

import (
	"fmt"
	"sync"
)

type chop struct{ sync.Mutex }

type philo struct {
	num, eaten  int
	left, right *chop
}

var host = make(chan struct{}, 2)

func (p *philo) eat(wg *sync.WaitGroup) string {
	defer wg.Done()

	if p.eaten == 3 {
		return "Eaten 3 times"
	}

	host <- struct{}{}

	p.left.Lock()
	p.right.Lock()

	fmt.Println("starting to eat:", p.num)
	fmt.Println("finishing to eat:", p.num)

	p.eaten++

	p.right.Unlock()
	p.left.Unlock()

	<-host

	return "ate"
}

func main() {
	var wg sync.WaitGroup

	chops := make([]*chop, 5)
	for i := 0; i < 5; i++ {
		chops[i] = &chop{}
	}

	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philo{
			num:   i + 1,
			left:  chops[i],
			right: chops[(i+1)%5],
		}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg)
	}

	wg.Wait()
	fmt.Println("All philosophers ate 3 times")
}
