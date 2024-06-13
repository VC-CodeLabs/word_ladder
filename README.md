# Word Ladder

**Problem Description**</br>
Given two words, **beginWord** and **endWord**, and a dictionary **wordList**, return all the shortest transformation sequences from **beginWord** to **endWord**, or an empty list if no such sequence exists. Each sequence should be returned as a list of the words **[beginWord, word_1, ..., endWord]**.

- The first word in the sequence is **beginWord**.
- The last word in the sequence is **endWord**.
- Only one letter can be changed at a time.
- Each transformed word must exist in the **wordList**.
</br>

**Constraints:**

- 1 <= beginWord.length <= 5
- endWord.length == beginWord.length
- 1 <= wordList.length <= 5000
- wordList[i].length == beginWord.length
- beginWord, endWord, and wordList[i] consist of lowercase English letters.
- beginWord != endWord
- The beginWord does not need to exist in the list.
- All the words in wordList are unique.
- Please provide variables for me to use to hold the beginWord, endWord, and wordList
- Choose any programming language you like
- You can use any external libraries that you feel will help you
- Feel free to use AI to assist your development - the only exception to this is you *CANNOT USE AI TO WRITE A SOLUTION* for you 
</br>

**Examples**</br>
Example 1:
- Input:
  - beginWord = "hit"
  - endWord = "cog"
  - wordList = ["hot","dot","dog","lot","log","cog"]
- Output:
```
[["hit","hot","dot","dog","cog"],
["hit","hot","lot","log","cog"]]
```
</br>

Example 2:
- Input:
  - beginWord = "lost"
  - endWord = "cost"
  - wordList = ["most","fost","cost","host","lost"]
- Output:
```
[["lost","cost"]]
```
</br>

Example 3:
- Input:
  - beginWord = "start"
  - endWord = "endit"
  - wordList = ["stark", "stack", "smack", "black", "endit", "blink", "bline", "cline"]
- Output:
```
  []
```
</br>

**Scoring**

- Correctness and Completeness (60 points): Solutions must pass all provided test cases, including edge cases.
- Code Quality and Style (30 points): Well-structured code with clear comments, a summary describing your algorithm, and good variable naming is expected.
- Performance and Efficiency (10 points): Fast execution times will be rewarded, but solutions must first and foremost be correct and complete.		
-	Performance	Points</br>	
![image](https://github.com/VC-CodeLabs/word_ladder/assets/154607214/a80c5043-3167-4f38-871d-248ea9dd5b97)


**Due Date**</br>
Please have the solution emailed to me or a PR submitted to the GitHub VC-CodeLabs/word_ladder repository by Sunday, June 23rd, by midnight.  

