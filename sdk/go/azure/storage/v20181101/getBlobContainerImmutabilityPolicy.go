


package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainerImmutabilityPolicy(ctx *pulumi.Context, args *LookupBlobContainerImmutabilityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerImmutabilityPolicyResult, error) {
	var rv LookupBlobContainerImmutabilityPolicyResult
	err := ctx.Invoke("azure-native:storage/v20181101:getBlobContainerImmutabilityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerImmutabilityPolicyArgs struct {
	AccountName            string `pulumi:"accountName"`
	ContainerName          string `pulumi:"containerName"`
	ImmutabilityPolicyName string `pulumi:"immutabilityPolicyName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupBlobContainerImmutabilityPolicyResult struct {
	Etag                                  string `pulumi:"etag"`
	Id                                    string `pulumi:"id"`
	ImmutabilityPeriodSinceCreationInDays int    `pulumi:"immutabilityPeriodSinceCreationInDays"`
	Name                                  string `pulumi:"name"`
	State                                 string `pulumi:"state"`
	Type                                  string `pulumi:"type"`
}
