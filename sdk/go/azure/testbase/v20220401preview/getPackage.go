


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPackage(ctx *pulumi.Context, args *LookupPackageArgs, opts ...pulumi.InvokeOption) (*LookupPackageResult, error) {
	var rv LookupPackageResult
	err := ctx.Invoke("azure-native:testbase/v20220401preview:getPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPackageArgs struct {
	PackageName         string `pulumi:"packageName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupPackageResult struct {
	ApplicationName   string                            `pulumi:"applicationName"`
	BlobPath          string                            `pulumi:"blobPath"`
	Etag              string                            `pulumi:"etag"`
	FlightingRing     string                            `pulumi:"flightingRing"`
	Id                string                            `pulumi:"id"`
	IsEnabled         bool                              `pulumi:"isEnabled"`
	LastModifiedTime  string                            `pulumi:"lastModifiedTime"`
	Location          string                            `pulumi:"location"`
	Name              string                            `pulumi:"name"`
	PackageStatus     string                            `pulumi:"packageStatus"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                `pulumi:"systemData"`
	Tags              map[string]string                 `pulumi:"tags"`
	TargetOSList      []TargetOSInfoResponse            `pulumi:"targetOSList"`
	TestTypes         []string                          `pulumi:"testTypes"`
	Tests             []TestResponse                    `pulumi:"tests"`
	Type              string                            `pulumi:"type"`
	ValidationResults []PackageValidationResultResponse `pulumi:"validationResults"`
	Version           string                            `pulumi:"version"`
}

func LookupPackageOutput(ctx *pulumi.Context, args LookupPackageOutputArgs, opts ...pulumi.InvokeOption) LookupPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPackageResult, error) {
			args := v.(LookupPackageArgs)
			r, err := LookupPackage(ctx, &args, opts...)
			var s LookupPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPackageResultOutput)
}

type LookupPackageOutputArgs struct {
	PackageName         pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackageArgs)(nil)).Elem()
}


type LookupPackageResultOutput struct{ *pulumi.OutputState }

func (LookupPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPackageResult)(nil)).Elem()
}

func (o LookupPackageResultOutput) ToLookupPackageResultOutput() LookupPackageResultOutput {
	return o
}

func (o LookupPackageResultOutput) ToLookupPackageResultOutputWithContext(ctx context.Context) LookupPackageResultOutput {
	return o
}

func (o LookupPackageResultOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) BlobPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.BlobPath }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) FlightingRing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.FlightingRing }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPackageResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupPackageResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) PackageStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.PackageStatus }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPackageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPackageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPackageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPackageResultOutput) TargetOSList() TargetOSInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupPackageResult) []TargetOSInfoResponse { return v.TargetOSList }).(TargetOSInfoResponseArrayOutput)
}

func (o LookupPackageResultOutput) TestTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPackageResult) []string { return v.TestTypes }).(pulumi.StringArrayOutput)
}

func (o LookupPackageResultOutput) Tests() TestResponseArrayOutput {
	return o.ApplyT(func(v LookupPackageResult) []TestResponse { return v.Tests }).(TestResponseArrayOutput)
}

func (o LookupPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPackageResultOutput) ValidationResults() PackageValidationResultResponseArrayOutput {
	return o.ApplyT(func(v LookupPackageResult) []PackageValidationResultResponse { return v.ValidationResults }).(PackageValidationResultResponseArrayOutput)
}

func (o LookupPackageResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPackageResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPackageResultOutput{})
}
