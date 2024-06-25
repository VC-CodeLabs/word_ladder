/* TomsWordLadder.cpp

Problem:
Given two words, beginWord and endWord, and a dictionary wordList, return all the shortest transformation sequences from beginWord to endWord,
or an empty list if no such sequence exists. Each sequence should be returned as a list of the words [beginWord, word_1, ..., endWord].
The first word in the sequence is beginWord.
The last word in the sequence is endWord.
Only one letter can be changed at a time.
Each transformed word must exist in the wordList.

Constraints:
1 <= beginWord.length <= 5
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
The beginWord does not need to exist in the list.
All the words in wordList are unique.
Choose any programming language you like
You can use any external libraries that you feel will help you
Feel free to use AI to assist your development - the only exception to this is you CANNOT USE AI TO WRITE A SOLUTION for you
 
 Method:

 1.  Create a mesh network of the dictionary with direct connections to all other words that are one char apart
 2.  Get beginWord
 3.  Either find beginWord in the dictionary or if it doesn't exist yet create a new word and direct connections
 4.  Find all the direct links between all the nodes and all the other nodes
 4.  Traverse the network to find the shortest path using Dijkstra's Algorithm or Breadth First Search.  Track the shortest path found.
    4a.  If start equals end then stop.  Mark the path as complete.  If this path is shorter than the shortest so far then note this.
    4b.  Loop through all links for start node
    4c.     If we are beyond the first iteration of the loop then create a new path
    4d.     Copy the previous path up to this start node
    4e.     If this new path is already too long then adandon this new path
    4f.     If the node does not exist in the path (new or existing) then add the node
    4f.     Recursively traverse the network with the current linked node being the start (step 4a)
 5.  Print results

*/

#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <time.h>
#include <atlstr.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_LEN 5
#define MAX_WORDS 5000
#define LARGE 10000

class Word
{
public:
    int index;
    char value[MAX_LEN + 1];
    int linkDistance[MAX_WORDS];
    BOOL visitedPath[MAX_WORDS];
    Word(char in[], int num);
    Word();
};

Word::Word(char in[], int num)
{
    index = num;
    strcpy_s(value, in);
    for (int i = 0; i < MAX_WORDS; i++)
    {
        linkDistance[i] = 0;
        visitedPath[i] = FALSE;
    }
    return;
}

Word::Word()
{
    index = MAX_WORDS;
    value[0] = '\0';
    for (int i = 0; i < MAX_WORDS; i++)
    {
        linkDistance[i] = 0;
        visitedPath[i] = FALSE;
    }
    return;
};

Word dictionary[MAX_WORDS];

class Path
{
public:
    int node[MAX_WORDS];
    int length;
    BOOL reachedEnd;

    Path() { for (int i = 0; i < MAX_WORDS; i++) node[i] = MAX_WORDS; length = 0; reachedEnd = FALSE; return; };
    CString GetName();
};

Path paths[MAX_WORDS*2];
int numPaths = 0;
int shortestPath = LARGE;
int numShortest = 0;

// Input
// My Small example
//char wordList[][MAX_LEN+1] = { "him","hit","hot","rit","rot","rob"};
//char beginWord[MAX_LEN+1] = { "him" };
//char endWord[MAX_LEN+1] = { "rob" };

// Example 1
//char wordList [][MAX_LEN+1] = {"hot","dot","dog","lot","log","cog"};
//char beginWord[MAX_LEN+1] =   {"hit"};
//char endWord[MAX_LEN+1] =     {"cog"};

// Example 2
//char wordList[][MAX_LEN+1] = { "most", "fost", "cost", "host", "lost" };
//char beginWord[MAX_LEN+1] = { "lost" };
//char endWord[MAX_LEN+1] = { "cost" };

// Example 3
//char wordList[][MAX_LEN + 1] = { "stark", "stack", "smack", "black", "endit", "blink", "bline", "cline" };
//char beginWord[MAX_LEN + 1] = { "start" };
//char endWord[MAX_LEN + 1] = { "endit" };

// Example 4
//char wordList[][MAX_LEN + 1] = { "cot", "dog", "cat", "cut", "mug", "fog", "fig", "mut", "dug"};
//char beginWord[MAX_LEN + 1] = { "cig" };
//char endWord[MAX_LEN + 1] = { "cot" };

