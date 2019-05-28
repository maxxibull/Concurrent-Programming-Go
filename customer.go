package main

import (
	"time"
)

func customer(index int, getProductFromStoreChannel chan<- readProduct) {
	for {
		nextProductRequest := readProduct{
			response: make(chan product)}

		getProductFromStoreChannel <- nextProductRequest
		response := <-nextProductRequest.response

		if isChattyMode {
			printCustomer(index, response)
		}

		time.Sleep(time.Millisecond * time.Duration(timeToSleepForCustomer))
	}
}
