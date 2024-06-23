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
	treeLadders := sanitizeInputAndFindWordLadders(beginWord, endWord, wordList)
	// emitWordLadder(wordLadder, wordList)
	fmt.Printf("\nladder(s): %v\n", treeLadders)
}

func sanitizeInputAndFindWordLadders(beginWord string, endWord string, wordList []string) [][]string {

	if len(beginWord) != len(endWord) {
		panic("begin/end length mismatch")
	}

	beginWordLC := strings.ToLower(beginWord)
	endWordLC := strings.ToLower(endWord)

	cleansedWordListLC := make([]string, 0)
	originalCase := make(map[string]string, 0)

	{
		uniqueWords := make(map[string]string, 0)

		uniqueWords[beginWordLC] = beginWord
		uniqueWords[endWordLC] = endWord

		for _, wordListItem := range wordList {

			if len(wordListItem) != len(beginWord) {
				panic(fmt.Sprintf("wordList item `%s` vs begin/end length mismatch", wordListItem))
			}
			wordListItemLC := strings.ToLower(wordListItem)
			if wordListItemLC != beginWordLC && wordListItemLC != endWordLC {
				_, already := uniqueWords[wordListItemLC]
				if !already {
					uniqueWords[wordListItemLC] = wordListItem
					if wordListItemLC != wordListItem {
						originalCase[wordListItemLC] = wordListItem
					}
					cleansedWordListLC = append(cleansedWordListLC, wordListItemLC)
				} else {
					if VERBOSE {
						fmt.Printf("duplicate wordList item `%s`\n", wordListItem)
					}
				}
			} else {
				if VERBOSE {
					fmt.Printf("wordList item `%s` matched beginWord or endWord\n", wordListItem)
				}
			}
		}

		if VERBOSE {
			fmt.Printf("uniqueWords: %v\n", uniqueWords)
			fmt.Printf("originalCase: %v\n", originalCase)
		}

		uniqueWords = nil
	}

	treeLadders := findWordLadders(beginWordLC, endWordLC, cleansedWordListLC)

	if len(originalCase) > 0 {
		if VERBOSE {
			fmt.Printf("\nLC ladder(s): %v\n", treeLadders)
		}

		originalWordLadders := make([][]string, 0)
		for _, ladder := range treeLadders {
			restoredCaseWordLadder := make([]string, 0)
			for _, rungWord := range ladder {
				_, exists := originalCase[rungWord]
				if exists {
					restoredCaseWordLadder = append(restoredCaseWordLadder, originalCase[rungWord])
				} else {
					restoredCaseWordLadder = append(restoredCaseWordLadder, rungWord)
				}

			}
			originalWordLadders = append(originalWordLadders, restoredCaseWordLadder)
		}
		treeLadders = originalWordLadders
	}

	return treeLadders
}

func findWordLadders(beginWord string, endWord string, wordList []string) [][]string {

	var treeLadders [][]string
	if !isOneLetterDiff(beginWord, endWord) {

		treeLadders = buildNextCandidateSteps(beginWord, endWord, wordList)

	} else {
		treeLadders = [][]string{{beginWord, endWord}}
	}

	return treeLadders
}

type Step struct {
	stepWord string

	nextSteps []Step
}

func buildNextCandidateSteps(beginWord string, endWord string, wordList []string) [][]string {

	wordTree := Step{beginWord, nil}

	addNextSteps(&wordTree, endWord, wordList)

	if VERBOSE {
		fmt.Printf("wordTree: %v\n", wordTree)
	}

	treeLadders := make([][]string, 0)

	getTreeLadders(&treeLadders, wordTree, endWord)

	if VERBOSE {
		for i, ladder := range treeLadders {
			fmt.Printf("Ladder[%d]: %v\n", i, ladder)
		}
	}

	return treeLadders
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
				// no need to keep looking- any amount more than one is enough
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
