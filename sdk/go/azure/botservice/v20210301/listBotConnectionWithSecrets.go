


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBotConnectionWithSecrets(ctx *pulumi.Context, args *ListBotConnectionWithSecretsArgs, opts ...pulumi.InvokeOption) (*ListBotConnectionWithSecretsResult, error) {
	var rv ListBotConnectionWithSecretsResult
	err := ctx.Invoke("azure-native:botservice/v20210301:listBotConnectionWithSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListBotConnectionWithSecretsArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListBotConnectionWithSecretsResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Kind       *string                             `pulumi:"kind"`
	Location   *string                             `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties ConnectionSettingPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                        `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
	Zones      []string                            `pulumi:"zones"`
}


func (val *ListBotConnectionWithSecretsResult) Defaults() *ListBotConnectionWithSecretsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func ListBotConnectionWithSecretsOutput(ctx *pulumi.Context, args ListBotConnectionWithSecretsOutputArgs, opts ...pulumi.InvokeOption) ListBotConnectionWithSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBotConnectionWithSecretsResult, error) {
			args := v.(ListBotConnectionWithSecretsArgs)
			r, err := ListBotConnectionWithSecrets(ctx, &args, opts...)
			var s ListBotConnectionWithSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBotConnectionWithSecretsResultOutput)
}

type ListBotConnectionWithSecretsOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListBotConnectionWithSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotConnectionWithSecretsArgs)(nil)).Elem()
}


type ListBotConnectionWithSecretsResultOutput struct{ *pulumi.OutputState }

func (ListBotConnectionWithSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotConnectionWithSecretsResult)(nil)).Elem()
}

func (o ListBotConnectionWithSecretsResultOutput) ToListBotConnectionWithSecretsResultOutput() ListBotConnectionWithSecretsResultOutput {
	return o
}

func (o ListBotConnectionWithSecretsResultOutput) ToListBotConnectionWithSecretsResultOutputWithContext(ctx context.Context) ListBotConnectionWithSecretsResultOutput {
	return o
}

func (o ListBotConnectionWithSecretsResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Properties() ConnectionSettingPropertiesResponseOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) ConnectionSettingPropertiesResponse { return v.Properties }).(ConnectionSettingPropertiesResponseOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o ListBotConnectionWithSecretsResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListBotConnectionWithSecretsResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBotConnectionWithSecretsResultOutput{})
}
