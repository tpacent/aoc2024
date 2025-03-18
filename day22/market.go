package day22

func NthStep(n uint32, count int) uint32 {
	for range count {
		n = Step(n)
	}

	return n
}

const pruneMagic = 16777216

func Step(n uint32) uint32 {
	n = (n*64 ^ n) % pruneMagic
	n = (n/32 ^ n) % pruneMagic
	n = (n*2048 ^ n) % pruneMagic
	return n
}

type Stocks struct {
	Deltas   []int8
	Prices   []int8
	Patterns map[uint32]int8
}

func MaxPattern(secrets []uint32, steps int) (maxValue int) {
	patterns := make(map[uint32]int, steps)

	for _, secret := range secrets {
		for key, value := range CalcStocks(secret, steps+1).Patterns {
			patterns[key] += int(value)
		}
	}

	for _, value := range patterns {
		if value > maxValue {
			maxValue = value
		}
	}

	return
}

func CalcStocks(secret uint32, steps int) Stocks {
	stocks := Stocks{
		Deltas: make([]int8, steps),
		Prices: make([]int8, steps),
	}

	for i := range steps {
		stocks.Prices[i] = int8(secret % 10)
		secret = Step(secret)
	}

	for i := 1; i < steps; i++ {
		stocks.Deltas[i] = stocks.Prices[i] - stocks.Prices[i-1]
	}

	stocks.Patterns = make(map[uint32]int8, steps)
	for i := steps - 1; i > 3; i-- {
		key := Key(
			stocks.Deltas[i-3],
			stocks.Deltas[i-2],
			stocks.Deltas[i-1],
			stocks.Deltas[i],
		)
		stocks.Patterns[key] = stocks.Prices[i]
	}

	return stocks
}

func Key(a, b, c, d int8) uint32 {
	return uint32(uint8(a)) | uint32(uint8(b))<<8 | uint32(uint8(c))<<16 | uint32(uint8(d))<<24
}
