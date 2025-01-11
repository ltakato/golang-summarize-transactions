package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunConcurrentTasks(t *testing.T) {
	tasks := []func() (interface{}, error){
		func() (interface{}, error) {
			return "Task 1 Result", nil
		},
		func() (interface{}, error) {
			return nil, errors.New("task 2 Error")
		},
		func() (interface{}, error) {
			return 42, nil
		},
	}

	results := RunConcurrentTasks(tasks)
	result1, err1 := results[0].Result, results[0].Error
	_, err2 := results[1].Result, results[1].Error
	result3, err3 := results[2].Result, results[2].Error

	assert.Equal(t, 3, len(results), "Number of results should match the number of tasks")

	assert.NoError(t, err1, "Task 1 should not return an error")
	assert.Equal(t, "Task 1 Result", result1, "Task 1 result should be correct")

	assert.Error(t, err2, "Task 2 should return an error")
	assert.Equal(t, "task 2 Error", err2.Error(), "Task 2 error message should match")

	assert.NoError(t, err3, "Task 3 should not return an error")
	assert.Equal(t, 42, result3, "Task 3 result should be correct")
}
