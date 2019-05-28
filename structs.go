package main

import "sync"

var isChattyMode = false
var mutex = &sync.Mutex{}
var workersInfo [numberOfWorkers]workerInfo

type task struct {
	firstArgument   int
	secondArgument  int
	operationSymbol string
	result          int
}

func (t *task) executeTask() {
	t.result = allowOperations[t.operationSymbol](t.firstArgument, t.secondArgument)
}

type machineChannels struct {
	isMachineReadyChannel chan bool
	fixMachineChannel     chan bool
	newMachineTaskChannel chan machineTask
}

type machineTask struct {
	currentTask     task
	workerIndex     int
	responseChannel chan *task
}

type readProduct struct {
	response chan product
}

type product struct {
	value int
}

type readTask struct {
	response chan task
}

type workerInfo struct {
	impatient     bool
	executedTasks int
}

type crashReport struct {
	machineType  string
	machineIndex int
}

type readReport struct {
	report chan crashReport
}
