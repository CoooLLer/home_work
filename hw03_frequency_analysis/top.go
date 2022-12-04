package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const TopCount = 10

type entry struct {
	count int
	word  string
}

func Top10(str string) []string {
	slice := strings.Fields(str)
	top := map[string]int{}

	for _, word := range slice {
		if _, isSet := top[word]; !isSet {
			top[word] = 1
		} else {
			top[word]++
		}
	}

	topSlice := make([]entry, len(top))

	index := 0
	for word, count := range top {
		topSlice[index] = entry{count, word}
		index++
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

	result := make([]string, len(topSlice))

	for pos, entry := range topSlice {
		result[pos] = entry.word
	}

	return result
}
