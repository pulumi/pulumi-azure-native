


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NspAccessRule struct {
	pulumi.CustomResourceState

	AddressPrefixes           pulumi.StringArrayOutput                    `pulumi:"addressPrefixes"`
	Direction                 pulumi.StringPtrOutput                      `pulumi:"direction"`
	FullyQualifiedDomainNames pulumi.StringArrayOutput                    `pulumi:"fullyQualifiedDomainNames"`
	Location                  pulumi.StringPtrOutput                      `pulumi:"location"`
	Name                      pulumi.StringOutput                         `pulumi:"name"`
	NetworkSecurityPerimeters PerimeterBasedAccessRuleResponseArrayOutput `pulumi:"networkSecurityPerimeters"`
	ProvisioningState         pulumi.StringOutput                         `pulumi:"provisioningState"`
	Subscriptions             SubscriptionIdResponseArrayOutput           `pulumi:"subscriptions"`
	Tags                      pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                      pulumi.StringOutput                         `pulumi:"type"`
}


func NewNspAccessRule(ctx *pulumi.Context,
	name string, args *NspAccessRuleArgs, opts ...pulumi.ResourceOption) (*NspAccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NspAccessRule"),
		},
	})
	opts = append(opts, aliases)
	var resource NspAccessRule
	err := ctx.RegisterResource("azure-native:network:NspAccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNspAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspAccessRuleState, opts ...pulumi.ResourceOption) (*NspAccessRule, error) {
	var resource NspAccessRule
	err := ctx.ReadResource("azure-native:network:NspAccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nspAccessRuleState struct {
}

type NspAccessRuleState struct {
}

func (NspAccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAccessRuleState)(nil)).Elem()
}

type nspAccessRuleArgs struct {
	AccessRuleName               *string                    `pulumi:"accessRuleName"`
	AddressPrefixes              []string                   `pulumi:"addressPrefixes"`
	Direction                    *string                    `pulumi:"direction"`
	FullyQualifiedDomainNames    []string                   `pulumi:"fullyQualifiedDomainNames"`
	Id                           *string                    `pulumi:"id"`
	Location                     *string                    `pulumi:"location"`
	Name                         *string                    `pulumi:"name"`
	NetworkSecurityPerimeterName string                     `pulumi:"networkSecurityPerimeterName"`
	NetworkSecurityPerimeters    []PerimeterBasedAccessRule `pulumi:"networkSecurityPerimeters"`
	ProfileName                  string                     `pulumi:"profileName"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	Subscriptions                []SubscriptionId           `pulumi:"subscriptions"`
	Tags                         map[string]string          `pulumi:"tags"`
}


type NspAccessRuleArgs struct {
	AccessRuleName               pulumi.StringPtrInput
	AddressPrefixes              pulumi.StringArrayInput
	Direction                    pulumi.StringPtrInput
	FullyQualifiedDomainNames    pulumi.StringArrayInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringInput
	NetworkSecurityPerimeters    PerimeterBasedAccessRuleArrayInput
	ProfileName                  pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	Subscriptions                SubscriptionIdArrayInput
	Tags                         pulumi.StringMapInput
}

func (NspAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAccessRuleArgs)(nil)).Elem()
}

type NspAccessRuleInput interface {
	pulumi.Input

	ToNspAccessRuleOutput() NspAccessRuleOutput
	ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput
}

func (*NspAccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAccessRule)(nil)).Elem()
}

func (i *NspAccessRule) ToNspAccessRuleOutput() NspAccessRuleOutput {
	return i.ToNspAccessRuleOutputWithContext(context.Background())
}

func (i *NspAccessRule) ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspAccessRuleOutput)
}

type NspAccessRuleOutput struct{ *pulumi.OutputState }

func (NspAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAccessRule)(nil)).Elem()
}

func (o NspAccessRuleOutput) ToNspAccessRuleOutput() NspAccessRuleOutput {
	return o
}

func (o NspAccessRuleOutput) ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput {
	return o
}

func (o NspAccessRuleOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NspAccessRuleOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringPtrOutput { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o NspAccessRuleOutput) FullyQualifiedDomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.FullyQualifiedDomainNames }).(pulumi.StringArrayOutput)
}

func (o NspAccessRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NspAccessRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NspAccessRuleOutput) NetworkSecurityPerimeters() PerimeterBasedAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) PerimeterBasedAccessRuleResponseArrayOutput { return v.NetworkSecurityPerimeters }).(PerimeterBasedAccessRuleResponseArrayOutput)
}

func (o NspAccessRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NspAccessRuleOutput) Subscriptions() SubscriptionIdResponseArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) SubscriptionIdResponseArrayOutput { return v.Subscriptions }).(SubscriptionIdResponseArrayOutput)
}

func (o NspAccessRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NspAccessRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspAccessRuleOutput{})
}
