package day7_test

import (
	"aoc24/day7"
	"aoc24/lib"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func TestPartOne(t *testing.T) {
	ops := []day7.Op{day7.OpAdd, day7.OpMul}
	lib.PrintResult(t, 7, 1, runMatch(t, ops), 2941973819040)
}

func TestPartTwo(t *testing.T) {
	ops := []day7.Op{day7.OpAdd, day7.OpMul, day7.OpCat}
	lib.PrintResult(t, 7, 2, runMatch(t, ops), 249943041417600)
}

func runMatch(t *testing.T, ops []day7.Op) (sum int) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	workers := runtime.NumCPU()
	taskC := make(chan day7.Task, workers)
	resultC := make(chan int, workers)
	wgTasks := &sync.WaitGroup{}
	wgTasks.Add(workers)
	wgResult := &sync.WaitGroup{}
	wgResult.Add(1)

	for range workers {
		go matchWorker(wgTasks, taskC, resultC, ops)
	}

	go func() {
		for value := range resultC {
			sum += value
		}
		wgResult.Done()
	}()

	for task := range lib.ReadInput(file, parseLine) {
		taskC <- task
	}

	close(taskC)
	wgTasks.Wait()
	close(resultC)
	wgResult.Wait()
	return
}

func matchWorker(wg *sync.WaitGroup, c chan day7.Task, r chan int, ops []day7.Op) {
	for task := range c {
		if day7.MatchExpr(task, ops) {
			r <- task.Result
		}
	}
	wg.Done()
}

const exampledata = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestExample(t *testing.T) {
	actual := 0
	for task := range lib.ReadInput(strings.NewReader(strings.TrimSpace(exampledata)), parseLine) {
		if day7.MatchExpr(task, []day7.Op{day7.OpAdd, day7.OpMul}) {
			actual += task.Result
		}
	}
	if actual != 3749 {
		t.Error("unexpected", actual)
	}
}

func TestExample2(t *testing.T) {
	actual := 0
	for task := range lib.ReadInput(strings.NewReader(strings.TrimSpace(exampledata)), parseLine) {
		if day7.MatchExpr(task, []day7.Op{day7.OpAdd, day7.OpMul, day7.OpCat}) {
			actual += task.Result
		}
	}
	if actual != 11387 {
		t.Error("unexpected", actual)
	}
}

func parseLine(s string) (task day7.Task) {
	result, numsLine, _ := strings.Cut(s, ":")
	task.Result = lib.MustParse(result)
	for _, value := range strings.Fields(numsLine) {
		task.Nums = append(task.Nums, lib.MustParse(value))
	}
	return
}
