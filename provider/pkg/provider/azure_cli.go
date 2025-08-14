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
)

// Subscription represents a Subscription from the Azure CLI
// https://github.com/Azure/go-autorest/blob/33e12ab7683c1c236a863ccfbfdd78c626f7fe28/autorest/azure/cli/profile.go#L35
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
var defaultAzSubscriptionProvider = func(ctx context.Context, subscriptionID string) (*Subscription, error) {
	// set a default timeout for this authentication iff the application hasn't done so already
	var cancel context.CancelFunc
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		ctx, cancel = context.WithTimeout(ctx, cliTimeout)
		defer cancel()
	}
	commandLine := "az account show -o json "
	if subscriptionID != "" {
		// subscription needs quotes because it may contain spaces
		commandLine += ` --subscription "` + subscriptionID + `"`
	}
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
	var stderr bytes.Buffer
	cliCmd.Stderr = &stderr

	output, err := cliCmd.Output()
	if err != nil {
		msg := stderr.String()
		var exErr *exec.ExitError
		if errors.As(err, &exErr) && exErr.ExitCode() == 127 || strings.HasPrefix(msg, "'az' is not recognized") {
			msg = "Azure CLI not found on path"
		}
		if msg == "" {
			msg = err.Error()
		}
		return nil, newSubscriptionUnavailableError(msg)
	}

	s := Subscription{}
	err = json.Unmarshal(output, &s)
	if err != nil {
		return nil, newSubscriptionUnavailableError(err.Error())
	}

	return &s, nil
}
