


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataPool(ctx *pulumi.Context, args *LookupDataPoolArgs, opts ...pulumi.InvokeOption) (*LookupDataPoolResult, error) {
	var rv LookupDataPoolResult
	err := ctx.Invoke("azure-native:autonomousdevelopmentplatform/v20210201preview:getDataPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	DataPoolName      string `pulumi:"dataPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataPoolResult struct {
	DataPoolId        string                     `pulumi:"dataPoolId"`
	Id                string                     `pulumi:"id"`
	Locations         []DataPoolLocationResponse `pulumi:"locations"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Type              string                     `pulumi:"type"`
}
