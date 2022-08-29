


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataFlow(ctx *pulumi.Context, args *LookupDataFlowArgs, opts ...pulumi.InvokeOption) (*LookupDataFlowResult, error) {
	var rv LookupDataFlowResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getDataFlow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataFlowArgs struct {
	DataFlowName      string `pulumi:"dataFlowName"`
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataFlowResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupDataFlowOutput(ctx *pulumi.Context, args LookupDataFlowOutputArgs, opts ...pulumi.InvokeOption) LookupDataFlowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataFlowResult, error) {
			args := v.(LookupDataFlowArgs)
			r, err := LookupDataFlow(ctx, &args, opts...)
			var s LookupDataFlowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataFlowResultOutput)
}

type LookupDataFlowOutputArgs struct {
	DataFlowName      pulumi.StringInput `pulumi:"dataFlowName"`
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataFlowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowArgs)(nil)).Elem()
}


type LookupDataFlowResultOutput struct{ *pulumi.OutputState }

func (LookupDataFlowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataFlowResult)(nil)).Elem()
}

func (o LookupDataFlowResultOutput) ToLookupDataFlowResultOutput() LookupDataFlowResultOutput {
	return o
}

func (o LookupDataFlowResultOutput) ToLookupDataFlowResultOutputWithContext(ctx context.Context) LookupDataFlowResultOutput {
	return o
}

func (o LookupDataFlowResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataFlowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataFlowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataFlowResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataFlowResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupDataFlowResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataFlowResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataFlowResultOutput{})
}
