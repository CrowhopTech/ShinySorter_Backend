package tickexecutor

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type tickExecutor struct {
	ctx    context.Context
	t      *time.Ticker
	onTick func(ctx context.Context) error
	onErr  func(err error)
}

func New(ctx context.Context, tickInterval time.Duration, onTick func(ctx context.Context) error, onErr func(err error)) {
	te := tickExecutor{
		ctx:    ctx,
		t:      time.NewTicker(tickInterval),
		onTick: onTick,
		onErr:  onErr,
	}
	go te.backgroundLoop()
}

func (te *tickExecutor) backgroundLoop() {
	done := false

	for !done {
		select {
		case <-te.ctx.Done():
			logrus.Debug("Ticker stopping due to context")
			done = true
		case <-te.t.C:
			if te.onTick != nil {
				err := te.onTick(te.ctx)
				if err != nil {
					if te.onErr != nil {
						te.onErr(err)
					} else {
						logrus.WithError(err).Error("Ticker function execution failed")
					}
				}
			} else {
				logrus.Warn("Ticker running with no function set")
			}
		}
	}
}
