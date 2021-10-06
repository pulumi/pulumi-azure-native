


package v20140801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:scheduler/v20140801preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobCollectionName string `pulumi:"jobCollectionName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobResult struct {
	Id         string                `pulumi:"id"`
	Name       string                `pulumi:"name"`
	Properties JobPropertiesResponse `pulumi:"properties"`
	Type       string                `pulumi:"type"`
}
