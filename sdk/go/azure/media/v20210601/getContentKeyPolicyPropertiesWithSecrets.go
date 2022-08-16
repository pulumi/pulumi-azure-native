


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetContentKeyPolicyPropertiesWithSecrets(ctx *pulumi.Context, args *GetContentKeyPolicyPropertiesWithSecretsArgs, opts ...pulumi.InvokeOption) (*GetContentKeyPolicyPropertiesWithSecretsResult, error) {
	var rv GetContentKeyPolicyPropertiesWithSecretsResult
	err := ctx.Invoke("azure-native:media/v20210601:getContentKeyPolicyPropertiesWithSecrets", args, &rv, opts...)
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

func GetContentKeyPolicyPropertiesWithSecretsOutput(ctx *pulumi.Context, args GetContentKeyPolicyPropertiesWithSecretsOutputArgs, opts ...pulumi.InvokeOption) GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContentKeyPolicyPropertiesWithSecretsResult, error) {
			args := v.(GetContentKeyPolicyPropertiesWithSecretsArgs)
			r, err := GetContentKeyPolicyPropertiesWithSecrets(ctx, &args, opts...)
			var s GetContentKeyPolicyPropertiesWithSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContentKeyPolicyPropertiesWithSecretsResultOutput)
}

type GetContentKeyPolicyPropertiesWithSecretsOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	ContentKeyPolicyName pulumi.StringInput `pulumi:"contentKeyPolicyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetContentKeyPolicyPropertiesWithSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContentKeyPolicyPropertiesWithSecretsArgs)(nil)).Elem()
}


type GetContentKeyPolicyPropertiesWithSecretsResultOutput struct{ *pulumi.OutputState }

func (GetContentKeyPolicyPropertiesWithSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContentKeyPolicyPropertiesWithSecretsResult)(nil)).Elem()
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) ToGetContentKeyPolicyPropertiesWithSecretsResultOutput() GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return o
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) ToGetContentKeyPolicyPropertiesWithSecretsResultOutputWithContext(ctx context.Context) GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return o
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Options() ContentKeyPolicyOptionResponseArrayOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) []ContentKeyPolicyOptionResponse {
		return v.Options
	}).(ContentKeyPolicyOptionResponseArrayOutput)
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContentKeyPolicyPropertiesWithSecretsResultOutput{})
}
