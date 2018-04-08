package utils

func DivMod(numerator int, denominator int) (quotient int, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
