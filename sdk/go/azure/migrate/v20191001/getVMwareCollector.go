


package v20191001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVMwareCollector(ctx *pulumi.Context, args *LookupVMwareCollectorArgs, opts ...pulumi.InvokeOption) (*LookupVMwareCollectorResult, error) {
	var rv LookupVMwareCollectorResult
	err := ctx.Invoke("azure-native:migrate/v20191001:getVMwareCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVMwareCollectorArgs struct {
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VmWareCollectorName string `pulumi:"vmWareCollectorName"`
}

type LookupVMwareCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}
