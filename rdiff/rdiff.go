package rdiff

import (
	"github.com/agugaillard/eiger/hash"
	"go.uber.org/zap"
)

type Rdiff struct {
	Logger     *zap.Logger
	StrongHash func([]byte) string
	WeakHash   func([]byte) (uint, uint, uint)
	WeakWindow func(l uint, r1 uint, r2 uint, out byte, in byte) (uint, uint, uint)
}

func NewDefaultRdiff(logger *zap.Logger) *Rdiff {
	return &Rdiff{
		Logger:     logger,
		StrongHash: hash.Md5,
		WeakHash:   hash.Rolling,
		WeakWindow: hash.RollingWindow,
	}
}
