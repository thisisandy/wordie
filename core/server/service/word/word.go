package word

import (
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"strconv"
	"wordie/core/db"
	"wordie/core/db/wordDB"
)

func removeDuplicate(words []string) []string {
	m := make(map[string]bool)
	for _, word := range words {
		m[word] = true
	}
	uniqueWords := make([]string, 0)
	seen := make(map[string]bool)
	for _, word := range words {
		//if the word is number, we don't need to check it
		if _, err := strconv.Atoi(word); err == nil {
			continue
		}
		if !seen[word] {
			seen[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}
	return uniqueWords
}

func FilterWordsWithSmallerFrequency(words []string, frequency int) []string {
	legitimatize, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	// deduplicate the words
	// Remove duplication words
	words = removeDuplicate(words)
	originalWords := make([]string, len(words))
	// legitimatize the words one by one
	for i, word := range words {
		originalWords[i] = legitimatize.Lemma(word)
	}
	database := db.Instance()
	wordsFrequency, _ := wordDB.GetFrequencies(database, originalWords)
	// filter the words with smaller frequency
	res := make([]string, 0)
	for i := range originalWords {
		if wordsFrequency[i] < frequency {
			res = append(res, words[i])
		}
	}
	return res
}
