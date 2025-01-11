package core

import "sync"

type Task = func() (interface{}, error)

type TaskResult struct {
	Result interface{}
	Error  error
}

func RunConcurrentTasks(tasks []Task) []TaskResult {
	var wg sync.WaitGroup
	results := make([]TaskResult, len(tasks))

	wg.Add(len(tasks))

	for i, task := range tasks {
		go func(idx int, t func() (interface{}, error)) {
			defer wg.Done()
			result, err := t()
			results[idx] = TaskResult{
				Result: result,
				Error:  err,
			}
		}(i, task)
	}

	wg.Wait()

	return results
}
