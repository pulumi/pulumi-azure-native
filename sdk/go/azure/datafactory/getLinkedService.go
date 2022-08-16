


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:datafactory:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkedServiceResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupLinkedServiceOutput(ctx *pulumi.Context, args LookupLinkedServiceOutputArgs, opts ...pulumi.InvokeOption) LookupLinkedServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkedServiceResult, error) {
			args := v.(LookupLinkedServiceArgs)
			r, err := LookupLinkedService(ctx, &args, opts...)
			var s LookupLinkedServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkedServiceResultOutput)
}

type LookupLinkedServiceOutputArgs struct {
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	LinkedServiceName pulumi.StringInput `pulumi:"linkedServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLinkedServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceArgs)(nil)).Elem()
}


type LookupLinkedServiceResultOutput struct{ *pulumi.OutputState }

func (LookupLinkedServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkedServiceResult)(nil)).Elem()
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutput() LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) ToLookupLinkedServiceResultOutputWithContext(ctx context.Context) LookupLinkedServiceResultOutput {
	return o
}

func (o LookupLinkedServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLinkedServiceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupLinkedServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkedServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkedServiceResultOutput{})
}
