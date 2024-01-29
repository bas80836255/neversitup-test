package odd_number

type OddNumber struct {
}

func (o *OddNumber) Find(input []int) int {
	mapResult := make(map[int]int)

	for _, n := range input {
		mapResult[n] = mapResult[n] + 1
	}

	for i, v := range mapResult {
		if v%2 != 0 {
			return i
		}
	}

	return 0
}
