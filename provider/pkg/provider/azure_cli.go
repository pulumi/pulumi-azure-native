// Copyright 2016-2025, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// Subscription represents a Subscription from the Azure CLI
type Subscription struct {
	EnvironmentName string `json:"environmentName"`
	ID              string `json:"id"`
	IsDefault       bool   `json:"isDefault"`
	Name            string `json:"name"`
	State           string `json:"state"`
	TenantID        string `json:"tenantId"`
	User            *User  `json:"user"`
}

// User represents a User from the Azure CLI
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type subscriptionUnavailableError struct {
	message string
}

func (e *subscriptionUnavailableError) Error() string {
	return e.message
}

func newSubscriptionUnavailableError(message string) error {
	return &subscriptionUnavailableError{message}
}

type azSubscriptionProvider func(ctx context.Context, subscriptionID string) (*Subscription, error)

// cliTimeout is the default timeout for authentication attempts via CLI tools
const cliTimeout = 10 * time.Second

// defaultAzSubscriptionProvider invokes the Azure CLI to acquire a subscription. It assumes
// callers have verified that all string arguments are safe to pass to the CLI.
// this code is derived from "CLI token provider" code in the Azure SDK for Go:
// https://github.com/Azure/azure-sdk-for-go/blob/519e8ab1a0e433b755c31ebaa6b177dfc83cb838/sdk/azidentity/azure_cli_credential.go#L117-L172
var defaultAzSubscriptionProvider = func(ctx context.Context, subscriptionID string) (*Subscription, error) {
	// set a default timeout for this authentication iff the application hasn't done so already
	// var cancel context.CancelFunc
	// if _, hasDeadline := ctx.Deadline(); !hasDeadline {
	// 	ctx, cancel = context.WithTimeout(ctx, cliTimeout)
	// 	defer cancel()
	// }
	// ctx = context.Background()

	commandLine := "az account show -o json "
	if subscriptionID != "" {
		// subscription needs quotes because it may contain spaces
		commandLine += ` --subscription "` + subscriptionID + `"`
	}
	logging.V(9).Infof("Running command: %s", commandLine)

	var cliCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		dir := os.Getenv("SYSTEMROOT")
		if dir == "" {
			return nil, newSubscriptionUnavailableError("environment variable 'SYSTEMROOT' has no value")
		}
		cliCmd = exec.CommandContext(ctx, "cmd.exe", "/c", commandLine)
		cliCmd.Dir = dir
	} else {
		cliCmd = exec.CommandContext(ctx, "/bin/sh", "-c", commandLine)
		cliCmd.Dir = "/bin"
	}
	cliCmd.Env = os.Environ()
	var stdout, stderr bytes.Buffer
	cliCmd.Stderr = &stderr
	cliCmd.Stdout = &stdout
	cliCmd.WaitDelay = 100 * time.Millisecond

	output, err := func() ([]byte, error) {
		err := cliCmd.Run()
		stdout := stdout.Bytes()
		if errors.Is(err, exec.ErrWaitDelay) && len(stdout) > 0 {
			// The child process wrote to stdout and exited without closing it.
			// Swallow this error and return stdout because it may contain output.
			return stdout, nil
		}
		return stdout, err
	}()
	if err != nil {
		msg := stderr.String()
		logging.Errorf("az command error %v: %s", err, msg)
		var exErr *exec.ExitError
		if errors.As(err, &exErr) && exErr.ExitCode() == 127 || strings.HasPrefix(msg, "'az' is not recognized") {
			msg = "Azure CLI not found on path"
		}
		if msg == "" {
			msg = err.Error()
		}
		return nil, newSubscriptionUnavailableError(msg)
	}
	logging.V(9).Infof("Command output: %s", output)
	s := Subscription{}
	err = json.Unmarshal(output, &s)
	if err != nil {
		return nil, newSubscriptionUnavailableError(err.Error())
	}

	return &s, nil
}
