package odd_number

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type OddNumberTest struct {
	suite.Suite
	oddNumber OddNumber
}

func (t *OddNumberTest) SetupTest() {
	t.oddNumber = OddNumber{}
}

func TestOddNumberTestSuite(t *testing.T) {
	suite.Run(t, new(OddNumberTest))
}

// input [7]
func (t *OddNumberTest) TestOddNumber_Find_WhenGivenInputCase1_ThenReturnSeven() {
	input := []int{7}
	result := t.oddNumber.Find(input)

	t.Equal(7, result)
}

// input [0]
func (t *OddNumberTest) TestOddNumber_Find_WhenGivenInputCase2_ThenReturnZero() {
	input := []int{0}
	result := t.oddNumber.Find(input)

	t.Equal(0, result)
}

// input [1,1,2]
func (t *OddNumberTest) TestOddNumber_Find_WhenGivenInputCase3_ThenReturnTwo() {
	input := []int{1, 1, 2}
	result := t.oddNumber.Find(input)

	t.Equal(2, result)
}

// input [0,1,0,1,0]
func (t *OddNumberTest) TestOddNumber_Find_WhenGivenInputCase4_ThenReturnZero() {
	input := []int{0, 1, 0, 1, 0}
	result := t.oddNumber.Find(input)

	t.Equal(0, result)
}

// input [1,2,2,3,3,3,4,3,3,3,2,2,1]
func (t *OddNumberTest) TestOddNumber_Find_WhenGivenInputCase5_ThenReturnFour() {
	input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	result := t.oddNumber.Find(input)

	t.Equal(4, result)
}
