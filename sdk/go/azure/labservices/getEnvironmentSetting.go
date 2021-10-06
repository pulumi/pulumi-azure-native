


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentSetting(ctx *pulumi.Context, args *LookupEnvironmentSettingArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSettingResult, error) {
	var rv LookupEnvironmentSettingResult
	err := ctx.Invoke("azure-native:labservices:getEnvironmentSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSettingArgs struct {
	EnvironmentSettingName string  `pulumi:"environmentSettingName"`
	Expand                 *string `pulumi:"expand"`
	LabAccountName         string  `pulumi:"labAccountName"`
	LabName                string  `pulumi:"labName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentSettingResult struct {
	ConfigurationState    *string                       `pulumi:"configurationState"`
	Description           *string                       `pulumi:"description"`
	Id                    string                        `pulumi:"id"`
	LastChanged           string                        `pulumi:"lastChanged"`
	LastPublished         string                        `pulumi:"lastPublished"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	ProvisioningState     *string                       `pulumi:"provisioningState"`
	PublishingState       string                        `pulumi:"publishingState"`
	ResourceSettings      ResourceSettingsResponse      `pulumi:"resourceSettings"`
	Tags                  map[string]string             `pulumi:"tags"`
	Title                 *string                       `pulumi:"title"`
	Type                  string                        `pulumi:"type"`
	UniqueIdentifier      *string                       `pulumi:"uniqueIdentifier"`
}
