


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetguestDiagnosticsSetting(ctx *pulumi.Context, args *GetguestDiagnosticsSettingArgs, opts ...pulumi.InvokeOption) (*GetguestDiagnosticsSettingResult, error) {
	var rv GetguestDiagnosticsSettingResult
	err := ctx.Invoke("azure-native:insights:getguestDiagnosticsSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetguestDiagnosticsSettingArgs struct {
	DiagnosticSettingsName string `pulumi:"diagnosticSettingsName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type GetguestDiagnosticsSettingResult struct {
	DataSources  []DataSourceResponse `pulumi:"dataSources"`
	Id           string               `pulumi:"id"`
	Location     string               `pulumi:"location"`
	Name         string               `pulumi:"name"`
	OsType       *string              `pulumi:"osType"`
	ProxySetting *string              `pulumi:"proxySetting"`
	Tags         map[string]string    `pulumi:"tags"`
	Type         string               `pulumi:"type"`
}

func GetguestDiagnosticsSettingOutput(ctx *pulumi.Context, args GetguestDiagnosticsSettingOutputArgs, opts ...pulumi.InvokeOption) GetguestDiagnosticsSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetguestDiagnosticsSettingResult, error) {
			args := v.(GetguestDiagnosticsSettingArgs)
			r, err := GetguestDiagnosticsSetting(ctx, &args, opts...)
			var s GetguestDiagnosticsSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetguestDiagnosticsSettingResultOutput)
}

type GetguestDiagnosticsSettingOutputArgs struct {
	DiagnosticSettingsName pulumi.StringInput `pulumi:"diagnosticSettingsName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetguestDiagnosticsSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetguestDiagnosticsSettingArgs)(nil)).Elem()
}


type GetguestDiagnosticsSettingResultOutput struct{ *pulumi.OutputState }

func (GetguestDiagnosticsSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetguestDiagnosticsSettingResult)(nil)).Elem()
}

func (o GetguestDiagnosticsSettingResultOutput) ToGetguestDiagnosticsSettingResultOutput() GetguestDiagnosticsSettingResultOutput {
	return o
}

func (o GetguestDiagnosticsSettingResultOutput) ToGetguestDiagnosticsSettingResultOutputWithContext(ctx context.Context) GetguestDiagnosticsSettingResultOutput {
	return o
}

func (o GetguestDiagnosticsSettingResultOutput) DataSources() DataSourceResponseArrayOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) []DataSourceResponse { return v.DataSources }).(DataSourceResponseArrayOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) ProxySetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) *string { return v.ProxySetting }).(pulumi.StringPtrOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetguestDiagnosticsSettingResultOutput{})
}
