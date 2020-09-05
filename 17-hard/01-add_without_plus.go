package hard

// this approach breaks if there is a carry between multiple
// bits
func addWithoutPlusFaulty(x, y int) int {
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

func addWithoutPlus(x, y int) int {
	for y != 0 {
		// carry now contains common
		// set bits of x and y
		carry := x & y
		// Sum of bits of x and
		// y where at least one
		// of the bits is not set
		x = x ^ y
		// Carry is shifted by
		// one so that adding it
		// to x gives the required sum
		y = carry << 1
	}
	return x
}
