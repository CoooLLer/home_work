package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

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

	var topSlice []entry

	for word, count := range top {
		topSlice = append(topSlice, entry{count, word})
	}

	sort.Slice(topSlice, func(i, j int) bool {
		if topSlice[i].count == topSlice[j].count {
			return topSlice[i].word < topSlice[j].word
		}
		return topSlice[i].count > topSlice[j].count
	})

	if len(topSlice) > 10 {
		topSlice = topSlice[:10]
	}

	var result []string

	for _, entry := range topSlice {
		result = append(result, entry.word)
	}

	return result
}
