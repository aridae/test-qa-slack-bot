package middleware

import qaslackbot "github.com/aridae/test-qa-slack-bot"

type Middleware func(service qaslackbot.Service) qaslackbot.Service
