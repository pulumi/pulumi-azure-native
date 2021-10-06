


package v20151031

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscConfiguration(ctx *pulumi.Context, args *LookupDscConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscConfigurationResult, error) {
	var rv LookupDscConfigurationResult
	err := ctx.Invoke("azure-native:automation/v20151031:getDscConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscConfigurationArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ConfigurationName     string `pulumi:"configurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscConfigurationResult struct {
	CreationTime           *string                                      `pulumi:"creationTime"`
	Description            *string                                      `pulumi:"description"`
	Etag                   *string                                      `pulumi:"etag"`
	Id                     string                                       `pulumi:"id"`
	JobCount               *int                                         `pulumi:"jobCount"`
	LastModifiedTime       *string                                      `pulumi:"lastModifiedTime"`
	Location               *string                                      `pulumi:"location"`
	LogVerbose             *bool                                        `pulumi:"logVerbose"`
	Name                   string                                       `pulumi:"name"`
	NodeConfigurationCount *int                                         `pulumi:"nodeConfigurationCount"`
	Parameters             map[string]DscConfigurationParameterResponse `pulumi:"parameters"`
	ProvisioningState      *string                                      `pulumi:"provisioningState"`
	Source                 *ContentSourceResponse                       `pulumi:"source"`
	State                  *string                                      `pulumi:"state"`
	Tags                   map[string]string                            `pulumi:"tags"`
	Type                   string                                       `pulumi:"type"`
}
