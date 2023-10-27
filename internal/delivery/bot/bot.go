package bot

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/nyaa-x-discord/internal/service"
	"github.com/rl404/nyaa-x-discord/internal/utils"
)

// Bot contains functions for bot.
type Bot struct {
	service service.Service
	prefix  string
}

// New to create new bot.
func New(service service.Service, prefix string) *Bot {
	return &Bot{
		service: service,
		prefix:  prefix,
	}
}

// Run to run bot.
func (b *Bot) Run() error {
	return b.service.Run()
}

// Stop to stop bot.
func (b *Bot) Stop() error {
	return b.service.Stop()
}

// RegisterReadyHandler to register ready handler.
func (b *Bot) RegisterHandler(nrApp *newrelic.Application) {
	b.service.RegisterReadyHandler(b.readyHandler())
	b.service.RegisterMessageHandler(b.messageHandler(nrApp))
}

func (b *Bot) log(ctx context.Context) {
	if rvr := recover(); rvr != nil {
		stack.Wrap(ctx, fmt.Errorf("%s", debug.Stack()), fmt.Errorf("%v", rvr), fmt.Errorf("panic"))
	}

	errStack := stack.Get(ctx)
	if len(errStack) > 0 {
		utils.Log(map[string]interface{}{
			"level": utils.ErrorLevel,
			"error": errStack,
		})
	}
}
