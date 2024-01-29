package manipulator

import "fmt"

type Manipulator struct {
}

//func (m *Manipulator) Shuffle(input string) []string {
//	var result []string
//	m.swap([]rune(input), 0, len(input)-1, &result)
//	return m.removeDuplicate(result)
//}
//
//func (m *Manipulator) swap(input []rune, left, right int, result *[]string) {
//	fmt.Println("left : ", left, " right: ", right, " input", string(input))
//	if left == right {
//		*result = append(*result, string(input))
//	} else {
//		fmt.Println("recursive")
//		for i := left; i <= right; i++ {
//			fmt.Println("input", string(input))
//			input[left], input[i] = input[i], input[left]
//			fmt.Println("before ", string(input))
//			m.swap(input, left+1, right, result)
//			input[left], input[i] = input[i], input[left]
//			fmt.Println("after ", string(input))
//		}
//	}
//}

func (m *Manipulator) Shuffle(input string) []string {
	result := m.swap([]rune(input), 0, len(input)-1)
	return m.removeDuplicate(result)
}

func (m *Manipulator) swap(input []rune, left, right int) []string {
	var result []string
	fmt.Println("left : ", left, " right: ", right, " input", string(input))
	if left == right {
		result = append(result, string(input))
	} else {
		fmt.Println("recursive")
		for i := left; i <= right; i++ {
			fmt.Println("input", string(input))
			input[left], input[i] = input[i], input[left]
			fmt.Println("before ", string(input))
			result = append(result, m.swap(input, left+1, right)...)
			input[left], input[i] = input[i], input[left]
			fmt.Println("after ", string(input))
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
