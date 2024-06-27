from collections import deque


# For keeping track of a word to it's parents
class Word_Node: 
    def __init__(self, word):
        self.word = word
        self.parents = []

# Use beginWord and endWord and use a queue to find all possible sequences
def word_ladder(beginWord: str, endWord: str, wordList: list) -> list:
    valid_ladders = []
    children_words = deque()
    seen_words = set()
    
    parent_to_children = check_one_char_away(wordList, beginWord) 
    root = Word_Node(beginWord)
    seen_words.add(beginWord)
    starting_children = parent_to_children.get(beginWord, [])

    if not starting_children or endWord not in wordList:
        return []

    for child in starting_children:
        child.parents.append(root)    
    
    children_words.append(starting_children) 

    solution_level_found = False
    while not solution_level_found:
        # Explore a level
        explored_children = []
        while children_words:
            children = children_words.popleft()
            for child in children:
                if child.word not in seen_words:
                    if child.word == endWord:
                        solution_level_found = True
                        path = [child.word]
                        
                        # Remove just one parent if more than one so other path can be constructed
                        if len(child.parents) > 1:
                            parent = child.parents.pop(0)
                        else:
                            parent = child.parents[0]
                        
                        path.append(parent.word)
                        while parent.parents:
                            if len(parent.parents) > 1:    
                                parent = parent.parents.pop()
                            else:
                                parent = parent.parents[0]
                            
                            path.append(parent.word)
                        
                        valid_ladders.append(path[::-1])
                        continue              
                    
                    explored_children.append(child)
        
        for child in explored_children:
            seen_words.add(child.word)
            new_children_nodes = parent_to_children[child.word]
            for new_child in new_children_nodes:
                new_child.parents.append(child)
            children_words.append(new_children_nodes)    
    
    return valid_ladders



# Preprocces by determing all possible alterations of a word and verifying if it's in wordList  
def check_one_char_away(wordList: list, beginWord: str) -> dict:
    one_char_away_words = {}
    word_set = set(wordList)
    word_set.add(beginWord)

    for word in word_set:
        for next_word in word_set:
            word_chars = list(next_word)
            i = 0 # Char index
            difference = 0       
            for char in word:
                if char != word_chars[i]:
                    difference += 1 
                i += 1    
            if difference == 1:
                child = Word_Node(next_word)
                # String of word corresponding to word_node object children
                if word in one_char_away_words:
                    one_char_away_words[word].append(child)
                else:
                    one_char_away_words[word] = [child]
        
    return one_char_away_words


# Test cases
beginWord = 'stone'
endWord = 'money'
wordList = ["stoke", "stony", "stome", "stomy", "stoey", "htoey", "htney", "hiney", "miney", "ttney", "toney", "itoey", "mtney", "soney", "store", "storm", "story", "monte", "monny", "monty", "money", "stane", "stine", "maney", "honey", "monde", "stnny", "mtone", "mtnne", "monne", "monee", "stnne", "mtnee"]

print(word_ladder(beginWord, endWord, wordList))