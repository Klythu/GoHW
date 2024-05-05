package main

import (
	"fmt"
	"sync"
)

func initc(wg *sync.WaitGroup, num int) chan int {
	intChanf := make(chan int)
	go func() {
		intChanf <- num
		wg.Done()
	}()
	return intChanf
}

func kvad(wg *sync.WaitGroup, intChan chan int) chan int {
	intChanf := make(chan int, 2)
	input := <-intChan
	num := input * input
	go func() {
		intChanf <- input
		intChanf <- num
		wg.Done()
	}()

	wg.Done()
	return intChanf

}
func mult(wg *sync.WaitGroup, intChan2 chan int) chan int {
	intChanf := make(chan int, 3)
	num := <-intChan2
	input := <-intChan2
	mult := num * input
	go func() {
		intChanf <- num
		intChanf <- input
		intChanf <- mult
		wg.Done()
	}()

	return intChanf
}
func main() {
	input := 0
	var wg sync.WaitGroup
	for {
		wg.Add(4)
		fmt.Scanf("%d\n", &input)
		intChan1 := initc(&wg, input)
		intChan2 := kvad(&wg, intChan1)
		intChan3 := mult(&wg, intChan2)
		wg.Wait()
		fmt.Println(<-intChan3, <-intChan3, <-intChan3)
	}

}
