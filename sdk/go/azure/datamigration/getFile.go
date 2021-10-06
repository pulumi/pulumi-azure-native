


package datamigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFile(ctx *pulumi.Context, args *LookupFileArgs, opts ...pulumi.InvokeOption) (*LookupFileResult, error) {
	var rv LookupFileResult
	err := ctx.Invoke("azure-native:datamigration:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileArgs struct {
	FileName    string `pulumi:"fileName"`
	GroupName   string `pulumi:"groupName"`
	ProjectName string `pulumi:"projectName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupFileResult struct {
	Etag       *string                       `pulumi:"etag"`
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties ProjectFilePropertiesResponse `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}
