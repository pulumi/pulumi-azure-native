


package v20150501

import (
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
