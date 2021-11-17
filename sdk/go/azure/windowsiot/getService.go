


package windowsiot

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:windowsiot:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupServiceResult struct {
	AdminDomainName   *string           `pulumi:"adminDomainName"`
	BillingDomainName *string           `pulumi:"billingDomainName"`
	Etag              *string           `pulumi:"etag"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Notes             *string           `pulumi:"notes"`
	Quantity          *float64          `pulumi:"quantity"`
	StartDate         string            `pulumi:"startDate"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
