


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountBatchConfiguration struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties BatchConfigurationPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, args *IntegrationAccountBatchConfigurationArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountBatchConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountBatchConfiguration
	err := ctx.RegisterResource("azure-native:logic/v20160601:IntegrationAccountBatchConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountBatchConfigurationState, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	var resource IntegrationAccountBatchConfiguration
	err := ctx.ReadResource("azure-native:logic/v20160601:IntegrationAccountBatchConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountBatchConfigurationState struct {
}

type IntegrationAccountBatchConfigurationState struct {
}

func (IntegrationAccountBatchConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationState)(nil)).Elem()
}

type integrationAccountBatchConfigurationArgs struct {
	BatchConfigurationName *string                      `pulumi:"batchConfigurationName"`
	IntegrationAccountName string                       `pulumi:"integrationAccountName"`
	Location               *string                      `pulumi:"location"`
	Properties             BatchConfigurationProperties `pulumi:"properties"`
	ResourceGroupName      string                       `pulumi:"resourceGroupName"`
	Tags                   map[string]string            `pulumi:"tags"`
}


type IntegrationAccountBatchConfigurationArgs struct {
	BatchConfigurationName pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	Properties             BatchConfigurationPropertiesInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (IntegrationAccountBatchConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationArgs)(nil)).Elem()
}

type IntegrationAccountBatchConfigurationInput interface {
	pulumi.Input

	ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput
	ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput
}

func (*IntegrationAccountBatchConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountBatchConfiguration)(nil)).Elem()
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return i.ToIntegrationAccountBatchConfigurationOutputWithContext(context.Background())
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountBatchConfigurationOutput)
}

type IntegrationAccountBatchConfigurationOutput struct{ *pulumi.OutputState }

func (IntegrationAccountBatchConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountBatchConfiguration)(nil)).Elem()
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return o
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return o
}

func (o IntegrationAccountBatchConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountBatchConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationAccountBatchConfigurationOutput) Properties() BatchConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) BatchConfigurationPropertiesResponseOutput {
		return v.Properties
	}).(BatchConfigurationPropertiesResponseOutput)
}

func (o IntegrationAccountBatchConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationAccountBatchConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountBatchConfigurationOutput{})
}
