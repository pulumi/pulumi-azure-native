


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTenantConfiguration(ctx *pulumi.Context, args *LookupTenantConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupTenantConfigurationResult, error) {
	var rv LookupTenantConfigurationResult
	err := ctx.Invoke("azure-native:portal/v20200901preview:getTenantConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTenantConfigurationArgs struct {
	ConfigurationName string `pulumi:"configurationName"`
}


type LookupTenantConfigurationResult struct {
	EnforcePrivateMarkdownStorage *bool  `pulumi:"enforcePrivateMarkdownStorage"`
	Id                            string `pulumi:"id"`
	Name                          string `pulumi:"name"`
	Type                          string `pulumi:"type"`
}

func LookupTenantConfigurationOutput(ctx *pulumi.Context, args LookupTenantConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupTenantConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTenantConfigurationResult, error) {
			args := v.(LookupTenantConfigurationArgs)
			r, err := LookupTenantConfiguration(ctx, &args, opts...)
			var s LookupTenantConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTenantConfigurationResultOutput)
}

type LookupTenantConfigurationOutputArgs struct {
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
}

func (LookupTenantConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantConfigurationArgs)(nil)).Elem()
}


type LookupTenantConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupTenantConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantConfigurationResult)(nil)).Elem()
}

func (o LookupTenantConfigurationResultOutput) ToLookupTenantConfigurationResultOutput() LookupTenantConfigurationResultOutput {
	return o
}

func (o LookupTenantConfigurationResultOutput) ToLookupTenantConfigurationResultOutputWithContext(ctx context.Context) LookupTenantConfigurationResultOutput {
	return o
}

func (o LookupTenantConfigurationResultOutput) EnforcePrivateMarkdownStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) *bool { return v.EnforcePrivateMarkdownStorage }).(pulumi.BoolPtrOutput)
}

func (o LookupTenantConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTenantConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTenantConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantConfigurationResultOutput{})
}
