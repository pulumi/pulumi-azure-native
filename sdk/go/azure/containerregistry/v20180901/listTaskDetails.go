


package v20180901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTaskDetails(ctx *pulumi.Context, args *ListTaskDetailsArgs, opts ...pulumi.InvokeOption) (*ListTaskDetailsResult, error) {
	var rv ListTaskDetailsResult
	err := ctx.Invoke("azure-native:containerregistry/v20180901:listTaskDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListTaskDetailsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TaskName          string `pulumi:"taskName"`
}



type ListTaskDetailsResult struct {
	AgentConfiguration *AgentPropertiesResponse   `pulumi:"agentConfiguration"`
	CreationDate       string                     `pulumi:"creationDate"`
	Credentials        *CredentialsResponse       `pulumi:"credentials"`
	Id                 string                     `pulumi:"id"`
	Location           string                     `pulumi:"location"`
	Name               string                     `pulumi:"name"`
	Platform           PlatformPropertiesResponse `pulumi:"platform"`
	ProvisioningState  string                     `pulumi:"provisioningState"`
	Status             *string                    `pulumi:"status"`
	Step               interface{}                `pulumi:"step"`
	Tags               map[string]string          `pulumi:"tags"`
	Timeout            *int                       `pulumi:"timeout"`
	Trigger            *TriggerPropertiesResponse `pulumi:"trigger"`
	Type               string                     `pulumi:"type"`
}


func (val *ListTaskDetailsResult) Defaults() *ListTaskDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Timeout) {
		timeout_ := 3600
		tmp.Timeout = &timeout_
	}
	tmp.Trigger = tmp.Trigger.Defaults()

	return &tmp
}
