


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfilesVersion struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ConfigurationProfilePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewConfigurationProfilesVersion(ctx *pulumi.Context,
	name string, args *ConfigurationProfilesVersionArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfilesVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ConfigurationProfilesVersion
	err := ctx.RegisterResource("azure-native:automanage/v20210430preview:ConfigurationProfilesVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfilesVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfilesVersionState, opts ...pulumi.ResourceOption) (*ConfigurationProfilesVersion, error) {
	var resource ConfigurationProfilesVersion
	err := ctx.ReadResource("azure-native:automanage/v20210430preview:ConfigurationProfilesVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfilesVersionState struct {
}

type ConfigurationProfilesVersionState struct {
}

func (ConfigurationProfilesVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilesVersionState)(nil)).Elem()
}

type configurationProfilesVersionArgs struct {
	ConfigurationProfileName string                          `pulumi:"configurationProfileName"`
	Location                 *string                         `pulumi:"location"`
	Properties               *ConfigurationProfileProperties `pulumi:"properties"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	Tags                     map[string]string               `pulumi:"tags"`
	VersionName              *string                         `pulumi:"versionName"`
}


type ConfigurationProfilesVersionArgs struct {
	ConfigurationProfileName pulumi.StringInput
	Location                 pulumi.StringPtrInput
	Properties               ConfigurationProfilePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
	VersionName              pulumi.StringPtrInput
}

func (ConfigurationProfilesVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilesVersionArgs)(nil)).Elem()
}

type ConfigurationProfilesVersionInput interface {
	pulumi.Input

	ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput
	ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput
}

func (*ConfigurationProfilesVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilesVersion)(nil))
}

func (i *ConfigurationProfilesVersion) ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput {
	return i.ToConfigurationProfilesVersionOutputWithContext(context.Background())
}

func (i *ConfigurationProfilesVersion) ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilesVersionOutput)
}

type ConfigurationProfilesVersionOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilesVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilesVersion)(nil))
}

func (o ConfigurationProfilesVersionOutput) ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput {
	return o
}

func (o ConfigurationProfilesVersionOutput) ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfilesVersionOutput{})
}
