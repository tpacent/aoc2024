package day22

func NthStep(n uint32, count int) uint32 {
	for range count {
		n = Step(n)
	}

	return n
}

const pruneNumber = 16777216

func Step(n uint32) uint32 {
	n = (n*64 ^ n) % pruneNumber
	n = (n/32 ^ n) % pruneNumber
	n = (n*2048 ^ n) % pruneNumber
	return n
}

type Stocks struct {
	Deltas   []int8
	Prices   []int8
	Patterns map[[4]int8]int8
}

func MaxPattern(secrets []uint32, steps int) (pattern [4]int8, maxValue int) {
	patterns := make(map[[4]int8]int, steps)

	for _, secret := range secrets {
		for key, value := range CalcStocks(secret, steps).Patterns {
			patterns[key] += int(value)
		}
	}

	for key, value := range patterns {
		if value > maxValue {
			pattern = key
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

	stocks.Patterns = make(map[[4]int8]int8, steps)
	for i := 1; i < steps-3; i++ {
		key := [4]int8{
			stocks.Deltas[i],
			stocks.Deltas[i+1],
			stocks.Deltas[i+2],
			stocks.Deltas[i+3],
		}

		if _, ok := stocks.Patterns[key]; ok {
			continue
		}

		stocks.Patterns[key] = stocks.Prices[i+3]
	}

	return stocks
}
