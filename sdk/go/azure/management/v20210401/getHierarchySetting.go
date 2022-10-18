


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHierarchySetting(ctx *pulumi.Context, args *LookupHierarchySettingArgs, opts ...pulumi.InvokeOption) (*LookupHierarchySettingResult, error) {
	var rv LookupHierarchySettingResult
	err := ctx.Invoke("azure-native:management/v20210401:getHierarchySetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHierarchySettingArgs struct {
	GroupId string `pulumi:"groupId"`
}


type LookupHierarchySettingResult struct {
	DefaultManagementGroup               *string `pulumi:"defaultManagementGroup"`
	Id                                   string  `pulumi:"id"`
	Name                                 string  `pulumi:"name"`
	RequireAuthorizationForGroupCreation *bool   `pulumi:"requireAuthorizationForGroupCreation"`
	TenantId                             *string `pulumi:"tenantId"`
	Type                                 string  `pulumi:"type"`
}

func LookupHierarchySettingOutput(ctx *pulumi.Context, args LookupHierarchySettingOutputArgs, opts ...pulumi.InvokeOption) LookupHierarchySettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHierarchySettingResult, error) {
			args := v.(LookupHierarchySettingArgs)
			r, err := LookupHierarchySetting(ctx, &args, opts...)
			var s LookupHierarchySettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHierarchySettingResultOutput)
}

type LookupHierarchySettingOutputArgs struct {
	GroupId pulumi.StringInput `pulumi:"groupId"`
}

func (LookupHierarchySettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHierarchySettingArgs)(nil)).Elem()
}


type LookupHierarchySettingResultOutput struct{ *pulumi.OutputState }

func (LookupHierarchySettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHierarchySettingResult)(nil)).Elem()
}

func (o LookupHierarchySettingResultOutput) ToLookupHierarchySettingResultOutput() LookupHierarchySettingResultOutput {
	return o
}

func (o LookupHierarchySettingResultOutput) ToLookupHierarchySettingResultOutputWithContext(ctx context.Context) LookupHierarchySettingResultOutput {
	return o
}

func (o LookupHierarchySettingResultOutput) DefaultManagementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *string { return v.DefaultManagementGroup }).(pulumi.StringPtrOutput)
}

func (o LookupHierarchySettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHierarchySettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHierarchySettingResultOutput) RequireAuthorizationForGroupCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *bool { return v.RequireAuthorizationForGroupCreation }).(pulumi.BoolPtrOutput)
}

func (o LookupHierarchySettingResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupHierarchySettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHierarchySettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHierarchySettingResultOutput{})
}
