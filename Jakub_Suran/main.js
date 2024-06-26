/* Solution to word ladder problem using bidirectional BFS algorithm.
 */

class Node {
  constructor(element, parent) {
    this.element = element;
    this.parent = parent;
  }
}

/**
 * Checks if two words differ by exactly one character.
 *
 * @param {string} word1 - The first word.
 * @param {string} word2 - The second word.
 * @returns {boolean} Returns true if the two words differ by exactly one character, otherwise returns false.
 */
function canStep(word1, word2) {
  if (word1.length != word2.length) return false;

  let diff = 0;
  for (let i = 0; i < word1.length; i++) {
    if (word1[i] != word2[i]) diff++;
  }
  return diff == 1;
}

/**
 * Creates a dictionary of possible steps for each word in the given word list.
 *
 * @param {string[]} wordList - The list of words.
 * @returns {Object} - The dictionary of possible steps for each word.
 */
function createPossibleStepDict(wordList) {
  const dict = {};
  for (let i = 0; i < wordList.length; i++) {
    dict[wordList[i]] = [];
    for (let j = 0; j < wordList.length; j++) {
      if (canStep(wordList[i], wordList[j])) {
        dict[wordList[i]].push(wordList[j]);
      }
    }
  }
  return dict;
}

/**
 * Expand the tree by one level by finding all the children of the nodes in the previous level.
 *
 * @param {Node[]} prevLevel - The level of nodes to process.
 * @param {Object} stepsDict - A dictionary containing the steps for each node.
 * @param {Set} visited - A set containing the visited nodes.
 * @returns {Node[]} - An array of nodes representing the next level of the tree.
 */
function expandTreeLevel(prevLevel, stepsDict, visited) {
  const nextLevel = [];
  for (const node of prevLevel) {
    for (const child of stepsDict[node.element]) {
      if (!visited.has(child)) nextLevel.push(new Node(child, node));
    }
  }
  for (const node of nextLevel) {
    visited.add(node.element);
  }
  return nextLevel;
}

/**
 * Checks the levels of forward and backward tree and returns all nodes that match.
 *
 * @param {Array} forwardLevel - The array of forward nodes.
 * @param {Array} backwardLevel - The array of backward nodes.
 * @returns {Array} - An array of matching pairs of nodes.
 */
function matchLevels(forwardLevel, backwardLevel) {
  let result = [];
  const map = new Map();
  for (const forNode of forwardLevel) {
    if (!map.has(forNode.element)) map.set(forNode.element, new Set([forNode]));
    else map.get(forNode.element).add(forNode);
  }
  for (const backNode of backwardLevel) {
    if (map.has(backNode.element)) {
      for (const forNode of map.get(backNode.element))
        result.push([forNode, backNode]);
    }
  }
  return result;
}

/**
 * Finds all shortest paths from a given beginWord to a given endWord using a word ladder.
 * Note: remove comments to see the time analysis.
 *
 * @param {string} beginWord - The starting word.
 * @param {string} endWord - The target word.
 * @param {string[]} wordList - The list of valid words.
 * @returns {string[][]} - An array of arrays representing the shortest paths from beginWord to endWord.
 */
function wordLadder(beginWord, endWord, wordList) {
  if (!wordList.includes(endWord)) return [];

  const forwardLevels = [];
  const backwardLevels = [];
  const visitedForward = new Set([beginWord]);
  const visitedBackward = new Set([endWord]);

  // add beginWord to queue with an empty path
  let forwardLevel = [new Node(beginWord, null)];
  let backwardLevel = [new Node(endWord, null)];

  let matches = [];
  let queueSwitch = false;

  if (!wordList.includes(beginWord)) wordList.push(beginWord);
  let stepsDict = createPossibleStepDict(wordList);

  // const startTime = Date.now();
  // const analysis = {
  //   expandTreeLevel: 0,
  //   matches: 0,
  // };
  // bfs
  while (forwardLevel.length > 0 && backwardLevel.length > 0) {
    // const processStart = Date.now();
    if (queueSwitch) {
      forwardLevels.push(forwardLevel);
      forwardLevel = expandTreeLevel(forwardLevel, stepsDict, visitedForward);
    } else {
      backwardLevels.push(backwardLevel);
      backwardLevel = expandTreeLevel(
        backwardLevel,
        stepsDict,
        visitedBackward
      );
    }
    // const processEnd = Date.now();
    // analysis.expandTreeLevel += processEnd - processStart;
    queueSwitch = !queueSwitch;

    // const matchStart = Date.now();
    matches = matchLevels(forwardLevel, backwardLevel);
    // const matchEnd = Date.now();
    // analysis.matches += matchEnd - matchStart;

    if (matches.length > 0) break;
  }
  // const endTime = Date.now();
  // const bfsTime = endTime - startTime;

  let results = [];
  // const pathStartTime = Date.now();
  for (const match of matches) {
    let forwardPath = [];
    let backwardPath = [];
    let forwardNode = match[0];
    let backwardNode = match[1];

    while (forwardNode) {
      forwardPath.push(forwardNode.element);
      forwardNode = forwardNode.parent;
    }

    while (backwardNode) {
      backwardPath.push(backwardNode.element);
      backwardNode = backwardNode.parent;
    }

    forwardPath.reverse();
    results.push([
      ...forwardPath.slice(0, forwardPath.length - 1),
      ...backwardPath,
    ]);
  }
  // const pathEndTime = Date.now();
  // const pathTime = pathEndTime - pathStartTime;

  // console.log("BFS Time:", bfsTime, "ms");
  // console.log("Process Tree Level Time:", analysis.expandTreeLevel, "ms");
  // console.log("Get Matches Time:", analysis.matches, "ms");
  // console.log("Path Time:", pathTime, "ms");

  return results;
}

// example
const beginWord = "hit";
const endWord = "cog";
const wordList = ["hot", "dot", "dog", "lot", "log", "cog"];
const result = wordLadder(beginWord, endWord, wordList);
console.log(result);
