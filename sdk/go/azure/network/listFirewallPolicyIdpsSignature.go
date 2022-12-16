


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFirewallPolicyIdpsSignature(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignatureArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignatureResult, error) {
	var rv ListFirewallPolicyIdpsSignatureResult
	err := ctx.Invoke("azure-native:network:listFirewallPolicyIdpsSignature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignatureArgs struct {
	Filters            []FilterItems `pulumi:"filters"`
	FirewallPolicyName string        `pulumi:"firewallPolicyName"`
	OrderBy            *OrderBy      `pulumi:"orderBy"`
	ResourceGroupName  string        `pulumi:"resourceGroupName"`
	ResultsPerPage     *int          `pulumi:"resultsPerPage"`
	Search             *string       `pulumi:"search"`
	Skip               *int          `pulumi:"skip"`
}


type ListFirewallPolicyIdpsSignatureResult struct {
	MatchingRecordsCount *float64                    `pulumi:"matchingRecordsCount"`
	Signatures           []SingleQueryResultResponse `pulumi:"signatures"`
}

func ListFirewallPolicyIdpsSignatureOutput(ctx *pulumi.Context, args ListFirewallPolicyIdpsSignatureOutputArgs, opts ...pulumi.InvokeOption) ListFirewallPolicyIdpsSignatureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFirewallPolicyIdpsSignatureResult, error) {
			args := v.(ListFirewallPolicyIdpsSignatureArgs)
			r, err := ListFirewallPolicyIdpsSignature(ctx, &args, opts...)
			var s ListFirewallPolicyIdpsSignatureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFirewallPolicyIdpsSignatureResultOutput)
}

type ListFirewallPolicyIdpsSignatureOutputArgs struct {
	Filters            FilterItemsArrayInput `pulumi:"filters"`
	FirewallPolicyName pulumi.StringInput    `pulumi:"firewallPolicyName"`
	OrderBy            OrderByPtrInput       `pulumi:"orderBy"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResultsPerPage     pulumi.IntPtrInput    `pulumi:"resultsPerPage"`
	Search             pulumi.StringPtrInput `pulumi:"search"`
	Skip               pulumi.IntPtrInput    `pulumi:"skip"`
}

func (ListFirewallPolicyIdpsSignatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignatureArgs)(nil)).Elem()
}


type ListFirewallPolicyIdpsSignatureResultOutput struct{ *pulumi.OutputState }

func (ListFirewallPolicyIdpsSignatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignatureResult)(nil)).Elem()
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) ToListFirewallPolicyIdpsSignatureResultOutput() ListFirewallPolicyIdpsSignatureResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) ToListFirewallPolicyIdpsSignatureResultOutputWithContext(ctx context.Context) ListFirewallPolicyIdpsSignatureResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) MatchingRecordsCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignatureResult) *float64 { return v.MatchingRecordsCount }).(pulumi.Float64PtrOutput)
}

func (o ListFirewallPolicyIdpsSignatureResultOutput) Signatures() SingleQueryResultResponseArrayOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignatureResult) []SingleQueryResultResponse { return v.Signatures }).(SingleQueryResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFirewallPolicyIdpsSignatureResultOutput{})
}