// Sample I found online
// https://runestone.academy/ns/books/published/pythonds3/Graphs/BuildingtheWordLadderGraph.html
// 4 solutions of length 7
//char wordList[][MAX_LEN + 1] = { "sage", "sale", "page", "pope", "pole", "pale", "pool", "poll", "fall", "pall", "fail", "fool", "cool", "foul", "foil"};
//char beginWord[MAX_LEN + 1] = { "sage" };
//char endWord[MAX_LEN + 1] = { "fool" };

// Test Case 1
//char wordList [][MAX_LEN+1] = {"hot","dot","dog","lot","log","cog"};
//char beginWord[MAX_LEN+1] =   {"hit"};
//char endWord[MAX_LEN+1] =     {"cog"};

// Test Case 2
//char wordList[][MAX_LEN + 1] = { "a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z" };
//char beginWord[MAX_LEN+1] =   {"a"};
//char endWord[MAX_LEN+1] =     {"c"};

// Test Case 3
//char wordList[][MAX_LEN + 1] = { "hot", "dot", "dog", "lot", "log" };
//char beginWord[MAX_LEN+1] =   {"hit"};
//char endWord[MAX_LEN+1] =     {"cog"};

// Test Case 4
//char wordList[][MAX_LEN + 1] = { "same", "came", "lame", "name" };
//char beginWord[MAX_LEN + 1] = { "lame" };
//char endWord[MAX_LEN + 1] = { "same" };

// Test Case 5
//char wordList[][MAX_LEN + 1] = { "aaaa", "abbb", "aaab", "aabb", "bbbb" };
//char beginWord[MAX_LEN + 1] = { "aaaa" };
//char endWord[MAX_LEN + 1] = { "bbbb" };

// Test Case 6
char wordList[][MAX_LEN + 1] = { "stoke", "stony", "stome", "stomy", "stoey", "htoey", "htney", "hiney", "miney", "ttney", "toney", "itoey", "mtney", "soney", "store", "storm", "story", "monte", "monny", "monty", "money", "stane", "stine", "maney", "honey", "monde", "stnny", "mtone", "mtnne", "monne", "monee", "stnne", "mtnee" };
char beginWord[MAX_LEN + 1] = { "stone" };
char endWord[MAX_LEN + 1] = { "money" };

int dictionaryLength = sizeof(wordList) / (MAX_LEN + 1);

BOOL Linked(char one[], char two[]);
void AddLinks(int index);
void TraverseMeshNetwork(int beginWordIndex, int endWordIndex, int path);

int main()
{
    int i, j;
    int beginWordIndex = MAX_WORDS;
    int endWordIndex = MAX_WORDS;

    // Create the dictionary array
    for (i = 0; i < dictionaryLength; i++)
        dictionary[i] = Word(wordList[i], i);

    // Turn the array into a mesh network by creating all the connecitons
    for (i = 0; i < dictionaryLength; i++)
        AddLinks(i);

    // Find the begin and end words in the dictionary
    for (i = 0; i < dictionaryLength; i++)
    {
        if (strcmp(dictionary[i].value, beginWord) == 0)
            beginWordIndex = i;
        if (strcmp(dictionary[i].value, endWord) == 0)
            endWordIndex = i;
    }

    // If the begin word is not in the dictionary then add it to the mesh.  End word is always in the dictionary.
    if (beginWordIndex == MAX_WORDS)
    {
        dictionary[dictionaryLength] = Word(beginWord, dictionaryLength);
        AddLinks(dictionaryLength);
        beginWordIndex = dictionaryLength;
        dictionaryLength++;
    }
    if (endWordIndex == MAX_WORDS)
    {
        printf("[]");
        exit(0);
    }

    // Traverse the mesh network to find all paths (start the first path)
    paths[0].node[0] = beginWordIndex;
    paths[0].length = 1;
    numPaths = 1;
    dictionary[beginWordIndex].visitedPath[0] = TRUE;
    TraverseMeshNetwork(beginWordIndex, endWordIndex, 0);
     
    // Print all the shortest paths
    printf("[");
    j = 0;
    for (i = 0; i < numPaths; i++)
    {
        if (paths[i].length == shortestPath && paths[i].reachedEnd)
        {
            printf("%S", (LPCTSTR)paths[i].GetName());
            j++;
            if (j<numShortest)
                printf(",\n");
        }
    }
    printf("]");
}


