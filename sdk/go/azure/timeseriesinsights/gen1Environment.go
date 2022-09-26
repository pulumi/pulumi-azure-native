


package timeseriesinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gen1Environment struct {
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


func NewGen1Environment(ctx *pulumi.Context,
	name string, args *Gen1EnvironmentArgs, opts ...pulumi.ResourceOption) (*Gen1Environment, error) {
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
	args.Kind = pulumi.String("Gen1")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:Gen1Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:Gen1Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:Gen1Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:Gen1Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:Gen1Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:Gen1Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Gen1Environment
	err := ctx.RegisterResource("azure-native:timeseriesinsights:Gen1Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGen1Environment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Gen1EnvironmentState, opts ...pulumi.ResourceOption) (*Gen1Environment, error) {
	var resource Gen1Environment
	err := ctx.ReadResource("azure-native:timeseriesinsights:Gen1Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gen1EnvironmentState struct {
}

type Gen1EnvironmentState struct {
}

func (Gen1EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*gen1EnvironmentState)(nil)).Elem()
}

type gen1EnvironmentArgs struct {
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


type Gen1EnvironmentArgs struct {
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

func (Gen1EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gen1EnvironmentArgs)(nil)).Elem()
}

type Gen1EnvironmentInput interface {
	pulumi.Input

	ToGen1EnvironmentOutput() Gen1EnvironmentOutput
	ToGen1EnvironmentOutputWithContext(ctx context.Context) Gen1EnvironmentOutput
}

func (*Gen1Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Gen1Environment)(nil)).Elem()
}

func (i *Gen1Environment) ToGen1EnvironmentOutput() Gen1EnvironmentOutput {
	return i.ToGen1EnvironmentOutputWithContext(context.Background())
}

func (i *Gen1Environment) ToGen1EnvironmentOutputWithContext(ctx context.Context) Gen1EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Gen1EnvironmentOutput)
}

type Gen1EnvironmentOutput struct{ *pulumi.OutputState }

func (Gen1EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gen1Environment)(nil)).Elem()
}

func (o Gen1EnvironmentOutput) ToGen1EnvironmentOutput() Gen1EnvironmentOutput {
	return o
}

func (o Gen1EnvironmentOutput) ToGen1EnvironmentOutputWithContext(ctx context.Context) Gen1EnvironmentOutput {
	return o
}

func (o Gen1EnvironmentOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.DataAccessId }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) DataRetentionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.DataRetentionTime }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) PartitionKeyProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v *Gen1Environment) TimeSeriesIdPropertyResponseArrayOutput { return v.PartitionKeyProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

func (o Gen1EnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o Gen1EnvironmentOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Gen1Environment) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o Gen1EnvironmentOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v *Gen1Environment) EnvironmentStatusResponseOutput { return v.Status }).(EnvironmentStatusResponseOutput)
}

func (o Gen1EnvironmentOutput) StorageLimitExceededBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringPtrOutput { return v.StorageLimitExceededBehavior }).(pulumi.StringPtrOutput)
}

func (o Gen1EnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o Gen1EnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen1Environment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Gen1EnvironmentOutput{})
}
