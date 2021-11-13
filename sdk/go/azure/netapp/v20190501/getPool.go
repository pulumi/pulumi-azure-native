


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:netapp/v20190501:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	Id                string      `pulumi:"id"`
	Location          string      `pulumi:"location"`
	Name              string      `pulumi:"name"`
	PoolId            string      `pulumi:"poolId"`
	ProvisioningState string      `pulumi:"provisioningState"`
	ServiceLevel      string      `pulumi:"serviceLevel"`
	Size              float64     `pulumi:"size"`
	Tags              interface{} `pulumi:"tags"`
	Type              string      `pulumi:"type"`
}
