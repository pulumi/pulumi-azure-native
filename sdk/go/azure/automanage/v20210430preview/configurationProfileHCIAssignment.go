


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfileHCIAssignment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                    `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                               `pulumi:"systemData"`
	Type       pulumi.StringOutput                                    `pulumi:"type"`
}


func NewConfigurationProfileHCIAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileHCIAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCIAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfileHCIAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfileHCIAssignment
	err := ctx.RegisterResource("azure-native:automanage/v20210430preview:ConfigurationProfileHCIAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfileHCIAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileHCIAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCIAssignment, error) {
	var resource ConfigurationProfileHCIAssignment
	err := ctx.ReadResource("azure-native:automanage/v20210430preview:ConfigurationProfileHCIAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfileHCIAssignmentState struct {
}

type ConfigurationProfileHCIAssignmentState struct {
}

func (ConfigurationProfileHCIAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCIAssignmentState)(nil)).Elem()
}

type configurationProfileHCIAssignmentArgs struct {
	ClusterName                        string                                    `pulumi:"clusterName"`
	ConfigurationProfileAssignmentName *string                                   `pulumi:"configurationProfileAssignmentName"`
	Properties                         *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	ResourceGroupName                  string                                    `pulumi:"resourceGroupName"`
}


type ConfigurationProfileHCIAssignmentArgs struct {
	ClusterName                        pulumi.StringInput
	ConfigurationProfileAssignmentName pulumi.StringPtrInput
	Properties                         ConfigurationProfileAssignmentPropertiesPtrInput
	ResourceGroupName                  pulumi.StringInput
}

func (ConfigurationProfileHCIAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCIAssignmentArgs)(nil)).Elem()
}

type ConfigurationProfileHCIAssignmentInput interface {
	pulumi.Input

	ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput
	ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput
}

func (*ConfigurationProfileHCIAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCIAssignment)(nil)).Elem()
}

func (i *ConfigurationProfileHCIAssignment) ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput {
	return i.ToConfigurationProfileHCIAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationProfileHCIAssignment) ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileHCIAssignmentOutput)
}

type ConfigurationProfileHCIAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileHCIAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCIAssignment)(nil)).Elem()
}

func (o ConfigurationProfileHCIAssignmentOutput) ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCIAssignmentOutput) ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCIAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationProfileHCIAssignmentOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) ConfigurationProfileAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o ConfigurationProfileHCIAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationProfileHCIAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileHCIAssignmentOutput{})
}
