package main

import (
	"sync"
	"testing/dsa/goroutine"
)

// var wg sync.WaitGroup
var Counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine.PrintNtimes("adarsh", 1, &wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goroutine.MutexConcept(&wg, &mutex, &Counter)
	}

	wg.Wait()
	/*var myName1, myName2, myName3 string
	fmt.Println("Enter 3 names..")
	fmt.Scan(&myName1, &myName2, &myName3)
	fmt.Print(myName1, myName2, myName3)
	myName2in, _ := fmt.Scanf("%%s, %s, %s", &myName1, &myName2, &myName3)
	fmt.Printf("%s, %s, %s, %d", myName1, myName2, myName3, myName2in)
	fmt.Scanln(&myName1, &myName2, &myName3)
	fmt.Println(myName1, myName2, myName3)
	fmt.Println(test.HelloWorld("aadi"))
	*/

}
