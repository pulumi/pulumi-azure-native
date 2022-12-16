


package v20220120preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Configuration struct {
	pulumi.CustomResourceState

	AllowedValues          pulumi.StringOutput      `pulumi:"allowedValues"`
	DataType               pulumi.StringOutput      `pulumi:"dataType"`
	DefaultValue           pulumi.StringOutput      `pulumi:"defaultValue"`
	Description            pulumi.StringOutput      `pulumi:"description"`
	DocumentationLink      pulumi.StringOutput      `pulumi:"documentationLink"`
	IsConfigPendingRestart pulumi.BoolOutput        `pulumi:"isConfigPendingRestart"`
	IsDynamicConfig        pulumi.BoolOutput        `pulumi:"isDynamicConfig"`
	IsReadOnly             pulumi.BoolOutput        `pulumi:"isReadOnly"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	Source                 pulumi.StringPtrOutput   `pulumi:"source"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
	Unit                   pulumi.StringOutput      `pulumi:"unit"`
	Value                  pulumi.StringPtrOutput   `pulumi:"value"`
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
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210410privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210615privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220308preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20221201:Configuration"),
		},
	})
	opts = append(opts, aliases)
	var resource Configuration
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20220120preview:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20220120preview:Configuration", name, id, state, &resource, opts...)
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

func (o ConfigurationOutput) AllowedValues() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.AllowedValues }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DataType }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DefaultValue }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) DocumentationLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DocumentationLink }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) IsConfigPendingRestart() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsConfigPendingRestart }).(pulumi.BoolOutput)
}

func (o ConfigurationOutput) IsDynamicConfig() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsDynamicConfig }).(pulumi.BoolOutput)
}

func (o ConfigurationOutput) IsReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsReadOnly }).(pulumi.BoolOutput)
}

func (o ConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o ConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Configuration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Unit }).(pulumi.StringOutput)
}

func (o ConfigurationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
}
