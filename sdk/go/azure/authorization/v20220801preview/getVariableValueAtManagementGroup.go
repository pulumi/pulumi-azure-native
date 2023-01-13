


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariableValueAtManagementGroup(ctx *pulumi.Context, args *LookupVariableValueAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupVariableValueAtManagementGroupResult, error) {
	var rv LookupVariableValueAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariableValueAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableValueAtManagementGroupArgs struct {
	ManagementGroupId string `pulumi:"managementGroupId"`
	VariableName      string `pulumi:"variableName"`
	VariableValueName string `pulumi:"variableValueName"`
}


type LookupVariableValueAtManagementGroupResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
	Values     []PolicyVariableValueColumnValueResponse `pulumi:"values"`
}

func LookupVariableValueAtManagementGroupOutput(ctx *pulumi.Context, args LookupVariableValueAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVariableValueAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableValueAtManagementGroupResult, error) {
			args := v.(LookupVariableValueAtManagementGroupArgs)
			r, err := LookupVariableValueAtManagementGroup(ctx, &args, opts...)
			var s LookupVariableValueAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableValueAtManagementGroupResultOutput)
}

type LookupVariableValueAtManagementGroupOutputArgs struct {
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	VariableName      pulumi.StringInput `pulumi:"variableName"`
	VariableValueName pulumi.StringInput `pulumi:"variableValueName"`
}

func (LookupVariableValueAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueAtManagementGroupArgs)(nil)).Elem()
}


type LookupVariableValueAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVariableValueAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueAtManagementGroupResult)(nil)).Elem()
}

func (o LookupVariableValueAtManagementGroupResultOutput) ToLookupVariableValueAtManagementGroupResultOutput() LookupVariableValueAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableValueAtManagementGroupResultOutput) ToLookupVariableValueAtManagementGroupResultOutputWithContext(ctx context.Context) LookupVariableValueAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableValueAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVariableValueAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableValueAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVariableValueAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVariableValueAtManagementGroupResultOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) []PolicyVariableValueColumnValueResponse {
		return v.Values
	}).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableValueAtManagementGroupResultOutput{})
}
