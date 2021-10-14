


package migrate

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerCollector(ctx *pulumi.Context, args *LookupServerCollectorArgs, opts ...pulumi.InvokeOption) (*LookupServerCollectorResult, error) {
	var rv LookupServerCollectorResult
	err := ctx.Invoke("azure-native:migrate:getServerCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCollectorArgs struct {
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServerCollectorName string `pulumi:"serverCollectorName"`
}

type LookupServerCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}
