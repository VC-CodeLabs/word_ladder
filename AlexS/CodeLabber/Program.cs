using Newtonsoft.Json;
using System.Runtime.CompilerServices;
using System.Text;

[assembly: InternalsVisibleTo("CodeLabberTests")]

namespace CodeLabber
{
    internal class Program
    {
        static void Main(string[] args)
        {
            //Input vars! Set stuff here!
            string beginWord = "hit";
            string endWord = "cog";
            string[] wordList = ["hot", "dot", "dog", "lot", "log", "cog"];

            SolveFor(beginWord, endWord, wordList);
        }

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

            //Use an inline function to do some recursive work
            void DoWords(string currentWord, ICollection<string> wordList, ICollection<string> path)
            {
                //Ignore paths that will be longer than a known shorter solution
                if (path.Count > minSolutionLength)
                    return;

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
            if (solutions.Count == 0) //LINQ barfs on empty collections, whoops
                Console.WriteLine(JsonConvert.SerializeObject(solutions));
            else
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Where(s => s.Count == minSolutionLength)));
        }
    }
}
