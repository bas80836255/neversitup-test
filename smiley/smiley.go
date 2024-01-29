package odd_number

import "fmt"

type Smiley struct {
}

func (s *Smiley) Count(input []string) int {
	result := 0
	for i, t := range input {
		isSmiley := false
		switch len(t) {
		case 2:
			isSmiley = s.eyeMouth(t)
			break
		case 3:
			isSmiley = s.eyeNoseMouth(t)
			break
		default:

		}
		fmt.Println("main loop :", i, " char : ", t, isSmiley)

		if isSmiley {
			result++
		}

	}

	return result
}

func (s *Smiley) eyeMouth(input string) bool {
	r := []rune(input)
	return s.eye(r[0]) && s.mouth(r[1])
}

func (s *Smiley) eyeNoseMouth(input string) bool {
	r := []rune(input)
	return s.eye(r[0]) && s.nose(r[1]) && s.mouth(r[2])
}

func (s *Smiley) eye(r rune) bool {
	return r == ':' || r == ';'
}
func (s *Smiley) nose(r rune) bool {
	return r == '-' || r == '~'
}

func (s *Smiley) mouth(r rune) bool {
	return r == ')' || r == 'D'
}
