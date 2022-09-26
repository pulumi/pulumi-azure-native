


package v20150101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagementLock(ctx *pulumi.Context, args *LookupManagementLockArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockResult, error) {
	var rv LookupManagementLockResult
	err := ctx.Invoke("azure-native:authorization/v20150101:getManagementLock", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockArgs struct {
	LockName string `pulumi:"lockName"`
}


type LookupManagementLockResult struct {
	Id    string  `pulumi:"id"`
	Level *string `pulumi:"level"`
	Name  *string `pulumi:"name"`
	Notes *string `pulumi:"notes"`
	Type  string  `pulumi:"type"`
}

func LookupManagementLockOutput(ctx *pulumi.Context, args LookupManagementLockOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockResult, error) {
			args := v.(LookupManagementLockArgs)
			r, err := LookupManagementLock(ctx, &args, opts...)
			var s LookupManagementLockResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockResultOutput)
}

type LookupManagementLockOutputArgs struct {
	LockName pulumi.StringInput `pulumi:"lockName"`
}

func (LookupManagementLockOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockArgs)(nil)).Elem()
}


type LookupManagementLockResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockResult)(nil)).Elem()
}

func (o LookupManagementLockResultOutput) ToLookupManagementLockResultOutput() LookupManagementLockResultOutput {
	return o
}

func (o LookupManagementLockResultOutput) ToLookupManagementLockResultOutputWithContext(ctx context.Context) LookupManagementLockResultOutput {
	return o
}

func (o LookupManagementLockResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementLockResultOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockResult) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockResultOutput{})
}
