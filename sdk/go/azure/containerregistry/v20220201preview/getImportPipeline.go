


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupImportPipeline(ctx *pulumi.Context, args *LookupImportPipelineArgs, opts ...pulumi.InvokeOption) (*LookupImportPipelineResult, error) {
	var rv LookupImportPipelineResult
	err := ctx.Invoke("azure-native:containerregistry/v20220201preview:getImportPipeline", args, &rv, opts...)
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

func LookupImportPipelineOutput(ctx *pulumi.Context, args LookupImportPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupImportPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportPipelineResult, error) {
			args := v.(LookupImportPipelineArgs)
			r, err := LookupImportPipeline(ctx, &args, opts...)
			var s LookupImportPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportPipelineResultOutput)
}

type LookupImportPipelineOutputArgs struct {
	ImportPipelineName pulumi.StringInput `pulumi:"importPipelineName"`
	RegistryName       pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupImportPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportPipelineArgs)(nil)).Elem()
}


type LookupImportPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupImportPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportPipelineResult)(nil)).Elem()
}

func (o LookupImportPipelineResultOutput) ToLookupImportPipelineResultOutput() LookupImportPipelineResultOutput {
	return o
}

func (o LookupImportPipelineResultOutput) ToLookupImportPipelineResultOutputWithContext(ctx context.Context) LookupImportPipelineResultOutput {
	return o
}

func (o LookupImportPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImportPipelineResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupImportPipelineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupImportPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupImportPipelineResultOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) []string { return v.Options }).(pulumi.StringArrayOutput)
}

func (o LookupImportPipelineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupImportPipelineResultOutput) Source() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) ImportPipelineSourcePropertiesResponse { return v.Source }).(ImportPipelineSourcePropertiesResponseOutput)
}

func (o LookupImportPipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupImportPipelineResultOutput) Trigger() PipelineTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) *PipelineTriggerPropertiesResponse { return v.Trigger }).(PipelineTriggerPropertiesResponsePtrOutput)
}

func (o LookupImportPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportPipelineResultOutput{})
}
