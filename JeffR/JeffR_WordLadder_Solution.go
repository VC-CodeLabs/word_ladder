package main

import (
	"cmp"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// var beginWord = "foo"
// var endWord = "bar"
// var wordList = []string{"boo", "bro", "brr", "bar"}

var beginWord = "stone"
var endWord = "money"
var wordList = []string{
	"stoke", "stony", "stome", "stomy", "stoey", "htoey", "htney", "hiney",
	"miney", "ttney", "toney", "itoey", "mtney", "soney", "store", "storm",
	"story", "monte", "monny", "monty", "money", "stane", "stine", "maney",
	"honey", "monde", "stnny", "mtone", "mtnne", "monne", "monee", "stnne", "mtnee"}

// NOTE expectedResult is optional, only used if -c option is spec'd to compare to actual result;
// otherwise, setting it to anything including empty set will have no impact on actual result
var expectedResult = [][]string{
	{"stone", "mtone", "mtnne", "monne", "monee", "money"},
	{"stone", "mtone", "mtnne", "monne", "monny", "money"},
	{"stone", "mtone", "mtnne", "mtnee", "monee", "money"},
	{"stone", "mtone", "mtnne", "mtnee", "mtney", "money"},
	{"stone", "stnne", "mtnne", "monne", "monee", "money"},
	{"stone", "stnne", "mtnne", "monne", "monny", "money"},
	{"stone", "stnne", "mtnne", "mtnee", "monee", "money"},
	{"stone", "stnne", "mtnne", "mtnee", "mtney", "money"}}

// feature flags, controlled by command-line parameters

var VERBOSE = false

var DEBUG = false

const SINGLE_PASS = true

var MULTI_THREADED = false

var OUTPUT_MODE = OUTPUT_PRETTY_PRINT

var INPUT_FILE = ""

var SANITIZE_INPUT = false

var CHECK_RESULT = false

func main() {

	// support numerous command-line parameters

	{ // scoping for vars related to parameter handling

		verboseParamPtr := flag.Bool("v", VERBOSE, "verbose output mode")
		debugParamPtr := flag.Bool("d", DEBUG, "debug output mode")
		beginWordParamPtr := flag.String("b", beginWord, "the beginWord")
		endWordParamPtr := flag.String("e", endWord, "the endWord")

		wordListAsString := //
			strings.Replace(
				strings.Replace(
					strings.ReplaceAll(fmt.Sprintf("%v", wordList), " ", ","),
					"[", "", 1),
				"]", "", 1)
		wordListParamPtr := flag.String("l", wordListAsString, "the wordList")

		expectedResultAsString := //
			strings.Replace(
				strings.Replace(
					strings.ReplaceAll(
						strings.ReplaceAll(fmt.Sprintf("%v", expectedResult), " ", ","),
						"],[", ";"),
					"[", "", 2),
				"]", "", 2)
		expectedResultParamPtr := flag.String("x", expectedResultAsString, "the expectedResult")

		inputFileParamPtr := flag.String("f", INPUT_FILE, "read input from json file (supersedes -b/-e/-l/-x)")
		outputParamPtr := flag.Int("o", OUTPUT_MODE, "output mode # - 0: raw, 1: pretty print, 2: graphical, 3: animated")
		checkParamPtr := flag.Bool("c", CHECK_RESULT, "check results, comparing expectResult/-x to actual result")

		threadedParamPtr := flag.Bool("t", MULTI_THREADED, "multi-threaded support")

		sanitizeParamPtr := flag.Bool("s", SANITIZE_INPUT, "sanitize input- unnecessary if input conforms to spec'd constraints")

		flag.Parse()

		if flag.NFlag() > 0 {
			if verboseParamPtr != nil {
				VERBOSE = *verboseParamPtr
				if VERBOSE {
					fmt.Println("VERBOSE mode enabled.")
				}
			}

			if debugParamPtr != nil {
				DEBUG = *debugParamPtr
				if DEBUG {
					log.Printf("DEBUG mode enabled.")
				}
			}

			if beginWordParamPtr != nil {
				beginWord = *beginWordParamPtr
			}

			if endWordParamPtr != nil {
				endWord = *endWordParamPtr
			}

			if expectedResultParamPtr != nil {
				expectedResult = make([][]string, 0)
				expectedResultAsString = *expectedResultParamPtr
				if len(strings.TrimSpace(expectedResultAsString)) > 0 {
					ladderWordsLists := strings.Split(expectedResultAsString, ";")
					for _, ladderWordsAsString := range ladderWordsLists {
						expectedLadder := strings.Split(ladderWordsAsString, ",")
						if len(expectedLadder) > 0 {
							expectedResult = append(expectedResult, expectedLadder)
						}
					}
				}
			}

			if wordListParamPtr != nil {
				wordListAsString = *wordListParamPtr
				wordList = strings.Split(wordListAsString, ",")
			}

			if outputParamPtr != nil {
				OUTPUT_MODE = *outputParamPtr
				switch OUTPUT_MODE {
				case OUTPUT_RAW:
				case OUTPUT_PRETTY_PRINT:
				case OUTPUT_GRAPHICAL:
				case OUTPUT_ANIMATED:
				default:
					fmt.Println("ERR: -o=<0-3> out of range")
					os.Exit(1)
				}
			}

			if inputFileParamPtr != nil {
				INPUT_FILE = *inputFileParamPtr
				if len(strings.TrimSpace(INPUT_FILE)) > 0 {
					if VERBOSE {
						fmt.Printf("inputFile: %s\n", INPUT_FILE)
					}
					loadJsonInput(INPUT_FILE)

				}
			}

			if checkParamPtr != nil {
				CHECK_RESULT = *checkParamPtr
			}

			if threadedParamPtr != nil {
				MULTI_THREADED = *threadedParamPtr
			}

			if sanitizeParamPtr != nil {
				SANITIZE_INPUT = *sanitizeParamPtr
			}
		}

	} // scoping for vars related to parameter handling

	// command-line flags processed, process the input

	if VERBOSE {
		fmt.Printf("beginWord: %s\n", beginWord)
		fmt.Printf("  endWord: %s\n", endWord)
		fmt.Printf(" wordList: %v\n", wordList)
	}

	var shortestWordLadders [][]string

	if SANITIZE_INPUT {
		shortestWordLadders = sanitizeInputAndFindShortestWordLadders(beginWord, endWord, wordList)
	} else {
		shortestWordLadders = findShortestWordLadders(beginWord, endWord, wordList)
	}

	outputWordLadders(shortestWordLadders, wordList)

	if CHECK_RESULT {
		verifyMatchingLadders(expectedResult, shortestWordLadders)
	}

}

// read inputs beginWord, endWord, wordList and expectedResult from json
func loadJsonInput(inputFile string) {
	inputBytes, err := os.ReadFile(inputFile)

	if err != nil {
		// bad file or access issue
		fmt.Println("ERR: " + err.Error())
		os.Exit(1)
	}

	// a quirk(?) of the json lib requires uppercase-first-char field names;
	// rather than update the "normal" Inputs struct, provide a json-specific version here
	type JsonInputs struct {
		BeginWord      string
		EndWord        string
		WordList       []string
		ExpectedResult [][]string
	}

	var inputs JsonInputs

	err = json.Unmarshal(inputBytes, &inputs)

	if err != nil {
		// some issue with the json
		fmt.Println("ERR: " + err.Error())
		os.Exit(1)

	}

	if DEBUG {
		log.Printf("%s: %v", inputFile, inputs)
	}

	beginWord = inputs.BeginWord
	endWord = inputs.EndWord
	wordList = inputs.WordList
	expectedResult = inputs.ExpectedResult
	// fmt.Printf("ExpectedLadders: %v\n", inputs.ExpectedLadders)

}

// trust, but verify
//
// converts all input to lowercase,
// validates consistent word lengths,
// ensures unique items in word list
func sanitizeInputAndFindShortestWordLadders(beginWord string, endWord string, wordList []string) [][]string {

	if len(beginWord) != len(endWord) {
		fmt.Println("ERR: begin/end length mismatch")
		os.Exit(1)
	}

	beginWordLC := strings.ToLower(beginWord)
	endWordLC := strings.ToLower(endWord)

	cleansedWordListLC := make([]string, 0)
	originalCase := make(map[string]string, 0)

	{
		uniqueWords := make(map[string]string, 0)

		uniqueWords[beginWordLC] = beginWord
		// uniqueWords[endWordLC] = endWord

		if beginWordLC != beginWord {
			originalCase[beginWordLC] = beginWord
		}

		/*
			if endWordLC != endWord {
				originalCase[endWordLC] = endWord
			}
		*/

		for _, wordListItem := range wordList {

			if len(wordListItem) != len(beginWord) {
				fmt.Printf("ERR: wordList item `%s` vs begin/end length mismatch\n", wordListItem)
				os.Exit(1)
			}
			wordListItemLC := strings.ToLower(wordListItem)
			if wordListItemLC != beginWordLC /* && wordListItemLC != endWordLC */ {
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

		if DEBUG {
			log.Printf("uniqueWords: %v\n", uniqueWords)
			log.Printf("originalCase: %v\n", originalCase)
		}

		uniqueWords = nil
	}

	if len(cleansedWordListLC) > 5000 {
		fmt.Printf("WARNING: wordList length %d exceeds spec'd maximum\n", len(cleansedWordListLC))
	}

	shortestWordLadders := findShortestWordLadders(beginWordLC, endWordLC, cleansedWordListLC)

	if len(originalCase) > 0 {

		// if the original case of words was different,
		// restore the original case for output

		if DEBUG {
			log.Printf("\nLC ladder(s): %v\n", shortestWordLadders)
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

	sortedWordList := append(make([]string, 0), wordList...)
	slices.Sort(sortedWordList)
	if DEBUG {
		log.Printf("sortedWordList: %v\n", sortedWordList)
	}
	_, endWordMatchFound := slices.BinarySearchFunc(sortedWordList, endWord, func(wordItem string, endWord string) int {
		return strings.Compare(strings.ToLower(wordItem), strings.ToLower(endWord))
	})

	var shortestWordLadders [][]string
	if !endWordMatchFound {
		// endWord must be in the wordList else return empty set!
		// see test case 3
	} else if !isOneLetterDiff(beginWord, endWord) {
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

// optional threading support
var mu sync.Mutex
var wg sync.WaitGroup

// builds out the step paths,
// then extracts the shortest ladders from those paths
func buildShortestLaddersFromStepPaths(beginWord string, endWord string, wordList []string) [][]string {

	shortestWordLadders := make([][]string, 0)

	wordTree := Step{beginWord, nil}

	if SINGLE_PASS {

		// original algorithm built out all the paths, then extracted ladders;
		// it's much more efficient to build ladders and keep/discard based on length as we go

		ladder := []string{wordTree.stepWord}
		paths++ // count the initial path; note we haven't initiated any threading yet

		// sort the list to bring all the words most like both beginWord and endWord to the top;
		// these presumably provide our best chance to build a complete ladder
		slices.SortFunc(wordList, func(a, b string) int {
			return cmp.Compare(wordDiff(a, beginWord, endWord), wordDiff(b, beginWord, endWord))
		})

		if DEBUG {
			log.Printf("weightedWordList: %v\n", wordList)
		}

		if MULTI_THREADED {
			wg.Add(1)
			var rootStepLock sync.Mutex // each step gets its' own lock for updates to it
			go buildStepPathsAndLadders(&shortestWordLadders, ladder, &wordTree, &rootStepLock, beginWord, endWord, wordList)
			wg.Wait()
		} else {
			buildStepPathsAndLadders(&shortestWordLadders, ladder, &wordTree, nil, beginWord, endWord, wordList)
		}

		if VERBOSE {
			fmt.Printf(" wordTree: %v\n", wordTree)
		} else if DEBUG {
			log.Printf("wordTree: %v\n", wordTree)
		}

		if VERBOSE {
			fmt.Printf("  path(s): %d\n", paths)
		} else if DEBUG {
			log.Printf("path(s): %d\n", paths)
		}

		if MULTI_THREADED {
			if VERBOSE {
				fmt.Printf("thread(s): %d\n", threads)
			} else if DEBUG {
				log.Printf("thread(s): %d\n", threads)
			}
		}

	} else {

		// the original algorithm- build out all the paths, then extract ladders;
		// not so bad for typical smallish test cases, abysmal for larger wordlists

		buildStepPaths(&wordTree, beginWord, endWord, wordList)

		if VERBOSE {
			fmt.Printf(" wordTree: %v\n", wordTree)
		} else if DEBUG {
			log.Printf("wordTree: %v\n", wordTree)
		}

		getShortestLaddersFromWordTree(&shortestWordLadders, wordTree, endWord)

	}

	if DEBUG {
		log.Printf("Found %d shortest ladder(s)\n", len(shortestWordLadders))
		for i, ladder := range shortestWordLadders {
			log.Printf("Ladder[%d]: %v\n", i, ladder)
		}
	}

	return shortestWordLadders
}

// determine how different a given word is from both beginWord and endWord-
// literally, how many letters (by position) are different?
// the highest scores are those w/ no common letters for either
func wordDiff(word, beginWord, endWord string) int {

	return getLettersDiff(word, beginWord) + getLettersDiff(word, endWord)
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
	// if len(*shortestWordLadders) == 0 || len(ladder) < len((*shortestWordLadders)[0])
	for _, nextStep := range step.nextSteps {
		getShortestLaddersFromStepPaths(shortestWordLadders, ladder, nextStep, endWord)
	}

}

// stat tracking
var shortestLadderLength int32 = 0
var paths int64 = 0

// thread-specific stat tracking
var threads int64 = 0     // how many threads we've started
var threadsDone int64 = 0 // how many threads have completed

// recursive method builds out the different subsequent step paths from current step,
// while also building ladders and tracking shortest of these
func buildStepPathsAndLadders(shortestWordLadders *[][]string, ladder []string, step *Step, stepLock *sync.Mutex, beginWord string, endWord string, wordList []string) {

	// NOTE ladder already includes step.stepWord

	if DEBUG {
		log.Printf("step: %v wordList: %v\n", step, wordList)
	}

	if MULTI_THREADED {
		mu.Lock()
	}

	// we haven't reached the end, yet we're already at or over the shortest ladder length- bail out
	if len(*shortestWordLadders) > 0 && len(ladder) >= len((*shortestWordLadders)[0]) {

		if MULTI_THREADED {
			wg.Done()
			atomic.AddInt64(&threadsDone, 1)
			mu.Unlock()
		}
		return
	}

	// check to see if we can get directly to the end word from the current step
	if isOneLetterDiff(step.stepWord, endWord) {
		// last step to the end word
		if DEBUG {
			log.Printf("step %s => %s [END]\n", step.stepWord, endWord)
		}
		// build the last step
		lastStep := Step{endWord, nil}
		// the one-and-only next step is the last one
		if MULTI_THREADED {
			(*stepLock).Lock()
		}
		(*step).nextSteps = []Step{lastStep}
		if MULTI_THREADED {
			(*stepLock).Unlock()
		}

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

		if MULTI_THREADED {
			wg.Done()
			atomic.AddInt64(&threadsDone, 1)
			mu.Unlock()
		}

		return
	}

	// we know we need at least one more word between current step and endWord;
	// this additional check is particularly important when we're using threading support
	if len(*shortestWordLadders) > 0 && len(ladder) >= len((*shortestWordLadders)[0])-1 {

		// can't get directly to end word, need at least one more step between,
		// but that means we'd be longer than the shortest known ladder- bail out
		if MULTI_THREADED {
			wg.Done()
			atomic.AddInt64(&threadsDone, 1)
			mu.Unlock()
		}
		return
	}

	if MULTI_THREADED {
		mu.Unlock()
	}

	//
	// we can't get directly to the end word,
	// so find the words that will work as an interim step
	// and fill out their paths of all possible subsequent steps
	//
	extended := 0 // track how many additional branches we created for this step
	for i, word := range wordList {

		if word == beginWord || word == endWord {
			// NOTE endWord special case is handled first in this method above
			continue
		}

		// will the word from the list work as a next step from the current step
		if isOneLetterDiff(step.stepWord, word) {

			if DEBUG {
				log.Printf("step %s => %s\n", step.stepWord, word)
			}

			// build our next step
			nextStep := Step{word, nil}

			// build our ladder for this branch,
			// while keeping the input ladder intact for use on a different branch
			stepLadder := append(append(make([]string, 0), ladder...), word)

			// extract the next step word from the wordList
			// to get remaining candidates for subsequent steps
			// on this particular path
			nextStepRemainingWords := append(append(make([]string, 0), wordList[:i]...), wordList[i+1:]...)

			var nextStepLock sync.Mutex // step-specific locker

			if MULTI_THREADED {
				wg.Add(1)
				threadsSoFar := atomic.AddInt64(&threads, 1)
				threadsDoneSoFar := atomic.LoadInt64(&threadsDone)
				if DEBUG {
					if threadsSoFar%100000 == 0 {
						log.Printf("threads so far: %d active: %d\n", threadsSoFar, threadsSoFar-threadsDoneSoFar)
					}
				}
				// find all the subsequent step paths
				go buildStepPathsAndLadders(shortestWordLadders, stepLadder, &nextStep, &nextStepLock, beginWord, endWord, nextStepRemainingWords)
			} else {
				// find all the subsequent step paths
				buildStepPathsAndLadders(shortestWordLadders, stepLadder, &nextStep, &nextStepLock, beginWord, endWord, nextStepRemainingWords)

			}

			if MULTI_THREADED {
				(*stepLock).Lock()
			}

			// add this next step to the prior parent step
			if step.nextSteps == nil {
				// first one is a continutation of existing branch, not a new path
				(*step).nextSteps = []Step{nextStep}
			} else {
				(*step).nextSteps = append(step.nextSteps, nextStep)
				var pathsSoFar int64

				// 2-N is a new branch
				if MULTI_THREADED {
					pathsSoFar = atomic.AddInt64(&paths, 1)
				} else {
					paths++
					pathsSoFar = paths
				}

				if DEBUG {
					if pathsSoFar%100000 == 0 {
						log.Printf("path(s) so far: %d\n", pathsSoFar)
					}
				}

				// fmt.Printf("paths: %d\n", paths)
			}

			if MULTI_THREADED {
				(*stepLock).Unlock()
			}

			extended++ // we added a(nother) branch to input step

		} else {
			if DEBUG {
				log.Printf("step %s <> %s\n", step.stepWord, word)
			}
		}
	}

	if len(wordList) == 0 || len(step.nextSteps) == 0 || extended == 0 {
		// no more steps available, we can't get from last word to end word- DEAD END
		if DEBUG {
			log.Printf("DEAD END: %v\n", step)
		}
	}

	if MULTI_THREADED {
		wg.Done()
		atomic.AddInt64(&threadsDone, 1)
	}
}

// recursive method builds out the different subsequent step paths from current step;
// original algo- by definition this builds out the entire word tree
func buildStepPaths(step *Step, beginWord string, endWord string, wordList []string) {

	if DEBUG {
		log.Printf("step: %v wordList: %v\n", step, wordList)
	}

	// check to see if we can get directly to the end word from the current step
	if isOneLetterDiff(step.stepWord, endWord) {
		// last step to the end word
		if DEBUG {
			log.Printf("step %s => %s [END]\n", step.stepWord, endWord)
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

			if DEBUG {
				log.Printf("step %s => %s\n", step.stepWord, word)
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
				// first one is a continutation of existing branch, not a new path
				(*step).nextSteps = []Step{nextStep}
			} else {
				(*step).nextSteps = append(step.nextSteps, nextStep)

				// 2-N is a new branch
				paths++

				if DEBUG {
					if paths%100000 == 0 {
						log.Printf("path(s) so far: %d\n", paths)
					}
				}
			}

			extended++

		} else {
			if DEBUG {
				log.Printf("step %s <> %s\n", step.stepWord, word)
			}
		}
	}

	if len(wordList) == 0 || len(step.nextSteps) == 0 || extended == 0 {
		// no more steps available, we can't get from last word to end word- DEAD END
		if DEBUG {
			log.Printf("DEAD END: %v\n", step)
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

// counts how many letters are differnt by position
func getLettersDiff(word1 string, word2 string) int {

	if word1 == word2 {
		return 0
	}

	totalDifferentLetters := 0

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			totalDifferentLetters++

		}
	}

	return totalDifferentLetters
}

const (
	OUTPUT_RAW = iota
	OUTPUT_PRETTY_PRINT
	OUTPUT_GRAPHICAL
	OUTPUT_ANIMATED
)

func sortLadders(ladders [][]string) [][]string {
	slices.SortFunc(ladders, func(a, b []string) int {
		for i := range a {
			cmp := cmp.Compare(a[i], b[i])
			if cmp != 0 {
				return cmp
			}
		}
		return 0
	})
	return ladders
}

// confirm expected results matches actual results; off by default, enable with -c parameter
func verifyMatchingLadders(expectedLadders [][]string, actualLadders [][]string) {

	if len(expectedLadders) != len(actualLadders) {
		fmt.Printf("Expected ladders count %d does not match actual ladders count %d\n",
			len(expectedLadders), len(actualLadders))
		os.Exit(1)
	}

	// we must sort both to help insure an accurate comparison!
	sortedExpectedLadders := sortLadders(expectedLadders)
	sortedActualLadders := sortLadders(actualLadders)

	// for the entire set of shortest ladders
	for i := range sortedExpectedLadders {
		expectedLadder := sortedExpectedLadders[i]
		actualLadder := sortedActualLadders[i]

		// compare length of expected vs actual ordinal ladder
		if len(expectedLadder) != len(actualLadder) {
			fmt.Printf("Expected ladder[%d] length %d does not match actual ladder length %d\n",
				i, len(expectedLadder), len(actualLadder))
			os.Exit(1)
		}

		// compare words by position across expected vs actual ladder
		for j := range expectedLadder {
			expectedWord := expectedLadder[j]
			actualWord := actualLadder[j]

			if expectedWord != actualWord {
				fmt.Printf("Expected ladder[%d] word[%d] `%s` does not match actual word `%s`\n",
					i, j, expectedWord, actualWord)
				os.Exit(1)
			}
		}
	}

	fmt.Println("*CHECKED*: Expected Ladder(s) == Actual")
}

// main driver for output- dispatches based on output mode -o=<0-3>
func outputWordLadders(wordLadders [][]string, wordList []string) {
	switch OUTPUT_MODE {
	case OUTPUT_RAW:
		if VERBOSE {
			fmt.Printf("\nladder(s): %v\n", wordLadders)
		} else {
			fmt.Printf("%v\n", wordLadders)
		}
		// break
	case OUTPUT_PRETTY_PRINT:
		emitWordLadders(wordLadders, wordList)
		// break
	case OUTPUT_GRAPHICAL:
		fallthrough
	case OUTPUT_ANIMATED:
		animateWordLadders(wordLadders, wordList)
	}
}

// pretty-print the ladders- similar to json pretty print
func emitWordLadders(wordLadders [][]string, wordList []string) {
	// hoping to have some fun "animating" the output

	slices.SortFunc(wordLadders, func(a, b []string) int {
		for i := range a {
			cmp := cmp.Compare(a[i], b[i])
			if cmp != 0 {
				return cmp
			}
		}
		return 0
	})

	multiLadder := len(wordLadders) > 1

	if VERBOSE {
		fmt.Printf("ladder(s): ")
		if multiLadder {
			fmt.Println()
		}
	}

	fmt.Printf("[")

	for i, wordLadder := range wordLadders {

		if i > 0 {
			fmt.Printf(",")
		}

		if multiLadder {
			fmt.Printf("\n\t")
		}

		if len(wordLadder) > 0 {
			fmt.Printf("[")
		}

		for ii, word := range wordLadder {
			if ii > 0 {
				fmt.Printf(",")
			}
			fmt.Printf(" \"%s\"", word)

		}

		if len(wordLadder) > 0 {
			fmt.Printf(" ]")
		}

	}

	if multiLadder {
		fmt.Println()
	}

	fmt.Printf("]\n")
}

// write out a graphical ladder w/ *OPTIONAL* animation-
// as each word is emitted, we first cycle thru wordList until we find the selected item.
// uses good ol' ASCII line-draw characters and ANSI escape codes to make it fancy
// @see https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
func animateWordLadders(wordLadders [][]string, wordList []string) {

	if len(wordLadders) == 0 {
		return
	}

	slices.SortFunc(wordLadders, func(a, b []string) int {
		for i := range a {
			cmp := cmp.Compare(a[i], b[i])
			if cmp != 0 {
				return cmp
			}
		}
		return 0
	})

	for l, ladder := range wordLadders {
		if l > 0 {
			fmt.Println()
		}
		fmt.Printf("Ladder #%d:\n\n", l+1)
		slices.Reverse(ladder) // beginWord is at the bottom of the ladder, endWord at the top
		for i, word := range ladder {
			if i > 0 {
				fmt.Print("\t\u2560\u2550\u2550") // left side w/ extension for rung
				for j := 0; j < len(word); j++ {
					fmt.Print("\u2550") // rung length same as word length
				}
				fmt.Print("\u2550\u2550\u2563\n") // right side w/ extension for rung
			}
			fmt.Printf("\t\u2551  ") // vertical left side
			if OUTPUT_ANIMATED == OUTPUT_MODE {
				fmt.Print("\u001b[7m\u001b[5m") // set blink and invert mode
				for _, candidate := range wordList {
					fmt.Printf("%s", candidate)
					time.Sleep(50 * time.Millisecond)
					fmt.Printf("\u001b[%dD", len(candidate)) // back up over word we just wrote
					if candidate == word {
						break
					}
				}
				fmt.Print("\u001b[27m\u001b[25m") // reset blink and invert mode
			}
			fmt.Printf("%s  \u2551\n", word) // vertical right side
		}
	}

	fmt.Println()
}

// MAX WORDS SAMPLE
type Inputs struct {
	beginWord string
	endWord   string
	wordList  []string
}
