


package datalakestore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTrustedIdProvider(ctx *pulumi.Context, args *LookupTrustedIdProviderArgs, opts ...pulumi.InvokeOption) (*LookupTrustedIdProviderResult, error) {
	var rv LookupTrustedIdProviderResult
	err := ctx.Invoke("azure-native:datalakestore:getTrustedIdProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustedIdProviderArgs struct {
	AccountName           string `pulumi:"accountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TrustedIdProviderName string `pulumi:"trustedIdProviderName"`
}


type LookupTrustedIdProviderResult struct {
	Id         string `pulumi:"id"`
	IdProvider string `pulumi:"idProvider"`
	Name       string `pulumi:"name"`
	Type       string `pulumi:"type"`
}

func LookupTrustedIdProviderOutput(ctx *pulumi.Context, args LookupTrustedIdProviderOutputArgs, opts ...pulumi.InvokeOption) LookupTrustedIdProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrustedIdProviderResult, error) {
			args := v.(LookupTrustedIdProviderArgs)
			r, err := LookupTrustedIdProvider(ctx, &args, opts...)
			var s LookupTrustedIdProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrustedIdProviderResultOutput)
}

type LookupTrustedIdProviderOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TrustedIdProviderName pulumi.StringInput `pulumi:"trustedIdProviderName"`
}

func (LookupTrustedIdProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedIdProviderArgs)(nil)).Elem()
}


type LookupTrustedIdProviderResultOutput struct{ *pulumi.OutputState }

func (LookupTrustedIdProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustedIdProviderResult)(nil)).Elem()
}

func (o LookupTrustedIdProviderResultOutput) ToLookupTrustedIdProviderResultOutput() LookupTrustedIdProviderResultOutput {
	return o
}

func (o LookupTrustedIdProviderResultOutput) ToLookupTrustedIdProviderResultOutputWithContext(ctx context.Context) LookupTrustedIdProviderResultOutput {
	return o
}

func (o LookupTrustedIdProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedIdProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTrustedIdProviderResultOutput) IdProvider() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedIdProviderResult) string { return v.IdProvider }).(pulumi.StringOutput)
}

func (o LookupTrustedIdProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedIdProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTrustedIdProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustedIdProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrustedIdProviderResultOutput{})
}
