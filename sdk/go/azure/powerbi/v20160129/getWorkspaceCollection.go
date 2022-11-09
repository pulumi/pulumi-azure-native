


package v20160129

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceCollection(ctx *pulumi.Context, args *LookupWorkspaceCollectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceCollectionResult, error) {
	var rv LookupWorkspaceCollectionResult
	err := ctx.Invoke("azure-native:powerbi/v20160129:getWorkspaceCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceCollectionArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type LookupWorkspaceCollectionResult struct {
	Id         *string           `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *AzureSkuResponse `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}

func LookupWorkspaceCollectionOutput(ctx *pulumi.Context, args LookupWorkspaceCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceCollectionResult, error) {
			args := v.(LookupWorkspaceCollectionArgs)
			r, err := LookupWorkspaceCollection(ctx, &args, opts...)
			var s LookupWorkspaceCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceCollectionResultOutput)
}

type LookupWorkspaceCollectionOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceCollectionName pulumi.StringInput `pulumi:"workspaceCollectionName"`
}

func (LookupWorkspaceCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceCollectionArgs)(nil)).Elem()
}

type LookupWorkspaceCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceCollectionResult)(nil)).Elem()
}

func (o LookupWorkspaceCollectionResultOutput) ToLookupWorkspaceCollectionResultOutput() LookupWorkspaceCollectionResultOutput {
	return o
}

func (o LookupWorkspaceCollectionResultOutput) ToLookupWorkspaceCollectionResultOutputWithContext(ctx context.Context) LookupWorkspaceCollectionResultOutput {
	return o
}

func (o LookupWorkspaceCollectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Sku() AzureSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *AzureSkuResponse { return v.Sku }).(AzureSkuResponsePtrOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceCollectionResultOutput{})
}
