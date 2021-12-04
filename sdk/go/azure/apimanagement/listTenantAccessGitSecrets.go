


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTenantAccessGitSecrets(ctx *pulumi.Context, args *ListTenantAccessGitSecretsArgs, opts ...pulumi.InvokeOption) (*ListTenantAccessGitSecretsResult, error) {
	var rv ListTenantAccessGitSecretsResult
	err := ctx.Invoke("azure-native:apimanagement:listTenantAccessGitSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTenantAccessGitSecretsArgs struct {
	AccessName        string `pulumi:"accessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListTenantAccessGitSecretsResult struct {
	Enabled      *bool   `pulumi:"enabled"`
	Id           *string `pulumi:"id"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
