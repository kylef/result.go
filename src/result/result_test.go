package result

import (
  "testing"
  "github.com/stvp/assert"
)


type errorString struct {
  value string
}

func (err *errorString) Error() string {
  return err.value
}


// Test Initialization

func TestFailure(t *testing.T) {
  err := &errorString{"testing error"}
  result := NewFailure(err)

  assert.Nil(t, result.Success)
  assert.Equal(t, result.Failure, err)
}

func TestSuccess(t *testing.T) {
  value := 5
  result := NewSuccess(value)

  assert.Equal(t, result.Success, value)
  assert.Nil(t, result.Failure)
}


// Test Analysis

func TestFailureAnalysis(t *testing.T) {
  err := &errorString{"testing error"}
  result := NewFailure(err)

  resultantErr := &errorString{"testing new error"}
  resultantResult := result.Analysis(func(value interface{}) Result { return NewSuccess(value) },
                                     func(err error) Result { return NewFailure(resultantErr) })

  assert.Equal(t, resultantResult.Failure, resultantErr)
  assert.Nil(t, resultantResult.Success)
}

func TestSuccessAnalysis(t *testing.T) {
  result := NewSuccess(5)
  resultantResult := result.Analysis(func(value interface{}) Result { return NewSuccess(value.(int) * 2) },
                                     func(err error) Result { return NewFailure(err) })

  assert.Equal(t, resultantResult.Success, 10)
  assert.Nil(t, resultantResult.Failure)
}


// Test FlatMap

func TestFlatMapOnSuccessReturnsNewValue(t *testing.T) {
  result := NewSuccess(5)
  resultantResult := result.FlatMap(func(value interface{}) Result { return NewSuccess(value.(int) * 2) })

  assert.Equal(t, resultantResult.Success, 10)
  assert.Nil(t, resultantResult.Failure)
}

func TestFlatMapOnFailureReturnsFailure(t *testing.T) {
  err := &errorString{"testing error"}
  result := NewFailure(err)
  resultantResult := result.FlatMap(func(value interface{}) Result { return NewSuccess(value.(int) * 2) })

  assert.Equal(t, resultantResult.Failure, err)
  assert.Nil(t, resultantResult.Success)
}

