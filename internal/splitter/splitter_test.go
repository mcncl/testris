package splitter

import (
	"testing"

	"github.com/mcncl/testris/internal/finder"
)

func TestGenerateTestPattern(t *testing.T) {
	t.Parallel()
	tests := []finder.TestInfo{
		{Name: "TestOne", Package: "example"},
		{Name: "TestTwo", Package: "example"},
		{Name: "TestThree", Package: "example"},
		{Name: "TestFour", Package: "example"},
	}

	testCases := []struct {
		name     string
		index    int
		total    int
		expected string
	}{
		{
			name:     "split 4 tests into 2 groups - first half",
			index:    0,
			total:    2,
			expected: "^(TestOne|TestTwo)$",
		},
		{
			name:     "split 4 tests into 2 groups - second half",
			index:    1,
			total:    2,
			expected: "^(TestThree|TestFour)$",
		},
		{
			name:     "all tests in one group",
			index:    0,
			total:    1,
			expected: "^(TestOne|TestTwo|TestThree|TestFour)$",
		},
		{
			name:     "empty pattern for out of range index",
			index:    2,
			total:    2,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pattern := GenerateTestPattern(tests, tc.index, tc.total)
			if pattern != tc.expected {
				t.Errorf("Expected pattern %q, got %q", tc.expected, pattern)
			}
		})
	}
}
