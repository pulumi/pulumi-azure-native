


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNspAccessRule(ctx *pulumi.Context, args *LookupNspAccessRuleArgs, opts ...pulumi.InvokeOption) (*LookupNspAccessRuleResult, error) {
	var rv LookupNspAccessRuleResult
	err := ctx.Invoke("azure-native:network:getNspAccessRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspAccessRuleArgs struct {
	AccessRuleName               string `pulumi:"accessRuleName"`
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	ProfileName                  string `pulumi:"profileName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type LookupNspAccessRuleResult struct {
	AddressPrefixes           []string                           `pulumi:"addressPrefixes"`
	Direction                 *string                            `pulumi:"direction"`
	FullyQualifiedDomainNames []string                           `pulumi:"fullyQualifiedDomainNames"`
	Id                        string                             `pulumi:"id"`
	Location                  *string                            `pulumi:"location"`
	Name                      string                             `pulumi:"name"`
	NetworkSecurityPerimeters []PerimeterBasedAccessRuleResponse `pulumi:"networkSecurityPerimeters"`
	ProvisioningState         string                             `pulumi:"provisioningState"`
	Subscriptions             []string                           `pulumi:"subscriptions"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Type                      string                             `pulumi:"type"`
}

func LookupNspAccessRuleOutput(ctx *pulumi.Context, args LookupNspAccessRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNspAccessRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspAccessRuleResult, error) {
			args := v.(LookupNspAccessRuleArgs)
			r, err := LookupNspAccessRule(ctx, &args, opts...)
			var s LookupNspAccessRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspAccessRuleResultOutput)
}

type LookupNspAccessRuleOutputArgs struct {
	AccessRuleName               pulumi.StringInput `pulumi:"accessRuleName"`
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	ProfileName                  pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspAccessRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAccessRuleArgs)(nil)).Elem()
}


type LookupNspAccessRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNspAccessRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAccessRuleResult)(nil)).Elem()
}

func (o LookupNspAccessRuleResultOutput) ToLookupNspAccessRuleResultOutput() LookupNspAccessRuleResultOutput {
	return o
}

func (o LookupNspAccessRuleResultOutput) ToLookupNspAccessRuleResultOutputWithContext(ctx context.Context) LookupNspAccessRuleResultOutput {
	return o
}

func (o LookupNspAccessRuleResultOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o LookupNspAccessRuleResultOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o LookupNspAccessRuleResultOutput) FullyQualifiedDomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.FullyQualifiedDomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupNspAccessRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNspAccessRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNspAccessRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNspAccessRuleResultOutput) NetworkSecurityPerimeters() PerimeterBasedAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []PerimeterBasedAccessRuleResponse {
		return v.NetworkSecurityPerimeters
	}).(PerimeterBasedAccessRuleResponseArrayOutput)
}

func (o LookupNspAccessRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNspAccessRuleResultOutput) Subscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.Subscriptions }).(pulumi.StringArrayOutput)
}

func (o LookupNspAccessRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNspAccessRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspAccessRuleResultOutput{})
}
