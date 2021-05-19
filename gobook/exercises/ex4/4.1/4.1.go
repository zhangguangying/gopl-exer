package __1

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b-1
	}
	return count
}

func bitDiff(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < len(c1) || i < len(c2); i++ {
		switch {
		case i >= len(c1):
			count += popCount(c2[i])
		case i >= len(c2):
			count += popCount(c1[i])
		default:
			count += popCount(c1[i] ^ c2[i])
		}
	}
	return count
}
