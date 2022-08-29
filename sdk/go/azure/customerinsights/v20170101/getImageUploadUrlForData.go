


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImageUploadUrlForData(ctx *pulumi.Context, args *GetImageUploadUrlForDataArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForDataResult, error) {
	var rv GetImageUploadUrlForDataResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getImageUploadUrlForData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForDataArgs struct {
	EntityType        *string `pulumi:"entityType"`
	EntityTypeName    *string `pulumi:"entityTypeName"`
	HubName           string  `pulumi:"hubName"`
	RelativePath      *string `pulumi:"relativePath"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GetImageUploadUrlForDataResult struct {
	ContentUrl   *string `pulumi:"contentUrl"`
	ImageExists  *bool   `pulumi:"imageExists"`
	RelativePath *string `pulumi:"relativePath"`
}

func GetImageUploadUrlForDataOutput(ctx *pulumi.Context, args GetImageUploadUrlForDataOutputArgs, opts ...pulumi.InvokeOption) GetImageUploadUrlForDataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageUploadUrlForDataResult, error) {
			args := v.(GetImageUploadUrlForDataArgs)
			r, err := GetImageUploadUrlForData(ctx, &args, opts...)
			var s GetImageUploadUrlForDataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageUploadUrlForDataResultOutput)
}

type GetImageUploadUrlForDataOutputArgs struct {
	EntityType        pulumi.StringPtrInput `pulumi:"entityType"`
	EntityTypeName    pulumi.StringPtrInput `pulumi:"entityTypeName"`
	HubName           pulumi.StringInput    `pulumi:"hubName"`
	RelativePath      pulumi.StringPtrInput `pulumi:"relativePath"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (GetImageUploadUrlForDataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForDataArgs)(nil)).Elem()
}


type GetImageUploadUrlForDataResultOutput struct{ *pulumi.OutputState }

func (GetImageUploadUrlForDataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForDataResult)(nil)).Elem()
}

func (o GetImageUploadUrlForDataResultOutput) ToGetImageUploadUrlForDataResultOutput() GetImageUploadUrlForDataResultOutput {
	return o
}

func (o GetImageUploadUrlForDataResultOutput) ToGetImageUploadUrlForDataResultOutputWithContext(ctx context.Context) GetImageUploadUrlForDataResultOutput {
	return o
}

func (o GetImageUploadUrlForDataResultOutput) ContentUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *string { return v.ContentUrl }).(pulumi.StringPtrOutput)
}

func (o GetImageUploadUrlForDataResultOutput) ImageExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *bool { return v.ImageExists }).(pulumi.BoolPtrOutput)
}

func (o GetImageUploadUrlForDataResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageUploadUrlForDataResultOutput{})
}
