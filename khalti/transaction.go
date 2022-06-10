package khalti

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Initiate tansaction khalti
// Endpoint: POST /api/v2/payment/initiate/
func (s *KhaltiService) InitiateTransaction(payload *InitiateTransactionRequest) (*InitiateTransactionResponse, *http.Response, error) {
	url := s.BaseUrl + "api/v2/payment/initiate/"
	req, err := http.NewRequest(http.MethodPost, url, Payload(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	data := &InitiateTransactionResponse{}
	err = json.Unmarshal(body, data)
	return data, nil, err
}

// Initiate tansaction khalti
// Endpoint: POST /api/v2/payment/confirm/
// func (s *KhaltiService) ConfirmationTransaction(ctx context.Context, payload *ConfirmTransactionRequest) (*ConformTransactionResponse, *http.Response, error) {
// 	req, err := s.Client.NewRequest(http.MethodPost, "api/v2/payment/confirm/", payload)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	data := &ConformTransactionResponse{}
// 	resp, err := s.client.Do(req, data)
// 	if err != nil {
// 		return nil, resp, err
// 	}
// 	return data, resp, nil
// }

// // Initiate tansaction khalti
// // Endpoint: POST /api/v2/payment/verify/
// func (s *KhaltiService) VerifyTransaction(ctx context.Context, payload *VerifyTransactionRequest) (*VerifyTransactionResponse, *http.Response, error) {
// 	req, err := s.client.NewRequest(http.MethodPost, "api/v2/payment/verify/", payload)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", fmt.Sprintf("key %s", s.client.Secret))
// 	data := &VerifyTransactionResponse{}
// 	resp, err := s.client.Do(req, data)
// 	if err != nil {
// 		return nil, resp, err
// 	}
// 	return data, resp, nil
// }

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
