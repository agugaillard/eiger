package rdiff

func getBlock(data []byte, offset, size uint) []byte {
	var block []byte
	if offset > uint(len(data))-size {
		block = data[offset:]
	} else {
		block = data[offset : offset+size]
	}
	return block
}
