// Copyright 2026, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	. "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const (
	namespaceName         = "namespaceName"
	topicName             = "topicName"
	subscriptionNameProp  = "subscriptionName"
	actionProp            = "action"
	correlationFilterProp = "correlationFilter"
	filterTypeProp        = "filterType"
	sqlFilterProp         = "sqlFilter"

	// defaultRuleName is the fixed name Azure gives to the rule it creates automatically
	// alongside every ServiceBus subscription.
	defaultRuleName = "$Default"
	defaultRuleTok  = "azure-native:servicebus:DefaultRule"
	defaultRulePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/" + defaultRuleName

	// defaultRuleFallbackApiVersion is used only if the default API version for servicebus:Rule
	// can't be looked up (e.g. in scoped/partial test builds).
	defaultRuleFallbackApiVersion = "2024-01-01"
	// defaultRuleSubscriptionFallbackApiVersion is the equivalent fallback for servicebus:Subscription,
	// used only when checking whether a DefaultRule's parent Subscription still exists.
	defaultRuleSubscriptionFallbackApiVersion = "2024-01-01"

	defaultRuleDescription = `The "$Default" rule that Azure automatically creates alongside every ServiceBus subscription.

By default this rule has a TrueFilter, which allows all messages through to the subscription. This resource lets
you manage that rule directly: update its filter (e.g. to a SqlFilter or CorrelationFilter), or delete this
resource to remove the rule entirely.

Unlike ` + "`Rule`" + `, this resource has no ` + "`ruleName`" + ` input since it always refers to the "$Default" rule that Azure
creates for you; there's no need to construct or import its resource ID by hand.

Note: deleting this resource permanently removes the "$Default" rule from the subscription. If no other rules
exist afterwards, the subscription will not receive any messages from the topic.

See [Issue #4489](https://github.com/pulumi/pulumi-azure-native/issues/4489) for more details on the problem this
resource solves.
`
)

// defaultRuleReadRetryDelay is the wait between retries in readDefaultRuleWithRetry.
// Tests are allowed to change this to a smaller value.
var defaultRuleReadRetryDelay = 2 * time.Second

// defaultRuleReadMaxAttempts bounds how long readDefaultRuleWithRetry waits before concluding
// that a 404 reflects a real deletion rather than transient ARM read-lag.
const defaultRuleReadMaxAttempts = 4

// readDefaultRuleWithRetry rides out ordinary ARM propagation delay right after a create or update
// of the "$Default" rule: a GET can 404 with "SubscriptionNotFound: Rule does not exist" for a
// brief window even though the write already succeeded. This retry is deliberately short - it is
// NOT a fix for the much longer-lived 404 Azure exhibits after the rule's filter is genuinely
// changed (confirmed to persist for multiple minutes, independent of client, connection, or retry
// count - see issue #4489). That longer-lived case is handled by the caller (defaultRule's Read),
// which falls back to checking whether the parent Subscription still exists before concluding the
// rule itself was deleted.
func readDefaultRuleWithRetry(ctx context.Context, read func(ctx context.Context) (map[string]any, error)) (map[string]any, bool, error) {
	for attempt := 0; ; attempt++ {
		response, err := read(ctx)
		if err == nil {
			return response, true, nil
		}
		if !azure.IsNotFound(err) {
			return nil, false, err
		}
		if attempt == defaultRuleReadMaxAttempts-1 {
			return nil, false, nil
		}
		select {
		case <-time.After(defaultRuleReadRetryDelay):
		case <-ctx.Done():
			return nil, false, ctx.Err()
		}
	}
}

