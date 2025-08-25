package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type FileRank struct {
	name string
	rank float64
}

func onErr(err any) {
	fmt.Println(err)
	os.Exit(1)
}

func getParams() (string, string) {
	inputs := os.Args[1:]

	if len(inputs) != 2 {
		onErr("Incorrect number of arguments passed")
	}

	searchTerm := inputs[0]
	path := inputs[1]

	return searchTerm, path
}

func main() {
	searchTerm, path := getParams()

	files := getFiles(path)

	ranks := getFileRanks(searchTerm, files)

	filteredRanks := filterFiles(ranks)

	sortedRanks := sortFiles(filteredRanks)

	printFiles(sortedRanks)
}

func getFiles(path string) []os.DirEntry {
	files, err := os.ReadDir(path)

	if err != nil {
		onErr(err)
	}

	return files
}

func getFileRanks(searchTerm string, files []os.DirEntry) []FileRank {
	rankings := []FileRank{}

	for _, file := range files {
		rankings = append(rankings, getRank(searchTerm, file.Name()))
	}

	return rankings
}

// Returns a ranking between 0 and 1. The higher the number the closer the match is.
func getRank(searchTerm string, name string) FileRank {
	letters := strings.SplitSeq(name, "")

	matchedLetterCount := 0

	highlightedName := ""

	for letter := range letters {
		if strings.Contains(searchTerm, letter) {
			matchedLetterCount++
			highlightedName += "\033[31m" + letter + "\033[97m"
		} else {
			highlightedName += letter
		}
	}

	return FileRank{highlightedName, float64(matchedLetterCount) / float64(len(name))}
}

// Sort the files so that the closest match is at the bottom.
func sortFiles(ranks []FileRank) []FileRank {
	slices.SortFunc(ranks, func(a, b FileRank) int {
		if a.rank > a.rank {
			return 1
		} else if a.rank < b.rank {
			return -1
		} else {
			return 0
		}
	})

	return ranks
}

func filterFiles(ranks []FileRank) []FileRank {
	filteredRanks := []FileRank{}

	for _, rank := range ranks {
		if rank.rank != 0 {
			filteredRanks = append(filteredRanks, rank)
		}
	}

	return filteredRanks
}

func printFiles(ranks []FileRank) {
	fmt.Println("Confidence\tName")

	for _, rank := range ranks {
		fmt.Printf("%.6f\t%s\n", rank.rank, rank.name)
	}
}
