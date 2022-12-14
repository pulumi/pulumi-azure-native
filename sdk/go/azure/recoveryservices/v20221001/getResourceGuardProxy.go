


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceGuardProxy(ctx *pulumi.Context, args *LookupResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardProxyResult, error) {
	var rv LookupResourceGuardProxyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20221001:getResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardProxyArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	VaultName              string `pulumi:"vaultName"`
}

type LookupResourceGuardProxyResult struct {
	ETag       *string                        `pulumi:"eTag"`
	Id         string                         `pulumi:"id"`
	Location   *string                        `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}

func LookupResourceGuardProxyOutput(ctx *pulumi.Context, args LookupResourceGuardProxyOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGuardProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGuardProxyResult, error) {
			args := v.(LookupResourceGuardProxyArgs)
			r, err := LookupResourceGuardProxy(ctx, &args, opts...)
			var s LookupResourceGuardProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceGuardProxyResultOutput)
}

type LookupResourceGuardProxyOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceGuardProxyName pulumi.StringInput `pulumi:"resourceGuardProxyName"`
	VaultName              pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupResourceGuardProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyArgs)(nil)).Elem()
}

type LookupResourceGuardProxyResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGuardProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyResult)(nil)).Elem()
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutput() LookupResourceGuardProxyResultOutput {
	return o
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutputWithContext(ctx context.Context) LookupResourceGuardProxyResultOutput {
	return o
}

func (o LookupResourceGuardProxyResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGuardProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceGuardProxyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGuardProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourceGuardProxyResultOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) ResourceGuardProxyBaseResponse { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

func (o LookupResourceGuardProxyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupResourceGuardProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGuardProxyResultOutput{})
}
