


package hybriddata

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobDefinition(ctx *pulumi.Context, args *LookupJobDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupJobDefinitionResult, error) {
	var rv LookupJobDefinitionResult
	err := ctx.Invoke("azure-native:hybriddata:getJobDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobDefinitionArgs struct {
	DataManagerName   string `pulumi:"dataManagerName"`
	DataServiceName   string `pulumi:"dataServiceName"`
	JobDefinitionName string `pulumi:"jobDefinitionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupJobDefinitionResult struct {
	CustomerSecrets  []CustomerSecretResponse `pulumi:"customerSecrets"`
	DataServiceInput interface{}              `pulumi:"dataServiceInput"`
	DataSinkId       string                   `pulumi:"dataSinkId"`
	DataSourceId     string                   `pulumi:"dataSourceId"`
	Id               string                   `pulumi:"id"`
	LastModifiedTime *string                  `pulumi:"lastModifiedTime"`
	Name             string                   `pulumi:"name"`
	RunLocation      *string                  `pulumi:"runLocation"`
	Schedules        []ScheduleResponse       `pulumi:"schedules"`
	State            string                   `pulumi:"state"`
	Type             string                   `pulumi:"type"`
	UserConfirmation *string                  `pulumi:"userConfirmation"`
}


func (val *LookupJobDefinitionResult) Defaults() *LookupJobDefinitionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UserConfirmation) {
		userConfirmation_ := "NotRequired"
		tmp.UserConfirmation = &userConfirmation_
	}
	return &tmp
}
