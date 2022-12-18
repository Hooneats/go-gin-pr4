package util

import (
	"context"
	"time"
)

const ModelTimeOut = 1 * time.Second
const DatabaseTimeOut = 2 * time.Second
const ControllerTimeOut = 1 * time.Second

func GetContext(t time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	return ctx, cancel
}
