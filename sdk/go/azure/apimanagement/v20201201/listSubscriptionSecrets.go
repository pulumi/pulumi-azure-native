


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubscriptionSecrets(ctx *pulumi.Context, args *ListSubscriptionSecretsArgs, opts ...pulumi.InvokeOption) (*ListSubscriptionSecretsResult, error) {
	var rv ListSubscriptionSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:listSubscriptionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubscriptionSecretsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	Sid               string `pulumi:"sid"`
}


type ListSubscriptionSecretsResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}
