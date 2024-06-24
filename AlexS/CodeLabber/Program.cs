using Newtonsoft.Json;
using System.Runtime.CompilerServices;
using System.Text;

[assembly: InternalsVisibleTo("CodeLabberTests")]

/*
 * Implementation for the "Word Ladder" challenge.
 * 
 * This solution attempts a recursive approach where, for a given "currentWord" (starting with beginWord):
 * 1. If currentWord equals endWord, add our "path" of words tested to our solution set, and return.
 * 2. Test currentWord against each word in the wordList for a valid match of one-character difference.
 * 3. For each valid word, repeat the above using a wordList of the remaining words (that did not match w/ a one-character difference).
 * 
 * This will continue until no further valid words are found, ultimately returning our complete solution set.
 * 
 * Some optimizations:
 * -- In Step 1, we also return early if our incoming wordList is longer than the length of a currently known solution.
 *    This allows us to begin ignoring unnecessarily long paths as we build our solution set.
 *    
 * -- In Step 2, we always load our incoming wordList into a new HashSet. This allows us to:
 *    A. Quickly lookup potential matches when comparing words, and remove them, yielding our list of remaining words.
 *    B. Avoid modifying the incoming wordList, as it may still be in use by later iterations of Step 3.
 * 
 * -- In Step 3, priority is given to valid words most similar to the endWord, to traverse likely shorter solutions early.
 *    This is computed once at the start of the routine, by comparing the number of character matches in each word to the endWord.
 *    The similarity counts of each word are then stuffed into a Dictionary "endWordSimilarity" for later lookup.
 */
namespace CodeLabber
{
    internal class Program
    {
        /// <summary>
        /// Whether to print extra statistics to output, including number of solutions and steps iterated.
        /// </summary>
        static bool printStatistics = false;

        /// <summary>
        /// Runs the program.
        /// </summary>
        /// <param name="args">Unused.</param>
        static void Main(string[] args)
        {
            //Input vars! Set stuff here!
            string beginWord = "hit";
            string endWord = "cog";
            string[] wordList = ["hot", "dot", "dog", "lot", "log", "cog"];
            printStatistics = true; //By default, we'll print some extra statistics when run via console. Disable if desired.

            SolveFor(beginWord, endWord, wordList);
        }

        /// <summary>
        /// Print a Word Ladder solution to console for a given begin / end word, using specified word list.
        /// </summary>
        /// <param name="beginWord">Begin word.</param>
        /// <param name="endWord">End word.</param>
        /// <param name="wordList">List of words used to hop from beginWord to endWord, changing only one letter at a time.</param>
        /// <exception cref="ArgumentException">Thrown on wordList of invalid word length. (differing from beginWord / endWord)</exception>
        internal static void SolveFor(string beginWord, string endWord, ICollection<string> wordList)
        {
            //List of solution sets, to be pretty-printed in output
            List<ICollection<string>> solutions = [];

            //Dictionary of wordList, with similarity values to endWord
            Dictionary<string, int> endWordSimilarity = [];
            foreach (string word in wordList)
                if (endWordSimilarity.TryAdd(word, 0))  //If we haven't already examined word...
                    for (int i = 0; i < word.Length; i++)
                        if (word[i] == endWord[i])      //Compare each character
                            endWordSimilarity[word]++;  //On match, increment our similarity value

            //Current min solution length - don't bother traversing paths longer than this
            int minSolutionLength = int.MaxValue;

            //Path counts for statistics
            int stepsTraversed = 0;
            int pathsIgnored = 0;

            //Use an inline function to do some recursive work
            void DoWords(string currentWord, ICollection<string> wordList, ICollection<string> path)
            {
                //Ignore paths that will be longer than a known shorter solution
                if (path.Count > minSolutionLength)
                {
                    pathsIgnored++;
                    return;
                }
                stepsTraversed++;

                //If we've found what we're looking for, stop here!
                if (currentWord == endWord)
                {
                    minSolutionLength = path.Count; //min by above check
                    solutions.Add(path);
                    return;
                }

                //Toss that list into a working set we can modify within loop
                HashSet<string> workingSet = [.. wordList];

                //Start comparing words
                HashSet<string> validWords = [];
                foreach (string word in wordList)
                {
                    //Sanity check
                    if (currentWord.Length != word.Length)
                        throw new ArgumentException($"{currentWord} and {word} are of differing lengths.");

                    //Start looking for valid words one character at a time
                    StringBuilder currentWordBuilder = new(currentWord);
                    for (int i = 0; i < currentWord.Length; i++)
                    {
                        currentWordBuilder[i] = word[i];
                        if (workingSet.Remove(currentWordBuilder.ToString()))
                            validWords.Add(currentWordBuilder.ToString());

                        currentWordBuilder[i] = currentWord[i]; //Reset for next
                    }
                }

                //Recurse our remaining workingSet for each possible validWord, preferring most similar words
                //This will naturally exhaust itself if no (further) solution exists
                foreach (string word in validWords.OrderByDescending(w => endWordSimilarity[w]))
                    DoWords(word, workingSet, [.. path, word]);
            }

            //Kick it!
            try
            {
                DoWords(beginWord, wordList, [beginWord]);
            }
            catch (ArgumentException ex)
            {
                Console.WriteLine($"ERROR: {ex.Message}");
            }

            //Print solutions (if any)
            Console.WriteLine(JsonConvert.SerializeObject(solutions.Where(s => s.Count == minSolutionLength)));

            //Print some statistics?
            if (printStatistics)
            {
                Console.WriteLine();
                Console.WriteLine($"We returned {solutions.Count(s => s.Count == minSolutionLength)} optimal solutions.");
                Console.WriteLine($"We threw away {solutions.Count(s => s.Count > minSolutionLength)} sub-optimal solutions.");
                Console.WriteLine($"We traversed {stepsTraversed} steps.");
                Console.WriteLine($"We ignored {pathsIgnored} paths longer than a known solution.");
                Console.WriteLine();
            }
        }
    }
}
