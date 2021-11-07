


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationTypeVersion(ctx *pulumi.Context, args *LookupApplicationTypeVersionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationTypeVersionResult, error) {
	var rv LookupApplicationTypeVersionResult
	err := ctx.Invoke("azure-native:servicefabric/v20211101preview:getApplicationTypeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationTypeVersionArgs struct {
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	ClusterName         string `pulumi:"clusterName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	Version             string `pulumi:"version"`
}


type LookupApplicationTypeVersionResult struct {
	AppPackageUrl     string             `pulumi:"appPackageUrl"`
	Id                string             `pulumi:"id"`
	Location          *string            `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}
