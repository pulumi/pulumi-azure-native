


package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSoftwareUpdateConfigurationByName(ctx *pulumi.Context, args *LookupSoftwareUpdateConfigurationByNameArgs, opts ...pulumi.InvokeOption) (*LookupSoftwareUpdateConfigurationByNameResult, error) {
	var rv LookupSoftwareUpdateConfigurationByNameResult
	err := ctx.Invoke("azure-native:automation:getSoftwareUpdateConfigurationByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSoftwareUpdateConfigurationByNameArgs struct {
	AutomationAccountName           string `pulumi:"automationAccountName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
	SoftwareUpdateConfigurationName string `pulumi:"softwareUpdateConfigurationName"`
}


type LookupSoftwareUpdateConfigurationByNameResult struct {
	CreatedBy           string                                    `pulumi:"createdBy"`
	CreationTime        string                                    `pulumi:"creationTime"`
	Error               *ErrorResponseResponse                    `pulumi:"error"`
	Id                  string                                    `pulumi:"id"`
	LastModifiedBy      string                                    `pulumi:"lastModifiedBy"`
	LastModifiedTime    string                                    `pulumi:"lastModifiedTime"`
	Name                string                                    `pulumi:"name"`
	ProvisioningState   string                                    `pulumi:"provisioningState"`
	ScheduleInfo        SUCSchedulePropertiesResponse             `pulumi:"scheduleInfo"`
	Tasks               *SoftwareUpdateConfigurationTasksResponse `pulumi:"tasks"`
	Type                string                                    `pulumi:"type"`
	UpdateConfiguration UpdateConfigurationResponse               `pulumi:"updateConfiguration"`
}


func (val *LookupSoftwareUpdateConfigurationByNameResult) Defaults() *LookupSoftwareUpdateConfigurationByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScheduleInfo = *tmp.ScheduleInfo.Defaults()

	return &tmp
}
