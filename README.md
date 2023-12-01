# Advent of Code 2023

Last year I decided I would try advent of code, solved the first two days and then sort of gave up.

This year will be different™️!

## Note the odd use of `pacakage main`

All cmd/puzzleXX directories are their own `main` package. Luckily Go is smart (apparently) and you can just cd into those directories and `go run .` and everything is scoped to that directory. I was a little worried that puzzle01::parseLine would leak into the main package of puzzle02, but it doesn't.

I'm pretty sure this is not how you're supposed to do things but the point was to learn, and this works for now. I'll refactor if I have to but it's AoC, not a production application.
