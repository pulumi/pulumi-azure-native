


package automanage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfileEnum struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewConfigurationProfileEnum(ctx *pulumi.Context,
	name string, args *ConfigurationProfileEnumArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileEnum, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automanage:ConfigurationProfile"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfile"),
		},
		{
			Type: pulumi.String("azure-nextgen:automanage/v20210430preview:ConfigurationProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfileEnum
	err := ctx.RegisterResource("azure-native:automanage:ConfigurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfileEnum(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileEnumState, opts ...pulumi.ResourceOption) (*ConfigurationProfileEnum, error) {
	var resource ConfigurationProfileEnum
	err := ctx.ReadResource("azure-native:automanage:ConfigurationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfileEnumState struct {
}

type ConfigurationProfileEnumState struct {
}

func (ConfigurationProfileEnumState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileEnumState)(nil)).Elem()
}

type configurationProfileEnumArgs struct {
	ConfigurationProfileName *string                         `pulumi:"configurationProfileName"`
	Location                 *string                         `pulumi:"location"`
	Properties               *ConfigurationProfileProperties `pulumi:"properties"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	Tags                     map[string]string               `pulumi:"tags"`
}


type ConfigurationProfileEnumArgs struct {
	ConfigurationProfileName pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Properties               ConfigurationProfilePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (ConfigurationProfileEnumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileEnumArgs)(nil)).Elem()
}

type ConfigurationProfileEnumInput interface {
	pulumi.Input

	ToConfigurationProfileEnumOutput() ConfigurationProfileEnumOutput
	ToConfigurationProfileEnumOutputWithContext(ctx context.Context) ConfigurationProfileEnumOutput
}

func (*ConfigurationProfileEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileEnum)(nil))
}

func (i *ConfigurationProfileEnum) ToConfigurationProfileEnumOutput() ConfigurationProfileEnumOutput {
	return i.ToConfigurationProfileEnumOutputWithContext(context.Background())
}

func (i *ConfigurationProfileEnum) ToConfigurationProfileEnumOutputWithContext(ctx context.Context) ConfigurationProfileEnumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileEnumOutput)
}

type ConfigurationProfileEnumOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileEnum)(nil))
}

func (o ConfigurationProfileEnumOutput) ToConfigurationProfileEnumOutput() ConfigurationProfileEnumOutput {
	return o
}

func (o ConfigurationProfileEnumOutput) ToConfigurationProfileEnumOutputWithContext(ctx context.Context) ConfigurationProfileEnumOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileEnumOutput{})
}
