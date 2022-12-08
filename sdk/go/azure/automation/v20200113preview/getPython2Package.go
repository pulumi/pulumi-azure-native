


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPython2Package(ctx *pulumi.Context, args *LookupPython2PackageArgs, opts ...pulumi.InvokeOption) (*LookupPython2PackageResult, error) {
	var rv LookupPython2PackageResult
	err := ctx.Invoke("azure-native:automation/v20200113preview:getPython2Package", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPython2PackageArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	PackageName           string `pulumi:"packageName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupPython2PackageResult struct {
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

func LookupPython2PackageOutput(ctx *pulumi.Context, args LookupPython2PackageOutputArgs, opts ...pulumi.InvokeOption) LookupPython2PackageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPython2PackageResult, error) {
			args := v.(LookupPython2PackageArgs)
			r, err := LookupPython2Package(ctx, &args, opts...)
			var s LookupPython2PackageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPython2PackageResultOutput)
}

type LookupPython2PackageOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	PackageName           pulumi.StringInput `pulumi:"packageName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPython2PackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython2PackageArgs)(nil)).Elem()
}


type LookupPython2PackageResultOutput struct{ *pulumi.OutputState }

func (LookupPython2PackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPython2PackageResult)(nil)).Elem()
}

func (o LookupPython2PackageResultOutput) ToLookupPython2PackageResultOutput() LookupPython2PackageResultOutput {
	return o
}

func (o LookupPython2PackageResultOutput) ToLookupPython2PackageResultOutputWithContext(ctx context.Context) LookupPython2PackageResultOutput {
	return o
}

func (o LookupPython2PackageResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o LookupPython2PackageResultOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupPython2PackageResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o LookupPython2PackageResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPython2PackageResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o LookupPython2PackageResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o LookupPython2PackageResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPython2PackageResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPython2PackageResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupPython2PackageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPython2PackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPython2PackageResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPython2PackageResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPython2PackageResultOutput{})
}
