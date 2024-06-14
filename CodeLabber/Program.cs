using Newtonsoft.Json;
using System.Text;

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

            //Test cases
            //SolveFor(beginWord, "derp", wordList); //endWord differs
            //SolveFor(beginWord, endWord, ["hot", "log", "cog"]);
            //SolveFor(beginWord, endWord, ["hot", "lg", "cog"]); //wordList differs
            //SolveFor(beginWord, endWord, ["hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog"]);
            //SolveFor(beginWord, endWord, ["hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog"]);
            //SolveFor("", "", []);
            //SolveFor("lost", "cost", ["most", "fost", "cost", "host", "lost"]);
            //SolveFor("start", "endit", ["stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"]);

        }

        static void SolveFor(string beginWord, string endWord, string[] wordList)
        {
            //Dictionary of solutions
            //Keys are string contents of the solution set to avoid duplicates
            //Values are simply the original solution set, to be pretty-printed in output
            Dictionary<string, HashSet<string>> solutions = [];

            //Use an inline function to do some recursive work
            void DoWords(string[] wordList)
            {
                //Sanity check
                if (beginWord.Length != endWord.Length)
                {
                    Console.WriteLine($"ERROR: {beginWord} and {endWord} are of differing lengths.");
                    return;
                }

                //Toss that list into a set
                HashSet<string> wordSet = [.. wordList];

                //Start comparing words
                HashSet<string> validWords = [beginWord];
                string currentWord = beginWord;
                foreach (string word in wordList)
                {
                    //Sanity check
                    if (beginWord.Length != word.Length)
                    {
                        Console.WriteLine($"ERROR: {beginWord} and {word} are of differing lengths.");
                        return;
                    }

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
                        solutions.TryAdd(JsonConvert.SerializeObject(validWords), validWords);
                        break;
                    }
                }

                //Again! Brute force a recursive "better solution"
                foreach (string word in wordList)
                    DoWords([.. wordList.Where(x => x != word)]);
            }

            //Kick it!
            DoWords(wordList);

            //Print solutions (if any)
            if (solutions.Count == 0) //LINQ barfs on empty collections, whoops
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Values));
            else
            {
                int minLength = solutions.Select(s => s.Value.Count).Min();
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Values.Where(x => x.Count == minLength)));
            }
        }
    }
}
