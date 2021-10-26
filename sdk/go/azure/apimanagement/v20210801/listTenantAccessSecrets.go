


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTenantAccessSecrets(ctx *pulumi.Context, args *ListTenantAccessSecretsArgs, opts ...pulumi.InvokeOption) (*ListTenantAccessSecretsResult, error) {
	var rv ListTenantAccessSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:listTenantAccessSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTenantAccessSecretsArgs struct {
	AccessName        string `pulumi:"accessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListTenantAccessSecretsResult struct {
	Enabled      *bool   `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	PrincipalId  *string `pulumi:"principalId"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
