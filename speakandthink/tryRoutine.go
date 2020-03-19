package main

import (
	"fmt"
	"sync"
	"time"
)

func add(x, y int, exitchan chan bool) {
	z := x + y
	if x == 6 {
		time.Sleep(time.Second * 5)
	}
	fmt.Printf("%d + %d =  %d\n", x, y, z)
	exitchan <- true
}

func addele(a []int, i int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Add ele fail")
		}
	}()
	a[i] = i
	fmt.Println(a)

}

func productor(mychan chan int, data int, wait *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wait.Done()
}

func consumer(mychan chan int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：", a)
	wait.Done()
}

func send(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
		fmt.Println("send data : ", i)
	}
}

func main() {
	resch := make(chan int, 20)
	strch := make(chan string, 10)
	go send(resch)
	//time.Sleep(time.Second * 2)
	//strch <- "wd"
	select {
	case a := <-resch:
		fmt.Println("get data : ", a)
	case b := <-strch:
		fmt.Println("get data : ", b)
	default:
		fmt.Println("no channel actvie")

	}

}
