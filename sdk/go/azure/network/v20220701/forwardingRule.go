


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ForwardingRule struct {
	pulumi.CustomResourceState

	DomainName          pulumi.StringOutput                `pulumi:"domainName"`
	Etag                pulumi.StringOutput                `pulumi:"etag"`
	ForwardingRuleState pulumi.StringPtrOutput             `pulumi:"forwardingRuleState"`
	Metadata            pulumi.StringMapOutput             `pulumi:"metadata"`
	Name                pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput           `pulumi:"systemData"`
	TargetDnsServers    TargetDnsServerResponseArrayOutput `pulumi:"targetDnsServers"`
	Type                pulumi.StringOutput                `pulumi:"type"`
}


func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsForwardingRulesetName == nil {
		return nil, errors.New("invalid value for required argument 'DnsForwardingRulesetName'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetDnsServers == nil {
		return nil, errors.New("invalid value for required argument 'TargetDnsServers'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ForwardingRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401preview:ForwardingRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ForwardingRule
	err := ctx.RegisterResource("azure-native:network/v20220701:ForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardingRuleState, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	var resource ForwardingRule
	err := ctx.ReadResource("azure-native:network/v20220701:ForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type forwardingRuleState struct {
}

type ForwardingRuleState struct {
}

func (ForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleState)(nil)).Elem()
}

type forwardingRuleArgs struct {
	DnsForwardingRulesetName string            `pulumi:"dnsForwardingRulesetName"`
	DomainName               string            `pulumi:"domainName"`
	ForwardingRuleName       *string           `pulumi:"forwardingRuleName"`
	ForwardingRuleState      *string           `pulumi:"forwardingRuleState"`
	Metadata                 map[string]string `pulumi:"metadata"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	TargetDnsServers         []TargetDnsServer `pulumi:"targetDnsServers"`
}


type ForwardingRuleArgs struct {
	DnsForwardingRulesetName pulumi.StringInput
	DomainName               pulumi.StringInput
	ForwardingRuleName       pulumi.StringPtrInput
	ForwardingRuleState      pulumi.StringPtrInput
	Metadata                 pulumi.StringMapInput
	ResourceGroupName        pulumi.StringInput
	TargetDnsServers         TargetDnsServerArrayInput
}

func (ForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleArgs)(nil)).Elem()
}

type ForwardingRuleInput interface {
	pulumi.Input

	ToForwardingRuleOutput() ForwardingRuleOutput
	ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput
}

func (*ForwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardingRule)(nil)).Elem()
}

func (i *ForwardingRule) ToForwardingRuleOutput() ForwardingRuleOutput {
	return i.ToForwardingRuleOutputWithContext(context.Background())
}

func (i *ForwardingRule) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRuleOutput)
}

type ForwardingRuleOutput struct{ *pulumi.OutputState }

func (ForwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardingRule)(nil)).Elem()
}

func (o ForwardingRuleOutput) ToForwardingRuleOutput() ForwardingRuleOutput {
	return o
}

func (o ForwardingRuleOutput) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return o
}

func (o ForwardingRuleOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o ForwardingRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ForwardingRuleOutput) ForwardingRuleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringPtrOutput { return v.ForwardingRuleState }).(pulumi.StringPtrOutput)
}

func (o ForwardingRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o ForwardingRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ForwardingRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ForwardingRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ForwardingRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ForwardingRuleOutput) TargetDnsServers() TargetDnsServerResponseArrayOutput {
	return o.ApplyT(func(v *ForwardingRule) TargetDnsServerResponseArrayOutput { return v.TargetDnsServers }).(TargetDnsServerResponseArrayOutput)
}

func (o ForwardingRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardingRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ForwardingRuleOutput{})
}
