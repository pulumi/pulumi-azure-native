


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupImportPipeline(ctx *pulumi.Context, args *LookupImportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupImportPipelineResult, error) {
	var rv LookupImportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry/v20201101preview:getImportPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupImportPipelineArgs struct {
	ImportPipelineName string `pulumi:"importPipelineName"`
	RegistryName       string `pulumi:"registryName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupImportPipelineResult struct {
	Id                string                                 `pulumi:"id"`
	Identity          *IdentityPropertiesResponse            `pulumi:"identity"`
	Location          *string                                `pulumi:"location"`
	Name              string                                 `pulumi:"name"`
	Options           []string                               `pulumi:"options"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	Source            ImportPipelineSourcePropertiesResponse `pulumi:"source"`
	SystemData        SystemDataResponse                     `pulumi:"systemData"`
	Trigger           *PipelineTriggerPropertiesResponse     `pulumi:"trigger"`
	Type              string                                 `pulumi:"type"`
}


func (val *LookupImportPipelineResult) Defaults() *LookupImportPipelineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = *tmp.Source.Defaults()

	tmp.Trigger = tmp.Trigger.Defaults()

	return &tmp
}
