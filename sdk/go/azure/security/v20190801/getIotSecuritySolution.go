


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotSecuritySolution(ctx *pulumi.Context, args *LookupIotSecuritySolutionArgs, opts ...pulumi.InvokeOption) (*LookupIotSecuritySolutionResult, error) {
	var rv LookupIotSecuritySolutionResult
	err := ctx.Invoke("azure-native:security/v20190801:getIotSecuritySolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIotSecuritySolutionArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SolutionName      string `pulumi:"solutionName"`
}


type LookupIotSecuritySolutionResult struct {
	AdditionalWorkspaces         []AdditionalWorkspacesPropertiesResponse        `pulumi:"additionalWorkspaces"`
	AutoDiscoveredResources      []string                                        `pulumi:"autoDiscoveredResources"`
	DisabledDataSources          []string                                        `pulumi:"disabledDataSources"`
	DisplayName                  string                                          `pulumi:"displayName"`
	Export                       []string                                        `pulumi:"export"`
	Id                           string                                          `pulumi:"id"`
	IotHubs                      []string                                        `pulumi:"iotHubs"`
	Location                     *string                                         `pulumi:"location"`
	Name                         string                                          `pulumi:"name"`
	RecommendationsConfiguration []RecommendationConfigurationPropertiesResponse `pulumi:"recommendationsConfiguration"`
	Status                       *string                                         `pulumi:"status"`
	SystemData                   SystemDataResponse                              `pulumi:"systemData"`
	Tags                         map[string]string                               `pulumi:"tags"`
	Type                         string                                          `pulumi:"type"`
	UnmaskedIpLoggingStatus      *string                                         `pulumi:"unmaskedIpLoggingStatus"`
	UserDefinedResources         *UserDefinedResourcesPropertiesResponse         `pulumi:"userDefinedResources"`
	Workspace                    *string                                         `pulumi:"workspace"`
}


func (val *LookupIotSecuritySolutionResult) Defaults() *LookupIotSecuritySolutionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	if isZero(tmp.UnmaskedIpLoggingStatus) {
		unmaskedIpLoggingStatus_ := "Disabled"
		tmp.UnmaskedIpLoggingStatus = &unmaskedIpLoggingStatus_
	}
	return &tmp
}
