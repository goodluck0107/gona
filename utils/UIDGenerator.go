package utils

// SN2UID converts serial number from redis to UID
func SN2UID(originSN int64) int64 {
	sn := int(originSN) + 10362409
	uid := 0

	serial := []int{5, 1, 7, 3, 6, 4, 2, 0}
	digits := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for _, pos := range serial {
		digits[pos] = sn % 10
		sn /= 10
	}
	for _, digit := range digits {
		uid *= 10
		uid += digit
	}

	return int64(uid)
}

func UID2SN(originSN int64) int64 {
	sn := int(originSN)
	uid := 0

	serial := []int{5, 3, 7, 2, 4, 1, 6, 0}
	digits := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for _, pos := range serial {
		digits[pos] = sn % 10
		sn /= 10
	}
	for _, digit := range digits {
		uid *= 10
		uid += digit
	}

	return int64(uid) - 10362409
}
