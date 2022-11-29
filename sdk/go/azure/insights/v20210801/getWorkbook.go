


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20210801:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookArgs struct {
	CanFetchContent   *bool  `pulumi:"canFetchContent"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookResult struct {
	Category       string                            `pulumi:"category"`
	Description    *string                           `pulumi:"description"`
	DisplayName    string                            `pulumi:"displayName"`
	Etag           map[string]string                 `pulumi:"etag"`
	Id             string                            `pulumi:"id"`
	Identity       *WorkbookResourceResponseIdentity `pulumi:"identity"`
	Kind           *string                           `pulumi:"kind"`
	Location       string                            `pulumi:"location"`
	Name           string                            `pulumi:"name"`
	Revision       string                            `pulumi:"revision"`
	SerializedData string                            `pulumi:"serializedData"`
	SourceId       *string                           `pulumi:"sourceId"`
	StorageUri     *string                           `pulumi:"storageUri"`
	SystemData     SystemDataResponse                `pulumi:"systemData"`
	Tags           map[string]string                 `pulumi:"tags"`
	TimeModified   string                            `pulumi:"timeModified"`
	Type           string                            `pulumi:"type"`
	UserId         string                            `pulumi:"userId"`
	Version        *string                           `pulumi:"version"`
}

func LookupWorkbookOutput(ctx *pulumi.Context, args LookupWorkbookOutputArgs, opts ...pulumi.InvokeOption) LookupWorkbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkbookResult, error) {
			args := v.(LookupWorkbookArgs)
			r, err := LookupWorkbook(ctx, &args, opts...)
			var s LookupWorkbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkbookResultOutput)
}

type LookupWorkbookOutputArgs struct {
	CanFetchContent   pulumi.BoolPtrInput `pulumi:"canFetchContent"`
	ResourceGroupName pulumi.StringInput  `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput  `pulumi:"resourceName"`
}

func (LookupWorkbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookArgs)(nil)).Elem()
}


type LookupWorkbookResultOutput struct{ *pulumi.OutputState }

func (LookupWorkbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookResult)(nil)).Elem()
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutput() LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutputWithContext(ctx context.Context) LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Etag() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Etag }).(pulumi.StringMapOutput)
}

func (o LookupWorkbookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Identity() WorkbookResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *WorkbookResourceResponseIdentity { return v.Identity }).(WorkbookResourceResponseIdentityPtrOutput)
}

func (o LookupWorkbookResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Revision }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.SerializedData }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkbookResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWorkbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkbookResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.UserId }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkbookResultOutput{})
}
