


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDelegationSettingSecrets(ctx *pulumi.Context, args *ListDelegationSettingSecretsArgs, opts ...pulumi.InvokeOption) (*ListDelegationSettingSecretsResult, error) {
	var rv ListDelegationSettingSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:listDelegationSettingSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDelegationSettingSecretsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListDelegationSettingSecretsResult struct {
	ValidationKey *string `pulumi:"validationKey"`
}
