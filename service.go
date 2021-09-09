package qaslackbot

import (
	"context"
	"errors"

	"github.com/aridae/test-qa-slack-bot/model"
)

var (
	ErrChanelAnavailable = errors.New("у бота нет прав на отправку сообщений в указанный канал")
	ErrInvalidToken      = errors.New("slack-токен недействительный")
)

type Service interface {
	SendDispatch(ctx context.Context, msgs model.Dispatch) ([]model.Message, error)
}
