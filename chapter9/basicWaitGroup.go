package main

import (
	"sync"
	"time"
)

func sleep5Seconds(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 5)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sleep5Seconds(&wg)
	wg.Wait()
}
