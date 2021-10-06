


package v20180331preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:datamigration/v20180331preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	GroupName   string `pulumi:"groupName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	Etag              *string             `pulumi:"etag"`
	Id                string              `pulumi:"id"`
	Kind              *string             `pulumi:"kind"`
	Location          string              `pulumi:"location"`
	Name              string              `pulumi:"name"`
	ProvisioningState string              `pulumi:"provisioningState"`
	PublicKey         *string             `pulumi:"publicKey"`
	Sku               *ServiceSkuResponse `pulumi:"sku"`
	Tags              map[string]string   `pulumi:"tags"`
	Type              string              `pulumi:"type"`
	VirtualSubnetId   string              `pulumi:"virtualSubnetId"`
}
