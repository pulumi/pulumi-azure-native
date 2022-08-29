


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockAtSubscriptionLevel(ctx *pulumi.Context, args *LookupManagementLockAtSubscriptionLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtSubscriptionLevelResult, error) {
	var rv LookupManagementLockAtSubscriptionLevelResult
	err := ctx.Invoke("azure-native:authorization/v20170401:getManagementLockAtSubscriptionLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtSubscriptionLevelArgs struct {
	LockName string `pulumi:"lockName"`
}


type LookupManagementLockAtSubscriptionLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}

func LookupManagementLockAtSubscriptionLevelOutput(ctx *pulumi.Context, args LookupManagementLockAtSubscriptionLevelOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockAtSubscriptionLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockAtSubscriptionLevelResult, error) {
			args := v.(LookupManagementLockAtSubscriptionLevelArgs)
			r, err := LookupManagementLockAtSubscriptionLevel(ctx, &args, opts...)
			var s LookupManagementLockAtSubscriptionLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockAtSubscriptionLevelResultOutput)
}

type LookupManagementLockAtSubscriptionLevelOutputArgs struct {
	LockName pulumi.StringInput `pulumi:"lockName"`
}

func (LookupManagementLockAtSubscriptionLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtSubscriptionLevelArgs)(nil)).Elem()
}


type LookupManagementLockAtSubscriptionLevelResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockAtSubscriptionLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockAtSubscriptionLevelResult)(nil)).Elem()
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) ToLookupManagementLockAtSubscriptionLevelResultOutput() LookupManagementLockAtSubscriptionLevelResultOutput {
	return o
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) ToLookupManagementLockAtSubscriptionLevelResultOutputWithContext(ctx context.Context) LookupManagementLockAtSubscriptionLevelResultOutput {
	return o
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) string { return v.Level }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

func (o LookupManagementLockAtSubscriptionLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockAtSubscriptionLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockAtSubscriptionLevelResultOutput{})
}
