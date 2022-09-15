


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockByScope(ctx *pulumi.Context, args *LookupManagementLockByScopeArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockByScopeResult, error) {
	var rv LookupManagementLockByScopeResult
	err := ctx.Invoke("azure-native:authorization/v20200501:getManagementLockByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockByScopeArgs struct {
	LockName string `pulumi:"lockName"`
	Scope    string `pulumi:"scope"`
}


type LookupManagementLockByScopeResult struct {
	Id         string                        `pulumi:"id"`
	Level      string                        `pulumi:"level"`
	Name       string                        `pulumi:"name"`
	Notes      *string                       `pulumi:"notes"`
	Owners     []ManagementLockOwnerResponse `pulumi:"owners"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}

func LookupManagementLockByScopeOutput(ctx *pulumi.Context, args LookupManagementLockByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupManagementLockByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementLockByScopeResult, error) {
			args := v.(LookupManagementLockByScopeArgs)
			r, err := LookupManagementLockByScope(ctx, &args, opts...)
			var s LookupManagementLockByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementLockByScopeResultOutput)
}

type LookupManagementLockByScopeOutputArgs struct {
	LockName pulumi.StringInput `pulumi:"lockName"`
	Scope    pulumi.StringInput `pulumi:"scope"`
}

func (LookupManagementLockByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockByScopeArgs)(nil)).Elem()
}


type LookupManagementLockByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupManagementLockByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementLockByScopeResult)(nil)).Elem()
}

func (o LookupManagementLockByScopeResultOutput) ToLookupManagementLockByScopeResultOutput() LookupManagementLockByScopeResultOutput {
	return o
}

func (o LookupManagementLockByScopeResultOutput) ToLookupManagementLockByScopeResultOutputWithContext(ctx context.Context) LookupManagementLockByScopeResultOutput {
	return o
}

func (o LookupManagementLockByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementLockByScopeResultOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) string { return v.Level }).(pulumi.StringOutput)
}

func (o LookupManagementLockByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementLockByScopeResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupManagementLockByScopeResultOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) []ManagementLockOwnerResponse { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

func (o LookupManagementLockByScopeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagementLockByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementLockByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementLockByScopeResultOutput{})
}
