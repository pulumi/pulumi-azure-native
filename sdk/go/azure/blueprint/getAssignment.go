


package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssignment(ctx *pulumi.Context, args *LookupAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupAssignmentResult, error) {
	var rv LookupAssignmentResult
	err := ctx.Invoke("azure-native:blueprint:getAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssignmentArgs struct {
	AssignmentName string `pulumi:"assignmentName"`
	ResourceScope  string `pulumi:"resourceScope"`
}


type LookupAssignmentResult struct {
	BlueprintId       *string                               `pulumi:"blueprintId"`
	Description       *string                               `pulumi:"description"`
	DisplayName       *string                               `pulumi:"displayName"`
	Id                string                                `pulumi:"id"`
	Identity          ManagedServiceIdentityResponse        `pulumi:"identity"`
	Location          string                                `pulumi:"location"`
	Locks             *AssignmentLockSettingsResponse       `pulumi:"locks"`
	Name              string                                `pulumi:"name"`
	Parameters        map[string]ParameterValueResponse     `pulumi:"parameters"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	ResourceGroups    map[string]ResourceGroupValueResponse `pulumi:"resourceGroups"`
	Scope             *string                               `pulumi:"scope"`
	Status            AssignmentStatusResponse              `pulumi:"status"`
	Type              string                                `pulumi:"type"`
}

func LookupAssignmentOutput(ctx *pulumi.Context, args LookupAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssignmentResult, error) {
			args := v.(LookupAssignmentArgs)
			r, err := LookupAssignment(ctx, &args, opts...)
			var s LookupAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssignmentResultOutput)
}

type LookupAssignmentOutputArgs struct {
	AssignmentName pulumi.StringInput `pulumi:"assignmentName"`
	ResourceScope  pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssignmentArgs)(nil)).Elem()
}


type LookupAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssignmentResult)(nil)).Elem()
}

func (o LookupAssignmentResultOutput) ToLookupAssignmentResultOutput() LookupAssignmentResultOutput {
	return o
}

func (o LookupAssignmentResultOutput) ToLookupAssignmentResultOutputWithContext(ctx context.Context) LookupAssignmentResultOutput {
	return o
}

func (o LookupAssignmentResultOutput) BlueprintId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.BlueprintId }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) Identity() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v LookupAssignmentResult) ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponseOutput)
}

func (o LookupAssignmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) Locks() AssignmentLockSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *AssignmentLockSettingsResponse { return v.Locks }).(AssignmentLockSettingsResponsePtrOutput)
}

func (o LookupAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v LookupAssignmentResult) map[string]ParameterValueResponse { return v.Parameters }).(ParameterValueResponseMapOutput)
}

func (o LookupAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) ResourceGroups() ResourceGroupValueResponseMapOutput {
	return o.ApplyT(func(v LookupAssignmentResult) map[string]ResourceGroupValueResponse { return v.ResourceGroups }).(ResourceGroupValueResponseMapOutput)
}

func (o LookupAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Status() AssignmentStatusResponseOutput {
	return o.ApplyT(func(v LookupAssignmentResult) AssignmentStatusResponse { return v.Status }).(AssignmentStatusResponseOutput)
}

func (o LookupAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssignmentResultOutput{})
}
