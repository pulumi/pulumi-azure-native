// Copyright 2026, Pulumi Corporation.  All rights reserved.

package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ruleServiceBusApiVersion matches the API version servicebus:Rule (and therefore DefaultRule)
// currently defaults to.
const ruleServiceBusApiVersion = "2024-01-01"

// ruleClient reads the "$Default" rule directly via the ARM REST API, at ruleServiceBusApiVersion.
//
// Note: Azure has a confirmed behavior where GET on the "$Default" rule 404s with
// "SubscriptionNotFound: Rule does not exist" for an extended period (multiple minutes, at
// least - reproduced independent of client, connection, or retry count) once its filter is
// genuinely changed. Because of that, this client is only used here to check the rule's state
// *before* any modification (reliable) and its *absence* after deletion (also reliable - a 404
// is the expected outcome either way). Verifying a filter change is instead done by reading
// DefaultRule's own properties back from Pulumi's state (see the stack outputs below), which
// isn't subject to this limitation since it doesn't re-read the rule from Azure at all.
type ruleClient struct {
	t              *testing.T
	subscriptionID string
}

func newRuleClient(t *testing.T) *ruleClient {
	return &ruleClient{
		t:              t,
		subscriptionID: os.Getenv("ARM_SUBSCRIPTION_ID"),
	}
}

// getDefaultRule returns the HTTP status code and the parsed JSON body (if any) of a GET on the
// "$Default" rule.
func (c *ruleClient) getDefaultRule(ctx context.Context, resourceGroupName, namespaceName, topicName, subscriptionName string) (int, map[string]any) {
	c.t.Helper()

	token, err := azureCredentials(c.t).GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{"https://management.azure.com/.default"}})
	require.NoError(c.t, err)

	url := fmt.Sprintf(
		"https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceBus/namespaces/%s/topics/%s/subscriptions/%s/rules/$Default?api-version=%s",
		c.subscriptionID, resourceGroupName, namespaceName, topicName, subscriptionName, ruleServiceBusApiVersion)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	require.NoError(c.t, err)
	req.Header.Set("Authorization", "Bearer "+token.Token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(c.t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(c.t, err)

	var parsed map[string]any
	if len(body) > 0 {
		require.NoError(c.t, json.Unmarshal(body, &parsed), "response body: %s", body)
	}
	return resp.StatusCode, parsed
}

// assertDefaultRuleDeletedWithRetry retries the GET until it observes a 404 for the "$Default"
// rule, tolerating ordinary ARM propagation delay right after the delete call.
func assertDefaultRuleDeletedWithRetry(t *testing.T, ctx context.Context, client *ruleClient, resourceGroupName, namespaceName, topicName, subscriptionName string) {
	t.Helper()
	var lastStatus int
	for i := 0; i < 5; i++ {
		status, _ := client.getDefaultRule(ctx, resourceGroupName, namespaceName, topicName, subscriptionName)
		if status == http.StatusNotFound {
			return
		}
		lastStatus = status
		time.Sleep(2 * time.Second)
	}
	t.Fatalf("$Default rule was not deleted in time, last status: %d", lastStatus)
}

// TestDefaultRuleAddThenRemove_YAML verifies that the "$Default" rule Azure automatically
// creates alongside a ServiceBus subscription can be managed via the DefaultRule resource:
// its SqlFilter can be set in place with no import step required, and removing the resource
// from the program actually deletes the rule from Azure.
// See https://github.com/pulumi/pulumi-azure-native/issues/4489.
func TestDefaultRuleAddThenRemove_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	resources := map[string]any{
		"resourcegroup": map[string]any{
			"type": "azure-native:resources:ResourceGroup",
		},
		"namespace": map[string]any{
			"type": "azure-native:servicebus:Namespace",
			"properties": map[string]any{
				"resourceGroupName": "${resourcegroup.name}",
				"sku": map[string]any{
					"name": "Standard",
					"tier": "Standard",
				},
			},
		},
		"topic": map[string]any{
			"type": "azure-native:servicebus:Topic",
			"properties": map[string]any{
				"resourceGroupName": "${resourcegroup.name}",
				"namespaceName":     "${namespace.name}",
			},
		},
		"subscription": map[string]any{
			"type": "azure-native:servicebus:Subscription",
			"properties": map[string]any{
				"resourceGroupName": "${resourcegroup.name}",
				"namespaceName":     "${namespace.name}",
				"topicName":         "${topic.name}",
			},
		},
	}

	outputs := map[string]any{
		"resourceGroupName": "${resourcegroup.name}",
		"namespaceName":     "${namespace.name}",
		"topicName":         "${topic.name}",
		"subscriptionName":  "${subscription.name}",
	}

	program := map[string]any{
		"name":      proj.name,
		"runtime":   "yaml",
		"resources": resources,
		"outputs":   outputs,
		"plugins":   plugins,
	}

	// First deploy: create the namespace/topic/subscription. Azure creates the "$Default" rule
	// automatically as a side effect of creating the subscription; DefaultRule isn't declared yet.
	updatePulumiYAML(t, test.WorkingDir(), program)
	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "first up should not have any errors")
	defer test.Destroy(t)

	resourceGroupName := upResult.Outputs["resourceGroupName"].Value.(string)
	namespaceName := upResult.Outputs["namespaceName"].Value.(string)
	topicName := upResult.Outputs["topicName"].Value.(string)
	subscriptionName := upResult.Outputs["subscriptionName"].Value.(string)

	ctx := context.Background()
	client := newRuleClient(t)

	// This GET is against the rule's untouched, Azure-managed default state, which (unlike a GET
	// right after a genuine modification) is reliably readable - see the ruleClient doc comment.
	status, _ := client.getDefaultRule(ctx, resourceGroupName, namespaceName, topicName, subscriptionName)
	require.Equal(t, http.StatusOK, status, "the $Default rule should already exist once the subscription is created")

	// Second deploy: adopt the existing "$Default" rule and set a SqlFilter on it directly,
	// with no import step required.
	resources["defaultRule"] = map[string]any{
		"type": "azure-native:servicebus:DefaultRule",
		"properties": map[string]any{
			"resourceGroupName": "${resourcegroup.name}",
			"namespaceName":     "${namespace.name}",
			"topicName":         "${topic.name}",
			"subscriptionName":  "${subscription.name}",
			"filterType":        "SqlFilter",
			"sqlFilter": map[string]any{
				"sqlExpression": "InsightId IS NOT NULL",
			},
		},
	}
	outputs["ruleSqlExpression"] = "${defaultRule.sqlFilter.sqlExpression}"
	updatePulumiYAML(t, test.WorkingDir(), program)
	upResult = test.Up(t)
	assert.Empty(t, upResult.StdErr, "second up should not have any errors")

	// Verified from Pulumi's own state (the PUT's response), not by re-reading the rule from
	// Azure: see the ruleClient doc comment for why an immediate re-read isn't reliable here.
	assert.Equal(t, "InsightId IS NOT NULL", upResult.Outputs["ruleSqlExpression"].Value)

	// A subsequent preview should show no changes, confirming the operation is idempotent.
	preview := test.Preview(t)
	t.Logf("Preview STDOUT: \n%s", preview.StdOut)
	assert.Equal(t, map[apitype.OpType]int{
		apitype.OpSame: 6, // stack + resourcegroup + namespace + topic + subscription + defaultRule
	}, preview.ChangeSummary)

	// Third deploy: remove the DefaultRule resource, which should actually delete the "$Default"
	// rule from Azure rather than just forgetting about it.
	delete(resources, "defaultRule")
	delete(outputs, "ruleSqlExpression")
	updatePulumiYAML(t, test.WorkingDir(), program)
	upResult = test.Up(t)
	assert.Empty(t, upResult.StdErr, "third up should not have any errors")

	assertDefaultRuleDeletedWithRetry(t, ctx, client, resourceGroupName, namespaceName, topicName, subscriptionName)
}

