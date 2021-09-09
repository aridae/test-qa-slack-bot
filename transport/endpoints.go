package transport

import (
	"context"

	qaslackbot "github.com/aridae/test-qa-slack-bot"
	"github.com/aridae/test-qa-slack-bot/model"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SendDispatch endpoint.Endpoint
}

func InitEndpoints(s qaslackbot.Service) Endpoints {
	return Endpoints{
		SendDispatch: initSendDispatch(s),
	}
}

func initSendDispatch(s qaslackbot.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendDispatchRequest)
		spoiledMsgs, err := s.SendDispatch(ctx, model.Dispatch{Token: req.Token, Messages: req.Messages})
		return SendDispatchResponse{
			Err: err,
			Stats: DispatchStats{
				Sent:            len(req.Messages),
				Delievered:      len(req.Messages) - len(spoiledMsgs),
				SpoiledMessages: spoiledMsgs,
			},
		}, nil
	}
}
