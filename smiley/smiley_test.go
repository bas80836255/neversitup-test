package odd_number

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SmileyTest struct {
	suite.Suite
	smiley Smiley
}

func (t *SmileyTest) SetupTest() {
	t.smiley = Smiley{}
}

func TestOddNumberTestSuite(t *testing.T) {
	suite.Run(t, new(SmileyTest))
}

// input [':)', ';(', ';}', ':-D']
func (t *SmileyTest) TestOddNumber_Find_WhenGivenInputCase1_ThenReturnTwo() {
	input := []string{":)", ";(", ";}", ":-D"}
	result := t.smiley.Count(input)

	t.Equal(2, result)
}

// input [';D', ':-(', ':-)', ';~)']
func (t *SmileyTest) TestOddNumber_Find_WhenGivenInputCase2_ThenReturnThree() {
	input := []string{";D", ":-(", ":-)", ";~)"}
	result := t.smiley.Count(input)

	t.Equal(3, result)
}

// input [';]', ':[', ';*', ':$', ';-D']
func (t *SmileyTest) TestOddNumber_Find_WhenGivenInputCase3_ThenReturnOne() {
	input := []string{";]", ":[", ";*", ":$", ";-D"}
	result := t.smiley.Count(input)

	t.Equal(1, result)
}

// input []
func (t *SmileyTest) TestOddNumber_Find_WhenGivenInputCase4_ThenReturnZero() {
	var input []string
	result := t.smiley.Count(input)

	t.Equal(0, result)
}
