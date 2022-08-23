


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupModule(ctx *pulumi.Context, args *LookupModuleArgs, opts ...pulumi.InvokeOption) (*LookupModuleResult, error) {
	var rv LookupModuleResult
	err := ctx.Invoke("azure-native:automation/v20190601:getModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModuleArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ModuleName            string `pulumi:"moduleName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupModuleResult struct {
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

func LookupModuleOutput(ctx *pulumi.Context, args LookupModuleOutputArgs, opts ...pulumi.InvokeOption) LookupModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModuleResult, error) {
			args := v.(LookupModuleArgs)
			r, err := LookupModule(ctx, &args, opts...)
			var s LookupModuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModuleResultOutput)
}

type LookupModuleOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ModuleName            pulumi.StringInput `pulumi:"moduleName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModuleArgs)(nil)).Elem()
}


type LookupModuleResultOutput struct{ *pulumi.OutputState }

func (LookupModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModuleResult)(nil)).Elem()
}

func (o LookupModuleResultOutput) ToLookupModuleResultOutput() LookupModuleResultOutput {
	return o
}

func (o LookupModuleResultOutput) ToLookupModuleResultOutputWithContext(ctx context.Context) LookupModuleResultOutput {
	return o
}

func (o LookupModuleResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

func (o LookupModuleResultOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupModuleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

func (o LookupModuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupModuleResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

func (o LookupModuleResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

func (o LookupModuleResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupModuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupModuleResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupModuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupModuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupModuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupModuleResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModuleResultOutput{})
}
