package main

import "fmt"
import "sync"

func routin1(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	for i := 1; i <= 26; i++ {
		<-ch1
		fmt.Print(i)
		ch2 <- struct{}{}
	}

	wg.Done()
}

func routin2(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 26; i++ {
		<-ch2
		fmt.Print(string('A' + i))
		if i < 25 {
			ch1 <- struct{}{}
		}
	}

	wg.Done()
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		routin1(ch1, ch2, &wg)
	}()

	go func() {
		routin2(ch1, ch2, &wg)
	}()

	ch1 <- struct{}{} // Start the first goroutine

	wg.Wait() // Wait for both goroutines to finish
	close(ch1)
	close(ch2)
	fmt.Println("    All done!")
}