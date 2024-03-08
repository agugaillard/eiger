package rdiff_test

import "go.uber.org/zap"

var (
	logger, _ = zap.Config{Development: true, Level: zap.NewAtomicLevelAt(zap.PanicLevel), Encoding: "json"}.Build()
)
