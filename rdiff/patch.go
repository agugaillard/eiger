package rdiff

import (
	"strconv"

	"go.uber.org/zap"
)

func (r *Rdiff) Patch(base []byte, delta Delta) (output string) {
	for _, change := range delta.Changes {
		if change.Kind == DELTA_KIND_PLAIN {
			output += change.Value
		} else {
			index, _ := strconv.Atoi(change.Value)
			output += string(getBlock(base, uint(index)*delta.ChunkSize, delta.ChunkSize))
		}
	}
	r.Logger.Debug("generating output", zap.String("output", output))
	return
}
