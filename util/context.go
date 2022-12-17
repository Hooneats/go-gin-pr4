package util

import (
	"context"
	"time"
)

const ModelTimeOut = 2 * time.Second
const DatabaseTimeOut = 3 * time.Second
const ControllerTimeOut = 5 * time.Second

func GetContext(t time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	return ctx, cancel
}
