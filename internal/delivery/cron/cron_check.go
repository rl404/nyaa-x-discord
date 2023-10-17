package cron

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/errors/stack"
)

// Check to check nyaa update.
func (c *Cron) Check(nrApp *newrelic.Application, interval time.Duration) error {
	ctx := stack.Init(context.Background())
	defer c.log(ctx)

	tx := nrApp.StartTransaction("Cron check")
	defer tx.End()

	ctx = newrelic.NewContext(ctx, tx)

	if err := c.service.CheckNyaa(ctx, nrApp, interval); err != nil {
		return stack.Wrap(ctx, err)
	}

	return nil
}
