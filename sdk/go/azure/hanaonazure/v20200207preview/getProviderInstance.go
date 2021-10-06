


package v20200207preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProviderInstance(ctx *pulumi.Context, args *LookupProviderInstanceArgs, opts ...pulumi.InvokeOption) (*LookupProviderInstanceResult, error) {
	var rv LookupProviderInstanceResult
	err := ctx.Invoke("azure-native:hanaonazure/v20200207preview:getProviderInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProviderInstanceArgs struct {
	ProviderInstanceName string `pulumi:"providerInstanceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	SapMonitorName       string `pulumi:"sapMonitorName"`
}


type LookupProviderInstanceResult struct {
	Id                string  `pulumi:"id"`
	Metadata          *string `pulumi:"metadata"`
	Name              string  `pulumi:"name"`
	Properties        string  `pulumi:"properties"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}
