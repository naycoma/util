package util_test

import (
	"fmt"
	"testing"

	"github.com/naycoma/util"
	"github.com/stretchr/testify/assert"
)

func TestErrorAs(t *testing.T) {
	a := assert.New(t)

	// Test case 1: Target error type matches
	err := &MyError{Code: 100, Message: "Something went wrong"}
	asErr, ok := util.ErrorAs[*MyError](err)
	a.True(ok)
	a.Equal(err, asErr)

	// Test case 2: Target error type does not match
	asErr2, ok2 := util.ErrorAs[*AnotherError](err)
	a.False(ok2)
	a.Nil(asErr2)

	// Test case 3: Error in chain matches
	wrapped := fmt.Errorf("wrapped error: %w", err)
	asErr3, ok3 := util.ErrorAs[*MyError](wrapped)
	a.True(ok3)
	a.Equal(err, asErr3)

	// Test case 4: Nil error
	asErr4, ok4 := util.ErrorAs[*MyError](nil)
	a.False(ok4)
	a.Nil(asErr4)
}

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

type AnotherError struct {
	Reason string
}

func (e *AnotherError) Error() string {
	return fmt.Sprintf("Reason: %s", e.Reason)
}