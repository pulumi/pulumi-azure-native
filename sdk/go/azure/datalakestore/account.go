


package datalakestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Account struct {
	pulumi.CustomResourceState

	AccountId                   pulumi.StringOutput                   `pulumi:"accountId"`
	CreationTime                pulumi.StringOutput                   `pulumi:"creationTime"`
	CurrentTier                 pulumi.StringOutput                   `pulumi:"currentTier"`
	DefaultGroup                pulumi.StringOutput                   `pulumi:"defaultGroup"`
	EncryptionConfig            EncryptionConfigResponseOutput        `pulumi:"encryptionConfig"`
	EncryptionProvisioningState pulumi.StringOutput                   `pulumi:"encryptionProvisioningState"`
	EncryptionState             pulumi.StringOutput                   `pulumi:"encryptionState"`
	Endpoint                    pulumi.StringOutput                   `pulumi:"endpoint"`
	FirewallAllowAzureIps       pulumi.StringOutput                   `pulumi:"firewallAllowAzureIps"`
	FirewallRules               FirewallRuleResponseArrayOutput       `pulumi:"firewallRules"`
	FirewallState               pulumi.StringOutput                   `pulumi:"firewallState"`
	Identity                    EncryptionIdentityResponseOutput      `pulumi:"identity"`
	LastModifiedTime            pulumi.StringOutput                   `pulumi:"lastModifiedTime"`
	Location                    pulumi.StringOutput                   `pulumi:"location"`
	Name                        pulumi.StringOutput                   `pulumi:"name"`
	NewTier                     pulumi.StringOutput                   `pulumi:"newTier"`
	ProvisioningState           pulumi.StringOutput                   `pulumi:"provisioningState"`
	State                       pulumi.StringOutput                   `pulumi:"state"`
	Tags                        pulumi.StringMapOutput                `pulumi:"tags"`
	TrustedIdProviderState      pulumi.StringOutput                   `pulumi:"trustedIdProviderState"`
	TrustedIdProviders          TrustedIdProviderResponseArrayOutput  `pulumi:"trustedIdProviders"`
	Type                        pulumi.StringOutput                   `pulumi:"type"`
	VirtualNetworkRules         VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
}


func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakestore/v20161101:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-native:datalakestore:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:datalakestore:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	AccountName            *string                                         `pulumi:"accountName"`
	DefaultGroup           *string                                         `pulumi:"defaultGroup"`
	EncryptionConfig       *EncryptionConfig                               `pulumi:"encryptionConfig"`
	EncryptionState        *EncryptionState                                `pulumi:"encryptionState"`
	FirewallAllowAzureIps  *FirewallAllowAzureIpsState                     `pulumi:"firewallAllowAzureIps"`
	FirewallRules          []CreateFirewallRuleWithAccountParameters       `pulumi:"firewallRules"`
	FirewallState          *FirewallState                                  `pulumi:"firewallState"`
	Identity               *EncryptionIdentity                             `pulumi:"identity"`
	Location               *string                                         `pulumi:"location"`
	NewTier                *TierType                                       `pulumi:"newTier"`
	ResourceGroupName      string                                          `pulumi:"resourceGroupName"`
	Tags                   map[string]string                               `pulumi:"tags"`
	TrustedIdProviderState *TrustedIdProviderStateEnum                     `pulumi:"trustedIdProviderState"`
	TrustedIdProviders     []CreateTrustedIdProviderWithAccountParameters  `pulumi:"trustedIdProviders"`
	VirtualNetworkRules    []CreateVirtualNetworkRuleWithAccountParameters `pulumi:"virtualNetworkRules"`
}


type AccountArgs struct {
	AccountName            pulumi.StringPtrInput
	DefaultGroup           pulumi.StringPtrInput
	EncryptionConfig       EncryptionConfigPtrInput
	EncryptionState        EncryptionStatePtrInput
	FirewallAllowAzureIps  FirewallAllowAzureIpsStatePtrInput
	FirewallRules          CreateFirewallRuleWithAccountParametersArrayInput
	FirewallState          FirewallStatePtrInput
	Identity               EncryptionIdentityPtrInput
	Location               pulumi.StringPtrInput
	NewTier                TierTypePtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	TrustedIdProviderState TrustedIdProviderStateEnumPtrInput
	TrustedIdProviders     CreateTrustedIdProviderWithAccountParametersArrayInput
	VirtualNetworkRules    CreateVirtualNetworkRuleWithAccountParametersArrayInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o AccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o AccountOutput) CurrentTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CurrentTier }).(pulumi.StringOutput)
}

func (o AccountOutput) DefaultGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.DefaultGroup }).(pulumi.StringOutput)
}

func (o AccountOutput) EncryptionConfig() EncryptionConfigResponseOutput {
	return o.ApplyT(func(v *Account) EncryptionConfigResponseOutput { return v.EncryptionConfig }).(EncryptionConfigResponseOutput)
}

func (o AccountOutput) EncryptionProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.EncryptionProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) EncryptionState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.EncryptionState }).(pulumi.StringOutput)
}

func (o AccountOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o AccountOutput) FirewallAllowAzureIps() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.FirewallAllowAzureIps }).(pulumi.StringOutput)
}

func (o AccountOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) FirewallRuleResponseArrayOutput { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o AccountOutput) FirewallState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.FirewallState }).(pulumi.StringOutput)
}

func (o AccountOutput) Identity() EncryptionIdentityResponseOutput {
	return o.ApplyT(func(v *Account) EncryptionIdentityResponseOutput { return v.Identity }).(EncryptionIdentityResponseOutput)
}

func (o AccountOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o AccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AccountOutput) NewTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.NewTier }).(pulumi.StringOutput)
}

func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AccountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AccountOutput) TrustedIdProviderState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.TrustedIdProviderState }).(pulumi.StringOutput)
}

func (o AccountOutput) TrustedIdProviders() TrustedIdProviderResponseArrayOutput {
	return o.ApplyT(func(v *Account) TrustedIdProviderResponseArrayOutput { return v.TrustedIdProviders }).(TrustedIdProviderResponseArrayOutput)
}

func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AccountOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *Account) VirtualNetworkRuleResponseArrayOutput { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
