package main

import (
	"fmt"
	"strings"
)

/** sample from Alek's email */
/** expected result: [[cig fig fog dog dug mug mut cut cot]] */
// var beginWord = "cig"
// var endWord = "cot"
// var wordList = []string{"cot", "dog", "cat", "cut", "mug", "fog", "fig", "mut", "dug"}

/** README Example 1 */
/** expected result: [[hit hot dot dog cog] [hit hot lot log cog]] */
var beginWord = "hit"
var endWord = "cog"
var wordList = []string{"hot", "dot", "dog", "lot", "log", "cog"}

/** README Example 2 */
/** expected result: [[lost cost]] */
// var beginWord = "lost"
// var endWord = "cost"
// var wordList = []string{"most", "fost", "cost", "host", "lost"}

/** README Example 3 */
/** expected result: [] */
// var beginWord = "start"
// var endWord = "endit"
// var wordList = []string{"stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"}

const VERBOSE = false

func main() {
	fmt.Printf("beginWord: %s\n", beginWord)
	fmt.Printf("  endWord: %s\n", endWord)
	fmt.Printf(" wordList: %v\n", wordList)
	/* wordLadder := */ buildWordLadder(beginWord, endWord, wordList)
	// emitWordLadder(wordLadder, wordList)
	fmt.Printf("\nladder(s): %v\n", treeLadders)
}

func buildWordLadder(beginWord string, endWord string, wordList []string) [][]string {

	beginWordLC := strings.ToLower(beginWord)
	endWordLC := strings.ToLower(endWord)

	uniqueWords := make(map[string]string, 0)

	uniqueWords[beginWordLC] = beginWord
	uniqueWords[endWordLC] = endWord

	cleansedWordListLC := make([]string, 0)

	for _, wordListItem := range wordList {
		wordListItemLC := strings.ToLower(wordListItem)
		if wordListItemLC != beginWordLC && wordListItemLC != endWordLC {
			_, already := uniqueWords[wordListItemLC]
			if !already {
				uniqueWords[wordListItemLC] = wordListItem
				cleansedWordListLC = append(cleansedWordListLC, wordListItemLC)
			}
		}
	}

	if VERBOSE {
		fmt.Printf("uniqueWords: %v]\n", uniqueWords)
	}

	/*
		wordLadder := make([]string, 0)

		wordLadder = append(wordLadder, beginWord)
	*/

	/* ladderSteps := */
	buildLadderSteps(beginWordLC, endWordLC, cleansedWordListLC)
	if VERBOSE {
		fmt.Printf("\nladder(s): %v\n", treeLadders)
	}
	originalWordLadders := make([][]string, 0)
	for _, ladder := range treeLadders {
		restoredCaseWordLadder := make([]string, 0)
		for _, rungWord := range ladder {
			_, exists := uniqueWords[rungWord]
			if !exists {
				fmt.Printf("Unique word not found: %s", rungWord)
			}
			restoredCaseWordLadder = append(restoredCaseWordLadder, uniqueWords[rungWord])
		}
		originalWordLadders = append(originalWordLadders, restoredCaseWordLadder)
	}
	treeLadders = originalWordLadders

	/*
		for _, stepWord := range ladderSteps {
			wordLadder = append(wordLadder, uniqueWords[stepWord])
		}

		wordLadder = append(wordLadder, endWord)
	*/

	return treeLadders // wordLadder
}

func buildLadderSteps(beginWord string, endWord string, wordList []string) [][]string {

	// ladderSteps := make([]string, 0)

	// shortestCandidateSteps := make([][]string, 0)

	if !isOneLetterDiff(beginWord, endWord) {

		/* candidateSteps := */
		buildNextCandidateSteps(beginWord, endWord, wordList)

		/*
			if len(shortestCandidateSteps) == 0 {
				shortestCandidateSteps = append(shortestCandidateSteps, candidateSteps)
			} else {

				if len(candidateSteps) == len(shortestCandidateSteps[0]) {
					shortestCandidateSteps = append(shortestCandidateSteps, candidateSteps)
				} else if len(candidateSteps) < len(shortestCandidateSteps[0]) {
					shortestCandidateSteps = [][]string{candidateSteps}
				}
			}
		*/

	} else {
		treeLadders = [][]string{{beginWord, endWord}}
	}

	/*
		if len(shortestCandidateSteps) > 0 {
			return shortestCandidateSteps[0]
		}
	*/

	return treeLadders // ladderSteps
}

