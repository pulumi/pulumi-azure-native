


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainer(ctx *pulumi.Context, args *LookupBlobContainerArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerResult, error) {
	var rv LookupBlobContainerResult
	err := ctx.Invoke("azure-native:storage/v20180201:getBlobContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBlobContainerResult struct {
	Etag                  string                               `pulumi:"etag"`
	HasImmutabilityPolicy bool                                 `pulumi:"hasImmutabilityPolicy"`
	HasLegalHold          bool                                 `pulumi:"hasLegalHold"`
	Id                    string                               `pulumi:"id"`
	ImmutabilityPolicy    ImmutabilityPolicyPropertiesResponse `pulumi:"immutabilityPolicy"`
	LastModifiedTime      string                               `pulumi:"lastModifiedTime"`
	LeaseDuration         string                               `pulumi:"leaseDuration"`
	LeaseState            string                               `pulumi:"leaseState"`
	LeaseStatus           string                               `pulumi:"leaseStatus"`
	LegalHold             LegalHoldPropertiesResponse          `pulumi:"legalHold"`
	Metadata              map[string]string                    `pulumi:"metadata"`
	Name                  string                               `pulumi:"name"`
	PublicAccess          *string                              `pulumi:"publicAccess"`
	Type                  string                               `pulumi:"type"`
}
