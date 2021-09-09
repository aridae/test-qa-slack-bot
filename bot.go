package qaslackbot

import (
	"context"

	"github.com/aridae/test-qa-slack-bot/model"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type bot struct {
	token  string
	logger log.Logger
}

func NewBot(token string, logger log.Logger) Service {
	return &bot{
		token:  token,
		logger: logger,
	}
}

func (b *bot) SendDispatch(ctx context.Context, msgs model.Dispatch) ([]model.Message, error) {
	for _, msg := range msgs.Messages {
		level.Debug(b.logger).Log("SendMessage function", msg.Channel, msg.Text)
	}
	return nil, nil
}
