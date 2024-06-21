package main

var beginWord = "cig"
var endWord = "cot"
var wordList = []string{"cot", "dog", "cat", "cut", "mug", "fog", "fig", "mut", "dug"}

func main() {
	wordLadder := buildWordLadder(beginWord, endWord, wordList)
	emitWordLadder(wordLadder, wordList)
}

func buildWordLadder(beginWord string, endWord string, wordList []string) []string {
	return []string{}
}

func emitWordLadder(wordLadder []string, wordList []string) {
	// hoping to have some fun "animating" the output
}
