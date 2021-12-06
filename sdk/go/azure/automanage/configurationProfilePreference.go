


package automanage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfilePreference struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                                    `pulumi:"location"`
	Name       pulumi.StringOutput                                    `pulumi:"name"`
	Properties ConfigurationProfilePreferencePropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type       pulumi.StringOutput                                    `pulumi:"type"`
}


func NewConfigurationProfilePreference(ctx *pulumi.Context,
	name string, args *ConfigurationProfilePreferenceArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20200630preview:ConfigurationProfilePreference"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfilePreference
	err := ctx.RegisterResource("azure-native:automanage:ConfigurationProfilePreference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfilePreference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfilePreferenceState, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	var resource ConfigurationProfilePreference
	err := ctx.ReadResource("azure-native:automanage:ConfigurationProfilePreference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfilePreferenceState struct {
}

type ConfigurationProfilePreferenceState struct {
}

func (ConfigurationProfilePreferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceState)(nil)).Elem()
}

type configurationProfilePreferenceArgs struct {
	ConfigurationProfilePreferenceName *string                                   `pulumi:"configurationProfilePreferenceName"`
	Location                           *string                                   `pulumi:"location"`
	Properties                         *ConfigurationProfilePreferenceProperties `pulumi:"properties"`
	ResourceGroupName                  string                                    `pulumi:"resourceGroupName"`
	Tags                               map[string]string                         `pulumi:"tags"`
}


type ConfigurationProfilePreferenceArgs struct {
	ConfigurationProfilePreferenceName pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	Properties                         ConfigurationProfilePreferencePropertiesPtrInput
	ResourceGroupName                  pulumi.StringInput
	Tags                               pulumi.StringMapInput
}

func (ConfigurationProfilePreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceArgs)(nil)).Elem()
}

type ConfigurationProfilePreferenceInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput
	ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput
}

func (*ConfigurationProfilePreference) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreference)(nil))
}

func (i *ConfigurationProfilePreference) ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput {
	return i.ToConfigurationProfilePreferenceOutputWithContext(context.Background())
}

func (i *ConfigurationProfilePreference) ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceOutput)
}

type ConfigurationProfilePreferenceOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePreference)(nil))
}

func (o ConfigurationProfilePreferenceOutput) ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput {
	return o
}

func (o ConfigurationProfilePreferenceOutput) ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceOutput{})
}
