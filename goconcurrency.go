package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done() //we use defer which runs a code line only and as soon as a function ends. 
	//we can use defer to code recover from panic
	for i:=0; i < 3; i++ {
		time.Sleep(100*time.Millisecond)
		if i == 2{
			panic("Panicker")
		}
		fmt.Println(s)
	}	
	//wg.Done() // what if the code throws an exception above? we will have to wait forever
}

func main121() {
	go say("Hey")
	wg.Add(1)
	go say("there")
	wg.Add(1)
	wg.Wait() //wait till weight-group's count becomes 0 again
}