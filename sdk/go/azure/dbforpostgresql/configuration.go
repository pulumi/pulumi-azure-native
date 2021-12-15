


package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Configuration struct {
	pulumi.CustomResourceState

	AllowedValues pulumi.StringOutput    `pulumi:"allowedValues"`
	DataType      pulumi.StringOutput    `pulumi:"dataType"`
	DefaultValue  pulumi.StringOutput    `pulumi:"defaultValue"`
	Description   pulumi.StringOutput    `pulumi:"description"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	Source        pulumi.StringPtrOutput `pulumi:"source"`
	Type          pulumi.StringOutput    `pulumi:"type"`
	Value         pulumi.StringPtrOutput `pulumi:"value"`
}


func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201preview:Configuration"),
		},
	})
	opts = append(opts, aliases)
	var resource Configuration
	err := ctx.RegisterResource("azure-native:dbforpostgresql:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("azure-native:dbforpostgresql:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationState struct {
}

type ConfigurationState struct {
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	ConfigurationName *string `pulumi:"configurationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServerName        string  `pulumi:"serverName"`
	Source            *string `pulumi:"source"`
	Value             *string `pulumi:"value"`
}


type ConfigurationArgs struct {
	ConfigurationName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	Source            pulumi.StringPtrInput
	Value             pulumi.StringPtrInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
}
