


package v20180315preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("azure-native:datamigration/v20180315preview:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	Expand      *string `pulumi:"expand"`
	GroupName   string  `pulumi:"groupName"`
	ProjectName string  `pulumi:"projectName"`
	ServiceName string  `pulumi:"serviceName"`
	TaskName    string  `pulumi:"taskName"`
}


type LookupTaskResult struct {
	Etag       *string     `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
