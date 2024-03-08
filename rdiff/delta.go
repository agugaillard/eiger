package rdiff

import (
	"strconv"

	"go.uber.org/zap"
)

const (
	DELTA_KIND_PLAIN = "P"
	DELTA_KIND_BLOCK = "B"
)

type Delta struct {
	ChunkSize uint
	Changes   []change `json:"changes"`
}

type change struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

func (d *Delta) addPlain(char string) {
	if len(d.Changes) > 0 && d.Changes[len(d.Changes)-1].Kind == DELTA_KIND_PLAIN {
		d.Changes[len(d.Changes)-1].Value += char
	} else {
		d.Changes = append(d.Changes, change{Kind: DELTA_KIND_PLAIN, Value: char})
	}
}

func (d *Delta) addBlock(offset uint) {
	d.Changes = append(d.Changes, change{Kind: DELTA_KIND_BLOCK, Value: strconv.Itoa(int(offset))})
}

func (r *Rdiff) Delta(input []byte, signature Signature) (delta Delta) {
	delta.ChunkSize = signature.ChunkSize
	var weak1, weak2, weak uint
	resetWindow := true
	for i := uint(0); i < uint(len(input)); i++ {
		block := getBlock(input, i, delta.ChunkSize)
		if resetWindow {
			weak1, weak2, weak = r.WeakHash(block)
			resetWindow = false
		} else {
			weak1, weak2, weak = r.WeakWindow(uint(len(block)), weak1, weak2, input[i-1], block[len(block)-1])
		}
		if strongHashes, weakFound := signature.HashMap[weak]; weakFound {
			strong := r.StrongHash(block)
			if offset, strongFound := strongHashes[strong]; strongFound {
				i += delta.ChunkSize - 1
				resetWindow = true
				delta.addBlock(offset)
				continue
			}
		}
		delta.addPlain(string(input[i]))
	}
	r.Logger.Debug("calculating deltas", zap.Uint("chunk_size", delta.ChunkSize), zap.Any("delta", delta.Changes))
	return
}
