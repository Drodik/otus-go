package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(inputString string) []string {
	if inputString == "" {
		return []string{}
	}

	frequencyMap := make(map[string]int)

	elements := strings.Fields(inputString)

	result := make([]string, 0, len(elements))

	for _, item := range elements {
		if frequencyMap[item] == 0 {
			result = append(result, item)
		}

		frequencyMap[item]++
	}

	sort.Slice(result, func(i, j int) bool {
		if frequencyMap[result[i]] == frequencyMap[result[j]] {
			return result[i] < result[j]
		}

		return frequencyMap[result[i]] > frequencyMap[result[j]]
	})

	return result[:10]
}
