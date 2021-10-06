


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StandardEnvironment struct {
	pulumi.CustomResourceState

	CreationTime                 pulumi.StringOutput                     `pulumi:"creationTime"`
	DataAccessFqdn               pulumi.StringOutput                     `pulumi:"dataAccessFqdn"`
	DataAccessId                 pulumi.StringOutput                     `pulumi:"dataAccessId"`
	DataRetentionTime            pulumi.StringOutput                     `pulumi:"dataRetentionTime"`
	Kind                         pulumi.StringOutput                     `pulumi:"kind"`
	Location                     pulumi.StringOutput                     `pulumi:"location"`
	Name                         pulumi.StringOutput                     `pulumi:"name"`
	PartitionKeyProperties       TimeSeriesIdPropertyResponseArrayOutput `pulumi:"partitionKeyProperties"`
	ProvisioningState            pulumi.StringOutput                     `pulumi:"provisioningState"`
	Sku                          SkuResponseOutput                       `pulumi:"sku"`
	Status                       EnvironmentStatusResponseOutput         `pulumi:"status"`
	StorageLimitExceededBehavior pulumi.StringPtrOutput                  `pulumi:"storageLimitExceededBehavior"`
	Tags                         pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                         pulumi.StringOutput                     `pulumi:"type"`
}


func NewStandardEnvironment(ctx *pulumi.Context,
	name string, args *StandardEnvironmentArgs, opts ...pulumi.ResourceOption) (*StandardEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataRetentionTime == nil {
		return nil, errors.New("invalid value for required argument 'DataRetentionTime'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	args.Kind = pulumi.String("Standard")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20180815preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20170228preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20171115:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20200515:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210331preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:StandardEnvironment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210630preview:StandardEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource StandardEnvironment
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20180815preview:StandardEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStandardEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StandardEnvironmentState, opts ...pulumi.ResourceOption) (*StandardEnvironment, error) {
	var resource StandardEnvironment
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20180815preview:StandardEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type standardEnvironmentState struct {
}

type StandardEnvironmentState struct {
}

func (StandardEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*standardEnvironmentState)(nil)).Elem()
}

type standardEnvironmentArgs struct {
	DataRetentionTime            string                 `pulumi:"dataRetentionTime"`
	EnvironmentName              *string                `pulumi:"environmentName"`
	Kind                         string                 `pulumi:"kind"`
	Location                     *string                `pulumi:"location"`
	PartitionKeyProperties       []TimeSeriesIdProperty `pulumi:"partitionKeyProperties"`
	ResourceGroupName            string                 `pulumi:"resourceGroupName"`
	Sku                          Sku                    `pulumi:"sku"`
	StorageLimitExceededBehavior *string                `pulumi:"storageLimitExceededBehavior"`
	Tags                         map[string]string      `pulumi:"tags"`
}


type StandardEnvironmentArgs struct {
	DataRetentionTime            pulumi.StringInput
	EnvironmentName              pulumi.StringPtrInput
	Kind                         pulumi.StringInput
	Location                     pulumi.StringPtrInput
	PartitionKeyProperties       TimeSeriesIdPropertyArrayInput
	ResourceGroupName            pulumi.StringInput
	Sku                          SkuInput
	StorageLimitExceededBehavior pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
}

func (StandardEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*standardEnvironmentArgs)(nil)).Elem()
}

type StandardEnvironmentInput interface {
	pulumi.Input

	ToStandardEnvironmentOutput() StandardEnvironmentOutput
	ToStandardEnvironmentOutputWithContext(ctx context.Context) StandardEnvironmentOutput
}

func (*StandardEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEnvironment)(nil))
}

func (i *StandardEnvironment) ToStandardEnvironmentOutput() StandardEnvironmentOutput {
	return i.ToStandardEnvironmentOutputWithContext(context.Background())
}

func (i *StandardEnvironment) ToStandardEnvironmentOutputWithContext(ctx context.Context) StandardEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardEnvironmentOutput)
}

type StandardEnvironmentOutput struct{ *pulumi.OutputState }

func (StandardEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardEnvironment)(nil))
}

func (o StandardEnvironmentOutput) ToStandardEnvironmentOutput() StandardEnvironmentOutput {
	return o
}

func (o StandardEnvironmentOutput) ToStandardEnvironmentOutputWithContext(ctx context.Context) StandardEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StandardEnvironmentOutput{})
}
