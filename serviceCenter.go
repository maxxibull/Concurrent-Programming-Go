package main

import (
	"time"
)

func checkBoard(condition bool, getNextReportChannel <-chan readReport) <-chan readReport {
	if condition {
		return getNextReportChannel
	}

	return nil
}

func serviceCenter(reportNewCrashChannel <-chan crashReport, getNextReportChannel <-chan readReport,
	reportFixedCrashChannel <-chan crashReport) {

	var newReportBoard []crashReport
	var sentReportBoard []crashReport

	for {
		select {
		case fixedCrashReport := <-reportFixedCrashChannel:
			for index, value := range sentReportBoard {
				if value.machineIndex == fixedCrashReport.machineIndex && value.machineType == fixedCrashReport.machineType {
					sentReportBoard = append(sentReportBoard[:index], sentReportBoard[index+1:]...)
					break
				}
			}
		case newReadReport := <-checkBoard(len(newReportBoard) > 0, getNextReportChannel):
			newReadReport.report <- newReportBoard[0]
			sentReportBoard = append(sentReportBoard, newReportBoard[0])
			newReportBoard = newReportBoard[1:]
		case newCrashReport := <-reportNewCrashChannel:
			shouldBeReportAdded := true
			for _, value := range append(newReportBoard, sentReportBoard...) {
				if value.machineIndex == newCrashReport.machineIndex && value.machineType == newCrashReport.machineType {
					shouldBeReportAdded = false
					break
				}
			}

			if shouldBeReportAdded {
				newReportBoard = append(newReportBoard, newCrashReport)
			}
		}
	}
}

func serviceCenterWorker(index int, getNextReportChannel chan<- readReport, reportFixedCrashChannel chan<- crashReport,
	machinesChannels map[string]([]machineChannels)) {

	for {
		nextReportRequest := readReport{
			make(chan crashReport)}

		getNextReportChannel <- nextReportRequest
		currentReport := <-nextReportRequest.report

		if isChattyMode {
			printServiceCenterWorkerIsGoingToFixMachine(index, currentReport)
		}

		time.Sleep(time.Millisecond * time.Duration(timeToSleepForServiceCenterWorker))
		machinesChannels[currentReport.machineType][currentReport.machineIndex].fixMachineChannel <- true
		<-machinesChannels[currentReport.machineType][currentReport.machineIndex].fixMachineChannel

		if isChattyMode {
			printServiceCenterWorkerFixedMachine(index, currentReport)
		}

		reportFixedCrashChannel <- currentReport
	}
}
