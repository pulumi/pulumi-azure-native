


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobResult struct {
	Description *string              `pulumi:"description"`
	Id          string               `pulumi:"id"`
	Name        string               `pulumi:"name"`
	Schedule    *JobScheduleResponse `pulumi:"schedule"`
	Type        string               `pulumi:"type"`
	Version     int                  `pulumi:"version"`
}
