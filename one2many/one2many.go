package one2many

import (
	"go_code/producer-consumer/out"
	"strconv"
	"sync"
)

var productionChan chan int = make(chan int)
var Wg sync.WaitGroup
var OutputBuffer chan string = make(chan string)

const taskNum int = 100
const consumerNum int = 3

func consume(id int) {
	defer Wg.Done()
	for i := range productionChan {
		OutputBuffer <- strconv.Itoa(id) + "号取走第" + strconv.Itoa(i) + "号产品"
	}
}

func producer() {
	defer Wg.Done()

	for i := 1; i <= taskNum; i++ {
		OutputBuffer <- "第" + strconv.Itoa(i) + "号产品生产完毕"
		productionChan <- i
	}
	close(productionChan)
}

func consumers() {
	for i := 1; i <= consumerNum; i++ {
		go consume(i)
	}
}

func Execute() {
	Wg.Add(consumerNum + 1)
	go producer()
	consumers()
	go out.Out(OutputBuffer)
	Wg.Wait()
}
