package main

func checkAmountOfTasks(condition bool, taskChannel <-chan task) <-chan task {
	if condition {
		return taskChannel
	}

	return nil
}

func checkTasksBoard(condition bool, taskChannel <-chan readTask) <-chan readTask {
	if condition {
		return taskChannel
	}

	return nil
}

func taskBoard(addTaskToBoardChannel <-chan task, getTaskFromBoardChannel <-chan readTask, printTasksFromBoardChannel <-chan bool) {
	var board []task

	for {
		select {
		case send := <-checkTasksBoard(len(board) > 0, getTaskFromBoardChannel):
			send.response <- board[0]
			board = board[1:]
		case add := <-checkAmountOfTasks(len(board) < maxAmountOfTasks, addTaskToBoardChannel):
			board = append(board, add)
		case <-printTasksFromBoardChannel:
			printTasksBoard(board)
		}
	}
}
