


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInput(ctx *pulumi.Context, args *LookupInputArgs, opts ...pulumi.InvokeOption) (*LookupInputResult, error) {
	var rv LookupInputResult
	err := ctx.Invoke("azure-native:streamanalytics/v20160301:getInput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputArgs struct {
	InputName         string `pulumi:"inputName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInputResult struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupInputOutput(ctx *pulumi.Context, args LookupInputOutputArgs, opts ...pulumi.InvokeOption) LookupInputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInputResult, error) {
			args := v.(LookupInputArgs)
			r, err := LookupInput(ctx, &args, opts...)
			var s LookupInputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInputResultOutput)
}

type LookupInputOutputArgs struct {
	InputName         pulumi.StringInput `pulumi:"inputName"`
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputArgs)(nil)).Elem()
}


type LookupInputResultOutput struct{ *pulumi.OutputState }

func (LookupInputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputResult)(nil)).Elem()
}

func (o LookupInputResultOutput) ToLookupInputResultOutput() LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) ToLookupInputResultOutputWithContext(ctx context.Context) LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInputResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInputResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInputResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupInputResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupInputResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInputResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInputResultOutput{})
}
