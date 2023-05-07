package word_processor

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Children  map[rune]*Node
	frequency int
}

type Trie struct {
	Root *Node
	add  func(string, int)
}

func (t *Trie) Add(word string, freq int) {
	node := t.Root
	for _, r := range word {
		if _, ok := node.Children[r]; !ok {
			node.Children[r] = &Node{}
		}
		node = node.Children[r]
	}
	node.frequency = freq
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{},
	}
}

func readVocalFile() (*os.File, error) {
	file, err := os.Open("./source/vocal.txt")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func parseLine(line string) (string, int, error) {
	parts := strings.Split(line, "\t")
	word := parts[0]
	freq, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}
	return word, freq, nil
}

func loadWordsToTrie(trie *Trie, file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		word, freq, err := parseLine(line)
		if err != nil {
			return err
		}
		trie.Add(word, freq)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func loadVocalToTrie(trie *Trie) error {
	file, err := readVocalFile()
	if err != nil {
		return err
	}
	defer file.Close()
	return loadWordsToTrie(trie, file)
}

func BuildTree() *Trie {
	// Create a new trie tree and load vocal.txt into it
	trie := NewTrie()

	if err := loadVocalToTrie(trie); err != nil {
		log.Fatal(err)
	}
	return trie
}
