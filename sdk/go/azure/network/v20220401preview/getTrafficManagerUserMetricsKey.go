


package v20220401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrafficManagerUserMetricsKey(ctx *pulumi.Context, args *LookupTrafficManagerUserMetricsKeyArgs, opts ...pulumi.InvokeOption) (*LookupTrafficManagerUserMetricsKeyResult, error) {
	var rv LookupTrafficManagerUserMetricsKeyResult
	err := ctx.Invoke("azure-native:network/v20220401preview:getTrafficManagerUserMetricsKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficManagerUserMetricsKeyArgs struct {
}


type LookupTrafficManagerUserMetricsKeyResult struct {
	Id   *string `pulumi:"id"`
	Key  *string `pulumi:"key"`
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}