// defaultRuleNotFoundOutcome decides what Read should report once the "$Default" rule's own GET
// has persistently 404'd (readDefaultRuleWithRetry gave up). Azure has a confirmed behavior where
// this rule becomes unreadable via GET for an extended period (multiple minutes, at least) once
// its filter is genuinely changed - reproduced independent of client, connection, or retry count,
// see issue #4489. A bare 404 here does not reliably mean the rule was deleted, so before
// concluding that, this checks whether the parent Subscription still exists: if it does, the rule
// almost certainly does too, and we don't have a safe way to guess its current property values (that's
// exactly what's unreadable) - so it fails instead of fabricating a result. A Create/Update's own
// read-after-write already tolerates and ignores Read errors (falling back to the PUT's own
// response), and an explicit `pulumi refresh` surfaces a clear error instead of silently corrupting
// or dropping state. Only when the Subscription is also gone is the rule reported as deleted.
func defaultRuleNotFoundOutcome(ctx context.Context, azureClient azure.AzureClient, subscriptionApiVersion, ruleID string) (map[string]any, bool, error) {
	subscriptionID := strings.TrimSuffix(ruleID, "/rules/"+defaultRuleName)
	_, err := azureClient.Get(ctx, subscriptionID, subscriptionApiVersion, nil)
	if err == nil {
		return nil, false, fmt.Errorf(
			"could not read $Default rule %q even though its parent subscription still exists; "+
				"this is a known Azure API limitation after modifying the rule's filter (see "+
				"https://github.com/pulumi/pulumi-azure-native/issues/4489) and should resolve "+
				"on its own - try again shortly", ruleID)
	}
	if azure.IsNotFound(err) {
		// The whole subscription (and therefore this rule) is genuinely gone.
		return nil, false, nil
	}
	return nil, false, err
}

