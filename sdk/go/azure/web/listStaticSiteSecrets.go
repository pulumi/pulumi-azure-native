


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteSecrets(ctx *pulumi.Context, args *ListStaticSiteSecretsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteSecretsResult, error) {
	var rv ListStaticSiteSecretsResult
	err := ctx.Invoke("azure-native:web:listStaticSiteSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteSecretsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteSecretsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListStaticSiteSecretsOutput(ctx *pulumi.Context, args ListStaticSiteSecretsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteSecretsResult, error) {
			args := v.(ListStaticSiteSecretsArgs)
			r, err := ListStaticSiteSecrets(ctx, &args, opts...)
			var s ListStaticSiteSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteSecretsResultOutput)
}

type ListStaticSiteSecretsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteSecretsArgs)(nil)).Elem()
}


type ListStaticSiteSecretsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteSecretsResult)(nil)).Elem()
}

func (o ListStaticSiteSecretsResultOutput) ToListStaticSiteSecretsResultOutput() ListStaticSiteSecretsResultOutput {
	return o
}

func (o ListStaticSiteSecretsResultOutput) ToListStaticSiteSecretsResultOutputWithContext(ctx context.Context) ListStaticSiteSecretsResultOutput {
	return o
}

func (o ListStaticSiteSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteSecretsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListStaticSiteSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteSecretsResultOutput{})
}
