


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationAssignmentParent struct {
	pulumi.CustomResourceState

	Location                   pulumi.StringPtrOutput   `pulumi:"location"`
	MaintenanceConfigurationId pulumi.StringPtrOutput   `pulumi:"maintenanceConfigurationId"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	ResourceId                 pulumi.StringPtrOutput   `pulumi:"resourceId"`
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
}


func NewConfigurationAssignmentParent(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentParentArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentParent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceParentName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceParentName'")
	}
	if args.ResourceParentType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceParentType'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance:ConfigurationAssignmentParent"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210901preview:ConfigurationAssignmentParent"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20220701preview:ConfigurationAssignmentParent"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20221101preview:ConfigurationAssignmentParent"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationAssignmentParent
	err := ctx.RegisterResource("azure-native:maintenance/v20210401preview:ConfigurationAssignmentParent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationAssignmentParent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentParentState, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentParent, error) {
	var resource ConfigurationAssignmentParent
	err := ctx.ReadResource("azure-native:maintenance/v20210401preview:ConfigurationAssignmentParent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationAssignmentParentState struct {
}

type ConfigurationAssignmentParentState struct {
}

func (ConfigurationAssignmentParentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentParentState)(nil)).Elem()
}

type configurationAssignmentParentArgs struct {
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	Location                    *string `pulumi:"location"`
	MaintenanceConfigurationId  *string `pulumi:"maintenanceConfigurationId"`
	ProviderName                string  `pulumi:"providerName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	ResourceId                  *string `pulumi:"resourceId"`
	ResourceName                string  `pulumi:"resourceName"`
	ResourceParentName          string  `pulumi:"resourceParentName"`
	ResourceParentType          string  `pulumi:"resourceParentType"`
	ResourceType                string  `pulumi:"resourceType"`
}


type ConfigurationAssignmentParentArgs struct {
	ConfigurationAssignmentName pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	MaintenanceConfigurationId  pulumi.StringPtrInput
	ProviderName                pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	ResourceId                  pulumi.StringPtrInput
	ResourceName                pulumi.StringInput
	ResourceParentName          pulumi.StringInput
	ResourceParentType          pulumi.StringInput
	ResourceType                pulumi.StringInput
}

func (ConfigurationAssignmentParentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentParentArgs)(nil)).Elem()
}

type ConfigurationAssignmentParentInput interface {
	pulumi.Input

	ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput
	ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput
}

func (*ConfigurationAssignmentParent) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentParent)(nil)).Elem()
}

func (i *ConfigurationAssignmentParent) ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput {
	return i.ToConfigurationAssignmentParentOutputWithContext(context.Background())
}

func (i *ConfigurationAssignmentParent) ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentParentOutput)
}

type ConfigurationAssignmentParentOutput struct{ *pulumi.OutputState }

func (ConfigurationAssignmentParentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentParent)(nil)).Elem()
}

func (o ConfigurationAssignmentParentOutput) ToConfigurationAssignmentParentOutput() ConfigurationAssignmentParentOutput {
	return o
}

func (o ConfigurationAssignmentParentOutput) ToConfigurationAssignmentParentOutputWithContext(ctx context.Context) ConfigurationAssignmentParentOutput {
	return o
}

func (o ConfigurationAssignmentParentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentParentOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentParentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationAssignmentParentOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentParentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationAssignmentParentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentParent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentParentOutput{})
}
