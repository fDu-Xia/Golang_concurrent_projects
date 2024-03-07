package many2one

import (
	"go_code/producer-consumer/out"
	"strconv"
	"sync"
)

var productionChan chan int = make(chan int)
var OutputBuffer chan string = make(chan string)
var producerWg sync.WaitGroup
var consumerWg sync.WaitGroup
var mutex sync.Mutex

const taskNum int = 100
const producerNum int = 5

var productionId int = 1

func produce(id int) {
	defer producerWg.Done()

	for {
		mutex.Lock()
		if productionId > taskNum {
			mutex.Unlock()
			break
		}
		OutputBuffer <- strconv.Itoa(id) + "号生产者：" + "第" + strconv.Itoa(productionId) + "号产品生产完毕"
		productionChan <- productionId
		productionId++
		mutex.Unlock()
	}
}

func producers() {
	defer close(productionChan)
	for i := 1; i <= producerNum; i++ {
		producerWg.Add(1)
		go produce(i)
	}
	producerWg.Wait()
}
func consumer() {
	defer consumerWg.Done()

	for i := range productionChan {
		OutputBuffer <- "取走第" + strconv.Itoa(i) + "号产品"
	}
	//close(out.OutputBuffer)
}

func Execute() {
	go out.Out(OutputBuffer)
	consumerWg.Add(1)
	go consumer()
	producers()
	consumerWg.Wait()
}
