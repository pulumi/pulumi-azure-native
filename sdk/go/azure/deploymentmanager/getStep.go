


package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStep(ctx *pulumi.Context, args *LookupStepArgs, opts ...pulumi.InvokeOption) (*LookupStepResult, error) {
	var rv LookupStepResult
	err := ctx.Invoke("azure-native:deploymentmanager:getStep", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStepArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StepName          string `pulumi:"stepName"`
}


type LookupStepResult struct {
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}

func LookupStepOutput(ctx *pulumi.Context, args LookupStepOutputArgs, opts ...pulumi.InvokeOption) LookupStepResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStepResult, error) {
			args := v.(LookupStepArgs)
			r, err := LookupStep(ctx, &args, opts...)
			var s LookupStepResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStepResultOutput)
}

type LookupStepOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StepName          pulumi.StringInput `pulumi:"stepName"`
}

func (LookupStepOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStepArgs)(nil)).Elem()
}


type LookupStepResultOutput struct{ *pulumi.OutputState }

func (LookupStepResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStepResult)(nil)).Elem()
}

func (o LookupStepResultOutput) ToLookupStepResultOutput() LookupStepResultOutput {
	return o
}

func (o LookupStepResultOutput) ToLookupStepResultOutputWithContext(ctx context.Context) LookupStepResultOutput {
	return o
}

func (o LookupStepResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStepResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStepResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStepResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStepResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStepResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStepResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStepResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupStepResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStepResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStepResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStepResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStepResultOutput{})
}
