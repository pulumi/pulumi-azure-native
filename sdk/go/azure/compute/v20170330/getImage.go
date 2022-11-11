


package v20170330

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("azure-native:compute/v20170330:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImageArgs struct {
	Expand            *string `pulumi:"expand"`
	ImageName         string  `pulumi:"imageName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupImageResult struct {
	Id                   string                       `pulumi:"id"`
	Location             string                       `pulumi:"location"`
	Name                 string                       `pulumi:"name"`
	ProvisioningState    string                       `pulumi:"provisioningState"`
	SourceVirtualMachine *SubResourceResponse         `pulumi:"sourceVirtualMachine"`
	StorageProfile       *ImageStorageProfileResponse `pulumi:"storageProfile"`
	Tags                 map[string]string            `pulumi:"tags"`
	Type                 string                       `pulumi:"type"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

type LookupImageOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ImageName         pulumi.StringInput    `pulumi:"imageName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}


type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) SourceVirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupImageResult) *SubResourceResponse { return v.SourceVirtualMachine }).(SubResourceResponsePtrOutput)
}

func (o LookupImageResultOutput) StorageProfile() ImageStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupImageResult) *ImageStorageProfileResponse { return v.StorageProfile }).(ImageStorageProfileResponsePtrOutput)
}

func (o LookupImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
