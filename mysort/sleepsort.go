package mysort

import (
	"time"
)

func SleepSort(inputSli []int) []int {
	var length = len(inputSli)
	var outputChan = make(chan int, length)
	var retSli = make([]int, 0)

	for _, v := range inputSli {
		go sleep(v, outputChan)
	}

	for {
		if len(retSli) == length {
			return retSli
		}
		select {
		case num := <-outputChan:
			retSli = append(retSli, num)
		}
	}
	return nil
}

func sleep(inputNum int, outputChan chan int) {
	time.Sleep(time.Millisecond * time.Duration(inputNum*10))

	outputChan <- inputNum
}
