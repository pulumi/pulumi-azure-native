


package v20151101preview

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
			Type: pulumi.String("azure-nextgen:operationsmanagement/v20151101preview:ManagementConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:operationsmanagement:ManagementConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationsmanagement:ManagementConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementConfiguration
	err := ctx.RegisterResource("azure-native:operationsmanagement/v20151101preview:ManagementConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementConfigurationState, opts ...pulumi.ResourceOption) (*ManagementConfiguration, error) {
	var resource ManagementConfiguration
	err := ctx.ReadResource("azure-native:operationsmanagement/v20151101preview:ManagementConfiguration", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*ManagementConfiguration)(nil))
}

func (i *ManagementConfiguration) ToManagementConfigurationOutput() ManagementConfigurationOutput {
	return i.ToManagementConfigurationOutputWithContext(context.Background())
}

func (i *ManagementConfiguration) ToManagementConfigurationOutputWithContext(ctx context.Context) ManagementConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementConfigurationOutput)
}

type ManagementConfigurationOutput struct{ *pulumi.OutputState }

func (ManagementConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementConfiguration)(nil))
}

func (o ManagementConfigurationOutput) ToManagementConfigurationOutput() ManagementConfigurationOutput {
	return o
}

func (o ManagementConfigurationOutput) ToManagementConfigurationOutputWithContext(ctx context.Context) ManagementConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagementConfigurationOutput{})
}
