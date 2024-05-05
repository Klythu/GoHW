package main

import (
	"fmt"
	"sync"
)

func gen(wg *sync.WaitGroup, intChan chan int) {
	i := 0
	for {
		wg.Add(1)
		i++
		intChan <- i
		wg.Done()
	}
}
func stop(wg *sync.WaitGroup, stop chan int) {
	input := ""
	fmt.Scanf("%s\n", &input)
	stop <- 1
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	stopChan := make(chan int, 1)
	intChan := make(chan int)
	i := 0
	go stop(&wg, stopChan)
	go gen(&wg, intChan)
	for {
		select {
		case <-intChan:
			i++
			fmt.Println(i * i)
		case <-stopChan:
			wg.Wait()
			a := <-intChan
			a = a * 0
			fmt.Println("unicorn used his power to stop this madness")
			break
		}
	}

}
