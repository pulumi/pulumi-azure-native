


package v20220210preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20220210preview:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	ApplicationName      string `pulumi:"applicationName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupApplicationResult struct {
	ApplicationType          *string            `pulumi:"applicationType"`
	CommandLineArguments     *string            `pulumi:"commandLineArguments"`
	CommandLineSetting       string             `pulumi:"commandLineSetting"`
	Description              *string            `pulumi:"description"`
	FilePath                 *string            `pulumi:"filePath"`
	FriendlyName             *string            `pulumi:"friendlyName"`
	IconContent              string             `pulumi:"iconContent"`
	IconHash                 string             `pulumi:"iconHash"`
	IconIndex                *int               `pulumi:"iconIndex"`
	IconPath                 *string            `pulumi:"iconPath"`
	Id                       string             `pulumi:"id"`
	MsixPackageApplicationId *string            `pulumi:"msixPackageApplicationId"`
	MsixPackageFamilyName    *string            `pulumi:"msixPackageFamilyName"`
	Name                     string             `pulumi:"name"`
	ObjectId                 string             `pulumi:"objectId"`
	ShowInPortal             *bool              `pulumi:"showInPortal"`
	SystemData               SystemDataResponse `pulumi:"systemData"`
	Type                     string             `pulumi:"type"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			var s LookupApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	ApplicationGroupName pulumi.StringInput `pulumi:"applicationGroupName"`
	ApplicationName      pulumi.StringInput `pulumi:"applicationName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}


type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) CommandLineArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.CommandLineArguments }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) CommandLineSetting() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.CommandLineSetting }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) IconContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.IconContent }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) IconHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.IconHash }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) IconIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *int { return v.IconIndex }).(pulumi.IntPtrOutput)
}

func (o LookupApplicationResultOutput) IconPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.IconPath }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) MsixPackageApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.MsixPackageApplicationId }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) MsixPackageFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.MsixPackageFamilyName }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) ShowInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *bool { return v.ShowInPortal }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
