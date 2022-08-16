


package v20161101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:datalakestore/v20161101:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	AccountId                   string                       `pulumi:"accountId"`
	CreationTime                string                       `pulumi:"creationTime"`
	CurrentTier                 string                       `pulumi:"currentTier"`
	DefaultGroup                string                       `pulumi:"defaultGroup"`
	EncryptionConfig            EncryptionConfigResponse     `pulumi:"encryptionConfig"`
	EncryptionProvisioningState string                       `pulumi:"encryptionProvisioningState"`
	EncryptionState             string                       `pulumi:"encryptionState"`
	Endpoint                    string                       `pulumi:"endpoint"`
	FirewallAllowAzureIps       string                       `pulumi:"firewallAllowAzureIps"`
	FirewallRules               []FirewallRuleResponse       `pulumi:"firewallRules"`
	FirewallState               string                       `pulumi:"firewallState"`
	Id                          string                       `pulumi:"id"`
	Identity                    EncryptionIdentityResponse   `pulumi:"identity"`
	LastModifiedTime            string                       `pulumi:"lastModifiedTime"`
	Location                    string                       `pulumi:"location"`
	Name                        string                       `pulumi:"name"`
	NewTier                     string                       `pulumi:"newTier"`
	ProvisioningState           string                       `pulumi:"provisioningState"`
	State                       string                       `pulumi:"state"`
	Tags                        map[string]string            `pulumi:"tags"`
	TrustedIdProviderState      string                       `pulumi:"trustedIdProviderState"`
	TrustedIdProviders          []TrustedIdProviderResponse  `pulumi:"trustedIdProviders"`
	Type                        string                       `pulumi:"type"`
	VirtualNetworkRules         []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}


type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) CurrentTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.CurrentTier }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) DefaultGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.DefaultGroup }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) EncryptionConfig() EncryptionConfigResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) EncryptionConfigResponse { return v.EncryptionConfig }).(EncryptionConfigResponseOutput)
}

func (o LookupAccountResultOutput) EncryptionProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.EncryptionProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) EncryptionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.EncryptionState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) FirewallAllowAzureIps() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.FirewallAllowAzureIps }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []FirewallRuleResponse { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o LookupAccountResultOutput) FirewallState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.FirewallState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Identity() EncryptionIdentityResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) EncryptionIdentityResponse { return v.Identity }).(EncryptionIdentityResponseOutput)
}

func (o LookupAccountResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) NewTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.NewTier }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccountResultOutput) TrustedIdProviderState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.TrustedIdProviderState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) TrustedIdProviders() TrustedIdProviderResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []TrustedIdProviderResponse { return v.TrustedIdProviders }).(TrustedIdProviderResponseArrayOutput)
}

func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupAccountResult) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
