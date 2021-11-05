


package v20180715preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceTask(ctx *pulumi.Context, args *LookupServiceTaskArgs, opts ...pulumi.InvokeOption) (*LookupServiceTaskResult, error) {
	var rv LookupServiceTaskResult
	err := ctx.Invoke("azure-native:datamigration/v20180715preview:getServiceTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceTaskArgs struct {
	Expand      *string `pulumi:"expand"`
	GroupName   string  `pulumi:"groupName"`
	ServiceName string  `pulumi:"serviceName"`
	TaskName    string  `pulumi:"taskName"`
}


type LookupServiceTaskResult struct {
	Etag       *string     `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
