package udprouter

import (
	"fmt"
	"time"
)

// Logger is just a simple Logger interface
// it should already be satisifed by many other loggers (e.g logrus)
type Logger interface {
	Info(msg string)
}

// WrapWithLogger wraps a given UDPRouter with a logger
// so incoming requests can be logged
func WrapWithLogger(r UDPRouter, logger Logger) UDPRouter {
	return &MockUDPRouter{
		HandleFn: func(route string, handleFn UDPHandleFunc) {
			r.Handle(route, loggingHandler(handleFn, logger, route))
		},
		ListenFn: func(addr string) error {
			return r.Listen(addr)
		},
	}
}

func loggingHandler(handleFn UDPHandleFunc, logger Logger, route string) UDPHandleFunc {
	return func(body string, respond UDPResponseFunc) {
		handleFn(body, respond)

		// Log request time, route and body
		msg := fmt.Sprintf("%s|%s|%s", time.Now(), route, body)
		logger.Info(msg)
	}
}
