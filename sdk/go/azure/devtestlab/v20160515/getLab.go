


package v20160515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	Expand            *string `pulumi:"expand"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	ArtifactsStorageAccount       string            `pulumi:"artifactsStorageAccount"`
	CreatedDate                   string            `pulumi:"createdDate"`
	DefaultPremiumStorageAccount  string            `pulumi:"defaultPremiumStorageAccount"`
	DefaultStorageAccount         string            `pulumi:"defaultStorageAccount"`
	Id                            string            `pulumi:"id"`
	LabStorageType                *string           `pulumi:"labStorageType"`
	Location                      *string           `pulumi:"location"`
	Name                          string            `pulumi:"name"`
	PremiumDataDiskStorageAccount string            `pulumi:"premiumDataDiskStorageAccount"`
	PremiumDataDisks              *string           `pulumi:"premiumDataDisks"`
	ProvisioningState             *string           `pulumi:"provisioningState"`
	Tags                          map[string]string `pulumi:"tags"`
	Type                          string            `pulumi:"type"`
	UniqueIdentifier              *string           `pulumi:"uniqueIdentifier"`
	VaultName                     string            `pulumi:"vaultName"`
}
