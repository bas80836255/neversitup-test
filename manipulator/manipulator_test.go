package manipulator

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ManipulatorTest struct {
	suite.Suite
	manipulator Manipulator
}

func (t *ManipulatorTest) SetupTest() {
	t.manipulator = Manipulator{}
}

func TestManipulatorTestSuite(t *testing.T) {
	suite.Run(t, new(ManipulatorTest))
}

// input "a"
func (t *ManipulatorTest) TestManipulator_Shuffle_WhenInputCase1_ThenReturnArraySizeOne() {
	input := "a"
	result := t.manipulator.Shuffle(input)

	expect := []string{"a"}

	t.Equal(1, len(result))
	t.ElementsMatch(expect, result)
}

// input "ab"
func (t *ManipulatorTest) TestManipulator_Shuffle_WhenInputCase2_ThenReturnArraySizeTwo() {
	input := "ab"
	result := t.manipulator.Shuffle(input)

	expect := []string{"ab", "ba"}

	t.Equal(2, len(result))
	t.ElementsMatch(expect, result)
}

// input "abc"
func (t *ManipulatorTest) TestManipulator_Shuffle_WhenInputCase3_ThenReturnArraySizeSix() {
	input := "abc"
	result := t.manipulator.Shuffle(input)

	expect := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

	t.Equal(6, len(result))
	t.ElementsMatch(expect, result)
	//t.Equal(expect, result)
}

// input "aabb"
func (t *ManipulatorTest) TestManipulator_Shuffle_WhenInputCase4_ThenReturnArraySizeSix() {
	input := "aabb"
	result := t.manipulator.Shuffle(input)

	expect := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}

	t.Equal(6, len(result))
	t.ElementsMatch(expect, result)
}
