using CodeLabber;

namespace CodeLabberTests
{
    [TestClass]
    public class ProgramTest
    {
        readonly string beginWord = "hit";
        readonly string endWord = "cog";
        readonly string[] wordList = ["hot", "dot", "dog", "lot", "log", "cog"];
        readonly string output = @"[[""hit"",""hot"",""dot"",""dog"",""cog""],[""hit"",""hot"",""lot"",""log"",""cog""]]";

        readonly StringWriter consoleWriter = new StringWriter();

        public ProgramTest()
        {
            Console.SetOut(consoleWriter);
        }

        [TestCleanup]
        public void Cleanup()
        {
            consoleWriter.Flush();
        }

        private void AssertOutput(string expected)
        {
            Assert.AreEqual(expected + Environment.NewLine, consoleWriter.ToString());
        }

        [TestMethod]
        public void Solution_Example1()
        {
            Program.SolveFor(beginWord, endWord, wordList);
            AssertOutput(output);
        }

        [TestMethod]
        public void Solution_Example2()
        {
            Program.SolveFor("lost", "cost", ["most", "fost", "cost", "host", "lost"]);
            AssertOutput(@"[[""lost"",""cost""]]");
        }

        [TestMethod]
        public void Solution_Example2_NoExtraWords()
        {
            Program.SolveFor("lost", "cost", ["cost"]);
            AssertOutput(@"[[""lost"",""cost""]]");
        }

        [TestMethod]
        public void NoSolution_Example3()
        {
            Program.SolveFor("start", "endit", ["stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"]);
            AssertOutput("[]");
        }

        [TestMethod]
        public void Error_WordLength_Differs()
        {
            Program.SolveFor(beginWord, endWord, ["hot", "lg", "cog"]);
            AssertOutput($"ERROR: hit and lg are of differing lengths.{Environment.NewLine}[]");
        }

        [TestMethod]
        public void NoSolution_Example1_ShortWordList()
        {
            Program.SolveFor(beginWord, endWord, ["hot", "log", "cog"]);
            AssertOutput("[]");
        }

        [TestMethod]
        public void NoSolution_EmptyArguments()
        {
            Program.SolveFor("", "", []);
            AssertOutput(@"[[""""]]"); //Technically correct? We don't need a path to go from "" to ""... :P
        }

        [TestMethod]
        public void NoSolution_EmptyWordList()
        {
            Program.SolveFor("a", "b", []);
            AssertOutput("[]");
        }

        [TestMethod]
        public void Solution_Example1_LongWordList_Dupes()
        {
            Program.SolveFor(beginWord, endWord, ["hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog"]);
            AssertOutput(output);
        }

        [TestMethod]
        public void Solution_Example1_LongerWordList_Dupes()
        {
            Program.SolveFor(beginWord, endWord, ["hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog", "hot", "dot", "dog", "lot", "log", "cog"]);
            AssertOutput(output);
        }

        [TestMethod]
        public void Solution_ExampleN_LongWordList()
        {
            List<string> list = [];
            for (int i = 0; i < 20; i++)
                list.Add(i.ToString("D2"));

            Program.SolveFor(list.First(), list.Last(), list);
            AssertOutput(@"[[""00"",""09"",""19""],[""00"",""10"",""19""]]");
        }

        [TestMethod]
        public void Solution_ExampleN_LongerWordList()
        {
            List<string> list = [];
            for (int i = 0; i < 200; i++)
                list.Add(i.ToString("D3"));

            Program.SolveFor(list.First(), list.Last(), list);
            AssertOutput(@"[[""000"",""009"",""099"",""199""],[""000"",""009"",""109"",""199""],[""000"",""090"",""099"",""199""],[""000"",""090"",""190"",""199""],[""000"",""100"",""109"",""199""],[""000"",""100"",""190"",""199""]]");
        }
    }
}
