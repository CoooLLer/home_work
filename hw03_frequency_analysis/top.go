package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const TopCount = 10

type topEntry struct {
	count int
	word  string
}

func Top10(str string) []string {
	slice := strings.Fields(str)
	top := map[string]int{}

	for _, word := range slice {
		top[word]++
	}

	topSlice := make([]topEntry, 0, len(top))

	for word, count := range top {
		topSlice = append(topSlice, topEntry{count, word})
	}

	sort.Slice(topSlice, func(i, j int) bool {
		if topSlice[i].count == topSlice[j].count {
			return topSlice[i].word < topSlice[j].word
		}
		return topSlice[i].count > topSlice[j].count
	})

	if len(topSlice) > TopCount {
		topSlice = topSlice[:TopCount]
	}

	result := make([]string, 0, cap(topSlice))

	for _, entry := range topSlice {
		result = append(result, entry.word)
	}

	return result
}
