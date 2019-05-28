package main

var operations = []string{
	"+",
	"*"}

var allowMachinesNumbers = map[string]int{
	"+": numberOfAddingMachines,
	"*": numberOfMultiplyingMachines}

var allowOperations = map[string]func(int, int) int{
	"+": func(x, y int) int { return x + y },
	"*": func(x, y int) int { return x * y }}

const minValueOfArgument = 0
const maxValueOfArgument = 10

const maxAmountOfTasks = 20
const maxAmountOfProducts = 20

const numberOfWorkers = 6
const numberOfCustomers = 3
const numberOfAddingMachines = 2
const numberOfMultiplyingMachines = 2
const numberOfServiceCenterWorkers = 3

const timeToSleepForWorker = 200
const timeToSleepForCustomer = 500
const timeToSleepForChief = 50
const timeToSleepForMachine = 300
const timeToSleepForServiceCenterWorker = 50

const timeToWaitForImpatientWorker = 150

const probabilityOfMachineCrash = 3 // from 0 to 10
