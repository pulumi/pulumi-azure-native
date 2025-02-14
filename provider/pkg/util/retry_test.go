// Copyright 2025, Pulumi Corporation.  All rights reserved.

package util

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPollOperation(t *testing.T) {
	t.Parallel()

	t.Run("SuccessBeforeTimeout", func(t *testing.T) {
		t.Parallel()
		operation := func() (bool, error) {
			return true, nil
		}

		err := RetryOperation(2*time.Second, 500*time.Millisecond, "test operation", operation)
		assert.NoError(t, err, "expected no error")
	})

	t.Run("TimesOut", func(t *testing.T) {
		t.Parallel()
		operation := func() (bool, error) {
			return false, nil
		}

		err := RetryOperation(1*time.Second, 500*time.Millisecond, "test operation", operation)
		assert.Error(t, err, "expected timeout error")
		assert.EqualError(t, err, "timed out during test operation after 1s")
	})

	t.Run("ErrorDuringOperation", func(t *testing.T) {
		t.Parallel()
		expectedErr := errors.New("operation error")
		operation := func() (bool, error) {
			return true, expectedErr
		}

		err := RetryOperation(2*time.Second, 500*time.Millisecond, "test operation", operation)
		assert.Error(t, err, "expected operation error")
		assert.True(t, errors.Is(err, expectedErr), "expected error to match")
	})
}
