package http

import (
	"net/http"

	"github.com/go-kit/kit/log"

	gokithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"

	"github.com/aridae/test-qa-slack-bot/transport"
)

func NewServer(
	endpoints transport.Endpoints,
	logger log.Logger,
	options []gokithttp.ServerOption,
) http.Handler {
	var (
		r           = mux.NewRouter()
		errorLogger = gokithttp.ServerErrorLogger(logger)
	)
	options = append(options, errorLogger)
	r.Methods("POST").Path("/dispatch").Handler(
		gokithttp.NewServer(
			endpoints.SendDispatch,
			transport.DecodeSendDispatchRequest,
			transport.EncodeSendDispatchResponse,
			options...,
		),
	)
	return r
}
