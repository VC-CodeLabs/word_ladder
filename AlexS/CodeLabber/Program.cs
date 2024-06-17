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
            //Dictionary of solutions
            //Keys are string contents of the solution set to avoid duplicates
            //Values are simply the original solution set, to be pretty-printed in output
            Dictionary<string, HashSet<string>> solutions = [];

            //HashSet of lists we've already processed
            HashSet<string> processed = [];

            //Use an inline function to do some recursive work
            void DoWords(ICollection<string> wordList)
            {
                //Skip if we've already done this list
                if (!processed.Add(string.Join(',', wordList)))
                    return;

                //Sanity check
                if (beginWord.Length != endWord.Length)
                    throw new ArgumentException($"{beginWord} and {endWord} are of differing lengths.");

                //Toss that list into a set
                HashSet<string> wordSet = [.. wordList];

                //Start comparing words
                HashSet<string> validWords = [beginWord];
                string currentWord = beginWord;
                foreach (string word in wordList)
                {
                    //Sanity check
                    if (beginWord.Length != word.Length)
                        throw new ArgumentException($"{beginWord} and {word} are of differing lengths.");

                    //Start looking for valid words one character at a time
                    StringBuilder currentWordBuilder = new(currentWord);
                    for (int i = 0; i < currentWord.Length; i++)
                    {
                        currentWordBuilder[i] = word[i];
                        if (wordSet.Remove(currentWordBuilder.ToString()))
                        {
                            currentWord = currentWordBuilder.ToString();
                            validWords.Add(currentWord);
                            break;
                        }

                        currentWordBuilder[i] = currentWord[i]; //Reset for next
                    }

                    //Bail here if we found our match
                    if (currentWord == endWord)
                    {
                        //But! Don't continue further if we've already "done" this solution tree
                        if (!solutions.TryAdd(string.Join(',', validWords), validWords))
                            return;

                        //Otherwise, resume regularly scheduled programming after this break :P
                        break;
                    }
                }

                //Again! Brute force a recursive "better solution"
                foreach (string word in wordList)
                    DoWords([.. wordList.Where(x => x != word)]);
            }

            //Kick it!
            try
            {
                DoWords(wordList);
            }
            catch (ArgumentException ex)
            {
                Console.WriteLine($"ERROR: {ex.Message}");
            }

            //Print solutions (if any)
            if (solutions.Count == 0) //LINQ barfs on empty collections, whoops
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Values));
            else
            {
                int minLength = solutions.Select(s => s.Value.Count).Min();
                Console.WriteLine(
                    JsonConvert.SerializeObject(
                        solutions.Values
                        .Where(x => x.Count == minLength)
                        .OrderBy(x => string.Join(',', x))
                    )
                );
            }
        }
    }
}
