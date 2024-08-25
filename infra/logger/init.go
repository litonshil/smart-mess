package logger

import "smart-mess/domain"

var client LogClient

func NewLogClient(lvl string) domain.ILogger {
	connectZap(lvl)

	return &LogClient{}
}

func Client() LogClient {
	return client
}
