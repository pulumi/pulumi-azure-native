


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariableValue(ctx *pulumi.Context, args *LookupVariableValueArgs, opts ...pulumi.InvokeOption) (*LookupVariableValueResult, error) {
	var rv LookupVariableValueResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariableValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableValueArgs struct {
	VariableName      string `pulumi:"variableName"`
	VariableValueName string `pulumi:"variableValueName"`
}


type LookupVariableValueResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
	Values     []PolicyVariableValueColumnValueResponse `pulumi:"values"`
}

func LookupVariableValueOutput(ctx *pulumi.Context, args LookupVariableValueOutputArgs, opts ...pulumi.InvokeOption) LookupVariableValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableValueResult, error) {
			args := v.(LookupVariableValueArgs)
			r, err := LookupVariableValue(ctx, &args, opts...)
			var s LookupVariableValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableValueResultOutput)
}

type LookupVariableValueOutputArgs struct {
	VariableName      pulumi.StringInput `pulumi:"variableName"`
	VariableValueName pulumi.StringInput `pulumi:"variableValueName"`
}

func (LookupVariableValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueArgs)(nil)).Elem()
}


type LookupVariableValueResultOutput struct{ *pulumi.OutputState }

func (LookupVariableValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueResult)(nil)).Elem()
}

func (o LookupVariableValueResultOutput) ToLookupVariableValueResultOutput() LookupVariableValueResultOutput {
	return o
}

func (o LookupVariableValueResultOutput) ToLookupVariableValueResultOutputWithContext(ctx context.Context) LookupVariableValueResultOutput {
	return o
}

func (o LookupVariableValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVariableValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableValueResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableValueResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVariableValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVariableValueResultOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableValueResult) []PolicyVariableValueColumnValueResponse { return v.Values }).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableValueResultOutput{})
}
