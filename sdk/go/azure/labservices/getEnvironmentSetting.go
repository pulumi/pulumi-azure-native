


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironmentSetting(ctx *pulumi.Context, args *LookupEnvironmentSettingArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSettingResult, error) {
	var rv LookupEnvironmentSettingResult
	err := ctx.Invoke("azure-native:labservices:getEnvironmentSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSettingArgs struct {
	EnvironmentSettingName string  `pulumi:"environmentSettingName"`
	Expand                 *string `pulumi:"expand"`
	LabAccountName         string  `pulumi:"labAccountName"`
	LabName                string  `pulumi:"labName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentSettingResult struct {
	ConfigurationState    *string                       `pulumi:"configurationState"`
	Description           *string                       `pulumi:"description"`
	Id                    string                        `pulumi:"id"`
	LastChanged           string                        `pulumi:"lastChanged"`
	LastPublished         string                        `pulumi:"lastPublished"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	ProvisioningState     *string                       `pulumi:"provisioningState"`
	PublishingState       string                        `pulumi:"publishingState"`
	ResourceSettings      ResourceSettingsResponse      `pulumi:"resourceSettings"`
	Tags                  map[string]string             `pulumi:"tags"`
	Title                 *string                       `pulumi:"title"`
	Type                  string                        `pulumi:"type"`
	UniqueIdentifier      *string                       `pulumi:"uniqueIdentifier"`
}

func LookupEnvironmentSettingOutput(ctx *pulumi.Context, args LookupEnvironmentSettingOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentSettingResult, error) {
			args := v.(LookupEnvironmentSettingArgs)
			r, err := LookupEnvironmentSetting(ctx, &args, opts...)
			var s LookupEnvironmentSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentSettingResultOutput)
}

type LookupEnvironmentSettingOutputArgs struct {
	EnvironmentSettingName pulumi.StringInput    `pulumi:"environmentSettingName"`
	Expand                 pulumi.StringPtrInput `pulumi:"expand"`
	LabAccountName         pulumi.StringInput    `pulumi:"labAccountName"`
	LabName                pulumi.StringInput    `pulumi:"labName"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupEnvironmentSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSettingArgs)(nil)).Elem()
}


type LookupEnvironmentSettingResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSettingResult)(nil)).Elem()
}

func (o LookupEnvironmentSettingResultOutput) ToLookupEnvironmentSettingResultOutput() LookupEnvironmentSettingResultOutput {
	return o
}

func (o LookupEnvironmentSettingResultOutput) ToLookupEnvironmentSettingResultOutputWithContext(ctx context.Context) LookupEnvironmentSettingResultOutput {
	return o
}

func (o LookupEnvironmentSettingResultOutput) ConfigurationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.ConfigurationState }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentSettingResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) LastChanged() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.LastChanged }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) LastPublished() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.LastPublished }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupEnvironmentSettingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentSettingResultOutput) PublishingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.PublishingState }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) ResourceSettings() ResourceSettingsResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) ResourceSettingsResponse { return v.ResourceSettings }).(ResourceSettingsResponseOutput)
}

func (o LookupEnvironmentSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEnvironmentSettingResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupEnvironmentSettingResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentSettingResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentSettingResultOutput{})
}
