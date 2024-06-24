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
	if VERBOSE {
		fmt.Printf("beginWord: %s\n", beginWord)
		fmt.Printf("  endWord: %s\n", endWord)
		fmt.Printf(" wordList: %v\n", wordList)
	}
	shortestWordLadders := sanitizeInputAndFindShortestWordLadders(beginWord, endWord, wordList)
	// emitWordLadder(wordLadder, wordList)
	if VERBOSE {
		fmt.Printf("\nladder(s): %v\n", shortestWordLadders)
	} else {
		fmt.Printf("%v\n", shortestWordLadders)
	}
}

// this method was conceived and mostly completed
// *before* I'd fully read the spec
// which theoretically makes this all obsolete
//
// converts all input to lowercase,
// validates consistent word lengths,
// ensures unique items in word list
func sanitizeInputAndFindShortestWordLadders(beginWord string, endWord string, wordList []string) [][]string {

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

		if beginWordLC != beginWord {
			originalCase[beginWordLC] = beginWord
		}

		if endWordLC != endWord {
			originalCase[endWordLC] = endWord
		}

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

	shortestWordLadders := findShortestWordLadders(beginWordLC, endWordLC, cleansedWordListLC)

	if len(originalCase) > 0 {

		// if the original case of words was different,
		// restore the original case for output

		if VERBOSE {
			fmt.Printf("\nLC ladder(s): %v\n", shortestWordLadders)
		}

		originalWordLadders := make([][]string, 0)
		for _, ladder := range shortestWordLadders {
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
		shortestWordLadders = originalWordLadders
	}

	return shortestWordLadders
}

// find the shortest word ladders
func findShortestWordLadders(beginWord string, endWord string, wordList []string) [][]string {

	var shortestWordLadders [][]string
	if !isOneLetterDiff(beginWord, endWord) {
		// standard sitch- can't get directly from beginWord to endWord,
		// see if we can get there by utilizing the word list
		shortestWordLadders = buildShortestLaddersFromStepPaths(beginWord, endWord, wordList)

	} else {
		// special case- no interim steps needed to get from beginWord to endWord
		shortestWordLadders = [][]string{{beginWord, endWord}}
	}

	return shortestWordLadders
}

// the Step structure is used to build out the possible paths
// to determine if they can lead us to the endWord
type Step struct {
	stepWord string

	nextSteps []Step
}

// builds out the step paths,
// then extracts the shortest ladders from those paths
func buildShortestLaddersFromStepPaths(beginWord string, endWord string, wordList []string) [][]string {

	wordTree := Step{beginWord, nil}

	buildStepPaths(&wordTree, beginWord, endWord, wordList)

	if VERBOSE {
		fmt.Printf("wordTree: %v\n", wordTree)
	}

	shortestWordLadders := make([][]string, 0)

	getShortestLaddersFromWordTree(&shortestWordLadders, wordTree, endWord)

	if VERBOSE {
		for i, ladder := range shortestWordLadders {
			fmt.Printf("Ladder[%d]: %v\n", i, ladder)
		}
	}

	return shortestWordLadders
}

// initialize each ladder from the root, the build out the ladder variations from step paths
func getShortestLaddersFromWordTree(shortestWordLadders *[][]string, root Step, endWord string) {

	for _, nextStep := range root.nextSteps {
		ladder := []string{root.stepWord}
		getShortestLaddersFromStepPaths(shortestWordLadders, ladder, nextStep, endWord)
	}

}

// recursive method builds out ladders from step paths;
// if a path/ladder culminates in the end word,
// and it is as short or shorter than any prior ladders, save it off
func getShortestLaddersFromStepPaths(shortestWordLadders *[][]string, ladder []string, step Step, endWord string) {

	if len(step.nextSteps) == 0 {
		if step.stepWord == endWord {
			ladder = append(ladder, endWord)

			if len(*shortestWordLadders) == 0 {
				// first valid ladder
				*shortestWordLadders = append(*shortestWordLadders, ladder)
			} else if len(ladder) < len((*shortestWordLadders)[0]) {
				// new ladder is shortest,
				// replace all prior ladders
				*shortestWordLadders = [][]string{ladder}
			} else if len(ladder) == len((*shortestWordLadders)[0]) {
				// new ladder is same length as prior shortest ladders,
				// add it to the list
				*shortestWordLadders = append(*shortestWordLadders, ladder)
			}
		}
		return
	}

	// we didn't have a step to endWord, keep building out the ladder for this path
	ladder = append(ladder, step.stepWord)
	for _, nextStep := range step.nextSteps {
		getShortestLaddersFromStepPaths(shortestWordLadders, ladder, nextStep, endWord)
	}

}

// recursive method builds out the different subsequent step paths from current step
func buildStepPaths(step *Step, beginWord string, endWord string, wordList []string) {

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
	extended := 0
	for i, word := range wordList {

		if word == beginWord || word == endWord {
			// NOTE endWord special case is handled first in this method above
			continue
		}

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
			buildStepPaths(&nextStep, beginWord, endWord, nextStepRemainingWords)

			// add this next step to the prior parent step
			if step.nextSteps == nil {
				(*step).nextSteps = []Step{nextStep}
			} else {
				(*step).nextSteps = append(step.nextSteps, nextStep)
			}

			extended++

		} else {
			if VERBOSE {
				fmt.Printf("%s <> %s\n", step.stepWord, word)
			}
		}
	}

	if len(wordList) == 0 || len(step.nextSteps) == 0 || extended == 0 {
		// no more steps available, we can't get from last word to end word- DEAD END
		fmt.Printf("DEAD END: %v\n", step)
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
