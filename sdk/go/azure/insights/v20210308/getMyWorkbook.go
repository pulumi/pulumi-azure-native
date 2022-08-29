


package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMyWorkbook(ctx *pulumi.Context, args *LookupMyWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupMyWorkbookResult, error) {
	var rv LookupMyWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20210308:getMyWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMyWorkbookArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMyWorkbookResult struct {
	Category       string                             `pulumi:"category"`
	DisplayName    string                             `pulumi:"displayName"`
	Etag           map[string]string                  `pulumi:"etag"`
	Id             *string                            `pulumi:"id"`
	Identity       *MyWorkbookManagedIdentityResponse `pulumi:"identity"`
	Kind           *string                            `pulumi:"kind"`
	Location       *string                            `pulumi:"location"`
	Name           *string                            `pulumi:"name"`
	SerializedData string                             `pulumi:"serializedData"`
	SourceId       *string                            `pulumi:"sourceId"`
	StorageUri     *string                            `pulumi:"storageUri"`
	SystemData     SystemDataResponse                 `pulumi:"systemData"`
	Tags           map[string]string                  `pulumi:"tags"`
	TimeModified   string                             `pulumi:"timeModified"`
	Type           *string                            `pulumi:"type"`
	UserId         string                             `pulumi:"userId"`
	Version        *string                            `pulumi:"version"`
}

func LookupMyWorkbookOutput(ctx *pulumi.Context, args LookupMyWorkbookOutputArgs, opts ...pulumi.InvokeOption) LookupMyWorkbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMyWorkbookResult, error) {
			args := v.(LookupMyWorkbookArgs)
			r, err := LookupMyWorkbook(ctx, &args, opts...)
			var s LookupMyWorkbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMyWorkbookResultOutput)
}

type LookupMyWorkbookOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMyWorkbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMyWorkbookArgs)(nil)).Elem()
}


type LookupMyWorkbookResultOutput struct{ *pulumi.OutputState }

func (LookupMyWorkbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMyWorkbookResult)(nil)).Elem()
}

func (o LookupMyWorkbookResultOutput) ToLookupMyWorkbookResultOutput() LookupMyWorkbookResultOutput {
	return o
}

func (o LookupMyWorkbookResultOutput) ToLookupMyWorkbookResultOutputWithContext(ctx context.Context) LookupMyWorkbookResultOutput {
	return o
}

func (o LookupMyWorkbookResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupMyWorkbookResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMyWorkbookResultOutput) Etag() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) map[string]string { return v.Etag }).(pulumi.StringMapOutput)
}

func (o LookupMyWorkbookResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) Identity() MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *MyWorkbookManagedIdentityResponse { return v.Identity }).(MyWorkbookManagedIdentityResponsePtrOutput)
}

func (o LookupMyWorkbookResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) string { return v.SerializedData }).(pulumi.StringOutput)
}

func (o LookupMyWorkbookResultOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMyWorkbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMyWorkbookResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupMyWorkbookResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupMyWorkbookResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) string { return v.UserId }).(pulumi.StringOutput)
}

func (o LookupMyWorkbookResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMyWorkbookResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMyWorkbookResultOutput{})
}