type Step struct {
	stepWord string

	nextSteps []Step
}

var wordTree = Step{"", nil}

var treeLadders [][]string

func buildNextCandidateSteps(beginWord string, endWord string, wordList []string) [][]string {

	if wordTree.stepWord == "" {

		// build the word tree

		wordTree.stepWord = beginWord

		addNextSteps(&wordTree, endWord, wordList)

		if VERBOSE {
			fmt.Printf("wordTree: %v\n", wordTree)
		}

		treeLadders = make([][]string, 0)

		getTreeLadders(&treeLadders, wordTree, endWord)

		if VERBOSE {
			for i, ladder := range treeLadders {
				fmt.Printf("Ladder[%d]: %v\n", i, ladder)
			}
		}
	}

	return treeLadders // []string{}
}

func getTreeLadders(ladders *[][]string, root Step, endWord string) {

	for _, nextStep := range root.nextSteps {
		ladder := []string{root.stepWord}
		getLadder(ladders, ladder, nextStep, endWord)
	}

}

func getLadder(ladders *[][]string, ladder []string, step Step, endWord string) {

	if len(step.nextSteps) == 0 {
		if step.stepWord == endWord {
			ladder = append(ladder, endWord)

			if len(*ladders) == 0 {
				*ladders = append(*ladders, ladder)
			} else if len(ladder) < len((*ladders)[0]) {
				*ladders = [][]string{ladder}
			} else if len(ladder) == len((*ladders)[0]) {
				*ladders = append(*ladders, ladder)
			}
		}
		return
	}

	ladder = append(ladder, step.stepWord)
	for _, nextStep := range step.nextSteps {
		getLadder(ladders, ladder, nextStep, endWord)
	}

}

func addNextSteps(step *Step, endWord string, wordList []string) {

	if VERBOSE {
		fmt.Printf("%v %v\n", step, wordList)
	}

	// check to see if we can get directly to the end word from the current step
	if isOneLetterDiff(step.stepWord, endWord) {
		// last step to the end word
		if VERBOSE {
			fmt.Printf("%s => %s [END]\n", step.stepWord, endWord)
		}
		// build the last step
		lastStep := Step{endWord, nil}
		// the one-and-only next step is the last one
		(*step).nextSteps = []Step{lastStep}
		return
	}

	//
	// we can't get directly to the end word,
	// so find the words that will work as an interim step
	// and fill out their paths of all possible subsequent steps
	//
	for i, word := range wordList {

		// will the word from the list work as a next step from the current step
		if isOneLetterDiff(step.stepWord, word) {

			if VERBOSE {
				fmt.Printf("%s => %s\n", step.stepWord, word)
			}

			// build our next step
			nextStep := Step{word, nil}

			// extract the next step word from the wordList
			// to get remaining candidates for subsequent steps
			// on this particular path
			nextStepRemainingWords := append(append(make([]string, 0), wordList[:i]...), wordList[i+1:]...)

			// find all the subsequent step paths
			addNextSteps(&nextStep, endWord, nextStepRemainingWords)

			// add this next step to the prior parent step
			if step.nextSteps == nil {
				(*step).nextSteps = []Step{nextStep}
			} else {
				(*step).nextSteps = append(step.nextSteps, nextStep)
			}

		} else {
			if VERBOSE {
				fmt.Printf("%s <> %s\n", step.stepWord, word)
			}
		}
	}
}

// determines if only one letter is different between two words
func isOneLetterDiff(word1 string, word2 string) bool {

	if word1 == word2 {
		return false
	}

	totalDifferentLetters := 0

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			totalDifferentLetters++

			if totalDifferentLetters > 1 {
				// no need to keep looking
				break
			}
		}
	}

	return totalDifferentLetters == 1
}

func emitWordLadder(wordLadder []string, wordList []string) {
	// hoping to have some fun "animating" the output

	for _, stepWord := range wordLadder {

		if VERBOSE {
			fmt.Println(stepWord)
		}
	}
}
