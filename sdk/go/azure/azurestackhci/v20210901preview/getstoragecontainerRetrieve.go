


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetstoragecontainerRetrieve(ctx *pulumi.Context, args *GetstoragecontainerRetrieveArgs, opts ...pulumi.InvokeOption) (*GetstoragecontainerRetrieveResult, error) {
	var rv GetstoragecontainerRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getstoragecontainerRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetstoragecontainerRetrieveArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	StoragecontainersName string `pulumi:"storagecontainersName"`
}


type GetstoragecontainerRetrieveResult struct {
	AvailableSizeMB   float64                                    `pulumi:"availableSizeMB"`
	ContainerSizeMB   float64                                    `pulumi:"containerSizeMB"`
	ExtendedLocation  *StoragecontainersResponseExtendedLocation `pulumi:"extendedLocation"`
	Id                string                                     `pulumi:"id"`
	Location          string                                     `pulumi:"location"`
	Name              string                                     `pulumi:"name"`
	Path              *string                                    `pulumi:"path"`
	ProvisioningState *string                                    `pulumi:"provisioningState"`
	ResourceName      *string                                    `pulumi:"resourceName"`
	Status            StorageContainerStatusResponse             `pulumi:"status"`
	SystemData        SystemDataResponse                         `pulumi:"systemData"`
	Tags              map[string]string                          `pulumi:"tags"`
	Type              string                                     `pulumi:"type"`
}

func GetstoragecontainerRetrieveOutput(ctx *pulumi.Context, args GetstoragecontainerRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetstoragecontainerRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetstoragecontainerRetrieveResult, error) {
			args := v.(GetstoragecontainerRetrieveArgs)
			r, err := GetstoragecontainerRetrieve(ctx, &args, opts...)
			var s GetstoragecontainerRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetstoragecontainerRetrieveResultOutput)
}

type GetstoragecontainerRetrieveOutputArgs struct {
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	StoragecontainersName pulumi.StringInput `pulumi:"storagecontainersName"`
}

func (GetstoragecontainerRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetstoragecontainerRetrieveArgs)(nil)).Elem()
}


type GetstoragecontainerRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetstoragecontainerRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetstoragecontainerRetrieveResult)(nil)).Elem()
}

func (o GetstoragecontainerRetrieveResultOutput) ToGetstoragecontainerRetrieveResultOutput() GetstoragecontainerRetrieveResultOutput {
	return o
}

func (o GetstoragecontainerRetrieveResultOutput) ToGetstoragecontainerRetrieveResultOutputWithContext(ctx context.Context) GetstoragecontainerRetrieveResultOutput {
	return o
}

func (o GetstoragecontainerRetrieveResultOutput) AvailableSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) float64 { return v.AvailableSizeMB }).(pulumi.Float64Output)
}

func (o GetstoragecontainerRetrieveResultOutput) ContainerSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) float64 { return v.ContainerSizeMB }).(pulumi.Float64Output)
}

func (o GetstoragecontainerRetrieveResultOutput) ExtendedLocation() StoragecontainersResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) *StoragecontainersResponseExtendedLocation {
		return v.ExtendedLocation
	}).(StoragecontainersResponseExtendedLocationPtrOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Status() StorageContainerStatusResponseOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) StorageContainerStatusResponse { return v.Status }).(StorageContainerStatusResponseOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetstoragecontainerRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetstoragecontainerRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetstoragecontainerRetrieveResultOutput{})
}
