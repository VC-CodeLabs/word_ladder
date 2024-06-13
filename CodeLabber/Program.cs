using Newtonsoft.Json;
using System.Text;

namespace CodeLabber
{
    internal class Program
    {
        static readonly Dictionary<string, HashSet<string>> solutions = [];

        static void Main(string[] args)
        {
            //Input vars! Set stuff here!
            string beginWord = "hit";
            string endWord = "cog";
            string[] wordList = ["hot", "dot", "dog", "lot", "log", "cog"];
            DoWords(beginWord, endWord, wordList);

            //Test cases
            //DoWords(beginWord, "derp", wordList);
            //DoWords(beginWord, endWord, ["hot", "log", "cog"]);
            //DoWords("", "", []);
            //DoWords("lost", "cost", ["most", "fost", "cost", "host", "lost"]);
            //DoWords("start", "endit", ["stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"]);

            //Print solutions (if any)
            if (solutions.Count == 0) //LINQ barfs on empty collections, whoops
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Values));
            else
            {
                int minLength = solutions.Select(s => s.Value.Count).Min();
                Console.WriteLine(JsonConvert.SerializeObject(solutions.Values.Where(x => x.Count == minLength)));
            }
        }

        static void DoWords(string beginWord, string endWord, string[] wordList)
        {
            //Sanity check
            if (beginWord.Length != endWord.Length)
            {
                Console.WriteLine("ERROR: beginWord and endWord are of differing lengths.");
                return;
            }

            //Toss that list into a set, including endWord
            HashSet<string> wordSet = [.. wordList, endWord];

            //Start comparing words
            HashSet<string> validWords = [beginWord];
            string currentWord = beginWord;
            foreach (string word in wordList)
            {
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
                DoWords(beginWord, endWord, [.. wordList.Where(x => x != word)]);
        }
    }
}
