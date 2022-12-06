


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationService struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ConfigurationServicePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewConfigurationService(ctx *pulumi.Context,
	name string, args *ConfigurationServiceArgs, opts ...pulumi.ResourceOption) (*ConfigurationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ConfigurationService"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ConfigurationService"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationService
	err := ctx.RegisterResource("azure-native:appplatform/v20220301preview:ConfigurationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationServiceState, opts ...pulumi.ResourceOption) (*ConfigurationService, error) {
	var resource ConfigurationService
	err := ctx.ReadResource("azure-native:appplatform/v20220301preview:ConfigurationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationServiceState struct {
}

type ConfigurationServiceState struct {
}

func (ConfigurationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationServiceState)(nil)).Elem()
}

type configurationServiceArgs struct {
	ConfigurationServiceName *string                         `pulumi:"configurationServiceName"`
	Properties               *ConfigurationServiceProperties `pulumi:"properties"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	ServiceName              string                          `pulumi:"serviceName"`
}


type ConfigurationServiceArgs struct {
	ConfigurationServiceName pulumi.StringPtrInput
	Properties               ConfigurationServicePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	ServiceName              pulumi.StringInput
}

func (ConfigurationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationServiceArgs)(nil)).Elem()
}

type ConfigurationServiceInput interface {
	pulumi.Input

	ToConfigurationServiceOutput() ConfigurationServiceOutput
	ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput
}

func (*ConfigurationService) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationService)(nil)).Elem()
}

func (i *ConfigurationService) ToConfigurationServiceOutput() ConfigurationServiceOutput {
	return i.ToConfigurationServiceOutputWithContext(context.Background())
}

func (i *ConfigurationService) ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationServiceOutput)
}

type ConfigurationServiceOutput struct{ *pulumi.OutputState }

func (ConfigurationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationService)(nil)).Elem()
}

func (o ConfigurationServiceOutput) ToConfigurationServiceOutput() ConfigurationServiceOutput {
	return o
}

func (o ConfigurationServiceOutput) ToConfigurationServiceOutputWithContext(ctx context.Context) ConfigurationServiceOutput {
	return o
}

func (o ConfigurationServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationServiceOutput) Properties() ConfigurationServicePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationService) ConfigurationServicePropertiesResponseOutput { return v.Properties }).(ConfigurationServicePropertiesResponseOutput)
}

func (o ConfigurationServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationServiceOutput{})
}
