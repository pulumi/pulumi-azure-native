


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PeriodicTimerEventTrigger struct {
	pulumi.CustomResourceState

	CustomContextTag pulumi.StringPtrOutput                `pulumi:"customContextTag"`
	Kind             pulumi.StringOutput                   `pulumi:"kind"`
	Name             pulumi.StringOutput                   `pulumi:"name"`
	SinkInfo         RoleSinkInfoResponseOutput            `pulumi:"sinkInfo"`
	SourceInfo       PeriodicTimerSourceInfoResponseOutput `pulumi:"sourceInfo"`
	SystemData       SystemDataResponseOutput              `pulumi:"systemData"`
	Type             pulumi.StringOutput                   `pulumi:"type"`
}


func NewPeriodicTimerEventTrigger(ctx *pulumi.Context,
	name string, args *PeriodicTimerEventTriggerArgs, opts ...pulumi.ResourceOption) (*PeriodicTimerEventTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SinkInfo == nil {
		return nil, errors.New("invalid value for required argument 'SinkInfo'")
	}
	if args.SourceInfo == nil {
		return nil, errors.New("invalid value for required argument 'SourceInfo'")
	}
	args.Kind = pulumi.String("PeriodicTimerEvent")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:PeriodicTimerEventTrigger"),
		},
	})
	opts = append(opts, aliases)
	var resource PeriodicTimerEventTrigger
	err := ctx.RegisterResource("azure-native:databoxedge:PeriodicTimerEventTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeriodicTimerEventTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeriodicTimerEventTriggerState, opts ...pulumi.ResourceOption) (*PeriodicTimerEventTrigger, error) {
	var resource PeriodicTimerEventTrigger
	err := ctx.ReadResource("azure-native:databoxedge:PeriodicTimerEventTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type periodicTimerEventTriggerState struct {
}

type PeriodicTimerEventTriggerState struct {
}

func (PeriodicTimerEventTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*periodicTimerEventTriggerState)(nil)).Elem()
}

type periodicTimerEventTriggerArgs struct {
	CustomContextTag  *string                 `pulumi:"customContextTag"`
	DeviceName        string                  `pulumi:"deviceName"`
	Kind              string                  `pulumi:"kind"`
	Name              *string                 `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	SinkInfo          RoleSinkInfo            `pulumi:"sinkInfo"`
	SourceInfo        PeriodicTimerSourceInfo `pulumi:"sourceInfo"`
}


type PeriodicTimerEventTriggerArgs struct {
	CustomContextTag  pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Kind              pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SinkInfo          RoleSinkInfoInput
	SourceInfo        PeriodicTimerSourceInfoInput
}

func (PeriodicTimerEventTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*periodicTimerEventTriggerArgs)(nil)).Elem()
}

type PeriodicTimerEventTriggerInput interface {
	pulumi.Input

	ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput
	ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput
}

func (*PeriodicTimerEventTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerEventTrigger)(nil)).Elem()
}

func (i *PeriodicTimerEventTrigger) ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput {
	return i.ToPeriodicTimerEventTriggerOutputWithContext(context.Background())
}

func (i *PeriodicTimerEventTrigger) ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerEventTriggerOutput)
}

type PeriodicTimerEventTriggerOutput struct{ *pulumi.OutputState }

func (PeriodicTimerEventTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerEventTrigger)(nil)).Elem()
}

func (o PeriodicTimerEventTriggerOutput) ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput {
	return o
}

func (o PeriodicTimerEventTriggerOutput) ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput {
	return o
}

func (o PeriodicTimerEventTriggerOutput) CustomContextTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringPtrOutput { return v.CustomContextTag }).(pulumi.StringPtrOutput)
}

func (o PeriodicTimerEventTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PeriodicTimerEventTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PeriodicTimerEventTriggerOutput) SinkInfo() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) RoleSinkInfoResponseOutput { return v.SinkInfo }).(RoleSinkInfoResponseOutput)
}

func (o PeriodicTimerEventTriggerOutput) SourceInfo() PeriodicTimerSourceInfoResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) PeriodicTimerSourceInfoResponseOutput { return v.SourceInfo }).(PeriodicTimerSourceInfoResponseOutput)
}

func (o PeriodicTimerEventTriggerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PeriodicTimerEventTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeriodicTimerEventTriggerOutput{})
}
