package main

import (
	"fmt"
	"sync"
)

func Square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num*num)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup


	for _, num := range arr {
		wg.Add(1)
		go Square(num, &wg)	
	}
	
	wg.Wait()
}