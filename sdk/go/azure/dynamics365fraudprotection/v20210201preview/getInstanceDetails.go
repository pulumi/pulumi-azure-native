


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstanceDetails(ctx *pulumi.Context, args *LookupInstanceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupInstanceDetailsResult, error) {
	var rv LookupInstanceDetailsResult
	err := ctx.Invoke("azure-native:dynamics365fraudprotection/v20210201preview:getInstanceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceDetailsArgs struct {
	InstanceName      string `pulumi:"instanceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceDetailsResult struct {
	Administration    *DFPInstanceAdministratorsResponse `pulumi:"administration"`
	Id                string                             `pulumi:"id"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                 `pulumi:"systemData"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              string                             `pulumi:"type"`
}
