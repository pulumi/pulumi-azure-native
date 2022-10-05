


package v20220504

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfile struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewConfigurationProfile(ctx *pulumi.Context,
	name string, args *ConfigurationProfileArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfile
	err := ctx.RegisterResource("azure-native:automanage/v20220504:ConfigurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileState, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	var resource ConfigurationProfile
	err := ctx.ReadResource("azure-native:automanage/v20220504:ConfigurationProfile", name, id, state, &resource, opts...)
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
	ConfigurationProfileName *string                         `pulumi:"configurationProfileName"`
	Location                 *string                         `pulumi:"location"`
	Properties               *ConfigurationProfileProperties `pulumi:"properties"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	Tags                     map[string]string               `pulumi:"tags"`
}


type ConfigurationProfileArgs struct {
	ConfigurationProfileName pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Properties               ConfigurationProfilePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
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
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (i *ConfigurationProfile) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return i.ToConfigurationProfileOutputWithContext(context.Background())
}

func (i *ConfigurationProfile) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileOutput)
}

type ConfigurationProfileOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ConfigurationProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationProfileOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfile) ConfigurationProfilePropertiesResponseOutput { return v.Properties }).(ConfigurationProfilePropertiesResponseOutput)
}

func (o ConfigurationProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConfigurationProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileOutput{})
}
