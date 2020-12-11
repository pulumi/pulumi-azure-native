// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestMessagingTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs1(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs2(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs3(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs4(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs5(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs6(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs7(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs8(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs9(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}
func TestMessagingTs10(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/azure-nextgen",
		},
	})

	return baseJS
}
