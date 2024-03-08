package rdiff_test

import (
	"testing"

	"github.com/agugaillard/eiger/hash"
	"github.com/agugaillard/eiger/rdiff"
	"github.com/stretchr/testify/assert"
)

func Test_Signature_whenChunkSizeNotDefined_SetsSqrt(t *testing.T) {
	rdiff := rdiff.NewDefaultRdiff(logger)
	bytes := make([]byte, 101)
	signature := rdiff.Signature(bytes)
	assert.Equal(t, uint(10), signature.ChunkSize)
}

func Test_Signature_whenWeakHashCollides_shouldAddBothStrongHashes(t *testing.T) {
	weak := uint(1)
	rdiff := rdiff.NewDefaultRdiff(logger)
	rdiff.WeakHash = func(b []byte) (uint, uint, uint) { return 0, 0, weak }
	signature := rdiff.Signature([]byte("1234"), 2)
	assert.Equal(t, 2, len(signature.HashMap[weak]))
	assert.Contains(t, signature.HashMap[weak], hash.Md5([]byte("12")))
	assert.Contains(t, signature.HashMap[weak], hash.Md5([]byte("34")))
}

func Test_Signature_whenChunkSizeIsNotMultipleOfInputLength_shouldRoundUpNumberOfSignatures(t *testing.T) {
	weak := uint(1)
	rdiff := rdiff.NewDefaultRdiff(logger)
	rdiff.WeakHash = func(b []byte) (uint, uint, uint) { return 0, 0, weak }
	signature := rdiff.Signature([]byte("12345"), 2)
	assert.Equal(t, 3, len(signature.HashMap[weak]))
	assert.Contains(t, signature.HashMap[weak], hash.Md5([]byte("12")))
	assert.Contains(t, signature.HashMap[weak], hash.Md5([]byte("34")))
	assert.Contains(t, signature.HashMap[weak], hash.Md5([]byte("5")))
}
