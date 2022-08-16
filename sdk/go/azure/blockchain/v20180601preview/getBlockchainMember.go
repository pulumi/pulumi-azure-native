


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlockchainMember(ctx *pulumi.Context, args *LookupBlockchainMemberArgs, opts ...pulumi.InvokeOption) (*LookupBlockchainMemberResult, error) {
	var rv LookupBlockchainMemberResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:getBlockchainMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlockchainMemberArgs struct {
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupBlockchainMemberResult struct {
	Consortium                          *string                           `pulumi:"consortium"`
	ConsortiumManagementAccountAddress  string                            `pulumi:"consortiumManagementAccountAddress"`
	ConsortiumManagementAccountPassword *string                           `pulumi:"consortiumManagementAccountPassword"`
	ConsortiumMemberDisplayName         *string                           `pulumi:"consortiumMemberDisplayName"`
	ConsortiumRole                      *string                           `pulumi:"consortiumRole"`
	Dns                                 string                            `pulumi:"dns"`
	FirewallRules                       []FirewallRuleResponse            `pulumi:"firewallRules"`
	Id                                  string                            `pulumi:"id"`
	Location                            *string                           `pulumi:"location"`
	Name                                string                            `pulumi:"name"`
	Password                            *string                           `pulumi:"password"`
	Protocol                            *string                           `pulumi:"protocol"`
	ProvisioningState                   string                            `pulumi:"provisioningState"`
	PublicKey                           string                            `pulumi:"publicKey"`
	RootContractAddress                 string                            `pulumi:"rootContractAddress"`
	Sku                                 *SkuResponse                      `pulumi:"sku"`
	Tags                                map[string]string                 `pulumi:"tags"`
	Type                                string                            `pulumi:"type"`
	UserName                            string                            `pulumi:"userName"`
	ValidatorNodesSku                   *BlockchainMemberNodesSkuResponse `pulumi:"validatorNodesSku"`
}

func LookupBlockchainMemberOutput(ctx *pulumi.Context, args LookupBlockchainMemberOutputArgs, opts ...pulumi.InvokeOption) LookupBlockchainMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlockchainMemberResult, error) {
			args := v.(LookupBlockchainMemberArgs)
			r, err := LookupBlockchainMember(ctx, &args, opts...)
			var s LookupBlockchainMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlockchainMemberResultOutput)
}

type LookupBlockchainMemberOutputArgs struct {
	BlockchainMemberName pulumi.StringInput `pulumi:"blockchainMemberName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlockchainMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockchainMemberArgs)(nil)).Elem()
}


type LookupBlockchainMemberResultOutput struct{ *pulumi.OutputState }

func (LookupBlockchainMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockchainMemberResult)(nil)).Elem()
}

func (o LookupBlockchainMemberResultOutput) ToLookupBlockchainMemberResultOutput() LookupBlockchainMemberResultOutput {
	return o
}

func (o LookupBlockchainMemberResultOutput) ToLookupBlockchainMemberResultOutputWithContext(ctx context.Context) LookupBlockchainMemberResultOutput {
	return o
}

func (o LookupBlockchainMemberResultOutput) Consortium() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.Consortium }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) ConsortiumManagementAccountAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.ConsortiumManagementAccountAddress }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) ConsortiumManagementAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.ConsortiumManagementAccountPassword }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) ConsortiumMemberDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.ConsortiumMemberDisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) ConsortiumRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.ConsortiumRole }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) Dns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.Dns }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) FirewallRules() FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) []FirewallRuleResponse { return v.FirewallRules }).(FirewallRuleResponseArrayOutput)
}

func (o LookupBlockchainMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LookupBlockchainMemberResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) RootContractAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.RootContractAddress }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupBlockchainMemberResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBlockchainMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) string { return v.UserName }).(pulumi.StringOutput)
}

func (o LookupBlockchainMemberResultOutput) ValidatorNodesSku() BlockchainMemberNodesSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBlockchainMemberResult) *BlockchainMemberNodesSkuResponse { return v.ValidatorNodesSku }).(BlockchainMemberNodesSkuResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlockchainMemberResultOutput{})
}
