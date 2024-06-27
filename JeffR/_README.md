requires GOTMPDIR setting  
bash: `export GOTMPDIR=~/Projects/tmp`  
cmd: `set GOTMPDIR=%USERPROFILE%\Projects\tmp`  
powershell: `$env:GOTMPDIR="$env:USERPROFILE\Projects\tmp"`

# Implementation Notes

The solution works by building out a tree of all possible word progression paths,
deriving the ladder(s) from the shortest paths that culminate in the end word.

As each path is built, only the words not already in that path are considered
for extending the path.

For example, with the following input:
```
beginWord: hit
  endWord: cog
 wordList: [hot dot dog lot log cog]
```
The word tree would be:
```
hit *
 ╚══hot *
     ╠══dot *
     ║   ╠══dog * 
     ║   ║   ╚══cog *
     ║   ║   
     ║   ╚══lot 
     ║       ╚══log 
     ║           ╚══cog
     ║ 
     ╚══lot *
         ╠══dot 
         ║   ╚══dog 
         ║       ╚══cog
         ║      
         ╚══log *
             ╚══cog *
```
...which produces two shortest ladders at 5 steps each (*'d in tree above):
```
 [[hit hot dot dog cog] [hit hot lot log cog]]
```

# Command-line Parameters

Command-line help is available with `-h` parameter.

Command-line options:

### -b={beginWord}
### -e={endWord}
### -l={wordList:word1,word2[...wordN]}

e.g.
`go run JeffR_WordLadder_Solution.go -b=foo -e=bar -l=bar,far,for`

Output: 
`[[ "foo", "for", "far", "bar" ]]`

NOTE for especially long wordLists or expectedResults, use `-f` for json input instead

### -x="{expectedResult:ladder1:word1,word2[...wordN][;ladder2:word1,word2[...wordN]]}"

Use `-x=""` to specify no ladders expected

quotes can be omitted if only a single ladder is spec'd

NOTE expectedResult only needs to be spec'd if `-c` is also used

`-x={anything}` won't *ever* change the actual result; omit from json or set `-x=""` in case there's any doubt

e.g.
`go run JeffR_WordLadder_Solution.go -b=foo -e=fir -l=for,fio,fir -x="foo,for,fir;foo,fio,fir" -c`

Output: 
    
`[[ "foo", "fio", "fir" ], [ "foo", "for", fir" ]]`

`*CHECKED*: ==`

NOTE for especially long wordLists or expectedResults, use `-f` for json input instead

### -c check expected vs actual result, report first issue found

### -f={jsonInputFileSpec}

see test*.json for syntax

e.g.
`go run JeffR_WordLadder_Solution.go -f=testCase2.json`

Output:
`[[ "a", "c" ]]`

### -s sanitize input

Ensures input conforms to spec'd constraints; off by default

### -v enable verbose mode

### -d enable debug mode

### -o={0-3} change output mode

See Output Modes section below

### -t enable threading support

Use with large wordLists. Disabled by default. *EXPERIMENTAL- proceed with caution!* 

# Output Modes

- 0: Raw 
    
    the found ladder(s) are emitted w/ no sorting or formatting

- 1: Pretty Print

    the found ladder(s) are emitted in order by content; if ladders > 1, each is indented starting on a new line; words are quoted.
    *this is the default output mode.*

- 2: Graphical

    the found ladder(s) are emitted inside a ladder-like structure with the beginWord at the bottom.

- 3: Animated

    same as graphical, except each step word cycles thru the list before settling on the correct step word.

# Test Cases

The testMaxWords.json provides a contrived example w/ generated words; most paths are dead ends.  Even with optimizations, ~4.5 millions paths in the word tree are evaluated.  On my system it generally takes around a minute and a half to complete processing of this test. Enabling threading support drops this time to around a minute.

all other test*.json were derived from provided examples or actual test cases 