// Find all the paths from begin to end
void TraverseMeshNetwork(int beginWordIndex, int endWordIndex, int path)
{
    int i, j, node, oldPath;
    int numLinks = 0;
    BOOL addNode;
    BOOL alreadyInPath;
    int pathIn = path;

    dictionary[beginWordIndex].visitedPath[path] = TRUE;
    if (beginWordIndex == endWordIndex)
    {
        paths[path].reachedEnd = TRUE;
        if (paths[path].length < shortestPath)
        {
            shortestPath = paths[path].length;
            numShortest = 1;
        }
        else if (paths[path].length == shortestPath)
        {
            numShortest++;
        }
        return;
    }

    for (i = 0; i < dictionaryLength; i++)
    {
        addNode = FALSE;

        // Add to path if there is a link and have not visited for this path
        if (dictionary[beginWordIndex].linkDistance[i] > 0 && !dictionary[i].visitedPath[path] && numLinks == 0 && !paths[path].reachedEnd)
        {
            addNode = TRUE;
        }
        // Create a new path for any additional links
        else if (dictionary[beginWordIndex].linkDistance[i] > 0 && numLinks > 0)
        {
            // Create a new path if we are beyond the first link and copy the whole path up to this point
            oldPath = path;
            path = numPaths;
            j = 0;
            node = paths[oldPath].node[j];
            alreadyInPath = FALSE;
            for (;;) // loop forever
            {
                paths[path].node[j] = node;
                paths[path].length++;
                if (node == i)
                {
                    alreadyInPath = TRUE;
                    break;
                }
                if (node == beginWordIndex)  // loop until the begin node is reached
                    break;
                j++;
                node = paths[oldPath].node[j];
            }

            // If the new path is already longer or equal to the shortest (even without adding a node) then just kill it
            // Or this path already had the node kill it
            if (paths[path].length >= shortestPath || alreadyInPath)
            {
                paths[path].length = 0;
                for (j = 0; j < paths[path].length; j++)
                    paths[path].node[j] = MAX_WORDS;
                path = pathIn;
            }
            else
            {
                // Mark nodes on the new path as visited
                for (j = 0; j < paths[path].length; j++)
                {
                    node = paths[path].node[j];
                    dictionary[node].visitedPath[path] = TRUE;
                }
                numPaths++;
                addNode = TRUE;
            }
        }

        // Add the node to the existing path
        if (addNode && paths[path].length < shortestPath)
        {
            paths[path].node[paths[path].length] = i;  // distance to adjacent node always 1 so length can be used for index
            paths[path].length++;
            dictionary[i].visitedPath[path] = TRUE;
            numLinks++;
            TraverseMeshNetwork(i, endWordIndex, path);
        }
    }
}

// Add all the links of the mesh to a word in the dictionary
void AddLinks(int index)
{
    int i;
    for (i = 0; i < dictionaryLength; i++)
    {
        if (i == index)
        {
            dictionary[index].linkDistance[i] = 0;  // distance to itself is zero
            continue; // skip to next node
        }
        if (Linked(dictionary[i].value, dictionary[index].value))  // check for link
        {
            // Add the link
            dictionary[index].linkDistance[i] = 1;  // all adjacent distances are one unit
        }
    }
}


// Compare two words to see if they are linked
// Definition of linked is that every letter of one word is the same except for a max on one letter
// Assumes that both words have the same length
BOOL Linked(char one[], char two[])
{
    int i, mismatch = 0, len = (int)strlen(one);

    for (i = 0; i < len; i++)
    {
        if (one[i] != two[i])
            mismatch++;
    }
    if (mismatch == 1)
        return(TRUE);
    else
        return(FALSE);
}


// Get output string for a path
CString Path::GetName()
{
    CString output = "[";
    for (int i = 0; i < length; i++)
    {
        output = output + "\"" + CString(dictionary[node[i]].value) + "\"";
        if (i == length - 1)
            output = output + "]";
        else
            output = output + ",";
    }
    return(output);
}
