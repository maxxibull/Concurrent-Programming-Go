package main

import (
	"math/rand"
	"time"
)

func generateRandom() int {
	return rand.Intn(maxValueOfArgument-minValueOfArgument) + minValueOfArgument
}

func chief(addTaskToBoardChannel chan<- task) {
	rand.Seed(time.Now().Unix())

	for {
		newTask := task{
			firstArgument:   generateRandom(),
			secondArgument:  generateRandom(),
			operationSymbol: operations[rand.Intn(len(operations))]}

		if isChattyMode {
			printChief(newTask)
		}

		addTaskToBoardChannel <- newTask
		time.Sleep(time.Millisecond * time.Duration(timeToSleepForChief))
	}
}
