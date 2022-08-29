


package managednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScopeAssignment(ctx *pulumi.Context, args *LookupScopeAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupScopeAssignmentResult, error) {
	var rv LookupScopeAssignmentResult
	err := ctx.Invoke("azure-native:managednetwork:getScopeAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScopeAssignmentArgs struct {
	Scope               string `pulumi:"scope"`
	ScopeAssignmentName string `pulumi:"scopeAssignmentName"`
}


type LookupScopeAssignmentResult struct {
	AssignedManagedNetwork *string `pulumi:"assignedManagedNetwork"`
	Etag                   string  `pulumi:"etag"`
	Id                     string  `pulumi:"id"`
	Location               *string `pulumi:"location"`
	Name                   string  `pulumi:"name"`
	ProvisioningState      string  `pulumi:"provisioningState"`
	Type                   string  `pulumi:"type"`
}

func LookupScopeAssignmentOutput(ctx *pulumi.Context, args LookupScopeAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupScopeAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScopeAssignmentResult, error) {
			args := v.(LookupScopeAssignmentArgs)
			r, err := LookupScopeAssignment(ctx, &args, opts...)
			var s LookupScopeAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScopeAssignmentResultOutput)
}

type LookupScopeAssignmentOutputArgs struct {
	Scope               pulumi.StringInput `pulumi:"scope"`
	ScopeAssignmentName pulumi.StringInput `pulumi:"scopeAssignmentName"`
}

func (LookupScopeAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAssignmentArgs)(nil)).Elem()
}


type LookupScopeAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupScopeAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScopeAssignmentResult)(nil)).Elem()
}

func (o LookupScopeAssignmentResultOutput) ToLookupScopeAssignmentResultOutput() LookupScopeAssignmentResultOutput {
	return o
}

func (o LookupScopeAssignmentResultOutput) ToLookupScopeAssignmentResultOutputWithContext(ctx context.Context) LookupScopeAssignmentResultOutput {
	return o
}

func (o LookupScopeAssignmentResultOutput) AssignedManagedNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) *string { return v.AssignedManagedNetwork }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAssignmentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupScopeAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScopeAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupScopeAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScopeAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScopeAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScopeAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScopeAssignmentResultOutput{})
}
