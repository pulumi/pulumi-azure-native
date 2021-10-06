


package scheduler

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobCollection(ctx *pulumi.Context, args *LookupJobCollectionArgs, opts ...pulumi.InvokeOption) (*LookupJobCollectionResult, error) {
	var rv LookupJobCollectionResult
	err := ctx.Invoke("azure-native:scheduler:getJobCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCollectionArgs struct {
	JobCollectionName string `pulumi:"jobCollectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupJobCollectionResult struct {
	Id         string                          `pulumi:"id"`
	Location   *string                         `pulumi:"location"`
	Name       *string                         `pulumi:"name"`
	Properties JobCollectionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string               `pulumi:"tags"`
	Type       string                          `pulumi:"type"`
}
