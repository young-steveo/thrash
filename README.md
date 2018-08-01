# Thrash

**Thrash** will eventually be an implementation of the [Rockstar](https://github.com/dylanbeattie/rockstar) programming language by [@dylanbeattie](https://github.com/dylanbeattie).  I'm thinking a bytecode VM written in Go is sufficiently metal enough, so that's the plan.

So far I've handwritten a Lexer that's mostly complete.  An AST parser is next, then I'll spit out some bytecode and write a simple VM to consume it.

## Installation

```bash
$ git checkout https://github.com/young-steveo/thrash.git
$ cd thrash
$ go build
$ thrash -f ~/path/to/lyrics.rock
```

## Flags

* `-f` rockstar script to execute. default: `./main.rock`
* `-r` start a REPL.  (TODO)
* `-d` debug output. Currently spits out tokens from the lexer.
