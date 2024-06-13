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
            //string[] wordList = ["hot", "log", "cog"];

            //string beginWord = "";
            //string endWord = "";
            //string[] wordList = [];

            //string beginWord = "lost";
            //string endWord = "cost";
            //string[] wordList = ["most", "fost", "cost", "host", "lost"];

            //string beginWord = "start";
            //string endWord = "endit";
            //string[] wordList = ["stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"];

            DoWords(beginWord, endWord, wordList);
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

            //Quick function to determine valid words
            //Just inline it here so we can use wordSet w/o passing it
            bool TryWord(string w1, string w2, out string retWord)
            {
                StringBuilder tempWord = new(w1);
                for (int i = 0; i < w1.Length; i++)
                {
                    tempWord[i] = w2[i];
                    if (wordSet.Remove(tempWord.ToString()))
                    {
                        retWord = tempWord.ToString();
                        return true;
                    }

                    tempWord[i] = w1[i]; //Reset for next
                }

                retWord = "";
                return false;
            }

            //Start comparing words
            HashSet<string> validWords = [];
            string curWord = beginWord;
            foreach (var word in wordList)
            {
                if (TryWord(curWord, endWord, out string retWord)
                    || TryWord(curWord, word, out retWord))
                {
                    validWords.Add(retWord);
                    Console.WriteLine(retWord);
                    curWord = retWord;
                    if (curWord == endWord)
                        wordSet.Clear();
                }
            }

            //Write the result
            string[] results = [beginWord, .. validWords];
            Console.WriteLine(string.Join(",", results));
        }
    }
}
