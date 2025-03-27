package day25

type Pins [5]uint8

func Fits(a, b Pins) bool {
	for i := range 5 {
		if a[i]+b[i] > 5 {
			return false
		}
	}
	return true
}

func NaiveFitPairs(locks, keys []Pins) (total int) {
	for _, lock := range locks {
		for _, key := range keys {
			if Fits(lock, key) {
				total++
			}
		}
	}

	return
}
