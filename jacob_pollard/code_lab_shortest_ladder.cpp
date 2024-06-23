#include <vector>
#include <string>
#include <unordered_set>
#include <queue>
#include <iostream>

using namespace std;

/**
 * @brief Finds all possible words that can be formed by changing one letter at a time from the given word
 * and are present in the provided set of distinct words.
 *
 * @param word The word from which to generate new words by changing one letter at a time.
 * @param distinctWords A set of distinct words to check against.
 * @return An unordered_set of valid words that can be formed.
 */
unordered_set<string> findNextWords(const string& word, const unordered_set<string>& distinctWords) {
    unordered_set<string> result;
    string nextWord = word;

    for (size_t i = 0; i < word.size(); ++i) {
        char originalChar = nextWord[i];
        for (char c = 'a'; c <= 'z'; ++c) {
            if (c == originalChar) {
                continue; // Skip the original character
            }
            nextWord[i] = c;
            if (distinctWords.find(nextWord) != distinctWords.end()) {
                result.insert(nextWord);
            }
        }
        nextWord[i] = originalChar; // Restore original character
    }

    return result;
}

/**
 * @brief Prints the vectors of strings in a formatted manner.
 *
 * @param ladders A 2D vector of strings to be printed.
 */
void printLadders(vector<vector<string>> &ladders)
{
  int size = ladders.size();

  if (size == 0)
  {
    cout << "[]\n";
    return;
  }

  cout << "[";

  for (int i = 0; i < size; ++i)
  {
    cout << "[";
    for (int j = 0; j < ladders[i].size(); ++j)
    {
      cout << "\"" << ladders[i][j] << "\"";
      if (j < ladders[i].size() - 1)
      {
        cout << ", ";
      }
    }
    cout << "]";
    if (i < size - 1)
    {
      cout << ",\n ";
    }
  }
  cout << "]\n";
}

/**
 * @brief Finds all the shortest transformation sequences from beginWord to endWord.
 *
 * @param wordList A list of words to be used in transformations.
 * @param beginWord The starting word of the transformation sequence.
 * @param endWord The ending word of the transformation sequence.
 * @return A vector of vector of strings, where each inner vector represents a shortest transformation sequence.
 */
vector<vector<string>> getShortestLadders(
    vector<string> &wordList,
    string beginWord,
    string endWord)
{
  vector<vector<string>> shortestLadders;

  // Visited words
  unordered_set<string> visited;

  // Queue for BFS
  queue<vector<string>> pathQueue;

  // Set of distinct words for quick lookup
  unordered_set<string> distinctWords(wordList.begin(), wordList.end());

  // Initialize BFS
  pathQueue.push({beginWord});

  // Flag to indicate if the shortest path has been found
  bool shortestPathFound = false;

  while (!pathQueue.empty())
  {
    int size = pathQueue.size();
    for (int i = 0; i < size; i++)
    {
      vector<string> current = pathQueue.front();
      pathQueue.pop();

      unordered_set<string> nextWords = findNextWords(current.back(), distinctWords);

      for (const string &nextWord : nextWords)
      {
        vector<string> temp(current.begin(), current.end());

        temp.push_back(nextWord);

        if (nextWord == endWord)
        {
          shortestPathFound = true;
          shortestLadders.push_back(temp);
        }

        visited.insert(nextWord);
        pathQueue.push(temp);
      }
    }

    if (shortestPathFound)
      break;

    for (auto it : visited)
    {
      distinctWords.erase(it);
    }

    visited.clear();
  }

  return shortestLadders;
}

int main()
{
  vector<string> wordList{"cot", "dog", "cat", "cut", "mug", "fog", "fig", "mut", "dug"};
  string beginWord = "cig";
  string endWord = "cot";

  vector<vector<string>> ladders = getShortestLadders(wordList, beginWord, endWord);

  printLadders(ladders);
}