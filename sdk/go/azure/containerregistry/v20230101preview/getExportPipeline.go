


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExportPipeline(ctx *pulumi.Context, args *LookupExportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupExportPipelineResult, error) {
	var rv LookupExportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:getExportPipeline", args, &rv, opts...)
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

func LookupExportPipelineOutput(ctx *pulumi.Context, args LookupExportPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupExportPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportPipelineResult, error) {
			args := v.(LookupExportPipelineArgs)
			r, err := LookupExportPipeline(ctx, &args, opts...)
			var s LookupExportPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportPipelineResultOutput)
}

type LookupExportPipelineOutputArgs struct {
	ExportPipelineName pulumi.StringInput `pulumi:"exportPipelineName"`
	RegistryName       pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExportPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportPipelineArgs)(nil)).Elem()
}


type LookupExportPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupExportPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportPipelineResult)(nil)).Elem()
}

func (o LookupExportPipelineResultOutput) ToLookupExportPipelineResultOutput() LookupExportPipelineResultOutput {
	return o
}

func (o LookupExportPipelineResultOutput) ToLookupExportPipelineResultOutputWithContext(ctx context.Context) LookupExportPipelineResultOutput {
	return o
}

func (o LookupExportPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupExportPipelineResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupExportPipelineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExportPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExportPipelineResultOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) []string { return v.Options }).(pulumi.StringArrayOutput)
}

func (o LookupExportPipelineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExportPipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupExportPipelineResultOutput) Target() ExportPipelineTargetPropertiesResponseOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) ExportPipelineTargetPropertiesResponse { return v.Target }).(ExportPipelineTargetPropertiesResponseOutput)
}

func (o LookupExportPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportPipelineResultOutput{})
}
