package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zaccone/goTrie"
)

var program = os.Args[0]

var dictionary string
var prefix string
var countWords bool

func init() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s [-flags] prefix\n", program)
		flag.PrintDefaults()
	}
	flag.StringVar(&dictionary, "dict", "/usr/share/dict/words",
		"Path to the file with words, one word per line. Defaults to /usr/share/dict/words")

	// option -c/--count
	flag.BoolVar(&countWords, "count", false,
		"Show number of matching words only.")
	flag.BoolVar(&countWords, "c", false,
		"Show number of matching words only.")

	flag.Parse()

	prefix = flag.Arg(0)

	if prefix == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func buildFromFile(dictionary string) *goTrie.Trie {
	trie := goTrie.New()
	if trie == nil {
		log.Fatal("Error while creating Trie object.")
	}

	file, err := os.Open(dictionary)

	if err != nil {
		log.Fatal(err)

	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		trie.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading dictionary: ", err)
	}

	if err := file.Close(); err != nil {
		log.Printf("Error while closing dictionary file: %s\n", err)
	}

	return trie

}

func countMatchingWords(root *goTrie.Trie, prefix string) {
	node := root.Get(prefix)
	fmt.Println(node.Children())
}

func printMatchingWords(root *goTrie.Trie, prefix string) {
	matchingWords := root.GetWordsFromPrefix(prefix)
	for _, word := range matchingWords {
		fmt.Println(word)
	}
}

func main() {
	root := buildFromFile(dictionary)
	if countWords {
		countMatchingWords(root, prefix)
	} else {
		printMatchingWords(root, prefix)
	}
}
