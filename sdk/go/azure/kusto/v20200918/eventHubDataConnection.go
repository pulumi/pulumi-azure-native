


package v20200918

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubDataConnection struct {
	pulumi.CustomResourceState

	Compression           pulumi.StringPtrOutput   `pulumi:"compression"`
	ConsumerGroup         pulumi.StringOutput      `pulumi:"consumerGroup"`
	DataFormat            pulumi.StringPtrOutput   `pulumi:"dataFormat"`
	EventHubResourceId    pulumi.StringOutput      `pulumi:"eventHubResourceId"`
	EventSystemProperties pulumi.StringArrayOutput `pulumi:"eventSystemProperties"`
	Kind                  pulumi.StringOutput      `pulumi:"kind"`
	Location              pulumi.StringPtrOutput   `pulumi:"location"`
	MappingRuleName       pulumi.StringPtrOutput   `pulumi:"mappingRuleName"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	TableName             pulumi.StringPtrOutput   `pulumi:"tableName"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewEventHubDataConnection(ctx *pulumi.Context,
	name string, args *EventHubDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventHubDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventHubResourceId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Kind = pulumi.String("EventHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:EventHubDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubDataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20200918:EventHubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubDataConnectionState, opts ...pulumi.ResourceOption) (*EventHubDataConnection, error) {
	var resource EventHubDataConnection
	err := ctx.ReadResource("azure-native:kusto/v20200918:EventHubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubDataConnectionState struct {
}

type EventHubDataConnectionState struct {
}

func (EventHubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubDataConnectionState)(nil)).Elem()
}

type eventHubDataConnectionArgs struct {
	ClusterName           string   `pulumi:"clusterName"`
	Compression           *string  `pulumi:"compression"`
	ConsumerGroup         string   `pulumi:"consumerGroup"`
	DataConnectionName    *string  `pulumi:"dataConnectionName"`
	DataFormat            *string  `pulumi:"dataFormat"`
	DatabaseName          string   `pulumi:"databaseName"`
	EventHubResourceId    string   `pulumi:"eventHubResourceId"`
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	Kind                  string   `pulumi:"kind"`
	Location              *string  `pulumi:"location"`
	MappingRuleName       *string  `pulumi:"mappingRuleName"`
	ResourceGroupName     string   `pulumi:"resourceGroupName"`
	TableName             *string  `pulumi:"tableName"`
}


type EventHubDataConnectionArgs struct {
	ClusterName           pulumi.StringInput
	Compression           pulumi.StringPtrInput
	ConsumerGroup         pulumi.StringInput
	DataConnectionName    pulumi.StringPtrInput
	DataFormat            pulumi.StringPtrInput
	DatabaseName          pulumi.StringInput
	EventHubResourceId    pulumi.StringInput
	EventSystemProperties pulumi.StringArrayInput
	Kind                  pulumi.StringInput
	Location              pulumi.StringPtrInput
	MappingRuleName       pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	TableName             pulumi.StringPtrInput
}

func (EventHubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubDataConnectionArgs)(nil)).Elem()
}

type EventHubDataConnectionInput interface {
	pulumi.Input

	ToEventHubDataConnectionOutput() EventHubDataConnectionOutput
	ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput
}

func (*EventHubDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubDataConnection)(nil)).Elem()
}

func (i *EventHubDataConnection) ToEventHubDataConnectionOutput() EventHubDataConnectionOutput {
	return i.ToEventHubDataConnectionOutputWithContext(context.Background())
}

func (i *EventHubDataConnection) ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubDataConnectionOutput)
}

type EventHubDataConnectionOutput struct{ *pulumi.OutputState }

func (EventHubDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubDataConnection)(nil)).Elem()
}

func (o EventHubDataConnectionOutput) ToEventHubDataConnectionOutput() EventHubDataConnectionOutput {
	return o
}

func (o EventHubDataConnectionOutput) ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput {
	return o
}

func (o EventHubDataConnectionOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o EventHubDataConnectionOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o EventHubDataConnectionOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.DataFormat }).(pulumi.StringPtrOutput)
}

func (o EventHubDataConnectionOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.EventHubResourceId }).(pulumi.StringOutput)
}

func (o EventHubDataConnectionOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringArrayOutput { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

func (o EventHubDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EventHubDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EventHubDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o EventHubDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubDataConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventHubDataConnectionOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o EventHubDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubDataConnectionOutput{})
}
