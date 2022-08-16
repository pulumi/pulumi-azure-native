


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssignment(ctx *pulumi.Context, args *LookupAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupAssignmentResult, error) {
	var rv LookupAssignmentResult
	err := ctx.Invoke("azure-native:security:getAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssignmentArgs struct {
	AssignmentId      string `pulumi:"assignmentId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssignmentResult struct {
	AdditionalData    *AssignmentPropertiesResponseAdditionalData `pulumi:"additionalData"`
	AssignedComponent *AssignedComponentItemResponse              `pulumi:"assignedComponent"`
	AssignedStandard  *AssignedStandardItemResponse               `pulumi:"assignedStandard"`
	Description       *string                                     `pulumi:"description"`
	DisplayName       *string                                     `pulumi:"displayName"`
	Effect            *string                                     `pulumi:"effect"`
	Etag              *string                                     `pulumi:"etag"`
	ExpiresOn         *string                                     `pulumi:"expiresOn"`
	Id                string                                      `pulumi:"id"`
	Kind              *string                                     `pulumi:"kind"`
	Location          *string                                     `pulumi:"location"`
	Metadata          interface{}                                 `pulumi:"metadata"`
	Name              string                                      `pulumi:"name"`
	Scope             *string                                     `pulumi:"scope"`
	SystemData        SystemDataResponse                          `pulumi:"systemData"`
	Tags              map[string]string                           `pulumi:"tags"`
	Type              string                                      `pulumi:"type"`
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
	AssignmentId      pulumi.StringInput `pulumi:"assignmentId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupAssignmentResultOutput) AdditionalData() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *AssignmentPropertiesResponseAdditionalData { return v.AdditionalData }).(AssignmentPropertiesResponseAdditionalDataPtrOutput)
}

func (o LookupAssignmentResultOutput) AssignedComponent() AssignedComponentItemResponsePtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *AssignedComponentItemResponse { return v.AssignedComponent }).(AssignedComponentItemResponsePtrOutput)
}

func (o LookupAssignmentResultOutput) AssignedStandard() AssignedStandardItemResponsePtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *AssignedStandardItemResponse { return v.AssignedStandard }).(AssignedStandardItemResponsePtrOutput)
}

func (o LookupAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Effect }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAssignmentResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAssignmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAssignmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssignmentResultOutput{})
}
