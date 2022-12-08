


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockAtResourceLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceLevelResult, error) {
	var rv LookupManagementLockAtResourceLevelResult
	err := ctx.Invoke("azure-native:authorization:getManagementLockAtResourceLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceLevelArgs struct {
	LockName                  string `pulumi:"lockName"`
	ParentResourcePath        string `pulumi:"parentResourcePath"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupManagementLockAtResourceLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}

func LookupManagementLockAtResourceLevelOutput(ctx *pulumi.Context, args LookupManagementLockAtResourceLevelOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockAtResourceLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockAtResourceLevelResult, error) {
			args := v.(LookupManagementLockAtResourceLevelArgs)
			r, err := LookupManagementLockAtResourceLevel(ctx, &args, opts...)
			var s LookupManagementLockAtResourceLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockAtResourceLevelResultOutput)
}

type LookupManagementLockAtResourceLevelOutputArgs struct {
	LockName                  pulumi.StringInput `pulumi:"lockName"`
	ParentResourcePath        pulumi.StringInput `pulumi:"parentResourcePath"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName              pulumi.StringInput `pulumi:"resourceName"`
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	ResourceType              pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupManagementLockAtResourceLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceLevelArgs)(nil)).Elem()
}


type LookupManagementLockAtResourceLevelResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockAtResourceLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtResourceLevelResult)(nil)).Elem()
}

func (o LookupManagementLockAtResourceLevelResultOutput) ToLookupManagementLockAtResourceLevelResultOutput() LookupManagementLockAtResourceLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceLevelResultOutput) ToLookupManagementLockAtResourceLevelResultOutputWithContext(ctx context.Context) LookupManagementLockAtResourceLevelResultOutput {
	return o
}

func (o LookupManagementLockAtResourceLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceLevelResultOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Level }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtResourceLevelResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockAtResourceLevelResultOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

func (o LookupManagementLockAtResourceLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtResourceLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockAtResourceLevelResultOutput{})
}
