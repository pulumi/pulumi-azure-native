


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPython3Package(ctx *pulumi.Context, args *LookupPython3PackageArgs, opts ...pulumi.InvokeOption) (*LookupPython3PackageResult, error) {
	var rv LookupPython3PackageResult
	err := ctx.Invoke("azure-native:automation/v20220808:getPython3Package", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPython3PackageArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	PackageName           string `pulumi:"packageName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupPython3PackageResult struct {
	ActivityCount     *int                     `pulumi:"activityCount"`
	ContentLink       *ContentLinkResponse     `pulumi:"contentLink"`
	CreationTime      *string                  `pulumi:"creationTime"`
	Description       *string                  `pulumi:"description"`
	Error             *ModuleErrorInfoResponse `pulumi:"error"`
	Etag              *string                  `pulumi:"etag"`
	Id                string                   `pulumi:"id"`
	IsComposite       *bool                    `pulumi:"isComposite"`
	IsGlobal          *bool                    `pulumi:"isGlobal"`
	LastModifiedTime  *string                  `pulumi:"lastModifiedTime"`
	Location          *string                  `pulumi:"location"`
	Name              string                   `pulumi:"name"`
	ProvisioningState *string                  `pulumi:"provisioningState"`
	SizeInBytes       *float64                 `pulumi:"sizeInBytes"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	Version           *string                  `pulumi:"version"`
}

func LookupPython3PackageOutput(ctx *pulumi.Context, args LookupPython3PackageOutputArgs, opts ...pulumi.InvokeOption) LookupPython3PackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPython3PackageResult, error) {
			args := v.(LookupPython3PackageArgs)
			r, err := LookupPython3Package(ctx, &args, opts...)
			var s LookupPython3PackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPython3PackageResultOutput)
}

type LookupPython3PackageOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	PackageName           pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPython3PackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython3PackageArgs)(nil)).Elem()
}


type LookupPython3PackageResultOutput struct{ *pulumi.OutputState }

func (LookupPython3PackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython3PackageResult)(nil)).Elem()
}

func (o LookupPython3PackageResultOutput) ToLookupPython3PackageResultOutput() LookupPython3PackageResultOutput {
	return o
}

func (o LookupPython3PackageResultOutput) ToLookupPython3PackageResultOutputWithContext(ctx context.Context) LookupPython3PackageResultOutput {
	return o
}

func (o LookupPython3PackageResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o LookupPython3PackageResultOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupPython3PackageResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o LookupPython3PackageResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPython3PackageResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o LookupPython3PackageResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o LookupPython3PackageResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPython3PackageResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPython3PackageResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupPython3PackageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPython3PackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPython3PackageResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython3PackageResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPython3PackageResultOutput{})
}
