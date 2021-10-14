


package v20200808preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDelegatedSubnetServiceDetails(ctx *pulumi.Context, args *LookupDelegatedSubnetServiceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupDelegatedSubnetServiceDetailsResult, error) {
	var rv LookupDelegatedSubnetServiceDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20200808preview:getDelegatedSubnetServiceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDelegatedSubnetServiceDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDelegatedSubnetServiceDetailsResult struct {
	ControllerDetails *ControllerDetailsResponse `pulumi:"controllerDetails"`
	Id                string                     `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	ResourceGuid      string                     `pulumi:"resourceGuid"`
	SubnetDetails     *SubnetDetailsResponse     `pulumi:"subnetDetails"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}
