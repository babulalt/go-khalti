package khalti

// InitiateTransactionRequest represents a Khalti Initaial Payment Request
type InitiateTransactionRequest struct {
	PubicKey        string `json:"public_key"`
	Mobile          string `json:"mobile"`
	TransactionPin  string `json:"transaction_pin"`
	Amount          uint64 `json:"amount"`
	ProductIdentity string `json:"product_identity"`
	ProductName     string `json:"product_name"`
	ProductUrl      string `json:"product_url"`
}

// InitiateTransactionResponse represents a Khalti Initaial Payment Response
type InitiateTransactionResponse struct {
	Token             string `json:"token"`
	PinCreated        bool   `json:"pin_created"`
	PinCreatedMessage string `json:"pin_created_message"`
}

// ConfirmTransactionRequest represents a Khalti Confirm Payment Request
type ConfirmTransactionRequest struct {
	PubicKey         string `json:"public_key"`
	Token            string `json:"token"`
	ConfirmationCode string `json:"confirmation_code"`
	TransactionPin   string `json:"transaction_pin"`
}

// ConfirmTransactionResponse represents a Khalti Confirm Payment Response
type ConformTransactionResponse struct {
	TransactionID   string `json:"idx"`
	Token           string `json:"token"`
	Amount          uint64 `json:"amount"`
	Mobile          string `json:"mobile"`
	ProductIdentity string `json:"product_identity"`
	ProductName     string `json:"product_name"`
	ProductUrl      string `json:"product_url"`
}

// VerifyTransactionRequest represents a Khalti Verify Payment Request
type VerifyTransactionRequest struct {
	Token  string `json:"token"`
	Amount uint64 `json:"amount"`
}

// VerifyTransactionResponse represents a Khalti Verify Payment Response
type VerifyTransactionResponse struct {
	TransactionID   string       `json:"idx"`
	Type            PaymentType  `json:"type"`
	State           PaymentState `json:"state"`
	User            User         `json:"user"`
	Marchent        Marchent     `json:"marchent"`
	CreatedOn       string       `json:"created_on"`
	Token           string       `json:"token"`
	Refunded        bool         `json:"refunded"`
	CashBack        uint64       `json:"cash_back"`
	ProductIdentity string       `json:"product_identity"`
	Reference       interface{}  `json:"reference"`
	Remarks         interface{}  `json:"remarks"`
}

type PaymentType struct {
	PaymentType string `json:"idx"`
	Name        string `json:"name"`
}

type PaymentState struct {
	PaymentStateID string `json:"idx"`
	Name           string `json:"name"`
	Templete       string `json:"template"`
	Amount         uint64 `json:"amount"`
	FeeAmount      uint64 `json:"fee_amount"`
}

type Marchent struct {
	ID     string `json:"idx"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type User struct {
	ID   string `json:"idx"`
	Name string `json:"name"`
}
