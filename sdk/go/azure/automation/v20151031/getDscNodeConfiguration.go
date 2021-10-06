


package v20151031

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDscNodeConfiguration(ctx *pulumi.Context, args *LookupDscNodeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscNodeConfigurationResult, error) {
	var rv LookupDscNodeConfigurationResult
	err := ctx.Invoke("azure-native:automation/v20151031:getDscNodeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscNodeConfigurationArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	NodeConfigurationName string `pulumi:"nodeConfigurationName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDscNodeConfigurationResult struct {
	Configuration    *DscConfigurationAssociationPropertyResponse `pulumi:"configuration"`
	CreationTime     *string                                      `pulumi:"creationTime"`
	Id               string                                       `pulumi:"id"`
	LastModifiedTime *string                                      `pulumi:"lastModifiedTime"`
	Name             string                                       `pulumi:"name"`
	Type             string                                       `pulumi:"type"`
}
