package __4

func PopCount(x uint64) int {
	n := 0
	for x > 0 {
		if x & 1 == 1 {
			n++
		}
		x = x >> 1
	}
	return n
}
