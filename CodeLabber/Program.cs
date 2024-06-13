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

            //Toss that list into a set, including endWord
            HashSet<string> wordSet = new(wordList) { endWord };

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
            string curWord = beginWord;
            foreach (var word in wordList)
            {
                if (TryWord(curWord, word, out string retWord))
                {
                    Console.WriteLine(retWord);
                    curWord = retWord;
                }
            }

            Console.WriteLine("---");
            Console.WriteLine(curWord);
            _ = Console.ReadLine();
        }
    }
}
