package main

func checkAmountOfProducts(condition bool, productChannel <-chan product) <-chan product {
	if condition {
		return productChannel
	}

	return nil
}

func checkProductsBoard(condition bool, taskChannel <-chan readProduct) <-chan readProduct {
	if condition {
		return taskChannel
	}

	return nil
}

func store(addProductToStoreChannel <-chan product, getProductFromStoreChannel <-chan readProduct,
	printProductsFromStoreChannel <-chan bool) {
	var board []product

	for {
		select {
		case sell := <-checkProductsBoard(len(board) > 0, getProductFromStoreChannel):
			sell.response <- board[0]
			board = board[1:]
		case add := <-checkAmountOfProducts(len(board) < maxAmountOfProducts, addProductToStoreChannel):
			board = append(board, add)
		case <-printProductsFromStoreChannel:
			printStore(board)
		}
	}
}
