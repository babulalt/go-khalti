package khalti

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Initiate tansaction khalti
// Endpoint: POST /api/v2/payment/initiate/
func (s *KhaltiService) InitiateTransaction(payload *InitiateTransactionRequest) (*InitiateTransactionResponse, error) {
	url := s.BaseUrl + "api/v2/payment/initiate/"
	req, err := http.NewRequest(http.MethodPost, url, Payload(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := &InitiateTransactionResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, err
}

//Confirm tansaction khalti
//Endpoint: POST /api/v2/payment/confirm/
func (s *KhaltiService) ConfirmationTransaction(ctx context.Context, payload *ConfirmTransactionRequest) (*ConformTransactionResponse, error) {
	url := s.BaseUrl + "api/v2/payment/confirm/"
	req, err := http.NewRequest(http.MethodPost, url, Payload(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	confirmTransactionResponse := &ConformTransactionResponse{}
	err = json.Unmarshal(body, confirmTransactionResponse)
	if err != nil {
		return nil, err
	}
	return confirmTransactionResponse, nil
}

// Initiate tansaction khalti
// Endpoint: POST /api/v2/payment/verify/
func (s *KhaltiService) VerifyTransaction(ctx context.Context, payload *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	url := s.BaseUrl + "api/v2/payment/verify/"
	req, err := http.NewRequest(http.MethodPost, url, Payload(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("key %s", s.SecretKey))
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	verifyTransactionResponse := &VerifyTransactionResponse{}
	err = json.Unmarshal(body, verifyTransactionResponse)
	if err != nil {
		return nil, err
	}
	return verifyTransactionResponse, nil
}

func Payload(payload interface{}) io.Reader {
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil
		}
		return bytes.NewBuffer(b)
	}
	return nil
}

func Response(res *http.Response, data interface{}) (interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//data := &InitiateTransactionResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
