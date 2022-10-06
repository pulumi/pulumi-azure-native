


package v20160901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SecurityRule struct {
	pulumi.CustomResourceState

	Access                   pulumi.StringOutput    `pulumi:"access"`
	Description              pulumi.StringPtrOutput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringOutput    `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringOutput    `pulumi:"direction"`
	Etag                     pulumi.StringPtrOutput `pulumi:"etag"`
	Name                     pulumi.StringPtrOutput `pulumi:"name"`
	Priority                 pulumi.IntPtrOutput    `pulumi:"priority"`
	Protocol                 pulumi.StringOutput    `pulumi:"protocol"`
	ProvisioningState        pulumi.StringPtrOutput `pulumi:"provisioningState"`
	SourceAddressPrefix      pulumi.StringOutput    `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
}


func NewSecurityRule(ctx *pulumi.Context,
	name string, args *SecurityRuleArgs, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Access == nil {
		return nil, errors.New("invalid value for required argument 'Access'")
	}
	if args.DestinationAddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAddressPrefix'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.NetworkSecurityGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceAddressPrefix == nil {
		return nil, errors.New("invalid value for required argument 'SourceAddressPrefix'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SecurityRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityRule
	err := ctx.RegisterResource("azure-native:network/v20160901:SecurityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityRuleState, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	var resource SecurityRule
	err := ctx.ReadResource("azure-native:network/v20160901:SecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityRuleState struct {
}

type SecurityRuleState struct {
}

func (SecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleState)(nil)).Elem()
}

type securityRuleArgs struct {
	Access                   string  `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix string  `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                string  `pulumi:"direction"`
	Id                       *string `pulumi:"id"`
	Name                     *string `pulumi:"name"`
	NetworkSecurityGroupName string  `pulumi:"networkSecurityGroupName"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 string  `pulumi:"protocol"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	SecurityRuleName         *string `pulumi:"securityRuleName"`
	SourceAddressPrefix      string  `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}


type SecurityRuleArgs struct {
	Access                   pulumi.StringInput
	Description              pulumi.StringPtrInput
	DestinationAddressPrefix pulumi.StringInput
	DestinationPortRange     pulumi.StringPtrInput
	Direction                pulumi.StringInput
	Id                       pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	NetworkSecurityGroupName pulumi.StringInput
	Priority                 pulumi.IntPtrInput
	Protocol                 pulumi.StringInput
	ProvisioningState        pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	SecurityRuleName         pulumi.StringPtrInput
	SourceAddressPrefix      pulumi.StringInput
	SourcePortRange          pulumi.StringPtrInput
}

func (SecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleArgs)(nil)).Elem()
}

type SecurityRuleInput interface {
	pulumi.Input

	ToSecurityRuleOutput() SecurityRuleOutput
	ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput
}

func (*SecurityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRule)(nil)).Elem()
}

func (i *SecurityRule) ToSecurityRuleOutput() SecurityRuleOutput {
	return i.ToSecurityRuleOutputWithContext(context.Background())
}

func (i *SecurityRule) ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleOutput)
}

type SecurityRuleOutput struct{ *pulumi.OutputState }

func (SecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRule)(nil)).Elem()
}

func (o SecurityRuleOutput) ToSecurityRuleOutput() SecurityRuleOutput {
	return o
}

func (o SecurityRuleOutput) ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput {
	return o
}

func (o SecurityRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityRuleOutput{})
}
