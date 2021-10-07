


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationServiceEnvironmentManagedApi(ctx *pulumi.Context, args *LookupIntegrationServiceEnvironmentManagedApiArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationServiceEnvironmentManagedApiResult, error) {
	var rv LookupIntegrationServiceEnvironmentManagedApiResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationServiceEnvironmentManagedApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationServiceEnvironmentManagedApiArgs struct {
	ApiName                           string `pulumi:"apiName"`
	IntegrationServiceEnvironmentName string `pulumi:"integrationServiceEnvironmentName"`
	ResourceGroup                     string `pulumi:"resourceGroup"`
}


type LookupIntegrationServiceEnvironmentManagedApiResult struct {
	Id         string                        `pulumi:"id"`
	Location   *string                       `pulumi:"location"`
	Name       string                        `pulumi:"name"`
	Properties ApiResourcePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string             `pulumi:"tags"`
	Type       string                        `pulumi:"type"`
}
