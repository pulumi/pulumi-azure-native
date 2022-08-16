


package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfileHCRPAssignment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                    `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                               `pulumi:"systemData"`
	Type       pulumi.StringOutput                                    `pulumi:"type"`
}


func NewConfigurationProfileHCRPAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileHCRPAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCRPAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfileHCRPAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfileHCRPAssignment
	err := ctx.RegisterResource("azure-native:automanage/v20210430preview:ConfigurationProfileHCRPAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfileHCRPAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileHCRPAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCRPAssignment, error) {
	var resource ConfigurationProfileHCRPAssignment
	err := ctx.ReadResource("azure-native:automanage/v20210430preview:ConfigurationProfileHCRPAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfileHCRPAssignmentState struct {
}

type ConfigurationProfileHCRPAssignmentState struct {
}

func (ConfigurationProfileHCRPAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCRPAssignmentState)(nil)).Elem()
}

type configurationProfileHCRPAssignmentArgs struct {
	ConfigurationProfileAssignmentName *string                                   `pulumi:"configurationProfileAssignmentName"`
	MachineName                        string                                    `pulumi:"machineName"`
	Properties                         *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	ResourceGroupName                  string                                    `pulumi:"resourceGroupName"`
}


type ConfigurationProfileHCRPAssignmentArgs struct {
	ConfigurationProfileAssignmentName pulumi.StringPtrInput
	MachineName                        pulumi.StringInput
	Properties                         ConfigurationProfileAssignmentPropertiesPtrInput
	ResourceGroupName                  pulumi.StringInput
}

func (ConfigurationProfileHCRPAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCRPAssignmentArgs)(nil)).Elem()
}

type ConfigurationProfileHCRPAssignmentInput interface {
	pulumi.Input

	ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput
	ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput
}

func (*ConfigurationProfileHCRPAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCRPAssignment)(nil)).Elem()
}

func (i *ConfigurationProfileHCRPAssignment) ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput {
	return i.ToConfigurationProfileHCRPAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationProfileHCRPAssignment) ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileHCRPAssignmentOutput)
}

type ConfigurationProfileHCRPAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileHCRPAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCRPAssignment)(nil)).Elem()
}

func (o ConfigurationProfileHCRPAssignmentOutput) ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCRPAssignmentOutput) ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCRPAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationProfileHCRPAssignmentOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) ConfigurationProfileAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o ConfigurationProfileHCRPAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationProfileHCRPAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileHCRPAssignmentOutput{})
}
