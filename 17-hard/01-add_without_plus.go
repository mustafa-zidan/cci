package hard

func add(x, y int) int {
	sum, carry, i := 0, 0, 0
	for x != 0 || y != 0 {
		xBit, yBit := x&1, y&1
		xor := (xBit ^ yBit) | carry
		carry = xBit & yBit
		sum |= (xor << i)
		x, y = x>>1, y>>1
		i++
	}
	if carry == 1 {
		sum |= carry << i
	}
	return sum
}
