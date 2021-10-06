


package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetContentKeyPolicyPropertiesWithSecrets(ctx *pulumi.Context, args *GetContentKeyPolicyPropertiesWithSecretsArgs, opts ...pulumi.InvokeOption) (*GetContentKeyPolicyPropertiesWithSecretsResult, error) {
	var rv GetContentKeyPolicyPropertiesWithSecretsResult
	err := ctx.Invoke("azure-native:media:getContentKeyPolicyPropertiesWithSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetContentKeyPolicyPropertiesWithSecretsArgs struct {
	AccountName          string `pulumi:"accountName"`
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type GetContentKeyPolicyPropertiesWithSecretsResult struct {
	Created      string                           `pulumi:"created"`
	Description  *string                          `pulumi:"description"`
	LastModified string                           `pulumi:"lastModified"`
	Options      []ContentKeyPolicyOptionResponse `pulumi:"options"`
	PolicyId     string                           `pulumi:"policyId"`
}
