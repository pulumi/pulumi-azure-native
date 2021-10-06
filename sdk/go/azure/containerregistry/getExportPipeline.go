


package containerregistry

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExportPipeline(ctx *pulumi.Context, args *LookupExportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupExportPipelineResult, error) {
	var rv LookupExportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry:getExportPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportPipelineArgs struct {
	ExportPipelineName string `pulumi:"exportPipelineName"`
	RegistryName       string `pulumi:"registryName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupExportPipelineResult struct {
	Id                string                                 `pulumi:"id"`
	Identity          *IdentityPropertiesResponse            `pulumi:"identity"`
	Location          *string                                `pulumi:"location"`
	Name              string                                 `pulumi:"name"`
	Options           []string                               `pulumi:"options"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                     `pulumi:"systemData"`
	Target            ExportPipelineTargetPropertiesResponse `pulumi:"target"`
	Type              string                                 `pulumi:"type"`
}
