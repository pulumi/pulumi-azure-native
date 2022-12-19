


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProactiveDetectionConfiguration(ctx *pulumi.Context, args *LookupProactiveDetectionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupProactiveDetectionConfigurationResult, error) {
	var rv LookupProactiveDetectionConfigurationResult
	err := ctx.Invoke("azure-native:insights/v20150501:getProactiveDetectionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProactiveDetectionConfigurationArgs struct {
	ConfigurationId   string `pulumi:"configurationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupProactiveDetectionConfigurationResult struct {
	CustomEmails                   []string                                                                            `pulumi:"customEmails"`
	Enabled                        *bool                                                                               `pulumi:"enabled"`
	LastUpdatedTime                *string                                                                             `pulumi:"lastUpdatedTime"`
	Name                           *string                                                                             `pulumi:"name"`
	RuleDefinitions                *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions `pulumi:"ruleDefinitions"`
	SendEmailsToSubscriptionOwners *bool                                                                               `pulumi:"sendEmailsToSubscriptionOwners"`
}

func LookupProactiveDetectionConfigurationOutput(ctx *pulumi.Context, args LookupProactiveDetectionConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupProactiveDetectionConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProactiveDetectionConfigurationResult, error) {
			args := v.(LookupProactiveDetectionConfigurationArgs)
			r, err := LookupProactiveDetectionConfiguration(ctx, &args, opts...)
			var s LookupProactiveDetectionConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProactiveDetectionConfigurationResultOutput)
}

type LookupProactiveDetectionConfigurationOutputArgs struct {
	ConfigurationId   pulumi.StringInput `pulumi:"configurationId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupProactiveDetectionConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveDetectionConfigurationArgs)(nil)).Elem()
}


type LookupProactiveDetectionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupProactiveDetectionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProactiveDetectionConfigurationResult)(nil)).Elem()
}

func (o LookupProactiveDetectionConfigurationResultOutput) ToLookupProactiveDetectionConfigurationResultOutput() LookupProactiveDetectionConfigurationResultOutput {
	return o
}

func (o LookupProactiveDetectionConfigurationResultOutput) ToLookupProactiveDetectionConfigurationResultOutputWithContext(ctx context.Context) LookupProactiveDetectionConfigurationResultOutput {
	return o
}

func (o LookupProactiveDetectionConfigurationResultOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o LookupProactiveDetectionConfigurationResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupProactiveDetectionConfigurationResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupProactiveDetectionConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupProactiveDetectionConfigurationResultOutput) RuleDefinitions() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions {
		return v.RuleDefinitions
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput)
}

func (o LookupProactiveDetectionConfigurationResultOutput) SendEmailsToSubscriptionOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProactiveDetectionConfigurationResult) *bool { return v.SendEmailsToSubscriptionOwners }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProactiveDetectionConfigurationResultOutput{})
}
