


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:hdinsight/v20210601:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	Etag       *string                       `pulumi:"etag"`
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties ApplicationPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Tags       map[string]string             `pulumi:"tags"`
	Type       string                        `pulumi:"type"`
}
