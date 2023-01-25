


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotSecuritySolution(ctx *pulumi.Context, args *LookupIotSecuritySolutionArgs, opts ...pulumi.InvokeOption) (*LookupIotSecuritySolutionResult, error) {
	var rv LookupIotSecuritySolutionResult
	err := ctx.Invoke("azure-native:security/v20170801preview:getIotSecuritySolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIotSecuritySolutionArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SolutionName      string `pulumi:"solutionName"`
}


type LookupIotSecuritySolutionResult struct {
	AutoDiscoveredResources      []string                                        `pulumi:"autoDiscoveredResources"`
	DisabledDataSources          []string                                        `pulumi:"disabledDataSources"`
	DisplayName                  string                                          `pulumi:"displayName"`
	Export                       []string                                        `pulumi:"export"`
	Id                           string                                          `pulumi:"id"`
	IotHubs                      []string                                        `pulumi:"iotHubs"`
	Location                     *string                                         `pulumi:"location"`
	Name                         string                                          `pulumi:"name"`
	RecommendationsConfiguration []RecommendationConfigurationPropertiesResponse `pulumi:"recommendationsConfiguration"`
	Status                       *string                                         `pulumi:"status"`
	Tags                         map[string]string                               `pulumi:"tags"`
	Type                         string                                          `pulumi:"type"`
	UserDefinedResources         *UserDefinedResourcesPropertiesResponse         `pulumi:"userDefinedResources"`
	Workspace                    string                                          `pulumi:"workspace"`
}


func (val *LookupIotSecuritySolutionResult) Defaults() *LookupIotSecuritySolutionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Enabled"
		tmp.Status = &status_
	}
	return &tmp
}

func LookupIotSecuritySolutionOutput(ctx *pulumi.Context, args LookupIotSecuritySolutionOutputArgs, opts ...pulumi.InvokeOption) LookupIotSecuritySolutionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotSecuritySolutionResult, error) {
			args := v.(LookupIotSecuritySolutionArgs)
			r, err := LookupIotSecuritySolution(ctx, &args, opts...)
			var s LookupIotSecuritySolutionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotSecuritySolutionResultOutput)
}

type LookupIotSecuritySolutionOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SolutionName      pulumi.StringInput `pulumi:"solutionName"`
}

func (LookupIotSecuritySolutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotSecuritySolutionArgs)(nil)).Elem()
}


type LookupIotSecuritySolutionResultOutput struct{ *pulumi.OutputState }

func (LookupIotSecuritySolutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotSecuritySolutionResult)(nil)).Elem()
}

func (o LookupIotSecuritySolutionResultOutput) ToLookupIotSecuritySolutionResultOutput() LookupIotSecuritySolutionResultOutput {
	return o
}

func (o LookupIotSecuritySolutionResultOutput) ToLookupIotSecuritySolutionResultOutputWithContext(ctx context.Context) LookupIotSecuritySolutionResultOutput {
	return o
}

func (o LookupIotSecuritySolutionResultOutput) AutoDiscoveredResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) []string { return v.AutoDiscoveredResources }).(pulumi.StringArrayOutput)
}

func (o LookupIotSecuritySolutionResultOutput) DisabledDataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) []string { return v.DisabledDataSources }).(pulumi.StringArrayOutput)
}

func (o LookupIotSecuritySolutionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Export() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) []string { return v.Export }).(pulumi.StringArrayOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotSecuritySolutionResultOutput) IotHubs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) []string { return v.IotHubs }).(pulumi.StringArrayOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotSecuritySolutionResultOutput) RecommendationsConfiguration() RecommendationConfigurationPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) []RecommendationConfigurationPropertiesResponse {
		return v.RecommendationsConfiguration
	}).(RecommendationConfigurationPropertiesResponseArrayOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupIotSecuritySolutionResultOutput) UserDefinedResources() UserDefinedResourcesPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) *UserDefinedResourcesPropertiesResponse {
		return v.UserDefinedResources
	}).(UserDefinedResourcesPropertiesResponsePtrOutput)
}

func (o LookupIotSecuritySolutionResultOutput) Workspace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotSecuritySolutionResult) string { return v.Workspace }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotSecuritySolutionResultOutput{})
}
