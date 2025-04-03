### Advent of Code 2024

Puzzle solutions are tests named `TestPartOne` and `TestPartTwo` inside each day directory;  
there is no main packages to run.

Run tests for a specific day and print results:

```
go test -v -run TestPart ./day1
```

Run all tests with pretty print:

```
go test -v ./... | grep Day | cut -d : -f 3- | sort -hk 2
```