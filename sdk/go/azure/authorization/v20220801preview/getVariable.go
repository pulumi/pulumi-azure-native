


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariable(ctx *pulumi.Context, args *LookupVariableArgs, opts ...pulumi.InvokeOption) (*LookupVariableResult, error) {
	var rv LookupVariableResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableArgs struct {
	VariableName string `pulumi:"variableName"`
}


type LookupVariableResult struct {
	Columns    []PolicyVariableColumnResponse `pulumi:"columns"`
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}

func LookupVariableOutput(ctx *pulumi.Context, args LookupVariableOutputArgs, opts ...pulumi.InvokeOption) LookupVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableResult, error) {
			args := v.(LookupVariableArgs)
			r, err := LookupVariable(ctx, &args, opts...)
			var s LookupVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableResultOutput)
}

type LookupVariableOutputArgs struct {
	VariableName pulumi.StringInput `pulumi:"variableName"`
}

func (LookupVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableArgs)(nil)).Elem()
}


type LookupVariableResultOutput struct{ *pulumi.OutputState }

func (LookupVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableResult)(nil)).Elem()
}

func (o LookupVariableResultOutput) ToLookupVariableResultOutput() LookupVariableResultOutput {
	return o
}

func (o LookupVariableResultOutput) ToLookupVariableResultOutputWithContext(ctx context.Context) LookupVariableResultOutput {
	return o
}

func (o LookupVariableResultOutput) Columns() PolicyVariableColumnResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableResult) []PolicyVariableColumnResponse { return v.Columns }).(PolicyVariableColumnResponseArrayOutput)
}

func (o LookupVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVariableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableResultOutput{})
}
