


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Assignment struct {
	pulumi.CustomResourceState

	AdditionalData    AssignmentPropertiesResponseAdditionalDataPtrOutput `pulumi:"additionalData"`
	AssignedComponent AssignedComponentItemResponsePtrOutput              `pulumi:"assignedComponent"`
	AssignedStandard  AssignedStandardItemResponsePtrOutput               `pulumi:"assignedStandard"`
	Description       pulumi.StringPtrOutput                              `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput                              `pulumi:"displayName"`
	Effect            pulumi.StringPtrOutput                              `pulumi:"effect"`
	Etag              pulumi.StringPtrOutput                              `pulumi:"etag"`
	ExpiresOn         pulumi.StringPtrOutput                              `pulumi:"expiresOn"`
	Kind              pulumi.StringPtrOutput                              `pulumi:"kind"`
	Location          pulumi.StringPtrOutput                              `pulumi:"location"`
	Metadata          pulumi.AnyOutput                                    `pulumi:"metadata"`
	Name              pulumi.StringOutput                                 `pulumi:"name"`
	Scope             pulumi.StringPtrOutput                              `pulumi:"scope"`
	SystemData        SystemDataResponseOutput                            `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                              `pulumi:"tags"`
	Type              pulumi.StringOutput                                 `pulumi:"type"`
}


func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20210801preview:Assignment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assignment
	err := ctx.RegisterResource("azure-native:security:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure-native:security:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assignmentState struct {
}

type AssignmentState struct {
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	AdditionalData    *AssignmentPropertiesAdditionalData `pulumi:"additionalData"`
	AssignedComponent *AssignedComponentItem              `pulumi:"assignedComponent"`
	AssignedStandard  *AssignedStandardItem               `pulumi:"assignedStandard"`
	AssignmentId      *string                             `pulumi:"assignmentId"`
	Description       *string                             `pulumi:"description"`
	DisplayName       *string                             `pulumi:"displayName"`
	Effect            *string                             `pulumi:"effect"`
	ExpiresOn         *string                             `pulumi:"expiresOn"`
	Kind              *string                             `pulumi:"kind"`
	Location          *string                             `pulumi:"location"`
	Metadata          interface{}                         `pulumi:"metadata"`
	ResourceGroupName string                              `pulumi:"resourceGroupName"`
	Scope             *string                             `pulumi:"scope"`
	Tags              map[string]string                   `pulumi:"tags"`
}


type AssignmentArgs struct {
	AdditionalData    AssignmentPropertiesAdditionalDataPtrInput
	AssignedComponent AssignedComponentItemPtrInput
	AssignedStandard  AssignedStandardItemPtrInput
	AssignmentId      pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Effect            pulumi.StringPtrInput
	ExpiresOn         pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Metadata          pulumi.Input
	ResourceGroupName pulumi.StringInput
	Scope             pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}

type AssignmentInput interface {
	pulumi.Input

	ToAssignmentOutput() AssignmentOutput
	ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput
}

func (*Assignment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (i *Assignment) ToAssignmentOutput() AssignmentOutput {
	return i.ToAssignmentOutputWithContext(context.Background())
}

func (i *Assignment) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentOutput)
}

type AssignmentOutput struct{ *pulumi.OutputState }

func (AssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assignment)(nil)).Elem()
}

func (o AssignmentOutput) ToAssignmentOutput() AssignmentOutput {
	return o
}

func (o AssignmentOutput) ToAssignmentOutputWithContext(ctx context.Context) AssignmentOutput {
	return o
}

func (o AssignmentOutput) AdditionalData() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o.ApplyT(func(v *Assignment) AssignmentPropertiesResponseAdditionalDataPtrOutput { return v.AdditionalData }).(AssignmentPropertiesResponseAdditionalDataPtrOutput)
}

func (o AssignmentOutput) AssignedComponent() AssignedComponentItemResponsePtrOutput {
	return o.ApplyT(func(v *Assignment) AssignedComponentItemResponsePtrOutput { return v.AssignedComponent }).(AssignedComponentItemResponsePtrOutput)
}

func (o AssignmentOutput) AssignedStandard() AssignedStandardItemResponsePtrOutput {
	return o.ApplyT(func(v *Assignment) AssignedStandardItemResponsePtrOutput { return v.AssignedStandard }).(AssignedStandardItemResponsePtrOutput)
}

func (o AssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Effect }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Assignment) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o AssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o AssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Assignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AssignmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentOutput{})
}
