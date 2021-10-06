


package v20190916preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMultipleActivationKey(ctx *pulumi.Context, args *LookupMultipleActivationKeyArgs, opts ...pulumi.InvokeOption) (*LookupMultipleActivationKeyResult, error) {
	var rv LookupMultipleActivationKeyResult
	err := ctx.Invoke("azure-native:windowsesu/v20190916preview:getMultipleActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMultipleActivationKeyArgs struct {
	MultipleActivationKeyName string `pulumi:"multipleActivationKeyName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupMultipleActivationKeyResult struct {
	AgreementNumber       *string           `pulumi:"agreementNumber"`
	ExpirationDate        string            `pulumi:"expirationDate"`
	Id                    string            `pulumi:"id"`
	InstalledServerNumber *int              `pulumi:"installedServerNumber"`
	IsEligible            *bool             `pulumi:"isEligible"`
	Location              string            `pulumi:"location"`
	MultipleActivationKey string            `pulumi:"multipleActivationKey"`
	Name                  string            `pulumi:"name"`
	OsType                *string           `pulumi:"osType"`
	ProvisioningState     string            `pulumi:"provisioningState"`
	SupportType           *string           `pulumi:"supportType"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
}
