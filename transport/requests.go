package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aridae/test-qa-slack-bot/model"
)

type SendDispatchRequest struct {
	Token    []byte
	Messages []model.Message
}

func DecodeSendDispatchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req SendDispatchRequest
	fmt.Print(r.Body, "\n")
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	fmt.Print(req, "\n")
	return req, nil
}
