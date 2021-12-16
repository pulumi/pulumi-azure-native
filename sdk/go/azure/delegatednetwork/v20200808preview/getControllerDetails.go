


package v20200808preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupControllerDetails(ctx *pulumi.Context, args *LookupControllerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupControllerDetailsResult, error) {
	var rv LookupControllerDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork/v20200808preview:getControllerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControllerDetailsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupControllerDetailsResult struct {
	DncAppId          string            `pulumi:"dncAppId"`
	DncEndpoint       string            `pulumi:"dncEndpoint"`
	DncTenantId       string            `pulumi:"dncTenantId"`
	Id                string            `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ResourceGuid      string            `pulumi:"resourceGuid"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
