package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	qaslackbot "github.com/aridae/test-qa-slack-bot"
	"github.com/aridae/test-qa-slack-bot/model"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next qaslackbot.Service) qaslackbot.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   qaslackbot.Service
	logger log.Logger
}

func (mw loggingMiddleware) SendDispatch(ctx context.Context, dispatch model.Dispatch) (_ []model.Message, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "SendDispatch", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.SendDispatch(ctx, dispatch)
}
