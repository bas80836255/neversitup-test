package manipulator

type Manipulator struct {
}

func (m *Manipulator) Shuffle(input string) []string {
	result := m.swap([]rune(input), 0, len(input)-1)
	return m.removeDuplicate(result)
}

func (m *Manipulator) swap(input []rune, left, right int) []string {
	var result []string
	if left == right {
		result = append(result, string(input))
	} else {
		for i := left; i <= right; i++ {
			input[left], input[i] = input[i], input[left]
			result = append(result, m.swap(input, left+1, right)...)
			input[left], input[i] = input[i], input[left]
		}
	}
	return result
}

func (m *Manipulator) removeDuplicate(result []string) []string {
	mapValue := make(map[string]bool)
	var newResult []string

	for _, s := range result {
		if mapValue[s] {
			continue
		}

		mapValue[s] = true
		newResult = append(newResult, s)
	}

	return newResult
}
