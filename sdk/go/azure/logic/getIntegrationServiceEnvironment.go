


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationServiceEnvironment(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentResult, error) {
	var rv LookupIntegrationServiceEnvironmentResult
	err := ctx.Invoke("azure-native:logic:getIntegrationServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationServiceEnvironmentArgs struct {
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     string `pulumi:"resourceGroup"`
}


type LookupIntegrationServiceEnvironmentResult struct {
	Id         string                                          `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse                 `pulumi:"identity"`
	Location   *string                                         `pulumi:"location"`
	Name       string                                          `pulumi:"name"`
	Properties IntegrationServiceEnvironmentPropertiesResponse `pulumi:"properties"`
	Sku        *IntegrationServiceEnvironmentSkuResponse       `pulumi:"sku"`
	Tags       map[string]string                               `pulumi:"tags"`
	Type       string                                          `pulumi:"type"`
}
