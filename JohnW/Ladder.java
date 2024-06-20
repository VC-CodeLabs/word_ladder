import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

public class Ladder {

   // Test cases. This allows quick swapping of one test case with another.
   // See the main() method.
   static Words simple = new Words("pow", "dog", List.of("cow", "cop", "dug", "dip", "dot", "cot", "dog"));
   static Words longWords = new Words("plane", "sling",
         List.of("plant", "plank", "blank", "blink", "slink", "sling", "cling", "clint", "flint"));

   static Words manyWords = new Words("Lane", "Bark",
         List.of("Bark", "Dark", "Darn", "Barn", "Bare", "Bake", "Cake", "Cane", "Lane", "Lame", "Lamp", "Damp", "Dame", "Game",
               "Gale", "Sale", "Sake", "Take", "Tale", "Tile", "File", "Fine", "Line", "Lime", "Time", "Tome", "Home", "Hope",
               "Rope", "Rope"));

   static Words repeatedWords = new Words("Bark", "Rope",
         List.of("Bark", "Dark", "Darn", "Barn", "Bare", "Bake", "Cake", "Cane", "Lane", "Lame", "Lamp", "Damp", "Dame", "Game",
               "Gale", "Sale", "Bark", "Bark", "Darn", "Cake", "Lame", "Lame", "Lame", "Lame", "Time", "Tome", "Home", "Hope",
               "Rope", "Rope"));

   public static void main(String[] args) {
      // Change testWords to point to one of the static test cases
      Words testWords = Ladder.manyWords;

      // If -v is passed on the command line, we'll print out all the intermediate steps
      boolean verbose = false;
      if (args.length > 0) {
         verbose = args[0].equals("-v");
      }
      System.out.println("\nLooking for path from " + testWords.startWord + " -> " + testWords.endWord);

      // Walk all the paths to find the shortest
      List<String> path = Ladder.walkWords(testWords.startWord, testWords.endWord, testWords.wordList, verbose);

      // Cleverly display the results
      if (path.isEmpty()) {
         System.out.println(
               "\nZoinks! The way is blocked!\nThere is no way to get from " + testWords.startWord + " to " + testWords.endWord +
                     ".");
      } else {
         System.out.println("\nYahoo! Here is your path:");
         StringBuilder buf = new StringBuilder();
         for (int i = 0; i < path.size(); i++) {
            buf.append(path.get(i));
            if (i < path.size() - 1) {
               buf.append(" -> ");
            }
         }
         System.out.println(buf);
      }
   }

   /**
    * Determines if word2 is only 1 letter off from word1
    *
    * @param word1 current word
    * @param word2 next word
    * @return true if only 1 letter difference
    */
   public static boolean isNextWord(String word1, String word2) {
      if (word1.length() != word2.length()) {
         return false;
      }
      int diff = 0;
      for (int i = 0; i < word1.length(); i++) {
         if (word1.charAt(i) != word2.charAt(i)) {
            diff++;
         }
      }
      return diff == 1;
   }

   /**
    * Walks through the word list to find the shortest path from start word to end word.
    *
    * @param startWord Starting word
    * @param endWord   Target end word
    * @param wordList  List of valid words
    * @return List of words to make a path from start to end
    */
   public static List<String> walkWords(String startWord, String endWord, List<String> wordList, boolean verbose) {
      List<WordPath> result = new LinkedList<>();

      // prime our result list with the start word
      WordPath path = new WordPath(startWord, List.of(startWord));
      result.add(path);

      // keep track of words we've seen
      Set<String> visitedWords = new HashSet<>();

      // Walk through the list of paths, adding words, until we get to the end word
      while (!result.isEmpty()) {
         // pop the current word and its path
         WordPath currentEntry = result.remove(0);
         // see if we're at the end
         if (currentEntry.word.equals(endWord)) {
            return currentEntry.path;
         }

         // Not at the end so find the next candidate word and add it to the current word's path
         for (String word : wordList) {
            if (isNextWord(currentEntry.word, word) && !visitedWords.contains(word)) {
               visitedWords.add(word);

               List<String> newPath = new ArrayList<>(currentEntry.path);

               newPath.add(word);
               WordPath wPath = new WordPath(word, newPath);
               if (verbose) {
                  System.out.println(wPath);
               }
               result.add(wPath);
            }
         }
      }
      return Collections.emptyList();
   }

   /**
    * A simple class to store the start, end, and valid word list for testing
    */
   static class Words {
      String startWord;
      String endWord;
      List<String> wordList;

      public Words(String startWord, String endWord, List<String> wordList) {
         this.startWord = startWord;
         this.endWord = endWord;
         this.wordList = wordList;
      }
   }

   /**
    * Contains a word and the current list of words heading toward the end word.
    * This is pushed onto a list during the word walking.
    */
   static class WordPath {
      public String word;
      public List<String> path;

      public WordPath(String word, List<String> path) {
         this.word = word;
         this.path = path;
      }

      public String toString() {
         return word + ": [" + path + "]";
      }
   }
}