// defaultRuleProperties are the filter-related properties shared between the auto-generated Rule
// resource and this hand-written one; they're not redefined here, only referenced by type token.
func defaultRule(
	lookupResource ResourceLookupFunc,
	crudClientFactory crud.ResourceCrudClientFactory,
	azureClient azure.AzureClient,
) (*CustomResource, error) {
	apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("ServiceBus", "Rule")
	if !ok {
		apiVersion = defaultRuleFallbackApiVersion
		logging.V(3).Infof("Warning: could not find default API version for servicebus:Rule. Using %s", apiVersion)
	}
	subscriptionApiVersion, ok := versionLookup.GetDefaultApiVersionForResource("ServiceBus", "Subscription")
	if !ok {
		subscriptionApiVersion = defaultRuleSubscriptionFallbackApiVersion
		logging.V(3).Infof("Warning: could not find default API version for servicebus:Subscription. Using %s", subscriptionApiVersion)
	}

	var client crud.ResourceCrudClient
	if lookupResource != nil && crudClientFactory != nil {
		res, found, err := lookupResource(defaultRuleTok)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("resource %q not found", defaultRuleTok)
		}
		client = crudClientFactory(&res)
	}

	return &CustomResource{
		tok:  defaultRuleTok,
		path: defaultRulePath,
		// DefaultRule has no equivalent in the Azure spec (unlike e.g. TagAtScope), so there's
		// nothing to derive the schema/metadata from - the `resource` argument is always nil.
		Schema: func(_ *ResourceDefinition) (*ResourceDefinition, error) {
			return &ResourceDefinition{
				Resource: schema.ResourceSpec{
					ObjectTypeSpec: schema.ObjectTypeSpec{
						Description: defaultRuleDescription,
						Type:        "object",
						Properties: map[string]schema.PropertySpec{
							actionProp: {
								TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:ActionResponse"},
								Description: "Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.",
							},
							"azureApiVersion": {
								TypeSpec:    schema.TypeSpec{Type: "string"},
								Description: "The Azure API version of the resource.",
							},
							correlationFilterProp: {
								TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:CorrelationFilterResponse"},
								Description: "Properties of correlationFilter",
							},
							filterTypeProp: {
								TypeSpec:    schema.TypeSpec{Type: "string"},
								Description: "Filter type that is evaluated against a BrokeredMessage.",
							},
							"location": {
								TypeSpec:    schema.TypeSpec{Type: "string"},
								Description: "The geo-location where the resource lives",
							},
							"name": {
								TypeSpec:    schema.TypeSpec{Type: "string"},
								Description: "The name of the resource",
							},
							sqlFilterProp: {
								TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:SqlFilterResponse"},
								Description: "Properties of sqlFilter",
							},
							"systemData": {
								TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:SystemDataResponse"},
								Description: "The system meta data relating to this resource.",
							},
							"type": {
								TypeSpec:    schema.TypeSpec{Type: "string"},
								Description: `The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"`,
							},
						},
						Required: []string{"azureApiVersion", "location", "name", "systemData", "type"},
					},
					InputProperties: map[string]schema.PropertySpec{
						actionProp: {
							TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:Action"},
							Description: "Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.",
						},
						correlationFilterProp: {
							TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:CorrelationFilter"},
							Description: "Properties of correlationFilter",
						},
						filterTypeProp: {
							TypeSpec:    schema.TypeSpec{Ref: "#/types/azure-native:servicebus:FilterType"},
							Description: "Filter type that is evaluated against a BrokeredMessage.",
						},
						namespaceName: {
							TypeSpec:             schema.TypeSpec{Type: "string"},
							Description:          "The namespace name",
							WillReplaceOnChanges: true,
						},
						resourceGroupName: {
							TypeSpec:             schema.TypeSpec{Type: "string"},
							Description:          "The name of the resource group. The name is case insensitive.",
							WillReplaceOnChanges: true,
						},
						sqlFilterProp: {
							TypeSpec:    schema.TypeSpec{Type: "object", Ref: "#/types/azure-native:servicebus:SqlFilter"},
							Description: "Properties of sqlFilter",
						},
						subscriptionNameProp: {
							TypeSpec:             schema.TypeSpec{Type: "string"},
							Description:          "The subscription name.",
							WillReplaceOnChanges: true,
						},
						topicName: {
							TypeSpec:             schema.TypeSpec{Type: "string"},
							Description:          "The topic name.",
							WillReplaceOnChanges: true,
						},
					},
					RequiredInputs: []string{namespaceName, resourceGroupName, subscriptionNameProp, topicName},
				},
				MetaResource: AzureAPIResource{
					APIVersion: apiVersion,
					Path:       defaultRulePath,
					// The "$Default" rule always exists once its parent Subscription does, so skip
					// the existence check on create and always PUT (adopt-in-place), like TagAtScope.
					Singleton: true,
					PutParameters: []AzureAPIParameter{
						{Name: resourceGroupName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
						{Name: namespaceName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
						{Name: topicName, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
						{Name: subscriptionNameProp, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
						{
							Name:       "parameters",
							Location:   "body",
							IsRequired: true,
							Value:      &AzureAPIProperty{},
							Body: &AzureAPIType{
								Properties: map[string]AzureAPIProperty{
									actionProp:            {Type: "object", Ref: "#/types/azure-native:servicebus:Action", Containers: []string{"properties"}},
									correlationFilterProp: {Type: "object", Ref: "#/types/azure-native:servicebus:CorrelationFilter", Containers: []string{"properties"}},
									filterTypeProp:        {Type: "string", Containers: []string{"properties"}},
									sqlFilterProp:         {Type: "object", Ref: "#/types/azure-native:servicebus:SqlFilter", Containers: []string{"properties"}},
								},
							},
						},
						{Name: subscriptionId, Location: "path", IsRequired: true, Value: &AzureAPIProperty{Type: "string"}},
					},
					Response: map[string]AzureAPIProperty{
						actionProp:            {Ref: "#/types/azure-native:servicebus:ActionResponse", Containers: []string{"properties"}},
						correlationFilterProp: {Ref: "#/types/azure-native:servicebus:CorrelationFilterResponse", Containers: []string{"properties"}},
						filterTypeProp:        {Containers: []string{"properties"}},
						"id":                  {},
						"location":            {},
						"name":                {},
						sqlFilterProp:         {Ref: "#/types/azure-native:servicebus:SqlFilterResponse", Containers: []string{"properties"}},
						"systemData":          {Ref: "#/types/azure-native:servicebus:SystemDataResponse"},
						"type":                {},
					},
				},
			}, nil
		},
		// Create and Update use the standard PUT-based flow (crud.ResourceCrudClient.CreateOrUpdate),
		// which always overwrites in place - there's nothing to override here.
		//
		// Read and Delete both need hand-written implementations to deal with Azure's confirmed
		// unreliable GET on the "$Default" rule (see #4489): the standard Read path would treat any
		// 404 as a genuine deletion, and the standard Delete path doesn't tolerate a 404 at all (the
		// rule already being gone).
		Read: func(ctx context.Context, id string, _ resource.PropertyMap) (map[string]any, bool, error) {
			outputs, found, err := readDefaultRuleWithRetry(ctx, func(ctx context.Context) (map[string]any, error) {
				response, err := client.Read(ctx, id, "")
				if err != nil {
					return nil, err
				}
				return client.ResponseBodyToSdkOutputs(response), nil
			})
			if err != nil || found {
				return outputs, found, err
			}
			return defaultRuleNotFoundOutcome(ctx, azureClient, subscriptionApiVersion, id)
		},
		Delete: func(ctx context.Context, id string, previousInputs, state resource.PropertyMap) error {
			err := azureClient.Delete(ctx, id, apiVersion, "", nil)
			if err != nil && !azure.IsNotFound(err) {
				return err
			}
			return nil
		},
	}, nil
}
