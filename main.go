package main

import (
	"cryptostocksaccounts/transaction"
	"fmt"
)

func main() {
	var transactionDatabase transaction.Database
	transactionDatabase.Init()
	defer transactionDatabase.Close()

	// transactionData := transaction.Transaction{UserId: 2, Type: "sell", Symbol: "ETH", Amount: 1.253, Price: 13768, Date: time.Now()}
	// transactionDatabase.InsertData(transactionData)
	transactions := transactionDatabase.GetAllTransactions()
	for _, v := range transactions {
		fmt.Println(v)
	}
}
