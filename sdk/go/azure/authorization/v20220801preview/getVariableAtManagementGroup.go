


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVariableAtManagementGroup(ctx *pulumi.Context, args *LookupVariableAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupVariableAtManagementGroupResult, error) {
	var rv LookupVariableAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariableAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableAtManagementGroupArgs struct {
	ManagementGroupId string `pulumi:"managementGroupId"`
	VariableName      string `pulumi:"variableName"`
}


type LookupVariableAtManagementGroupResult struct {
	Columns    []PolicyVariableColumnResponse `pulumi:"columns"`
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}

func LookupVariableAtManagementGroupOutput(ctx *pulumi.Context, args LookupVariableAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVariableAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableAtManagementGroupResult, error) {
			args := v.(LookupVariableAtManagementGroupArgs)
			r, err := LookupVariableAtManagementGroup(ctx, &args, opts...)
			var s LookupVariableAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableAtManagementGroupResultOutput)
}

type LookupVariableAtManagementGroupOutputArgs struct {
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	VariableName      pulumi.StringInput `pulumi:"variableName"`
}

func (LookupVariableAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableAtManagementGroupArgs)(nil)).Elem()
}


type LookupVariableAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVariableAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableAtManagementGroupResult)(nil)).Elem()
}

func (o LookupVariableAtManagementGroupResultOutput) ToLookupVariableAtManagementGroupResultOutput() LookupVariableAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableAtManagementGroupResultOutput) ToLookupVariableAtManagementGroupResultOutputWithContext(ctx context.Context) LookupVariableAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableAtManagementGroupResultOutput) Columns() PolicyVariableColumnResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableAtManagementGroupResult) []PolicyVariableColumnResponse { return v.Columns }).(PolicyVariableColumnResponseArrayOutput)
}

func (o LookupVariableAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVariableAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVariableAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVariableAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableAtManagementGroupResultOutput{})
}
