


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceEndpointPolicyDefinition(ctx *pulumi.Context, args *LookupServiceEndpointPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyDefinitionResult, error) {
	var rv LookupServiceEndpointPolicyDefinitionResult
	err := ctx.Invoke("azure-native:network/v20210501:getServiceEndpointPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyDefinitionArgs struct {
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyDefinitionName string `pulumi:"serviceEndpointPolicyDefinitionName"`
	ServiceEndpointPolicyName           string `pulumi:"serviceEndpointPolicyName"`
}


type LookupServiceEndpointPolicyDefinitionResult struct {
	Description       *string  `pulumi:"description"`
	Etag              string   `pulumi:"etag"`
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Service           *string  `pulumi:"service"`
	ServiceResources  []string `pulumi:"serviceResources"`
	Type              *string  `pulumi:"type"`
}
