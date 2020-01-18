# rabn

A tool for storing frequency of selections from a list, intended to be used with [fzf], for example:

```
rabn -p <path> -H <history-file> | fzf +s
```

to list the contents of `path` sorted by the selection counts in `history-file` and displaying the list via fzf.

[fzf]: (https://github.com/junegunn/fzf)
