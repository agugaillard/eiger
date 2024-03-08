package rdiff

import (
	"fmt"
	"math"

	"go.uber.org/zap"
)

type Signature struct {
	ChunkSize uint
	HashMap   HashMap
}

type HashMap map[uint]map[string]uint

func (r *Rdiff) Signature(input []byte, chunkSize ...uint) (signature Signature) {
	if len(chunkSize) == 0 || chunkSize[0] == 0 {
		signature.ChunkSize = uint(math.Floor(math.Sqrt(float64(len(input)))))
		r.Logger.Info("calculating chunk_size", zap.Uint("chunk_size", signature.ChunkSize))
	} else {
		signature.ChunkSize = chunkSize[0]
	}

	lenBlocks := uint(math.Ceil(float64(len(input)) / float64(signature.ChunkSize)))

	// Worst case: All weak signatures are different
	signature.HashMap = make(HashMap, lenBlocks)

	debugSignature := make(map[string]string)
	for i := uint(0); i < lenBlocks; i++ {
		block := getBlock(input, i*signature.ChunkSize, signature.ChunkSize)
		_, _, weak := r.WeakHash(block)
		strong := r.StrongHash(block)
		if _, found := signature.HashMap[weak]; found {
			signature.HashMap[weak][strong] = i
		} else {
			signature.HashMap[weak] = map[string]uint{strong: i}
		}
		debugSignature[string(block)] = fmt.Sprintf("%d#%s", weak, strong)
	}
	r.Logger.Debug("creating signature", zap.Uint("chunk_size", signature.ChunkSize), zap.Any("hashmap", debugSignature))
	return
}
