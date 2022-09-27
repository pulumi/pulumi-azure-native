


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBigDataPool(ctx *pulumi.Context, args *LookupBigDataPoolArgs, opts ...pulumi.InvokeOption) (*LookupBigDataPoolResult, error) {
	var rv LookupBigDataPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210501:getBigDataPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBigDataPoolArgs struct {
	BigDataPoolName   string `pulumi:"bigDataPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupBigDataPoolResult struct {
	AutoPause                   *AutoPausePropertiesResponse       `pulumi:"autoPause"`
	AutoScale                   *AutoScalePropertiesResponse       `pulumi:"autoScale"`
	CacheSize                   *int                               `pulumi:"cacheSize"`
	CreationDate                *string                            `pulumi:"creationDate"`
	CustomLibraries             []LibraryInfoResponse              `pulumi:"customLibraries"`
	DefaultSparkLogFolder       *string                            `pulumi:"defaultSparkLogFolder"`
	DynamicExecutorAllocation   *DynamicExecutorAllocationResponse `pulumi:"dynamicExecutorAllocation"`
	Id                          string                             `pulumi:"id"`
	IsComputeIsolationEnabled   *bool                              `pulumi:"isComputeIsolationEnabled"`
	LastSucceededTimestamp      string                             `pulumi:"lastSucceededTimestamp"`
	LibraryRequirements         *LibraryRequirementsResponse       `pulumi:"libraryRequirements"`
	Location                    string                             `pulumi:"location"`
	Name                        string                             `pulumi:"name"`
	NodeCount                   *int                               `pulumi:"nodeCount"`
	NodeSize                    *string                            `pulumi:"nodeSize"`
	NodeSizeFamily              *string                            `pulumi:"nodeSizeFamily"`
	ProvisioningState           *string                            `pulumi:"provisioningState"`
	SessionLevelPackagesEnabled *bool                              `pulumi:"sessionLevelPackagesEnabled"`
	SparkConfigProperties       *LibraryRequirementsResponse       `pulumi:"sparkConfigProperties"`
	SparkEventsFolder           *string                            `pulumi:"sparkEventsFolder"`
	SparkVersion                *string                            `pulumi:"sparkVersion"`
	Tags                        map[string]string                  `pulumi:"tags"`
	Type                        string                             `pulumi:"type"`
}

func LookupBigDataPoolOutput(ctx *pulumi.Context, args LookupBigDataPoolOutputArgs, opts ...pulumi.InvokeOption) LookupBigDataPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBigDataPoolResult, error) {
			args := v.(LookupBigDataPoolArgs)
			r, err := LookupBigDataPool(ctx, &args, opts...)
			var s LookupBigDataPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBigDataPoolResultOutput)
}

type LookupBigDataPoolOutputArgs struct {
	BigDataPoolName   pulumi.StringInput `pulumi:"bigDataPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBigDataPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigDataPoolArgs)(nil)).Elem()
}


type LookupBigDataPoolResultOutput struct{ *pulumi.OutputState }

func (LookupBigDataPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigDataPoolResult)(nil)).Elem()
}

func (o LookupBigDataPoolResultOutput) ToLookupBigDataPoolResultOutput() LookupBigDataPoolResultOutput {
	return o
}

func (o LookupBigDataPoolResultOutput) ToLookupBigDataPoolResultOutputWithContext(ctx context.Context) LookupBigDataPoolResultOutput {
	return o
}

func (o LookupBigDataPoolResultOutput) AutoPause() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *AutoPausePropertiesResponse { return v.AutoPause }).(AutoPausePropertiesResponsePtrOutput)
}

func (o LookupBigDataPoolResultOutput) AutoScale() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *AutoScalePropertiesResponse { return v.AutoScale }).(AutoScalePropertiesResponsePtrOutput)
}

func (o LookupBigDataPoolResultOutput) CacheSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *int { return v.CacheSize }).(pulumi.IntPtrOutput)
}

func (o LookupBigDataPoolResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) CustomLibraries() LibraryInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) []LibraryInfoResponse { return v.CustomLibraries }).(LibraryInfoResponseArrayOutput)
}

func (o LookupBigDataPoolResultOutput) DefaultSparkLogFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.DefaultSparkLogFolder }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) DynamicExecutorAllocation() DynamicExecutorAllocationResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *DynamicExecutorAllocationResponse { return v.DynamicExecutorAllocation }).(DynamicExecutorAllocationResponsePtrOutput)
}

func (o LookupBigDataPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBigDataPoolResultOutput) IsComputeIsolationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *bool { return v.IsComputeIsolationEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupBigDataPoolResultOutput) LastSucceededTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.LastSucceededTimestamp }).(pulumi.StringOutput)
}

func (o LookupBigDataPoolResultOutput) LibraryRequirements() LibraryRequirementsResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *LibraryRequirementsResponse { return v.LibraryRequirements }).(LibraryRequirementsResponsePtrOutput)
}

func (o LookupBigDataPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBigDataPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBigDataPoolResultOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o LookupBigDataPoolResultOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.NodeSize }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) SessionLevelPackagesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *bool { return v.SessionLevelPackagesEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupBigDataPoolResultOutput) SparkConfigProperties() LibraryRequirementsResponsePtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *LibraryRequirementsResponse { return v.SparkConfigProperties }).(LibraryRequirementsResponsePtrOutput)
}

func (o LookupBigDataPoolResultOutput) SparkEventsFolder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.SparkEventsFolder }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) *string { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

func (o LookupBigDataPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBigDataPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigDataPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBigDataPoolResultOutput{})
}
