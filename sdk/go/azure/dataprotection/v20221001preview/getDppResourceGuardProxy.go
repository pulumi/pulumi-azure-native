


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDppResourceGuardProxy(ctx *pulumi.Context, args *LookupDppResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupDppResourceGuardProxyResult, error) {
	var rv LookupDppResourceGuardProxyResult
	err := ctx.Invoke("azure-native:dataprotection/v20221001preview:getDppResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDppResourceGuardProxyArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	VaultName              string `pulumi:"vaultName"`
}

type LookupDppResourceGuardProxyResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}

func LookupDppResourceGuardProxyOutput(ctx *pulumi.Context, args LookupDppResourceGuardProxyOutputArgs, opts ...pulumi.InvokeOption) LookupDppResourceGuardProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDppResourceGuardProxyResult, error) {
			args := v.(LookupDppResourceGuardProxyArgs)
			r, err := LookupDppResourceGuardProxy(ctx, &args, opts...)
			var s LookupDppResourceGuardProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDppResourceGuardProxyResultOutput)
}

type LookupDppResourceGuardProxyOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceGuardProxyName pulumi.StringInput `pulumi:"resourceGuardProxyName"`
	VaultName              pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupDppResourceGuardProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDppResourceGuardProxyArgs)(nil)).Elem()
}

type LookupDppResourceGuardProxyResultOutput struct{ *pulumi.OutputState }

func (LookupDppResourceGuardProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDppResourceGuardProxyResult)(nil)).Elem()
}

func (o LookupDppResourceGuardProxyResultOutput) ToLookupDppResourceGuardProxyResultOutput() LookupDppResourceGuardProxyResultOutput {
	return o
}

func (o LookupDppResourceGuardProxyResultOutput) ToLookupDppResourceGuardProxyResultOutputWithContext(ctx context.Context) LookupDppResourceGuardProxyResultOutput {
	return o
}

func (o LookupDppResourceGuardProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDppResourceGuardProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDppResourceGuardProxyResultOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) ResourceGuardProxyBaseResponse { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

func (o LookupDppResourceGuardProxyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDppResourceGuardProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDppResourceGuardProxyResultOutput{})
}
