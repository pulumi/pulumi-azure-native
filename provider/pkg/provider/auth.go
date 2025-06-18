// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	goversion "github.com/hashicorp/go-version"
)

func getAzVersion() (*goversion.Version, error) {
	_, err := exec.LookPath("az")
	if err != nil {
		return nil, fmt.Errorf("could not find `az`: %w", err)
	}

	var azVersion struct {
		Cli string `json:"azure-cli"`
	}
	err = runAzCmd(&azVersion, "version")
	if err != nil {
		return nil, fmt.Errorf("could not determine az version: %w", err)
	}

	actual, err := goversion.NewVersion(azVersion.Cli)
	if err != nil {
		return nil, fmt.Errorf("could not parse az version \"%q\": %w", azVersion.Cli, err)
	}

	return actual, nil
}

func assertAzVersion(version *goversion.Version) error {
	const versionHint = `Please make sure that the Azure CLI is installed in a version of at least 2.37 but less than
3.x; or configure another authentication method. See
https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#credentials
for more information.`

	// We need this version because it doesn't print the error of #1565
	versionRange := goversion.MustConstraints(goversion.NewConstraint(">=2.37.0, <3"))

	if !versionRange.Check(version) {
		return fmt.Errorf("found incompatible az version %s. %s", version, versionHint)
	}
	return nil
}

// Run `az` with the given args and unmarshal the output into the given target struct.
// The `az` output is always in JSON, requested via `--output json`.
func runAzCmd(target interface{}, arg ...string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	arg = append(arg, "--output", "json")

	cmd := exec.Command("az", arg...)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		err := fmt.Errorf("running az: %w", err)
		if stdErrStr := stderr.String(); stdErrStr != "" {
			err = fmt.Errorf("%s: %s", err, strings.TrimSpace(stdErrStr))
		}
		return err
	}

	if err := json.Unmarshal(stdout.Bytes(), target); err != nil {
		return fmt.Errorf("unmarshaling the result of Azure CLI: %w", err)
	}

	return nil
}
