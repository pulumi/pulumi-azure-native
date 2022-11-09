


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConfigurationPolicyGroup(ctx *pulumi.Context, args *LookupConfigurationPolicyGroupArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationPolicyGroupResult, error) {
	var rv LookupConfigurationPolicyGroupResult
	err := ctx.Invoke("azure-native:network/v20220501:getConfigurationPolicyGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationPolicyGroupArgs struct {
	ConfigurationPolicyGroupName string `pulumi:"configurationPolicyGroupName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
	VpnServerConfigurationName   string `pulumi:"vpnServerConfigurationName"`
}


type LookupConfigurationPolicyGroupResult struct {
	Etag                        string                                            `pulumi:"etag"`
	Id                          *string                                           `pulumi:"id"`
	IsDefault                   *bool                                             `pulumi:"isDefault"`
	Name                        *string                                           `pulumi:"name"`
	P2SConnectionConfigurations []SubResourceResponse                             `pulumi:"p2SConnectionConfigurations"`
	PolicyMembers               []VpnServerConfigurationPolicyGroupMemberResponse `pulumi:"policyMembers"`
	Priority                    *int                                              `pulumi:"priority"`
	ProvisioningState           string                                            `pulumi:"provisioningState"`
	Type                        string                                            `pulumi:"type"`
}

func LookupConfigurationPolicyGroupOutput(ctx *pulumi.Context, args LookupConfigurationPolicyGroupOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationPolicyGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationPolicyGroupResult, error) {
			args := v.(LookupConfigurationPolicyGroupArgs)
			r, err := LookupConfigurationPolicyGroup(ctx, &args, opts...)
			var s LookupConfigurationPolicyGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationPolicyGroupResultOutput)
}

type LookupConfigurationPolicyGroupOutputArgs struct {
	ConfigurationPolicyGroupName pulumi.StringInput `pulumi:"configurationPolicyGroupName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
	VpnServerConfigurationName   pulumi.StringInput `pulumi:"vpnServerConfigurationName"`
}

func (LookupConfigurationPolicyGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationPolicyGroupArgs)(nil)).Elem()
}


type LookupConfigurationPolicyGroupResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationPolicyGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationPolicyGroupResult)(nil)).Elem()
}

func (o LookupConfigurationPolicyGroupResultOutput) ToLookupConfigurationPolicyGroupResultOutput() LookupConfigurationPolicyGroupResultOutput {
	return o
}

func (o LookupConfigurationPolicyGroupResultOutput) ToLookupConfigurationPolicyGroupResultOutputWithContext(ctx context.Context) LookupConfigurationPolicyGroupResultOutput {
	return o
}

func (o LookupConfigurationPolicyGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) P2SConnectionConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) []SubResourceResponse {
		return v.P2SConnectionConfigurations
	}).(SubResourceResponseArrayOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) []VpnServerConfigurationPolicyGroupMemberResponse {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConfigurationPolicyGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationPolicyGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationPolicyGroupResultOutput{})
}
