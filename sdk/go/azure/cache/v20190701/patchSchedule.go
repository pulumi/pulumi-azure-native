


package v20190701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PatchSchedule struct {
	pulumi.CustomResourceState

	Name            pulumi.StringOutput              `pulumi:"name"`
	ScheduleEntries ScheduleEntryResponseArrayOutput `pulumi:"scheduleEntries"`
	Type            pulumi.StringOutput              `pulumi:"type"`
}


func NewPatchSchedule(ctx *pulumi.Context,
	name string, args *PatchScheduleArgs, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleEntries == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleEntries'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:PatchSchedule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:PatchSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource PatchSchedule
	err := ctx.RegisterResource("azure-native:cache/v20190701:PatchSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPatchSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchScheduleState, opts ...pulumi.ResourceOption) (*PatchSchedule, error) {
	var resource PatchSchedule
	err := ctx.ReadResource("azure-native:cache/v20190701:PatchSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type patchScheduleState struct {
}

type PatchScheduleState struct {
}

func (PatchScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleState)(nil)).Elem()
}

type patchScheduleArgs struct {
	Default           *string         `pulumi:"default"`
	Name              string          `pulumi:"name"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
	ScheduleEntries   []ScheduleEntry `pulumi:"scheduleEntries"`
}


type PatchScheduleArgs struct {
	Default           pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ScheduleEntries   ScheduleEntryArrayInput
}

func (PatchScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchScheduleArgs)(nil)).Elem()
}

type PatchScheduleInput interface {
	pulumi.Input

	ToPatchScheduleOutput() PatchScheduleOutput
	ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput
}

func (*PatchSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSchedule)(nil)).Elem()
}

func (i *PatchSchedule) ToPatchScheduleOutput() PatchScheduleOutput {
	return i.ToPatchScheduleOutputWithContext(context.Background())
}

func (i *PatchSchedule) ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchScheduleOutput)
}

type PatchScheduleOutput struct{ *pulumi.OutputState }

func (PatchScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSchedule)(nil)).Elem()
}

func (o PatchScheduleOutput) ToPatchScheduleOutput() PatchScheduleOutput {
	return o
}

func (o PatchScheduleOutput) ToPatchScheduleOutputWithContext(ctx context.Context) PatchScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PatchScheduleOutput{})
}
