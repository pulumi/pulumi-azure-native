// Copyright 2021, Pulumi Corporation.  All rights reserved.

package debug

import "fmt"

// TODO This is a hack. Consider using glog or another logging library instead.

// Debug determines if debug logging is turned on
var Debug *bool

// Log logs debug info to stdout if debug.Debug is set to true
func Log(format string, args ...interface{}) {
	if Debug != nil && *Debug {
		fmt.Printf(format, args...)
	}
}
