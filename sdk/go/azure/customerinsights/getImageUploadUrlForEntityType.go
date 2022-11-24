


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetImageUploadUrlForEntityType(ctx *pulumi.Context, args *GetImageUploadUrlForEntityTypeArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForEntityTypeResult, error) {
	var rv GetImageUploadUrlForEntityTypeResult
	err := ctx.Invoke("azure-native:customerinsights:getImageUploadUrlForEntityType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForEntityTypeArgs struct {
	EntityType        *string `pulumi:"entityType"`
	EntityTypeName    *string `pulumi:"entityTypeName"`
	HubName           string  `pulumi:"hubName"`
	RelativePath      *string `pulumi:"relativePath"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GetImageUploadUrlForEntityTypeResult struct {
	ContentUrl   *string `pulumi:"contentUrl"`
	ImageExists  *bool   `pulumi:"imageExists"`
	RelativePath *string `pulumi:"relativePath"`
}

func GetImageUploadUrlForEntityTypeOutput(ctx *pulumi.Context, args GetImageUploadUrlForEntityTypeOutputArgs, opts ...pulumi.InvokeOption) GetImageUploadUrlForEntityTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageUploadUrlForEntityTypeResult, error) {
			args := v.(GetImageUploadUrlForEntityTypeArgs)
			r, err := GetImageUploadUrlForEntityType(ctx, &args, opts...)
			var s GetImageUploadUrlForEntityTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageUploadUrlForEntityTypeResultOutput)
}

type GetImageUploadUrlForEntityTypeOutputArgs struct {
	EntityType        pulumi.StringPtrInput `pulumi:"entityType"`
	EntityTypeName    pulumi.StringPtrInput `pulumi:"entityTypeName"`
	HubName           pulumi.StringInput    `pulumi:"hubName"`
	RelativePath      pulumi.StringPtrInput `pulumi:"relativePath"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (GetImageUploadUrlForEntityTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForEntityTypeArgs)(nil)).Elem()
}


type GetImageUploadUrlForEntityTypeResultOutput struct{ *pulumi.OutputState }

func (GetImageUploadUrlForEntityTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForEntityTypeResult)(nil)).Elem()
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ToGetImageUploadUrlForEntityTypeResultOutput() GetImageUploadUrlForEntityTypeResultOutput {
	return o
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ToGetImageUploadUrlForEntityTypeResultOutputWithContext(ctx context.Context) GetImageUploadUrlForEntityTypeResultOutput {
	return o
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ContentUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *string { return v.ContentUrl }).(pulumi.StringPtrOutput)
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ImageExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *bool { return v.ImageExists }).(pulumi.BoolPtrOutput)
}

func (o GetImageUploadUrlForEntityTypeResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageUploadUrlForEntityTypeResultOutput{})
}
