package cron

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

// Check to check nyaa update.
func (c *Cron) Check(nrApp *newrelic.Application, interval time.Duration) error {
	ctx := errors.Init(context.Background())
	defer c.log(ctx)

	tx := nrApp.StartTransaction("Cron check")
	defer tx.End()

	ctx = newrelic.NewContext(ctx, tx)

	if err := c.service.CheckNyaa(ctx, nrApp, interval); err != nil {
		return errors.Wrap(ctx, err)
	}

	return nil
}
