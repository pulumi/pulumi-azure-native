


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationPolicyGroup struct {
	pulumi.CustomResourceState

	Etag                        pulumi.StringOutput                                        `pulumi:"etag"`
	IsDefault                   pulumi.BoolPtrOutput                                       `pulumi:"isDefault"`
	Name                        pulumi.StringPtrOutput                                     `pulumi:"name"`
	P2SConnectionConfigurations SubResourceResponseArrayOutput                             `pulumi:"p2SConnectionConfigurations"`
	PolicyMembers               VpnServerConfigurationPolicyGroupMemberResponseArrayOutput `pulumi:"policyMembers"`
	Priority                    pulumi.IntPtrOutput                                        `pulumi:"priority"`
	ProvisioningState           pulumi.StringOutput                                        `pulumi:"provisioningState"`
	Type                        pulumi.StringOutput                                        `pulumi:"type"`
}


func NewConfigurationPolicyGroup(ctx *pulumi.Context,
	name string, args *ConfigurationPolicyGroupArgs, opts ...pulumi.ResourceOption) (*ConfigurationPolicyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VpnServerConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'VpnServerConfigurationName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ConfigurationPolicyGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ConfigurationPolicyGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationPolicyGroup
	err := ctx.RegisterResource("azure-native:network/v20220501:ConfigurationPolicyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationPolicyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationPolicyGroupState, opts ...pulumi.ResourceOption) (*ConfigurationPolicyGroup, error) {
	var resource ConfigurationPolicyGroup
	err := ctx.ReadResource("azure-native:network/v20220501:ConfigurationPolicyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationPolicyGroupState struct {
}

type ConfigurationPolicyGroupState struct {
}

func (ConfigurationPolicyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyGroupState)(nil)).Elem()
}

type configurationPolicyGroupArgs struct {
	ConfigurationPolicyGroupName *string                                   `pulumi:"configurationPolicyGroupName"`
	Id                           *string                                   `pulumi:"id"`
	IsDefault                    *bool                                     `pulumi:"isDefault"`
	Name                         *string                                   `pulumi:"name"`
	PolicyMembers                []VpnServerConfigurationPolicyGroupMember `pulumi:"policyMembers"`
	Priority                     *int                                      `pulumi:"priority"`
	ResourceGroupName            string                                    `pulumi:"resourceGroupName"`
	VpnServerConfigurationName   string                                    `pulumi:"vpnServerConfigurationName"`
}


type ConfigurationPolicyGroupArgs struct {
	ConfigurationPolicyGroupName pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	IsDefault                    pulumi.BoolPtrInput
	Name                         pulumi.StringPtrInput
	PolicyMembers                VpnServerConfigurationPolicyGroupMemberArrayInput
	Priority                     pulumi.IntPtrInput
	ResourceGroupName            pulumi.StringInput
	VpnServerConfigurationName   pulumi.StringInput
}

func (ConfigurationPolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationPolicyGroupArgs)(nil)).Elem()
}

type ConfigurationPolicyGroupInput interface {
	pulumi.Input

	ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput
	ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput
}

func (*ConfigurationPolicyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicyGroup)(nil)).Elem()
}

func (i *ConfigurationPolicyGroup) ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput {
	return i.ToConfigurationPolicyGroupOutputWithContext(context.Background())
}

func (i *ConfigurationPolicyGroup) ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPolicyGroupOutput)
}

type ConfigurationPolicyGroupOutput struct{ *pulumi.OutputState }

func (ConfigurationPolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationPolicyGroup)(nil)).Elem()
}

func (o ConfigurationPolicyGroupOutput) ToConfigurationPolicyGroupOutput() ConfigurationPolicyGroupOutput {
	return o
}

func (o ConfigurationPolicyGroupOutput) ToConfigurationPolicyGroupOutputWithContext(ctx context.Context) ConfigurationPolicyGroupOutput {
	return o
}

func (o ConfigurationPolicyGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ConfigurationPolicyGroupOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationPolicyGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationPolicyGroupOutput) P2SConnectionConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) SubResourceResponseArrayOutput { return v.P2SConnectionConfigurations }).(SubResourceResponseArrayOutput)
}

func (o ConfigurationPolicyGroupOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput)
}

func (o ConfigurationPolicyGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o ConfigurationPolicyGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConfigurationPolicyGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationPolicyGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationPolicyGroupOutput{})
}
