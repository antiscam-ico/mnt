package api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

type SendResponse struct {
	Result struct {
		Code int    `json:"code"`
		Data string `json:"data"`
		Log  string `json:"log"`
		Hash string `json:"hash"`
	} `json:"result"`
	Error struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		TxResult struct {
			Code int    `json:"code"`
			Log  string `json:"log"`
		} `json:"tx_result"`
	} `json:"error"`
}

func (a *Api) Send(transaction transaction.SignedTransaction) (*SendResponse, error) {
	bytes, err := transaction.Encode()
	if err != nil {
		return nil, err
	}
	response := new(SendResponse)
	res, err := a.client.R().SetResult(response).Get(fmt.Sprintf("/send_transaction?tx=%s", bytes))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		err = json.Unmarshal(res.Body(), response) //todo edit response from server
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}