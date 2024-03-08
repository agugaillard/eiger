package hash_test

import (
	"testing"

	"github.com/agugaillard/eiger/hash"
	"github.com/stretchr/testify/assert"
)

func Test_RollingHash(t *testing.T) {
	base := []byte("123456789")
	baseHash1, baseHash2, _ := hash.Rolling(base)

	next := []byte("234567890")
	_, _, nextHash := hash.Rolling(next)

	_, _, nextRollingHash := hash.RollingWindow(uint(len(base)), baseHash1, baseHash2, '1', '0')

	assert.Equal(t, nextHash, nextRollingHash)
}
