package hash

const (
	mod = 1 << 16
)

func Rolling(data []byte) (uint, uint, uint) {
	var a, b uint
	l := uint(len(data))
	for i, value := range data {
		a += uint(value)
		b += (l - uint(i)) * uint(value)
	}
	r1 := a % mod
	r2 := b % mod
	r := r1 + mod*r2
	return r1, r2, r
}

func RollingWindow(l, r1, r2 uint, out, in byte) (uint, uint, uint) {
	r1 = (r1 - uint(out) + uint(in)) % mod
	r2 = (r2 - l*uint(out) + r1) % mod
	r := r1 + mod*r2
	return r1, r2, r
}
