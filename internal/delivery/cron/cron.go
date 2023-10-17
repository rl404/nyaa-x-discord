package cron

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/nyaa-x-discord/internal/service"
	"github.com/rl404/nyaa-x-discord/internal/utils"
	"github.com/rl404/nyaa-x-discord/pkg/log"
)

// Cron contains functions for cron.
type Cron struct {
	service service.Service
}

// New to create new cron.
func New(service service.Service) *Cron {
	return &Cron{
		service: service,
	}
}

func (c *Cron) log(ctx context.Context) {
	if rvr := recover(); rvr != nil {
		stack.Wrap(ctx, fmt.Errorf("%s", debug.Stack()), fmt.Errorf("%v", rvr), fmt.Errorf("panic"))
	}

	errStack := stack.Get(ctx)
	if len(errStack) > 0 {
		utils.Log(map[string]interface{}{
			"level": log.ErrorLevel,
			"error": errStack,
		})
	}
}
