


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotHubResourceEventHubConsumerGroup struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput    `pulumi:"etag"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, args *IotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubEndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devices/v20200801:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20160203:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20160203:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170119:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20170119:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20170701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20180122:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20180401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20181201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190322:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20190701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20191104:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200301:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200615:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200831:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200831preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210201preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210303preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210331:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210702:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20210702preview:IotHubResourceEventHubConsumerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.RegisterResource("azure-native:devices/v20200801:IotHubResourceEventHubConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceEventHubConsumerGroupState, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.ReadResource("azure-native:devices/v20200801:IotHubResourceEventHubConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotHubResourceEventHubConsumerGroupState struct {
}

type IotHubResourceEventHubConsumerGroupState struct {
}

func (IotHubResourceEventHubConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupState)(nil)).Elem()
}

type iotHubResourceEventHubConsumerGroupArgs struct {
	EventHubEndpointName string                     `pulumi:"eventHubEndpointName"`
	Name                 *string                    `pulumi:"name"`
	Properties           *EventHubConsumerGroupName `pulumi:"properties"`
	ResourceGroupName    string                     `pulumi:"resourceGroupName"`
	ResourceName         string                     `pulumi:"resourceName"`
}


type IotHubResourceEventHubConsumerGroupArgs struct {
	EventHubEndpointName pulumi.StringInput
	Name                 pulumi.StringPtrInput
	Properties           EventHubConsumerGroupNamePtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceName         pulumi.StringInput
}

func (IotHubResourceEventHubConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupArgs)(nil)).Elem()
}

type IotHubResourceEventHubConsumerGroupInput interface {
	pulumi.Input

	ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput
	ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput
}

func (*IotHubResourceEventHubConsumerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubResourceEventHubConsumerGroup)(nil))
}

func (i *IotHubResourceEventHubConsumerGroup) ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput {
	return i.ToIotHubResourceEventHubConsumerGroupOutputWithContext(context.Background())
}

func (i *IotHubResourceEventHubConsumerGroup) ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubResourceEventHubConsumerGroupOutput)
}

type IotHubResourceEventHubConsumerGroupOutput struct{ *pulumi.OutputState }

func (IotHubResourceEventHubConsumerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubResourceEventHubConsumerGroup)(nil))
}

func (o IotHubResourceEventHubConsumerGroupOutput) ToIotHubResourceEventHubConsumerGroupOutput() IotHubResourceEventHubConsumerGroupOutput {
	return o
}

func (o IotHubResourceEventHubConsumerGroupOutput) ToIotHubResourceEventHubConsumerGroupOutputWithContext(ctx context.Context) IotHubResourceEventHubConsumerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotHubResourceEventHubConsumerGroupOutput{})
}
