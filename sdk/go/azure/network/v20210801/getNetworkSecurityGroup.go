


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkSecurityGroup(ctx *pulumi.Context, args *LookupNetworkSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSecurityGroupResult, error) {
	var rv LookupNetworkSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20210801:getNetworkSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSecurityGroupArgs struct {
	Expand                   *string `pulumi:"expand"`
	NetworkSecurityGroupName string  `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
}


type LookupNetworkSecurityGroupResult struct {
	DefaultSecurityRules []SecurityRuleResponse     `pulumi:"defaultSecurityRules"`
	Etag                 string                     `pulumi:"etag"`
	FlowLogs             []FlowLogResponse          `pulumi:"flowLogs"`
	Id                   *string                    `pulumi:"id"`
	Location             *string                    `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	NetworkInterfaces    []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	ProvisioningState    string                     `pulumi:"provisioningState"`
	ResourceGuid         string                     `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleResponse     `pulumi:"securityRules"`
	Subnets              []SubnetResponse           `pulumi:"subnets"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
}

func LookupNetworkSecurityGroupOutput(ctx *pulumi.Context, args LookupNetworkSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkSecurityGroupResult, error) {
			args := v.(LookupNetworkSecurityGroupArgs)
			r, err := LookupNetworkSecurityGroup(ctx, &args, opts...)
			var s LookupNetworkSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkSecurityGroupResultOutput)
}

type LookupNetworkSecurityGroupOutputArgs struct {
	Expand                   pulumi.StringPtrInput `pulumi:"expand"`
	NetworkSecurityGroupName pulumi.StringInput    `pulumi:"networkSecurityGroupName"`
	ResourceGroupName        pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNetworkSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityGroupArgs)(nil)).Elem()
}


type LookupNetworkSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityGroupResult)(nil)).Elem()
}

func (o LookupNetworkSecurityGroupResultOutput) ToLookupNetworkSecurityGroupResultOutput() LookupNetworkSecurityGroupResultOutput {
	return o
}

func (o LookupNetworkSecurityGroupResultOutput) ToLookupNetworkSecurityGroupResultOutputWithContext(ctx context.Context) LookupNetworkSecurityGroupResultOutput {
	return o
}

func (o LookupNetworkSecurityGroupResultOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SecurityRuleResponse { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) FlowLogs() FlowLogResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []FlowLogResponse { return v.FlowLogs }).(FlowLogResponseArrayOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SecurityRuleResponse { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkSecurityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkSecurityGroupResultOutput{})
}
