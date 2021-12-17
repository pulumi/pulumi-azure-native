


package v20211030preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:datamigration/v20211030preview:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	GroupName   string `pulumi:"groupName"`
	ProjectName string `pulumi:"projectName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupProjectResult struct {
	AzureAuthenticationInfo *string                `pulumi:"azureAuthenticationInfo"`
	CreationTime            string                 `pulumi:"creationTime"`
	DatabasesInfo           []DatabaseInfoResponse `pulumi:"databasesInfo"`
	ETag                    *string                `pulumi:"eTag"`
	Id                      string                 `pulumi:"id"`
	Location                *string                `pulumi:"location"`
	Name                    string                 `pulumi:"name"`
	ProvisioningState       string                 `pulumi:"provisioningState"`
	SourceConnectionInfo    interface{}            `pulumi:"sourceConnectionInfo"`
	SourcePlatform          string                 `pulumi:"sourcePlatform"`
	SystemData              SystemDataResponse     `pulumi:"systemData"`
	Tags                    map[string]string      `pulumi:"tags"`
	TargetConnectionInfo    interface{}            `pulumi:"targetConnectionInfo"`
	TargetPlatform          string                 `pulumi:"targetPlatform"`
	Type                    string                 `pulumi:"type"`
}
