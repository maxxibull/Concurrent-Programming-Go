package main

import (
	"math/rand"
	"time"
)

func machine(index int, op string, theMachineChannels machineChannels) {
	isMachineBroken := false
	theMachineChannels.isMachineReadyChannel <- true

	for {
		if !isMachineBroken {
			isMachineBroken = crashMachine()

			if isMachineBroken && isChattyMode {
				printMachineIsBroken(index, op)
			}
		}

		select {
		case <-theMachineChannels.fixMachineChannel:
			isMachineBroken = false

			if isChattyMode {
				printMachineIsOK(index, op)
			}

			theMachineChannels.fixMachineChannel <- true
		case newMachineTask := <-theMachineChannels.newMachineTaskChannel:
			if isMachineBroken {
				time.Sleep(time.Millisecond * time.Duration(timeToSleepForMachine))
				newMachineTask.responseChannel <- nil
			} else {
				newTask := newMachineTask.currentTask

				if isChattyMode {
					printMachineStart(index, newMachineTask.workerIndex, newTask)
				}

				newTask.executeTask()
				time.Sleep(time.Millisecond * time.Duration(timeToSleepForMachine))

				if isChattyMode {
					printMachineFinish(index, newMachineTask.workerIndex, newTask)
				}

				newMachineTask.responseChannel <- &newTask
			}

			theMachineChannels.isMachineReadyChannel <- true
		}
	}
}

func crashMachine() bool {
	if rand.Intn(10) < probabilityOfMachineCrash {
		return true
	}

	return false
}
