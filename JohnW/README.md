# Word Ladder

This does a breadth-first search to find the shortest path between 2 words, given a list of valid words.

To compile:

```bash
javac Ladder.java
```

To run:
```bash
java -cp . Ladder
```

### Modifying start and end words
At the top of the Ladder class is a list of candidate tests. Each test is an instatiation of a `Words` object.
A `Words` object contains a start word, end word, and the list of valid words.

You can easily change which test runs by changing the 'testWords' 
variable in the main() method:

```java
      Words testWords = Ladder.manyWords;
```
