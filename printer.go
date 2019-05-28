package main

import (
	"fmt"
	"os"
	"time"
)

func printAvailableModesAndChooseOne() {
	var mode int
	fmt.Println("Hello! Which mode do you choose?")
	fmt.Println("\u001b[31m[1]\u001b[0m Chatty mode")
	fmt.Println("\u001b[31m[2]\u001b[0m Normal mode")
	fmt.Println("================")
	fmt.Print("Your choice: \u001b[31m")
	fmt.Scan(&mode)
	fmt.Println("\u001b[0m================")

	if mode == 1 {
		isChattyMode = true
	} else if mode != 2 {
		fmt.Println("Wrong input!")
		os.Exit(0)
	}
}

func printAvailableOptionsAndChooseOne(printTasksFromBoardChannel chan bool, printProductsFromStoreChannel chan bool) {
	for {
		var mode int
		time.Sleep(time.Second)
		fmt.Println("What do you want to check?")
		fmt.Println("\u001b[33m[0]\u001b[0m Exit")
		fmt.Println("\u001b[33m[1]\u001b[0m Tasks Board")
		fmt.Println("\u001b[33m[2]\u001b[0m Store Board")
		fmt.Println("\u001b[33m[3]\u001b[0m Workers")
		fmt.Println("================")
		fmt.Print("Your choice: \u001b[33m")
		fmt.Scan(&mode)
		fmt.Println("\u001b[0m================")

		switch mode {
		case 0:
			fmt.Println("Bye bye!")
			os.Exit(0)
		case 1:
			printTasksFromBoardChannel <- true
		case 2:
			printProductsFromStoreChannel <- true
		case 3:
			fmt.Println("\u001b[32mWorkers\u001b[0m")
			mutex.Lock()
			fmt.Println(workersInfo)
			mutex.Unlock()
			fmt.Println("================")
		}
	}
}

func printChief(newTask task) {
	fmt.Println("\u001b[33mchief\u001b[0m", newTask.firstArgument, newTask.operationSymbol, newTask.secondArgument)
}

func printCustomer(index int, newProduct product) {
	fmt.Println("\u001b[36mcustomer", index, ":\u001b[0m", newProduct.value)
}

func printMachineStart(machineIndex int, workerIndex int, newTask task) {
	fmt.Println("\u001b[31m[", newTask.operationSymbol, "] machine", machineIndex, ":\u001b[0m", newTask.firstArgument,
		newTask.operationSymbol, newTask.secondArgument, "[ from worker", workerIndex, "]")
}

func printMachineFinish(machineIndex int, workerIndex int, newTask task) {
	fmt.Println("\u001b[31m[", newTask.operationSymbol, "] machine", machineIndex, ":\u001b[0m", newTask.firstArgument,
		newTask.operationSymbol, newTask.secondArgument, "=", newTask.result, "[ for worker", workerIndex, "]")
}

func printMachineIsBroken(index int, op string) {
	fmt.Println("\u001b[31m[", op, "] machine", index, ":\u001b[0m BROKEN")
}

func printMachineIsOK(index int, op string) {
	fmt.Println("\u001b[31m[", op, "] machine", index, ":\u001b[0m FIXED")
}

func printStore(board []product) {
	fmt.Println("\u001b[32mStore Board\u001b[0m")
	fmt.Println(board)
	fmt.Println("================")
}

func printTasksBoard(board []task) {
	fmt.Println("\u001b[32mTasks Board\u001b[0m")
	fmt.Println(board)
	fmt.Println("================")
}

func printWorkerStart(index int, newTask task) {
	fmt.Println("\u001b[32mworker", index, ":\u001b[0m", newTask.firstArgument,
		newTask.operationSymbol, newTask.secondArgument)
}

func printWorkerFinish(index int, newTask task, newProduct product) {
	fmt.Println("\u001b[32mworker", index, ":\u001b[0m", newTask.firstArgument,
		newTask.operationSymbol, newTask.secondArgument, "=", newProduct.value)
}

func printWorkerChoseMachine(workerIndex int, machineIndex int, typeOfWorker string, newTask task) {
	fmt.Println("\u001b[32mworker", workerIndex, typeOfWorker, ":\u001b[0m [", newTask.operationSymbol, "] machine",
		machineIndex, "has been chosen")
}

func printWorkerAbandonedMachine(workerIndex int, machineIndex int, newTask task) {
	fmt.Println("\u001b[32mworker", workerIndex, ":\u001b[0m [", newTask.operationSymbol, "] machine",
		machineIndex, "was abandoned")
}

func printWorkerFoundBrokenMachine(workerIndex int, machineIndex int, newTask task) {
	fmt.Println("\u001b[32mworker", workerIndex, ":\u001b[0m [", newTask.operationSymbol, "] machine",
		machineIndex, "is broken")
}

func printServiceCenterWorkerIsGoingToFixMachine(index int, newCrashReport crashReport) {
	fmt.Println("\u001b[34mservice center worker", index, ":\u001b[0m going to fix [", newCrashReport.machineType, "] machine",
		newCrashReport.machineIndex)
}

func printServiceCenterWorkerFixedMachine(index int, newCrashReport crashReport) {
	fmt.Println("\u001b[34mservice center worker", index, ":\u001b[0m [", newCrashReport.machineType, "] machine",
		newCrashReport.machineIndex, "was fixed")
}
