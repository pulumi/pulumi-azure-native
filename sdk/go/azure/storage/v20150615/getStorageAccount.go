


package v20150615

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20150615:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountResult struct {
	AccountType         *string               `pulumi:"accountType"`
	CreationTime        *string               `pulumi:"creationTime"`
	CustomDomain        *CustomDomainResponse `pulumi:"customDomain"`
	Id                  string                `pulumi:"id"`
	LastGeoFailoverTime *string               `pulumi:"lastGeoFailoverTime"`
	Location            *string               `pulumi:"location"`
	Name                string                `pulumi:"name"`
	PrimaryEndpoints    *EndpointsResponse    `pulumi:"primaryEndpoints"`
	PrimaryLocation     *string               `pulumi:"primaryLocation"`
	ProvisioningState   *string               `pulumi:"provisioningState"`
	SecondaryEndpoints  *EndpointsResponse    `pulumi:"secondaryEndpoints"`
	SecondaryLocation   *string               `pulumi:"secondaryLocation"`
	StatusOfPrimary     *string               `pulumi:"statusOfPrimary"`
	StatusOfSecondary   *string               `pulumi:"statusOfSecondary"`
	Tags                map[string]string     `pulumi:"tags"`
	Type                string                `pulumi:"type"`
}
