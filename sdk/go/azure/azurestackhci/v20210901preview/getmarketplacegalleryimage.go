


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Getmarketplacegalleryimage(ctx *pulumi.Context, args *GetmarketplacegalleryimageArgs, opts ...pulumi.InvokeOption) (*GetmarketplacegalleryimageResult, error) {
	var rv GetmarketplacegalleryimageResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getmarketplacegalleryimage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetmarketplacegalleryimageArgs struct {
	MarketplacegalleryimagesName string `pulumi:"marketplacegalleryimagesName"`
	ResourceGroupName            string `pulumi:"resourceGroupName"`
}


type GetmarketplacegalleryimageResult struct {
	CloudInitDataSource *string                               `pulumi:"cloudInitDataSource"`
	ContainerName       *string                               `pulumi:"containerName"`
	ExtendedLocation    *ExtendedLocationResponse             `pulumi:"extendedLocation"`
	HyperVGeneration    *string                               `pulumi:"hyperVGeneration"`
	Id                  string                                `pulumi:"id"`
	Identifier          *GalleryImageIdentifierResponse       `pulumi:"identifier"`
	Location            string                                `pulumi:"location"`
	Name                string                                `pulumi:"name"`
	OsType              *string                               `pulumi:"osType"`
	ProvisioningState   string                                `pulumi:"provisioningState"`
	ResourceName        *string                               `pulumi:"resourceName"`
	Status              MarketplaceGalleryImageStatusResponse `pulumi:"status"`
	SystemData          SystemDataResponse                    `pulumi:"systemData"`
	Tags                map[string]string                     `pulumi:"tags"`
	Type                string                                `pulumi:"type"`
	Version             *GalleryImageVersionResponse          `pulumi:"version"`
}

func GetmarketplacegalleryimageOutput(ctx *pulumi.Context, args GetmarketplacegalleryimageOutputArgs, opts ...pulumi.InvokeOption) GetmarketplacegalleryimageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetmarketplacegalleryimageResult, error) {
			args := v.(GetmarketplacegalleryimageArgs)
			r, err := Getmarketplacegalleryimage(ctx, &args, opts...)
			var s GetmarketplacegalleryimageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetmarketplacegalleryimageResultOutput)
}

type GetmarketplacegalleryimageOutputArgs struct {
	MarketplacegalleryimagesName pulumi.StringInput `pulumi:"marketplacegalleryimagesName"`
	ResourceGroupName            pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetmarketplacegalleryimageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmarketplacegalleryimageArgs)(nil)).Elem()
}


type GetmarketplacegalleryimageResultOutput struct{ *pulumi.OutputState }

func (GetmarketplacegalleryimageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmarketplacegalleryimageResult)(nil)).Elem()
}

func (o GetmarketplacegalleryimageResultOutput) ToGetmarketplacegalleryimageResultOutput() GetmarketplacegalleryimageResultOutput {
	return o
}

func (o GetmarketplacegalleryimageResultOutput) ToGetmarketplacegalleryimageResultOutputWithContext(ctx context.Context) GetmarketplacegalleryimageResultOutput {
	return o
}

func (o GetmarketplacegalleryimageResultOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *string { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *GalleryImageIdentifierResponse { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetmarketplacegalleryimageResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetmarketplacegalleryimageResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Status() MarketplaceGalleryImageStatusResponseOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) MarketplaceGalleryImageStatusResponse { return v.Status }).(MarketplaceGalleryImageStatusResponseOutput)
}

func (o GetmarketplacegalleryimageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetmarketplacegalleryimageResultOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v GetmarketplacegalleryimageResult) *GalleryImageVersionResponse { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetmarketplacegalleryimageResultOutput{})
}
