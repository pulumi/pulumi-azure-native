// Copyright 2016-2025, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
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

// defaultAzSubscriptionProvider invokes the Azure CLI to acquire a subscription. It assumes
// callers have verified that all string arguments are safe to pass to the CLI.
var defaultAzSubscriptionProvider = func(ctx context.Context, subscriptionID string) (*Subscription, error) {
	// Build command arguments as array to avoid shell quoting issues
	args := []string{"account", "show", "-o", "json"}
	if subscriptionID != "" {
		args = append(args, "--subscription", subscriptionID)
	}
	output, err := runAzCommand(ctx, args...)
	if err != nil {
		return nil, newSubscriptionUnavailableError(err.Error())
	}
	s := Subscription{}
	err = json.Unmarshal(output, &s)
	if err != nil {
		return nil, newSubscriptionUnavailableError(err.Error())
	}

	return &s, nil
}

type azCloudProvider func(ctx context.Context, cloudName string) (*cloud.Configuration, error)

// defaultAzCloudProvider invokes the Azure CLI to describe one of its clouds, such as a custom cloud registered with
// `az cloud register` (e.g. an Azure Stack Hub). It assumes callers have verified that the cloud name is safe to pass
// to the CLI.
var defaultAzCloudProvider = func(ctx context.Context, cloudName string) (*cloud.Configuration, error) {
	output, err := runAzCommand(ctx, "cloud", "show", "--name", cloudName, "-o", "json")
	if err != nil {
		return nil, err
	}
	return cloudConfigurationFromAzureCLI(output)
}

// azureCLICloud is the part of the `az cloud show` output that describes how to authenticate to and manage a cloud.
type azureCLICloud struct {
	Name      string `json:"name"`
	Endpoints struct {
		ActiveDirectory           string `json:"activeDirectory"`
		ActiveDirectoryResourceID string `json:"activeDirectoryResourceId"`
		ResourceManager           string `json:"resourceManager"`
		MicrosoftGraphResourceID  string `json:"microsoftGraphResourceId"`
	} `json:"endpoints"`
	Suffixes struct {
		StorageEndpoint string `json:"storageEndpoint"`
		KeyVaultDNS     string `json:"keyvaultDns"`
	} `json:"suffixes"`
}

// cloudConfigurationFromAzureCLI converts the output of `az cloud show -o json` into a cloud configuration. The
// endpoints are normalized to the form of the azcore cloud configurations, so that the public cloud read from the
// Azure CLI equals cloud.AzurePublic.
func cloudConfigurationFromAzureCLI(output []byte) (*cloud.Configuration, error) {
	var cliCloud azureCLICloud
	if err := json.Unmarshal(output, &cliCloud); err != nil {
		return nil, fmt.Errorf("reading the cloud described by the Azure CLI: %w", err)
	}
	for _, required := range []struct{ name, value string }{
		{"activeDirectory", cliCloud.Endpoints.ActiveDirectory},
		{"activeDirectoryResourceId", cliCloud.Endpoints.ActiveDirectoryResourceID},
		{"resourceManager", cliCloud.Endpoints.ResourceManager},
	} {
		if required.value == "" {
			return nil, fmt.Errorf("the Azure CLI cloud %q has no %q endpoint", cliCloud.Name, required.name)
		}
	}

	return &cloud.Configuration{
		Name: cliCloud.Name,
		Configuration: azcloud.Configuration{
			ActiveDirectoryAuthorityHost: strings.TrimSuffix(cliCloud.Endpoints.ActiveDirectory, "/") + "/",
			Services: map[azcloud.ServiceName]azcloud.ServiceConfiguration{
				azcloud.ResourceManager: {
					Audience: cliCloud.Endpoints.ActiveDirectoryResourceID,
					Endpoint: strings.TrimSuffix(cliCloud.Endpoints.ResourceManager, "/"),
				},
			},
		},
		Endpoints: cloud.ConfigurationEndpoints{
			MicrosoftGraph: cliCloud.Endpoints.MicrosoftGraphResourceID,
		},
		Suffixes: cloud.ConfigurationSuffixes{
			StorageEndpoint: cliCloud.Suffixes.StorageEndpoint,
			KeyVaultDNS:     cliCloud.Suffixes.KeyVaultDNS,
		},
	}, nil
}

// runAzCommand invokes the Azure CLI with the given arguments and returns its standard output.
// This code is derived from "CLI token provider" code in the Azure SDK for Go:
// https://github.com/Azure/azure-sdk-for-go/blob/519e8ab1a0e433b755c31ebaa6b177dfc83cb838/sdk/azidentity/azure_cli_credential.go#L117-L172
func runAzCommand(ctx context.Context, args ...string) ([]byte, error) {
	logging.V(9).Infof("Running command: az %s", strings.Join(args, " "))

	cliCmd := exec.CommandContext(ctx, "az", args...)
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
		return nil, errors.New(msg)
	}
	logging.V(9).Infof("Command output: %s", output)
	return output, nil
}
