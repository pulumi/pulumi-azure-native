


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceNetworkRuleSet struct {
	pulumi.CustomResourceState

	DefaultAction               pulumi.StringPtrOutput                          `pulumi:"defaultAction"`
	IpRules                     NWRuleSetIpRulesResponseArrayOutput             `pulumi:"ipRules"`
	Location                    pulumi.StringOutput                             `pulumi:"location"`
	Name                        pulumi.StringOutput                             `pulumi:"name"`
	PublicNetworkAccess         pulumi.StringPtrOutput                          `pulumi:"publicNetworkAccess"`
	SystemData                  SystemDataResponseOutput                        `pulumi:"systemData"`
	TrustedServiceAccessEnabled pulumi.BoolPtrOutput                            `pulumi:"trustedServiceAccessEnabled"`
	Type                        pulumi.StringOutput                             `pulumi:"type"`
	VirtualNetworkRules         NWRuleSetVirtualNetworkRulesResponseArrayOutput `pulumi:"virtualNetworkRules"`
}


func NewNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, args *NamespaceNetworkRuleSetArgs, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:NamespaceNetworkRuleSet"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:NamespaceNetworkRuleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource NamespaceNetworkRuleSet
	err := ctx.RegisterResource("azure-native:eventhub/v20220101preview:NamespaceNetworkRuleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamespaceNetworkRuleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceNetworkRuleSetState, opts ...pulumi.ResourceOption) (*NamespaceNetworkRuleSet, error) {
	var resource NamespaceNetworkRuleSet
	err := ctx.ReadResource("azure-native:eventhub/v20220101preview:NamespaceNetworkRuleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namespaceNetworkRuleSetState struct {
}

type NamespaceNetworkRuleSetState struct {
}

func (NamespaceNetworkRuleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetState)(nil)).Elem()
}

type namespaceNetworkRuleSetArgs struct {
	DefaultAction               *string                        `pulumi:"defaultAction"`
	IpRules                     []NWRuleSetIpRules             `pulumi:"ipRules"`
	NamespaceName               string                         `pulumi:"namespaceName"`
	PublicNetworkAccess         *string                        `pulumi:"publicNetworkAccess"`
	ResourceGroupName           string                         `pulumi:"resourceGroupName"`
	TrustedServiceAccessEnabled *bool                          `pulumi:"trustedServiceAccessEnabled"`
	VirtualNetworkRules         []NWRuleSetVirtualNetworkRules `pulumi:"virtualNetworkRules"`
}


type NamespaceNetworkRuleSetArgs struct {
	DefaultAction               pulumi.StringPtrInput
	IpRules                     NWRuleSetIpRulesArrayInput
	NamespaceName               pulumi.StringInput
	PublicNetworkAccess         pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	TrustedServiceAccessEnabled pulumi.BoolPtrInput
	VirtualNetworkRules         NWRuleSetVirtualNetworkRulesArrayInput
}

func (NamespaceNetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceNetworkRuleSetArgs)(nil)).Elem()
}

type NamespaceNetworkRuleSetInput interface {
	pulumi.Input

	ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput
	ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput
}

func (*NamespaceNetworkRuleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceNetworkRuleSet)(nil)).Elem()
}

func (i *NamespaceNetworkRuleSet) ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput {
	return i.ToNamespaceNetworkRuleSetOutputWithContext(context.Background())
}

func (i *NamespaceNetworkRuleSet) ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceNetworkRuleSetOutput)
}

type NamespaceNetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NamespaceNetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceNetworkRuleSet)(nil)).Elem()
}

func (o NamespaceNetworkRuleSetOutput) ToNamespaceNetworkRuleSetOutput() NamespaceNetworkRuleSetOutput {
	return o
}

func (o NamespaceNetworkRuleSetOutput) ToNamespaceNetworkRuleSetOutputWithContext(ctx context.Context) NamespaceNetworkRuleSetOutput {
	return o
}

func (o NamespaceNetworkRuleSetOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.StringPtrOutput { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NamespaceNetworkRuleSetOutput) IpRules() NWRuleSetIpRulesResponseArrayOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) NWRuleSetIpRulesResponseArrayOutput { return v.IpRules }).(NWRuleSetIpRulesResponseArrayOutput)
}

func (o NamespaceNetworkRuleSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NamespaceNetworkRuleSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceNetworkRuleSetOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o NamespaceNetworkRuleSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NamespaceNetworkRuleSetOutput) TrustedServiceAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.BoolPtrOutput { return v.TrustedServiceAccessEnabled }).(pulumi.BoolPtrOutput)
}

func (o NamespaceNetworkRuleSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NamespaceNetworkRuleSetOutput) VirtualNetworkRules() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o.ApplyT(func(v *NamespaceNetworkRuleSet) NWRuleSetVirtualNetworkRulesResponseArrayOutput {
		return v.VirtualNetworkRules
	}).(NWRuleSetVirtualNetworkRulesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceNetworkRuleSetOutput{})
}
