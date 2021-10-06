


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobTargetGroup(ctx *pulumi.Context, args *LookupJobTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupJobTargetGroupResult, error) {
	var rv LookupJobTargetGroupResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getJobTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobTargetGroupArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	TargetGroupName   string `pulumi:"targetGroupName"`
}


type LookupJobTargetGroupResult struct {
	Id      string              `pulumi:"id"`
	Members []JobTargetResponse `pulumi:"members"`
	Name    string              `pulumi:"name"`
	Type    string              `pulumi:"type"`
}
