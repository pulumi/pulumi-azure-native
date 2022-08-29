


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("azure-native:streamanalytics/v20160301:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionArgs struct {
	FunctionName      string `pulumi:"functionName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFunctionResult struct {
	Id         string                           `pulumi:"id"`
	Name       *string                          `pulumi:"name"`
	Properties ScalarFunctionPropertiesResponse `pulumi:"properties"`
	Type       string                           `pulumi:"type"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionResult, error) {
			args := v.(LookupFunctionArgs)
			r, err := LookupFunction(ctx, &args, opts...)
			var s LookupFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionResultOutput)
}

type LookupFunctionOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}


type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionResultOutput) Properties() ScalarFunctionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupFunctionResult) ScalarFunctionPropertiesResponse { return v.Properties }).(ScalarFunctionPropertiesResponseOutput)
}

func (o LookupFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}
