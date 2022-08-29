


package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagementLockAtResourceGroupLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceGroupLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceGroupLevelResult, error) {
	var rv LookupManagementLockAtResourceGroupLevelResult
	err := ctx.Invoke("azure-native:authorization/v20160901:getManagementLockAtResourceGroupLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceGroupLevelArgs struct {
	LockName          string `pulumi:"lockName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagementLockAtResourceGroupLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}

func LookupManagementLockAtResourceGroupLevelOutput(ctx *pulumi.Context, args LookupManagementLockAtResourceGroupLevelOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockAtResourceGroupLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockAtResourceGroupLevelResult, error) {
			args := v.(LookupManagementLockAtResourceGroupLevelArgs)
			r, err := LookupManagementLockAtResourceGroupLevel(ctx, &args, opts...)
			var s LookupManagementLockAtResourceGroupLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockAtResourceGroupLevelResultOutput)
}

type LookupManagementLockAtResourceGroupLevelOutputArgs struct {
	LockName          pulumi.StringInput `pulumi:"lockName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagementLockAtResourceGroupLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceGroupLevelArgs)(nil)).Elem()
}


type LookupManagementLockAtResourceGroupLevelResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockAtResourceGroupLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceGroupLevelResult)(nil)).Elem()
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) ToLookupManagementLockAtResourceGroupLevelResultOutput() LookupManagementLockAtResourceGroupLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) ToLookupManagementLockAtResourceGroupLevelResultOutputWithContext(ctx context.Context) LookupManagementLockAtResourceGroupLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) string { return v.Level }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

func (o LookupManagementLockAtResourceGroupLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceGroupLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockAtResourceGroupLevelResultOutput{})
}
