package goroutine

import (
	"fmt"
	"sync"
)

func PrintNtimes(input string, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		fmt.Println(input)
	}
	wg.Done()
}

func MutexConcept(wg *sync.WaitGroup, mutex *sync.Mutex, counter *int) {
	mutex.Lock()
	// LIFO concept of stack is used
	defer wg.Done()
	defer mutex.Unlock()
	*counter++
	fmt.Println("value := counter", *counter)
}
