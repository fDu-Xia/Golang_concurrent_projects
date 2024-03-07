package one2one

import (
	"go_code/producer-consumer/out"
	"strconv"
	"sync"
)

var productionChan chan int = make(chan int)
var OutputBuffer chan string = make(chan string)
var Wg sync.WaitGroup

const taskNum int = 10

func producer() {
	defer Wg.Done()

	for i := 1; i <= taskNum; i++ {
		OutputBuffer <- "第" + strconv.Itoa(i) + "号产品生产完毕"
		productionChan <- i
	}
	close(productionChan)
}

func consumer() {
	defer Wg.Done()

	for i := range productionChan {
		OutputBuffer <- "取走第" + strconv.Itoa(i) + "号产品"
	}
	//close(out.OutputBuffer)
}

func Execute() {
	Wg.Add(2)
	go producer()
	go consumer()
	go out.Out(OutputBuffer)
	Wg.Wait()
}
