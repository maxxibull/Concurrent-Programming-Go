package main

import (
	"math/rand"
	"time"
)

func worker(index int, getTaskChannel chan<- readTask, newProductChannel chan<- product,
	reportNewCrashChannel chan<- crashReport, machinesChannels map[string]([]machineChannels)) {

	chooseWorkerType := map[bool]func(int, task, chan<- crashReport, []machineChannels) task{
		true:  impatientWorker,
		false: patientWorker}

	impatient := randPatient()
	findMachine := chooseWorkerType[impatient]

	mutex.Lock()
	workersInfo[index].executedTasks = 0
	workersInfo[index].impatient = impatient
	mutex.Unlock()

	for {
		nextTaskRequest := readTask{
			response: make(chan task)}

		getTaskChannel <- nextTaskRequest
		currentTask := <-nextTaskRequest.response

		if isChattyMode {
			printWorkerStart(index, currentTask)
		}

		currentTask = findMachine(index, currentTask, reportNewCrashChannel, machinesChannels[currentTask.operationSymbol])
		time.Sleep(time.Millisecond * time.Duration(timeToSleepForWorker))
		newProduct := product{
			value: currentTask.result}

		if isChattyMode {
			printWorkerFinish(index, currentTask, newProduct)
		}

		newProductChannel <- newProduct
		mutex.Lock()
		workersInfo[index].executedTasks++
		mutex.Unlock()
	}
}

func impatientWorker(index int, currentTask task, reportNewCrashChannel chan<- crashReport,
	currentMachinesChannels []machineChannels) task {

	for {
		machineIndex := rand.Intn(len(currentMachinesChannels))

		if isChattyMode {
			printWorkerChoseMachine(index, machineIndex, "impatient", currentTask)
		}

		theMachineChannels := currentMachinesChannels[machineIndex]

		select {
		case <-theMachineChannels.isMachineReadyChannel:
			newMachineTask := machineTask{
				currentTask,
				index,
				make(chan *task)}
			theMachineChannels.newMachineTaskChannel <- newMachineTask
			resultTask := <-newMachineTask.responseChannel

			if resultTask == nil {
				reportNewCrashChannel <- crashReport{
					currentTask.operationSymbol,
					machineIndex}

				if isChattyMode {
					printWorkerFoundBrokenMachine(index, machineIndex, currentTask)
				}

				break
			}

			return *resultTask
		case <-time.After(time.Millisecond * time.Duration(timeToWaitForImpatientWorker)):
			if isChattyMode {
				printWorkerAbandonedMachine(index, machineIndex, currentTask)
			}
			break
		}
	}
}

func patientWorker(index int, currentTask task, reportNewCrashChannel chan<- crashReport,
	currentMachinesChannels []machineChannels) task {

	for {
		machineIndex := rand.Intn(len(currentMachinesChannels))

		if isChattyMode {
			printWorkerChoseMachine(index, machineIndex, "patient", currentTask)
		}

		theMachineChannels := currentMachinesChannels[machineIndex]

		select {
		case <-theMachineChannels.isMachineReadyChannel:
			newMachineTask := machineTask{
				currentTask,
				index,
				make(chan *task)}
			theMachineChannels.newMachineTaskChannel <- newMachineTask
			resultTask := <-newMachineTask.responseChannel

			if resultTask == nil {
				reportNewCrashChannel <- crashReport{
					currentTask.operationSymbol,
					machineIndex}

				if isChattyMode {
					printWorkerFoundBrokenMachine(index, machineIndex, currentTask)
				}

				break
			}

			return *resultTask
		}
	}
}

func randPatient() bool {
	if rand.Intn(10) < 5 {
		return false
	}

	return true
}
