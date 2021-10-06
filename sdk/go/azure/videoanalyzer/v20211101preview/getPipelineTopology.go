


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipelineTopology(ctx *pulumi.Context, args *LookupPipelineTopologyArgs, opts ...pulumi.InvokeOption) (*LookupPipelineTopologyResult, error) {
	var rv LookupPipelineTopologyResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:getPipelineTopology", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineTopologyArgs struct {
	AccountName          string `pulumi:"accountName"`
	PipelineTopologyName string `pulumi:"pipelineTopologyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}







type LookupPipelineTopologyResult struct {
	Description *string                        `pulumi:"description"`
	Id          string                         `pulumi:"id"`
	Kind        string                         `pulumi:"kind"`
	Name        string                         `pulumi:"name"`
	Parameters  []ParameterDeclarationResponse `pulumi:"parameters"`
	Processors  []EncoderProcessorResponse     `pulumi:"processors"`
	Sinks       []VideoSinkResponse            `pulumi:"sinks"`
	Sku         SkuResponse                    `pulumi:"sku"`
	Sources     []interface{}                  `pulumi:"sources"`
	SystemData  SystemDataResponse             `pulumi:"systemData"`
	Type        string                         `pulumi:"type"`
}
