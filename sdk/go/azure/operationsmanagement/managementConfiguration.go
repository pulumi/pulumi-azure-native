


package operationsmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementConfiguration struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                             `pulumi:"name"`
	Properties ManagementConfigurationPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                             `pulumi:"type"`
}


func NewManagementConfiguration(ctx *pulumi.Context,
	name string, args *ManagementConfigurationArgs, opts ...pulumi.ResourceOption) (*ManagementConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationsmanagement/v20151101preview:ManagementConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementConfiguration
	err := ctx.RegisterResource("azure-native:operationsmanagement:ManagementConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementConfigurationState, opts ...pulumi.ResourceOption) (*ManagementConfiguration, error) {
	var resource ManagementConfiguration
	err := ctx.ReadResource("azure-native:operationsmanagement:ManagementConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementConfigurationState struct {
}

type ManagementConfigurationState struct {
}

func (ManagementConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementConfigurationState)(nil)).Elem()
}

type managementConfigurationArgs struct {
	Location                    *string                            `pulumi:"location"`
	ManagementConfigurationName *string                            `pulumi:"managementConfigurationName"`
	Properties                  *ManagementConfigurationProperties `pulumi:"properties"`
	ResourceGroupName           string                             `pulumi:"resourceGroupName"`
}


type ManagementConfigurationArgs struct {
	Location                    pulumi.StringPtrInput
	ManagementConfigurationName pulumi.StringPtrInput
	Properties                  ManagementConfigurationPropertiesPtrInput
	ResourceGroupName           pulumi.StringInput
}

func (ManagementConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementConfigurationArgs)(nil)).Elem()
}

type ManagementConfigurationInput interface {
	pulumi.Input

	ToManagementConfigurationOutput() ManagementConfigurationOutput
	ToManagementConfigurationOutputWithContext(ctx context.Context) ManagementConfigurationOutput
}

func (*ManagementConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementConfiguration)(nil)).Elem()
}

func (i *ManagementConfiguration) ToManagementConfigurationOutput() ManagementConfigurationOutput {
	return i.ToManagementConfigurationOutputWithContext(context.Background())
}

func (i *ManagementConfiguration) ToManagementConfigurationOutputWithContext(ctx context.Context) ManagementConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementConfigurationOutput)
}

type ManagementConfigurationOutput struct{ *pulumi.OutputState }

func (ManagementConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementConfiguration)(nil)).Elem()
}

func (o ManagementConfigurationOutput) ToManagementConfigurationOutput() ManagementConfigurationOutput {
	return o
}

func (o ManagementConfigurationOutput) ToManagementConfigurationOutputWithContext(ctx context.Context) ManagementConfigurationOutput {
	return o
}

func (o ManagementConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ManagementConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementConfigurationOutput) Properties() ManagementConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagementConfiguration) ManagementConfigurationPropertiesResponseOutput { return v.Properties }).(ManagementConfigurationPropertiesResponseOutput)
}

func (o ManagementConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementConfigurationOutput{})
}
