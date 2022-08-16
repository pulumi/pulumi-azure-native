


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineGroup struct {
	pulumi.CustomResourceState

	Count       pulumi.IntPtrOutput                          `pulumi:"count"`
	DisplayName pulumi.StringOutput                          `pulumi:"displayName"`
	Etag        pulumi.StringPtrOutput                       `pulumi:"etag"`
	GroupType   pulumi.StringPtrOutput                       `pulumi:"groupType"`
	Kind        pulumi.StringOutput                          `pulumi:"kind"`
	Machines    MachineReferenceWithHintsResponseArrayOutput `pulumi:"machines"`
	Name        pulumi.StringOutput                          `pulumi:"name"`
	Type        pulumi.StringOutput                          `pulumi:"type"`
}


func NewMachineGroup(ctx *pulumi.Context,
	name string, args *MachineGroupArgs, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("machineGroup")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:MachineGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineGroup
	err := ctx.RegisterResource("azure-native:operationalinsights/v20151101preview:MachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineGroupState, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	var resource MachineGroup
	err := ctx.ReadResource("azure-native:operationalinsights/v20151101preview:MachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineGroupState struct {
}

type MachineGroupState struct {
}

func (MachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupState)(nil)).Elem()
}

type machineGroupArgs struct {
	Count             *int                        `pulumi:"count"`
	DisplayName       string                      `pulumi:"displayName"`
	GroupType         *string                     `pulumi:"groupType"`
	Kind              string                      `pulumi:"kind"`
	MachineGroupName  *string                     `pulumi:"machineGroupName"`
	Machines          []MachineReferenceWithHints `pulumi:"machines"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	WorkspaceName     string                      `pulumi:"workspaceName"`
}


type MachineGroupArgs struct {
	Count             pulumi.IntPtrInput
	DisplayName       pulumi.StringInput
	GroupType         pulumi.StringPtrInput
	Kind              pulumi.StringInput
	MachineGroupName  pulumi.StringPtrInput
	Machines          MachineReferenceWithHintsArrayInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (MachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupArgs)(nil)).Elem()
}

type MachineGroupInput interface {
	pulumi.Input

	ToMachineGroupOutput() MachineGroupOutput
	ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput
}

func (*MachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (i *MachineGroup) ToMachineGroupOutput() MachineGroupOutput {
	return i.ToMachineGroupOutputWithContext(context.Background())
}

func (i *MachineGroup) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineGroupOutput)
}

type MachineGroupOutput struct{ *pulumi.OutputState }

func (MachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroup)(nil)).Elem()
}

func (o MachineGroupOutput) ToMachineGroupOutput() MachineGroupOutput {
	return o
}

func (o MachineGroupOutput) ToMachineGroupOutputWithContext(ctx context.Context) MachineGroupOutput {
	return o
}

func (o MachineGroupOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o MachineGroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MachineGroupOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MachineGroupOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringPtrOutput { return v.GroupType }).(pulumi.StringPtrOutput)
}

func (o MachineGroupOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MachineGroupOutput) Machines() MachineReferenceWithHintsResponseArrayOutput {
	return o.ApplyT(func(v *MachineGroup) MachineReferenceWithHintsResponseArrayOutput { return v.Machines }).(MachineReferenceWithHintsResponseArrayOutput)
}

func (o MachineGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineGroupOutput{})
}
