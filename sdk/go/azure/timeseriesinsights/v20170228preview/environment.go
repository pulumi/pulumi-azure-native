


package v20170228preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Environment struct {
	pulumi.CustomResourceState

	CreationTime                 pulumi.StringOutput    `pulumi:"creationTime"`
	DataAccessFqdn               pulumi.StringOutput    `pulumi:"dataAccessFqdn"`
	DataAccessId                 pulumi.StringOutput    `pulumi:"dataAccessId"`
	DataRetentionTime            pulumi.StringOutput    `pulumi:"dataRetentionTime"`
	Location                     pulumi.StringOutput    `pulumi:"location"`
	Name                         pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState            pulumi.StringOutput    `pulumi:"provisioningState"`
	Sku                          SkuResponsePtrOutput   `pulumi:"sku"`
	StorageLimitExceededBehavior pulumi.StringPtrOutput `pulumi:"storageLimitExceededBehavior"`
	Tags                         pulumi.StringMapOutput `pulumi:"tags"`
	Type                         pulumi.StringOutput    `pulumi:"type"`
}


func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataRetentionTime == nil {
		return nil, errors.New("invalid value for required argument 'DataRetentionTime'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20170228preview:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20171115:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20180815preview:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20200515:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210331preview:Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:Environment"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210630preview:Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Environment
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20170228preview:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20170228preview:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	DataRetentionTime            string                        `pulumi:"dataRetentionTime"`
	EnvironmentName              *string                       `pulumi:"environmentName"`
	Location                     *string                       `pulumi:"location"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Sku                          Sku                           `pulumi:"sku"`
	StorageLimitExceededBehavior *StorageLimitExceededBehavior `pulumi:"storageLimitExceededBehavior"`
	Tags                         map[string]string             `pulumi:"tags"`
}


type EnvironmentArgs struct {
	DataRetentionTime            pulumi.StringInput
	EnvironmentName              pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          SkuInput
	StorageLimitExceededBehavior StorageLimitExceededBehaviorPtrInput
	Tags                         pulumi.StringMapInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
