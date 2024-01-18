package logging

import (
	"sync"

	"github.com/hashicorp/go-hclog"
)

var (
	logger hclog.Logger
	once   sync.Once
)

func GetLogger() hclog.Logger {
	once.Do(func() {
		opts := &hclog.LoggerOptions{
			Name:  "gokey-logger",
			Level: 1,
		}
		logger = hclog.New(opts)

	})
	return logger
}
