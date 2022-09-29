


package automanage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfileAssignment struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                    `pulumi:"name"`
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                    `pulumi:"type"`
}


func NewConfigurationProfileAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20200630preview:ConfigurationProfileAssignment"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfileAssignment"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfileAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfileAssignment
	err := ctx.RegisterResource("azure-native:automanage:ConfigurationProfileAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationProfileAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileAssignment, error) {
	var resource ConfigurationProfileAssignment
	err := ctx.ReadResource("azure-native:automanage:ConfigurationProfileAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationProfileAssignmentState struct {
}

type ConfigurationProfileAssignmentState struct {
}

func (ConfigurationProfileAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileAssignmentState)(nil)).Elem()
}

type configurationProfileAssignmentArgs struct {
	ConfigurationProfileAssignmentName *string                                   `pulumi:"configurationProfileAssignmentName"`
	Properties                         *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	ResourceGroupName                  string                                    `pulumi:"resourceGroupName"`
	VmName                             string                                    `pulumi:"vmName"`
}


type ConfigurationProfileAssignmentArgs struct {
	ConfigurationProfileAssignmentName pulumi.StringPtrInput
	Properties                         ConfigurationProfileAssignmentPropertiesPtrInput
	ResourceGroupName                  pulumi.StringInput
	VmName                             pulumi.StringInput
}

func (ConfigurationProfileAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileAssignmentArgs)(nil)).Elem()
}

type ConfigurationProfileAssignmentInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentOutput() ConfigurationProfileAssignmentOutput
	ToConfigurationProfileAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentOutput
}

func (*ConfigurationProfileAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignment)(nil)).Elem()
}

func (i *ConfigurationProfileAssignment) ToConfigurationProfileAssignmentOutput() ConfigurationProfileAssignmentOutput {
	return i.ToConfigurationProfileAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationProfileAssignment) ToConfigurationProfileAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentOutput)
}

type ConfigurationProfileAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignment)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentOutput) ToConfigurationProfileAssignmentOutput() ConfigurationProfileAssignmentOutput {
	return o
}

func (o ConfigurationProfileAssignmentOutput) ToConfigurationProfileAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentOutput {
	return o
}

func (o ConfigurationProfileAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationProfileAssignmentOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignment) ConfigurationProfileAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

func (o ConfigurationProfileAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentOutput{})
}
