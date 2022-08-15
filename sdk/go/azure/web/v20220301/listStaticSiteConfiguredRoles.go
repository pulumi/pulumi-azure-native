


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteConfiguredRoles(ctx *pulumi.Context, args *ListStaticSiteConfiguredRolesArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteConfiguredRolesResult, error) {
	var rv ListStaticSiteConfiguredRolesResult
	err := ctx.Invoke("azure-native:web/v20220301:listStaticSiteConfiguredRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteConfiguredRolesArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteConfiguredRolesResult struct {
	Id         string   `pulumi:"id"`
	Kind       *string  `pulumi:"kind"`
	Name       string   `pulumi:"name"`
	Properties []string `pulumi:"properties"`
	Type       string   `pulumi:"type"`
}

func ListStaticSiteConfiguredRolesOutput(ctx *pulumi.Context, args ListStaticSiteConfiguredRolesOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteConfiguredRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteConfiguredRolesResult, error) {
			args := v.(ListStaticSiteConfiguredRolesArgs)
			r, err := ListStaticSiteConfiguredRoles(ctx, &args, opts...)
			var s ListStaticSiteConfiguredRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteConfiguredRolesResultOutput)
}

type ListStaticSiteConfiguredRolesOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteConfiguredRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteConfiguredRolesArgs)(nil)).Elem()
}


type ListStaticSiteConfiguredRolesResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteConfiguredRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteConfiguredRolesResult)(nil)).Elem()
}

func (o ListStaticSiteConfiguredRolesResultOutput) ToListStaticSiteConfiguredRolesResultOutput() ListStaticSiteConfiguredRolesResultOutput {
	return o
}

func (o ListStaticSiteConfiguredRolesResultOutput) ToListStaticSiteConfiguredRolesResultOutputWithContext(ctx context.Context) ListStaticSiteConfiguredRolesResultOutput {
	return o
}

func (o ListStaticSiteConfiguredRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteConfiguredRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteConfiguredRolesResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteConfiguredRolesResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteConfiguredRolesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteConfiguredRolesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteConfiguredRolesResultOutput) Properties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListStaticSiteConfiguredRolesResult) []string { return v.Properties }).(pulumi.StringArrayOutput)
}

func (o ListStaticSiteConfiguredRolesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteConfiguredRolesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteConfiguredRolesResultOutput{})
}
