


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetgalleryimageRetrieve(ctx *pulumi.Context, args *GetgalleryimageRetrieveArgs, opts ...pulumi.InvokeOption) (*GetgalleryimageRetrieveResult, error) {
	var rv GetgalleryimageRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210701preview:getgalleryimageRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetgalleryimageRetrieveArgs struct {
	GalleryimagesName string `pulumi:"galleryimagesName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetgalleryimageRetrieveResult struct {
	ContainerName     *string                    `pulumi:"containerName"`
	ExtendedLocation  *ExtendedLocationResponse  `pulumi:"extendedLocation"`
	Id                string                     `pulumi:"id"`
	ImagePath         *string                    `pulumi:"imagePath"`
	Location          string                     `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	OsType            *string                    `pulumi:"osType"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	ResourceName      *string                    `pulumi:"resourceName"`
	Status            GalleryImageStatusResponse `pulumi:"status"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

func GetgalleryimageRetrieveOutput(ctx *pulumi.Context, args GetgalleryimageRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetgalleryimageRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetgalleryimageRetrieveResult, error) {
			args := v.(GetgalleryimageRetrieveArgs)
			r, err := GetgalleryimageRetrieve(ctx, &args, opts...)
			var s GetgalleryimageRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetgalleryimageRetrieveResultOutput)
}

type GetgalleryimageRetrieveOutputArgs struct {
	GalleryimagesName pulumi.StringInput `pulumi:"galleryimagesName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetgalleryimageRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetgalleryimageRetrieveArgs)(nil)).Elem()
}


type GetgalleryimageRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetgalleryimageRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetgalleryimageRetrieveResult)(nil)).Elem()
}

func (o GetgalleryimageRetrieveResultOutput) ToGetgalleryimageRetrieveResultOutput() GetgalleryimageRetrieveResultOutput {
	return o
}

func (o GetgalleryimageRetrieveResultOutput) ToGetgalleryimageRetrieveResultOutputWithContext(ctx context.Context) GetgalleryimageRetrieveResultOutput {
	return o
}

func (o GetgalleryimageRetrieveResultOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o GetgalleryimageRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetgalleryimageRetrieveResultOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) *string { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetgalleryimageRetrieveResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetgalleryimageRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetgalleryimageRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Status() GalleryImageStatusResponseOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) GalleryImageStatusResponse { return v.Status }).(GalleryImageStatusResponseOutput)
}

func (o GetgalleryimageRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetgalleryimageRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetgalleryimageRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetgalleryimageRetrieveResultOutput{})
}
