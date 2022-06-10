package main

import (
	"fmt"

	"github.com/babulalt/go-khalti/v1/khalti"
)

func main() {
	clientId := "PUBLIC_KEY"
	secretId := "SECRET_KEY"
	khaltiService, err := khalti.NewKhaltiClient(clientId, secretId, nil)
	if err != nil {
		fmt.Println("error in client:: ", err)
		return
	}
	trasaction := &khalti.InitiateTransactionRequest{
		PubicKey:        khaltiService.ClientID,
		Mobile:          "XXXXXXXXXX",
		TransactionPin:  "XXXX",
		Amount:          1000,
		ProductIdentity: "Test",
		ProductName:     "test",
		ProductUrl:      "product_url",
	}
	data, _, err := khaltiService.InitiateTransaction(trasaction)
	if err != nil {
		fmt.Println("initiate payment error ::", err)
		return
	}
	fmt.Println("data", data)
}
