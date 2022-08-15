


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20210801preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryResult struct {
	AdminUserEnabled           *bool                               `pulumi:"adminUserEnabled"`
	AnonymousPullEnabled       *bool                               `pulumi:"anonymousPullEnabled"`
	CreationDate               string                              `pulumi:"creationDate"`
	DataEndpointEnabled        *bool                               `pulumi:"dataEndpointEnabled"`
	DataEndpointHostNames      []string                            `pulumi:"dataEndpointHostNames"`
	Encryption                 *EncryptionPropertyResponse         `pulumi:"encryption"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityPropertiesResponse         `pulumi:"identity"`
	Location                   string                              `pulumi:"location"`
	LoginServer                string                              `pulumi:"loginServer"`
	Name                       string                              `pulumi:"name"`
	NetworkRuleBypassOptions   *string                             `pulumi:"networkRuleBypassOptions"`
	NetworkRuleSet             *NetworkRuleSetResponse             `pulumi:"networkRuleSet"`
	Policies                   *PoliciesResponse                   `pulumi:"policies"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	Sku                        SkuResponse                         `pulumi:"sku"`
	Status                     StatusResponse                      `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	ZoneRedundancy             *string                             `pulumi:"zoneRedundancy"`
}


func (val *LookupRegistryResult) Defaults() *LookupRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdminUserEnabled) {
		adminUserEnabled_ := false
		tmp.AdminUserEnabled = &adminUserEnabled_
	}
	if isZero(tmp.AnonymousPullEnabled) {
		anonymousPullEnabled_ := false
		tmp.AnonymousPullEnabled = &anonymousPullEnabled_
	}
	if isZero(tmp.NetworkRuleBypassOptions) {
		networkRuleBypassOptions_ := "AzureServices"
		tmp.NetworkRuleBypassOptions = &networkRuleBypassOptions_
	}
	tmp.NetworkRuleSet = tmp.NetworkRuleSet.Defaults()

	tmp.Policies = tmp.Policies.Defaults()

	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if isZero(tmp.ZoneRedundancy) {
		zoneRedundancy_ := "Disabled"
		tmp.ZoneRedundancy = &zoneRedundancy_
	}
	return &tmp
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}


type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRegistryResultOutput) AnonymousPullEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AnonymousPullEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRegistryResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) DataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.DataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRegistryResultOutput) DataEndpointHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []string { return v.DataEndpointHostNames }).(pulumi.StringArrayOutput)
}

func (o LookupRegistryResultOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *EncryptionPropertyResponse { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.LoginServer }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) NetworkRuleBypassOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.NetworkRuleBypassOptions }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

func (o LookupRegistryResultOutput) Policies() PoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *PoliciesResponse { return v.Policies }).(PoliciesResponsePtrOutput)
}

func (o LookupRegistryResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupRegistryResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

func (o LookupRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
