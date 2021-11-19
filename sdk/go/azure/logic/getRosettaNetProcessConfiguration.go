


package logic

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRosettaNetProcessConfiguration(ctx *pulumi.Context, args *LookupRosettaNetProcessConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupRosettaNetProcessConfigurationResult, error) {
	var rv LookupRosettaNetProcessConfigurationResult
	err := ctx.Invoke("azure-native:logic:getRosettaNetProcessConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRosettaNetProcessConfigurationArgs struct {
	IntegrationAccountName             string `pulumi:"integrationAccountName"`
	ResourceGroupName                  string `pulumi:"resourceGroupName"`
	RosettaNetProcessConfigurationName string `pulumi:"rosettaNetProcessConfigurationName"`
}


type LookupRosettaNetProcessConfigurationResult struct {
	ActivitySettings      RosettaNetPipActivitySettingsResponse `pulumi:"activitySettings"`
	ChangedTime           string                                `pulumi:"changedTime"`
	CreatedTime           string                                `pulumi:"createdTime"`
	Description           *string                               `pulumi:"description"`
	Id                    string                                `pulumi:"id"`
	InitiatorRoleSettings RosettaNetPipRoleSettingsResponse     `pulumi:"initiatorRoleSettings"`
	Location              *string                               `pulumi:"location"`
	Metadata              map[string]string                     `pulumi:"metadata"`
	Name                  string                                `pulumi:"name"`
	ProcessCode           string                                `pulumi:"processCode"`
	ProcessName           string                                `pulumi:"processName"`
	ProcessVersion        string                                `pulumi:"processVersion"`
	ResponderRoleSettings RosettaNetPipRoleSettingsResponse     `pulumi:"responderRoleSettings"`
	Tags                  map[string]string                     `pulumi:"tags"`
	Type                  string                                `pulumi:"type"`
}
