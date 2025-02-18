// Copyright 2025, Pulumi Corporation.  All rights reserved.

package util

import (
	"fmt"
	"time"
)

// RetryOperation repeats the given operation until it succeeds, fails, or times out.
// The given operation should return true if the operation is done, false if it should be retried.
func RetryOperation(timeout time.Duration, interval time.Duration, description string, operation func() (bool, error)) error {
	timeoutChan := time.After(timeout)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-timeoutChan:
			return fmt.Errorf("timed out during %s after %s", description, timeout)
		case <-ticker.C:
			done, err := operation()
			if err != nil {
				return fmt.Errorf("error during %s: %w", description, err)
			}
			if done {
				return nil
			}
		}
	}
}
