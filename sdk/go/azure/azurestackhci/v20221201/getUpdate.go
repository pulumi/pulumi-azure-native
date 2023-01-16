


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUpdate(ctx *pulumi.Context, args *LookupUpdateArgs, opts ...pulumi.InvokeOption) (*LookupUpdateResult, error) {
	var rv LookupUpdateResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221201:getUpdate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UpdateName        string `pulumi:"updateName"`
}


type LookupUpdateResult struct {
	AdditionalProperties *string                      `pulumi:"additionalProperties"`
	AvailabilityType     *string                      `pulumi:"availabilityType"`
	Description          *string                      `pulumi:"description"`
	DisplayName          *string                      `pulumi:"displayName"`
	HealthCheckDate      *string                      `pulumi:"healthCheckDate"`
	Id                   string                       `pulumi:"id"`
	InstalledDate        *string                      `pulumi:"installedDate"`
	Location             *string                      `pulumi:"location"`
	Name                 string                       `pulumi:"name"`
	NotifyMessage        *string                      `pulumi:"notifyMessage"`
	PackagePath          *string                      `pulumi:"packagePath"`
	PackageSizeInMb      *float64                     `pulumi:"packageSizeInMb"`
	PackageType          *string                      `pulumi:"packageType"`
	Prerequisites        []UpdatePrerequisiteResponse `pulumi:"prerequisites"`
	ProgressPercentage   *float64                     `pulumi:"progressPercentage"`
	ProvisioningState    string                       `pulumi:"provisioningState"`
	Publisher            *string                      `pulumi:"publisher"`
	ReleaseLink          *string                      `pulumi:"releaseLink"`
	State                *string                      `pulumi:"state"`
	SystemData           SystemDataResponse           `pulumi:"systemData"`
	Type                 string                       `pulumi:"type"`
	Version              *string                      `pulumi:"version"`
}

func LookupUpdateOutput(ctx *pulumi.Context, args LookupUpdateOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUpdateResult, error) {
			args := v.(LookupUpdateArgs)
			r, err := LookupUpdate(ctx, &args, opts...)
			var s LookupUpdateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUpdateResultOutput)
}

type LookupUpdateOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	UpdateName        pulumi.StringInput `pulumi:"updateName"`
}

func (LookupUpdateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateArgs)(nil)).Elem()
}


type LookupUpdateResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateResult)(nil)).Elem()
}

func (o LookupUpdateResultOutput) ToLookupUpdateResultOutput() LookupUpdateResultOutput {
	return o
}

func (o LookupUpdateResultOutput) ToLookupUpdateResultOutputWithContext(ctx context.Context) LookupUpdateResultOutput {
	return o
}

func (o LookupUpdateResultOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) AvailabilityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.AvailabilityType }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUpdateResultOutput) InstalledDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.InstalledDate }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUpdateResultOutput) NotifyMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.NotifyMessage }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) PackagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.PackagePath }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) PackageSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *float64 { return v.PackageSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o LookupUpdateResultOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.PackageType }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) Prerequisites() UpdatePrerequisiteResponseArrayOutput {
	return o.ApplyT(func(v LookupUpdateResult) []UpdatePrerequisiteResponse { return v.Prerequisites }).(UpdatePrerequisiteResponseArrayOutput)
}

func (o LookupUpdateResultOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *float64 { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

func (o LookupUpdateResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupUpdateResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) ReleaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.ReleaseLink }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupUpdateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUpdateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupUpdateResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateResultOutput{})
}
