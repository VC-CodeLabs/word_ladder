requires GOTMPDIR setting  
bash: `export GOTMPDIR=~/Projects/tmp`  
cmd: `set GOTMPDIR=%USERPROFILE%\Projects\tmp`  
powershell: `$env:GOTMPDIR="$env:USERPROFILE\Projects\tmp"`

# Implementation Notes

The solution works by building out a tree of all possible word progression paths,
deriving the ladder(s) from the shortest paths that culminate in the end word.

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