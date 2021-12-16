


package v20180315preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:datamigration/v20180315preview:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProjectArgs struct {
	GroupName   string `pulumi:"groupName"`
	ProjectName string `pulumi:"projectName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupProjectResult struct {
	CreationTime         string                     `pulumi:"creationTime"`
	DatabasesInfo        []DatabaseInfoResponse     `pulumi:"databasesInfo"`
	Id                   string                     `pulumi:"id"`
	Location             string                     `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	ProvisioningState    string                     `pulumi:"provisioningState"`
	SourceConnectionInfo *SqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	SourcePlatform       string                     `pulumi:"sourcePlatform"`
	Tags                 map[string]string          `pulumi:"tags"`
	TargetConnectionInfo *SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
	TargetPlatform       string                     `pulumi:"targetPlatform"`
	Type                 string                     `pulumi:"type"`
}


func (val *LookupProjectResult) Defaults() *LookupProjectResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = tmp.TargetConnectionInfo.Defaults()

	return &tmp
}
