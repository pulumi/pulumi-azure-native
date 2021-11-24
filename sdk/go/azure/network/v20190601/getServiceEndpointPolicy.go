


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceEndpointPolicy(ctx *pulumi.Context, args *LookupServiceEndpointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyResult, error) {
	var rv LookupServiceEndpointPolicyResult
	err := ctx.Invoke("azure-native:network/v20190601:getServiceEndpointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyArgs struct {
	Expand                    *string `pulumi:"expand"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	ServiceEndpointPolicyName string  `pulumi:"serviceEndpointPolicyName"`
}


type LookupServiceEndpointPolicyResult struct {
	Etag                             *string                                   `pulumi:"etag"`
	Id                               *string                                   `pulumi:"id"`
	Location                         *string                                   `pulumi:"location"`
	Name                             string                                    `pulumi:"name"`
	ProvisioningState                string                                    `pulumi:"provisioningState"`
	ResourceGuid                     string                                    `pulumi:"resourceGuid"`
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionResponse `pulumi:"serviceEndpointPolicyDefinitions"`
	Subnets                          []SubnetResponse                          `pulumi:"subnets"`
	Tags                             map[string]string                         `pulumi:"tags"`
	Type                             string                                    `pulumi:"type"`
}
