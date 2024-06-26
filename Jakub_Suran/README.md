# Word Ladder Solution

I solved the problem by using bidirectional bredth-first search. Here are the main ideas:

### Main ideas

- **Preprocessing:** all possible steps (one character modifications) are constructed for each word in the wordList and stored as a dictionary for O(1) lookup
- **Bidirectional BFS:** since the word ladder is a shortest path in unweighted graph problem, a simple BFS algorith can be used to solve it. Aditionally it can be optimized by using a second BFS from the target which significantly cuts down the search space
- **Matching by levels:** to make sure that the solution finds **all** shorest paths I grew both BFS trees by the whole level at a time, one tree at a time. After each expansion I checked for all matches in the top/bottom most levels of the trees. If more then 0 matches occured the solution was found.
