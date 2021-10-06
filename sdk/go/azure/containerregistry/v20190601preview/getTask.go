


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
