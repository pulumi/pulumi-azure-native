


package v20210309preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMSIXPackage(ctx *pulumi.Context, args *LookupMSIXPackageArgs, opts ...pulumi.InvokeOption) (*LookupMSIXPackageResult, error) {
	var rv LookupMSIXPackageResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210309preview:getMSIXPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMSIXPackageArgs struct {
	HostPoolName        string `pulumi:"hostPoolName"`
	MsixPackageFullName string `pulumi:"msixPackageFullName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupMSIXPackageResult struct {
	DisplayName           *string                           `pulumi:"displayName"`
	Id                    string                            `pulumi:"id"`
	ImagePath             *string                           `pulumi:"imagePath"`
	IsActive              *bool                             `pulumi:"isActive"`
	IsRegularRegistration *bool                             `pulumi:"isRegularRegistration"`
	LastUpdated           *string                           `pulumi:"lastUpdated"`
	Name                  string                            `pulumi:"name"`
	PackageApplications   []MsixPackageApplicationsResponse `pulumi:"packageApplications"`
	PackageDependencies   []MsixPackageDependenciesResponse `pulumi:"packageDependencies"`
	PackageFamilyName     *string                           `pulumi:"packageFamilyName"`
	PackageName           *string                           `pulumi:"packageName"`
	PackageRelativePath   *string                           `pulumi:"packageRelativePath"`
	Type                  string                            `pulumi:"type"`
	Version               *string                           `pulumi:"version"`
}

func LookupMSIXPackageOutput(ctx *pulumi.Context, args LookupMSIXPackageOutputArgs, opts ...pulumi.InvokeOption) LookupMSIXPackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMSIXPackageResult, error) {
			args := v.(LookupMSIXPackageArgs)
			r, err := LookupMSIXPackage(ctx, &args, opts...)
			var s LookupMSIXPackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMSIXPackageResultOutput)
}

type LookupMSIXPackageOutputArgs struct {
	HostPoolName        pulumi.StringInput `pulumi:"hostPoolName"`
	MsixPackageFullName pulumi.StringInput `pulumi:"msixPackageFullName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMSIXPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSIXPackageArgs)(nil)).Elem()
}


type LookupMSIXPackageResultOutput struct{ *pulumi.OutputState }

func (LookupMSIXPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMSIXPackageResult)(nil)).Elem()
}

func (o LookupMSIXPackageResultOutput) ToLookupMSIXPackageResultOutput() LookupMSIXPackageResultOutput {
	return o
}

func (o LookupMSIXPackageResultOutput) ToLookupMSIXPackageResultOutputWithContext(ctx context.Context) LookupMSIXPackageResultOutput {
	return o
}

func (o LookupMSIXPackageResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMSIXPackageResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o LookupMSIXPackageResultOutput) IsRegularRegistration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *bool { return v.IsRegularRegistration }).(pulumi.BoolPtrOutput)
}

func (o LookupMSIXPackageResultOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMSIXPackageResultOutput) PackageApplications() MsixPackageApplicationsResponseArrayOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) []MsixPackageApplicationsResponse { return v.PackageApplications }).(MsixPackageApplicationsResponseArrayOutput)
}

func (o LookupMSIXPackageResultOutput) PackageDependencies() MsixPackageDependenciesResponseArrayOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) []MsixPackageDependenciesResponse { return v.PackageDependencies }).(MsixPackageDependenciesResponseArrayOutput)
}

func (o LookupMSIXPackageResultOutput) PackageFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageFamilyName }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) PackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageName }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) PackageRelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.PackageRelativePath }).(pulumi.StringPtrOutput)
}

func (o LookupMSIXPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMSIXPackageResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMSIXPackageResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMSIXPackageResultOutput{})
}