// TestDefaultRuleCreatedWithSubscription_YAML verifies that DefaultRule can be declared
// up front, in the same deploy as the Namespace/Topic/Subscription that it belongs to,
// rather than only being addable on a subsequent `pulumi up` once the subscription (and its
// auto-created "$Default" rule) already exist. Since Azure creates the "$Default" rule as a
// side effect of creating the Subscription, DefaultRule's Create still needs to adopt an
// already-existing rule even though, from Pulumi's point of view, this is the very first
// deploy of the stack.
// See https://github.com/pulumi/pulumi-azure-native/issues/4489.
func TestDefaultRuleCreatedWithSubscription_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"resources": map[string]any{
			"resourcegroup": map[string]any{
				"type": "azure-native:resources:ResourceGroup",
			},
			"namespace": map[string]any{
				"type": "azure-native:servicebus:Namespace",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"sku": map[string]any{
						"name": "Standard",
						"tier": "Standard",
					},
				},
			},
			"topic": map[string]any{
				"type": "azure-native:servicebus:Topic",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"namespaceName":     "${namespace.name}",
				},
			},
			"subscription": map[string]any{
				"type": "azure-native:servicebus:Subscription",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"namespaceName":     "${namespace.name}",
					"topicName":         "${topic.name}",
				},
			},
			"defaultRule": map[string]any{
				"type": "azure-native:servicebus:DefaultRule",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"namespaceName":     "${namespace.name}",
					"topicName":         "${topic.name}",
					"subscriptionName":  "${subscription.name}",
					"filterType":        "SqlFilter",
					"sqlFilter": map[string]any{
						"sqlExpression": "InsightId IS NOT NULL",
					},
				},
			},
		},
		"outputs": map[string]any{
			"ruleSqlExpression": "${defaultRule.sqlFilter.sqlExpression}",
		},
		"plugins": plugins,
	}

	// Single deploy: the namespace/topic/subscription and the DefaultRule are all declared from
	// the start. Even though this is the first `pulumi up` for this stack, DefaultRule's Create
	// still has to adopt the "$Default" rule that Azure creates automatically alongside the
	// subscription, rather than trying (and failing) to create it as if from scratch.
	updatePulumiYAML(t, test.WorkingDir(), program)
	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "up should not have any errors")
	defer test.Destroy(t)

	// Verified from Pulumi's own state (the PUT's response), not by re-reading the rule from
	// Azure: see the ruleClient doc comment above for why an immediate re-read isn't reliable.
	assert.Equal(t, "InsightId IS NOT NULL", upResult.Outputs["ruleSqlExpression"].Value)

	// A subsequent preview should show no changes, confirming the operation is idempotent.
	preview := test.Preview(t)
	t.Logf("Preview STDOUT: \n%s", preview.StdOut)
	assert.Equal(t, map[apitype.OpType]int{
		apitype.OpSame: 6, // stack + resourcegroup + namespace + topic + subscription + defaultRule
	}, preview.ChangeSummary)
}
