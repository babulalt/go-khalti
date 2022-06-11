# Go client for Khalti REST API

## Coverage

###  /api/v2/payment
* POST /api/v2/payment/initiate/
* POST /api/v2/payment/confirm/
* POST /api/v2/payment/verify/

## Usage

```go
import (
	"github.com/babulalt/go-khalti/khalti"
)

    // Create a Khalti Client instance
	clientId := "PUBLIC_KEY"
	secretId := "SERECT_KEY"
	khaltiClient, err := khalti.NewKhaltiClient(clientId, secretId, nil)
```
## Initiate Khalti Tansaction 

```go
initiateTrasaction := &khalti.InitiateTransactionRequest{
		PubicKey:        khaltiClient.ClientID,
		Mobile:          "XXXXXXXXXX",
		TransactionPin:  "XXXX",
		Amount:          1000,
		ProductIdentity: "Test",
		ProductName:     "Test",
		ProductUrl:      "product-url",
	}
	initiateTrasactionResponse, err := khaltiService.InitiateTransaction(initiateTrasaction)
```

## Confirmation Khalti Tansaction 

```go
confirmTansactionRequest := &khalti.ConfirmTransactionRequest{
		PubicKey:         "",
		Token:            "",
		ConfirmationCode: "",
		TransactionPin:   "XXXX",
	}
	confirmTansactionResponse, err := khaltiService.ConfirmationTransaction(context.Background(),confirmTansactionRequest)
```

## Verify Khalti Tansaction 

```go
verifyTransactionRequest := &khalti.VerifyTransactionRequest{
		Token:  "XXXXXXXXXXXXXX",
		Amount: 1000,
	}
verrifyTansactionResponse, err := khaltiService.VerifyTransaction(context.Background(), verifyTransactionRequest)
```

## How to Contribute

* Fork a repository
* Add/Fix something
* Create PR
