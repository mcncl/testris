package splitter

import (
	"regexp"
	"strings"

	"github.com/mcncl/testris/internal/finder"
)

func GenerateTestPattern(tests []finder.TestInfo, index, total int) string {
	if len(tests) == 0 {
		return ""
	}

	testsPerNode := (len(tests) + total - 1) / total
	start := index * testsPerNode
	end := start + testsPerNode
	if end > len(tests) {
		end = len(tests)
	}

	if start >= len(tests) {
		return ""
	}

	patterns := make([]string, 0, end-start)
	for _, test := range tests[start:end] {
		patterns = append(patterns, regexp.QuoteMeta(test.Name))
	}

	return "^(" + strings.Join(patterns, "|") + ")$"
}
