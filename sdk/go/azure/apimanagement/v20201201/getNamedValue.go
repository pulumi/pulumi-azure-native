


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamedValue(ctx *pulumi.Context, args *LookupNamedValueArgs, opts ...pulumi.InvokeOption) (*LookupNamedValueResult, error) {
	var rv LookupNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamedValueArgs struct {
	NamedValueId      string `pulumi:"namedValueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupNamedValueResult struct {
	DisplayName string                              `pulumi:"displayName"`
	Id          string                              `pulumi:"id"`
	KeyVault    *KeyVaultContractPropertiesResponse `pulumi:"keyVault"`
	Name        string                              `pulumi:"name"`
	Secret      *bool                               `pulumi:"secret"`
	Tags        []string                            `pulumi:"tags"`
	Type        string                              `pulumi:"type"`
	Value       *string                             `pulumi:"value"`
}
