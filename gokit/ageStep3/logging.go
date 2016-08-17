package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   AgeService
}

func (mw loggingMiddleware) CalculateAge(y int) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "calculateAge",
			"input", y,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.CalculateAge(y)
	return
}
