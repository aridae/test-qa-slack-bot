package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aridae/test-qa-slack-bot/model"
)

type DispatchStats struct {
	Sent            int
	Delievered      int
	SpoiledMessages []model.Message
}

type SendDispatchResponse struct {
	Err   error
	Stats DispatchStats
}

func EncodeSendDispatchResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
