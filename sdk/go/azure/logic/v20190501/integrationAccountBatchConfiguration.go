


package v20190501

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
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountBatchConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountBatchConfiguration
	err := ctx.RegisterResource("azure-native:logic/v20190501:IntegrationAccountBatchConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountBatchConfigurationState, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	var resource IntegrationAccountBatchConfiguration
	err := ctx.ReadResource("azure-native:logic/v20190501:IntegrationAccountBatchConfiguration", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*IntegrationAccountBatchConfiguration)(nil))
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return i.ToIntegrationAccountBatchConfigurationOutputWithContext(context.Background())
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountBatchConfigurationOutput)
}

type IntegrationAccountBatchConfigurationOutput struct{ *pulumi.OutputState }

func (IntegrationAccountBatchConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountBatchConfiguration)(nil))
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return o
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountBatchConfigurationOutput{})
}
