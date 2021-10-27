


package v20200515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gen2Environment struct {
	pulumi.CustomResourceState

	CreationTime           pulumi.StringOutput                               `pulumi:"creationTime"`
	DataAccessFqdn         pulumi.StringOutput                               `pulumi:"dataAccessFqdn"`
	DataAccessId           pulumi.StringOutput                               `pulumi:"dataAccessId"`
	Kind                   pulumi.StringOutput                               `pulumi:"kind"`
	Location               pulumi.StringOutput                               `pulumi:"location"`
	Name                   pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                               `pulumi:"provisioningState"`
	Sku                    SkuResponseOutput                                 `pulumi:"sku"`
	Status                 EnvironmentStatusResponseOutput                   `pulumi:"status"`
	StorageConfiguration   Gen2StorageConfigurationOutputResponseOutput      `pulumi:"storageConfiguration"`
	Tags                   pulumi.StringMapOutput                            `pulumi:"tags"`
	TimeSeriesIdProperties TimeSeriesIdPropertyResponseArrayOutput           `pulumi:"timeSeriesIdProperties"`
	Type                   pulumi.StringOutput                               `pulumi:"type"`
	WarmStoreConfiguration WarmStoreConfigurationPropertiesResponsePtrOutput `pulumi:"warmStoreConfiguration"`
}


func NewGen2Environment(ctx *pulumi.Context,
	name string, args *Gen2EnvironmentArgs, opts ...pulumi.ResourceOption) (*Gen2Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.StorageConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'StorageConfiguration'")
	}
	if args.TimeSeriesIdProperties == nil {
		return nil, errors.New("invalid value for required argument 'TimeSeriesIdProperties'")
	}
	args.Kind = pulumi.String("Gen2")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20200515:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20170228preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20171115:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20180815preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210331preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210630preview:Gen2Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Gen2Environment
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20200515:Gen2Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGen2Environment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Gen2EnvironmentState, opts ...pulumi.ResourceOption) (*Gen2Environment, error) {
	var resource Gen2Environment
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20200515:Gen2Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gen2EnvironmentState struct {
}

type Gen2EnvironmentState struct {
}

func (Gen2EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*gen2EnvironmentState)(nil)).Elem()
}

type gen2EnvironmentArgs struct {
	EnvironmentName        *string                           `pulumi:"environmentName"`
	Kind                   string                            `pulumi:"kind"`
	Location               *string                           `pulumi:"location"`
	ResourceGroupName      string                            `pulumi:"resourceGroupName"`
	Sku                    Sku                               `pulumi:"sku"`
	StorageConfiguration   Gen2StorageConfigurationInput     `pulumi:"storageConfiguration"`
	Tags                   map[string]string                 `pulumi:"tags"`
	TimeSeriesIdProperties []TimeSeriesIdProperty            `pulumi:"timeSeriesIdProperties"`
	WarmStoreConfiguration *WarmStoreConfigurationProperties `pulumi:"warmStoreConfiguration"`
}


type Gen2EnvironmentArgs struct {
	EnvironmentName        pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuInput
	StorageConfiguration   Gen2StorageConfigurationInputInput
	Tags                   pulumi.StringMapInput
	TimeSeriesIdProperties TimeSeriesIdPropertyArrayInput
	WarmStoreConfiguration WarmStoreConfigurationPropertiesPtrInput
}

func (Gen2EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gen2EnvironmentArgs)(nil)).Elem()
}

type Gen2EnvironmentInput interface {
	pulumi.Input

	ToGen2EnvironmentOutput() Gen2EnvironmentOutput
	ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput
}

func (*Gen2Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*Gen2Environment)(nil))
}

func (i *Gen2Environment) ToGen2EnvironmentOutput() Gen2EnvironmentOutput {
	return i.ToGen2EnvironmentOutputWithContext(context.Background())
}

func (i *Gen2Environment) ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Gen2EnvironmentOutput)
}

type Gen2EnvironmentOutput struct{ *pulumi.OutputState }

func (Gen2EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gen2Environment)(nil))
}

func (o Gen2EnvironmentOutput) ToGen2EnvironmentOutput() Gen2EnvironmentOutput {
	return o
}

func (o Gen2EnvironmentOutput) ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Gen2EnvironmentInput)(nil)).Elem(), &Gen2Environment{})
	pulumi.RegisterOutputType(Gen2EnvironmentOutput{})
}
