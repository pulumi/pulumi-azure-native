


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceTopology(ctx *pulumi.Context, args *LookupServiceTopologyArgs, opts ...pulumi.InvokeOption) (*LookupServiceTopologyResult, error) {
	var rv LookupServiceTopologyResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20191101preview:getServiceTopology", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceTopologyArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
}


type LookupServiceTopologyResult struct {
	ArtifactSourceId *string           `pulumi:"artifactSourceId"`
	Id               string            `pulumi:"id"`
	Location         string            `pulumi:"location"`
	Name             string            `pulumi:"name"`
	Tags             map[string]string `pulumi:"tags"`
	Type             string            `pulumi:"type"`
}
