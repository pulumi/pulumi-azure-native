


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfile struct {
	pulumi.CustomResourceState

	Identity   ResourceIdentityResponsePtrOutput                    `pulumi:"identity"`
	Location   pulumi.StringPtrOutput                               `pulumi:"location"`
	Name       pulumi.StringOutput                                  `pulumi:"name"`
	Properties ConfigurationProfileResourcePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponsePtrOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                                  `pulumi:"type"`
}


func NewConfigurationProfile(ctx *pulumi.Context,
	name string, args *ConfigurationProfileArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	if args == nil {
		args = &ConfigurationProfileArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:changeanalysis/v20200401preview:ConfigurationProfile"),
		},
		{
			Type: pulumi.String("azure-native:changeanalysis:ConfigurationProfile"),
		},
		{
			Type: pulumi.String("azure-nextgen:changeanalysis:ConfigurationProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfile
	err := ctx.RegisterResource("azure-native:changeanalysis/v20200401preview:ConfigurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileState, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	var resource ConfigurationProfile
	err := ctx.ReadResource("azure-native:changeanalysis/v20200401preview:ConfigurationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfileState struct {
}

type ConfigurationProfileState struct {
}

func (ConfigurationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileState)(nil)).Elem()
}

type configurationProfileArgs struct {
	Identity    *ResourceIdentity                       `pulumi:"identity"`
	Location    *string                                 `pulumi:"location"`
	ProfileName *string                                 `pulumi:"profileName"`
	Properties  *ConfigurationProfileResourceProperties `pulumi:"properties"`
}


type ConfigurationProfileArgs struct {
	Identity    ResourceIdentityPtrInput
	Location    pulumi.StringPtrInput
	ProfileName pulumi.StringPtrInput
	Properties  ConfigurationProfileResourcePropertiesPtrInput
}

func (ConfigurationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileArgs)(nil)).Elem()
}

type ConfigurationProfileInput interface {
	pulumi.Input

	ToConfigurationProfileOutput() ConfigurationProfileOutput
	ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput
}

func (*ConfigurationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfile)(nil))
}

func (i *ConfigurationProfile) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return i.ToConfigurationProfileOutputWithContext(context.Background())
}

func (i *ConfigurationProfile) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileOutput)
}

type ConfigurationProfileOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfile)(nil))
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileOutput{})
}
