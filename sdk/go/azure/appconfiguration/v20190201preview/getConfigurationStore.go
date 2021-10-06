


package v20190201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationStore(ctx *pulumi.Context, args *LookupConfigurationStoreArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationStoreResult, error) {
	var rv LookupConfigurationStoreResult
	err := ctx.Invoke("azure-native:appconfiguration/v20190201preview:getConfigurationStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationStoreArgs struct {
	ConfigStoreName   string `pulumi:"configStoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupConfigurationStoreResult struct {
	CreationDate      string            `pulumi:"creationDate"`
	Endpoint          string            `pulumi:"endpoint"`
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
