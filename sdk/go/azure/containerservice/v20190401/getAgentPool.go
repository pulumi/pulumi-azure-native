


package v20190401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerservice/v20190401:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAgentPoolArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupAgentPoolResult struct {
	AvailabilityZones   []string `pulumi:"availabilityZones"`
	Count               int      `pulumi:"count"`
	EnableAutoScaling   *bool    `pulumi:"enableAutoScaling"`
	Id                  string   `pulumi:"id"`
	MaxCount            *int     `pulumi:"maxCount"`
	MaxPods             *int     `pulumi:"maxPods"`
	MinCount            *int     `pulumi:"minCount"`
	Name                string   `pulumi:"name"`
	OrchestratorVersion *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        *int     `pulumi:"osDiskSizeGB"`
	OsType              *string  `pulumi:"osType"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	Type                string   `pulumi:"type"`
	VmSize              string   `pulumi:"vmSize"`
	VnetSubnetID        *string  `pulumi:"vnetSubnetID"`
}


func (val *LookupAgentPoolResult) Defaults() *LookupAgentPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
}
