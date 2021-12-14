


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LongTermEnvironment struct {
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
	StorageConfiguration   LongTermStorageConfigurationOutputResponseOutput  `pulumi:"storageConfiguration"`
	Tags                   pulumi.StringMapOutput                            `pulumi:"tags"`
	TimeSeriesIdProperties TimeSeriesIdPropertyResponseArrayOutput           `pulumi:"timeSeriesIdProperties"`
	Type                   pulumi.StringOutput                               `pulumi:"type"`
	WarmStoreConfiguration WarmStoreConfigurationPropertiesResponsePtrOutput `pulumi:"warmStoreConfiguration"`
}


func NewLongTermEnvironment(ctx *pulumi.Context,
	name string, args *LongTermEnvironmentArgs, opts ...pulumi.ResourceOption) (*LongTermEnvironment, error) {
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
	args.Kind = pulumi.String("LongTerm")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:LongTermEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:LongTermEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:LongTermEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:LongTermEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:LongTermEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:LongTermEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource LongTermEnvironment
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20180815preview:LongTermEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLongTermEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LongTermEnvironmentState, opts ...pulumi.ResourceOption) (*LongTermEnvironment, error) {
	var resource LongTermEnvironment
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20180815preview:LongTermEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type longTermEnvironmentState struct {
}

type LongTermEnvironmentState struct {
}

func (LongTermEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermEnvironmentState)(nil)).Elem()
}

type longTermEnvironmentArgs struct {
	EnvironmentName        *string                           `pulumi:"environmentName"`
	Kind                   string                            `pulumi:"kind"`
	Location               *string                           `pulumi:"location"`
	ResourceGroupName      string                            `pulumi:"resourceGroupName"`
	Sku                    Sku                               `pulumi:"sku"`
	StorageConfiguration   LongTermStorageConfigurationInput `pulumi:"storageConfiguration"`
	Tags                   map[string]string                 `pulumi:"tags"`
	TimeSeriesIdProperties []TimeSeriesIdProperty            `pulumi:"timeSeriesIdProperties"`
	WarmStoreConfiguration *WarmStoreConfigurationProperties `pulumi:"warmStoreConfiguration"`
}


type LongTermEnvironmentArgs struct {
	EnvironmentName        pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuInput
	StorageConfiguration   LongTermStorageConfigurationInputInput
	Tags                   pulumi.StringMapInput
	TimeSeriesIdProperties TimeSeriesIdPropertyArrayInput
	WarmStoreConfiguration WarmStoreConfigurationPropertiesPtrInput
}

func (LongTermEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*longTermEnvironmentArgs)(nil)).Elem()
}

type LongTermEnvironmentInput interface {
	pulumi.Input

	ToLongTermEnvironmentOutput() LongTermEnvironmentOutput
	ToLongTermEnvironmentOutputWithContext(ctx context.Context) LongTermEnvironmentOutput
}

func (*LongTermEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermEnvironment)(nil))
}

func (i *LongTermEnvironment) ToLongTermEnvironmentOutput() LongTermEnvironmentOutput {
	return i.ToLongTermEnvironmentOutputWithContext(context.Background())
}

func (i *LongTermEnvironment) ToLongTermEnvironmentOutputWithContext(ctx context.Context) LongTermEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermEnvironmentOutput)
}

type LongTermEnvironmentOutput struct{ *pulumi.OutputState }

func (LongTermEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermEnvironment)(nil))
}

func (o LongTermEnvironmentOutput) ToLongTermEnvironmentOutput() LongTermEnvironmentOutput {
	return o
}

func (o LongTermEnvironmentOutput) ToLongTermEnvironmentOutputWithContext(ctx context.Context) LongTermEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LongTermEnvironmentOutput{})
}
