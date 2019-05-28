package main

import (
	"fmt"
	"os"
)

func main() {
	printAvailableModesAndChooseOne()
	addTaskToBoardChannel := make(chan task)
	getTaskFromBoardChannel := make(chan readTask)
	printTasksFromBoardChannel := make(chan bool)
	addProductToStoreChannel := make(chan product)
	getProductFromStoreChannel := make(chan readProduct)
	printProductsFromStoreChannel := make(chan bool)
	reportNewCrashChannel := make(chan crashReport)
	getNextReportChannel := make(chan readReport)
	reportFixedCrashChannel := make(chan crashReport)
	machinesChannels := make(map[string]([]machineChannels))

	for typeOfMachine, numberOfMachines := range allowMachinesNumbers {
		var currentTypeMachinesChannels []machineChannels

		for i := 0; i < numberOfMachines; i++ {
			currentTypeMachinesChannels = append(
				currentTypeMachinesChannels,
				machineChannels{
					make(chan bool),
					make(chan bool),
					make(chan machineTask)})
			go machine(i, typeOfMachine, currentTypeMachinesChannels[i])
		}

		machinesChannels[typeOfMachine] = currentTypeMachinesChannels
	}

	go taskBoard(addTaskToBoardChannel, getTaskFromBoardChannel, printTasksFromBoardChannel)
	go store(addProductToStoreChannel, getProductFromStoreChannel, printProductsFromStoreChannel)
	go serviceCenter(reportNewCrashChannel, getNextReportChannel, reportFixedCrashChannel)
	go chief(addTaskToBoardChannel)

	for i := 0; i < numberOfServiceCenterWorkers; i++ {
		go serviceCenterWorker(i, getNextReportChannel, reportFixedCrashChannel, machinesChannels)
	}

	for i := 0; i < numberOfWorkers; i++ {
		go worker(i, getTaskFromBoardChannel, addProductToStoreChannel, reportNewCrashChannel, machinesChannels)
	}

	for i := 0; i < numberOfCustomers; i++ {
		go customer(i, getProductFromStoreChannel)
	}

	if isChattyMode == false {
		printAvailableOptionsAndChooseOne(printTasksFromBoardChannel, printProductsFromStoreChannel)
	} else {
		var mode int
		fmt.Scan(&mode)
		if mode == 0 {
			os.Exit(0)
		}
	}
}
