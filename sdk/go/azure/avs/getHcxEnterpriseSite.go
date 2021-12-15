


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHcxEnterpriseSite(ctx *pulumi.Context, args *LookupHcxEnterpriseSiteArgs, opts ...pulumi.InvokeOption) (*LookupHcxEnterpriseSiteResult, error) {
	var rv LookupHcxEnterpriseSiteResult
	err := ctx.Invoke("azure-native:avs:getHcxEnterpriseSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHcxEnterpriseSiteArgs struct {
	HcxEnterpriseSiteName string `pulumi:"hcxEnterpriseSiteName"`
	PrivateCloudName      string `pulumi:"privateCloudName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupHcxEnterpriseSiteResult struct {
	ActivationKey string `pulumi:"activationKey"`
	Id            string `pulumi:"id"`
	Name          string `pulumi:"name"`
	Status        string `pulumi:"status"`
	Type          string `pulumi:"type"`
}
