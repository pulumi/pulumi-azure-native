


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:containerregistry:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTaskArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskName          string `pulumi:"taskName"`
}



type LookupTaskResult struct {
	AgentConfiguration *AgentPropertiesResponse    `pulumi:"agentConfiguration"`
	AgentPoolName      *string                     `pulumi:"agentPoolName"`
	CreationDate       string                      `pulumi:"creationDate"`
	Credentials        *CredentialsResponse        `pulumi:"credentials"`
	Id                 string                      `pulumi:"id"`
	Identity           *IdentityPropertiesResponse `pulumi:"identity"`
	IsSystemTask       *bool                       `pulumi:"isSystemTask"`
	Location           string                      `pulumi:"location"`
	LogTemplate        *string                     `pulumi:"logTemplate"`
	Name               string                      `pulumi:"name"`
	Platform           *PlatformPropertiesResponse `pulumi:"platform"`
	ProvisioningState  string                      `pulumi:"provisioningState"`
	Status             *string                     `pulumi:"status"`
	Step               interface{}                 `pulumi:"step"`
	SystemData         SystemDataResponse          `pulumi:"systemData"`
	Tags               map[string]string           `pulumi:"tags"`
	Timeout            *int                        `pulumi:"timeout"`
	Trigger            *TriggerPropertiesResponse  `pulumi:"trigger"`
	Type               string                      `pulumi:"type"`
}


func (val *LookupTaskResult) Defaults() *LookupTaskResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSystemTask) {
		isSystemTask_ := false
		tmp.IsSystemTask = &isSystemTask_
	}
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	tmp.Trigger = tmp.Trigger.Defaults()

	return &tmp
}
