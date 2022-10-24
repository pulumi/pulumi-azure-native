


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetstorageSpaceRetrieve(ctx *pulumi.Context, args *GetstorageSpaceRetrieveArgs, opts ...pulumi.InvokeOption) (*GetstorageSpaceRetrieveResult, error) {
	var rv GetstorageSpaceRetrieveResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20220501preview:getstorageSpaceRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetstorageSpaceRetrieveArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StorageSpacesName string `pulumi:"storageSpacesName"`
}


type GetstorageSpaceRetrieveResult struct {
	ExtendedLocation *StorageSpacesResponseExtendedLocation `pulumi:"extendedLocation"`
	Id               string                                 `pulumi:"id"`
	Location         string                                 `pulumi:"location"`
	Name             string                                 `pulumi:"name"`
	Properties       StorageSpacesPropertiesResponse        `pulumi:"properties"`
	SystemData       SystemDataResponse                     `pulumi:"systemData"`
	Tags             map[string]string                      `pulumi:"tags"`
	Type             string                                 `pulumi:"type"`
}

func GetstorageSpaceRetrieveOutput(ctx *pulumi.Context, args GetstorageSpaceRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetstorageSpaceRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetstorageSpaceRetrieveResult, error) {
			args := v.(GetstorageSpaceRetrieveArgs)
			r, err := GetstorageSpaceRetrieve(ctx, &args, opts...)
			var s GetstorageSpaceRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetstorageSpaceRetrieveResultOutput)
}

type GetstorageSpaceRetrieveOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageSpacesName pulumi.StringInput `pulumi:"storageSpacesName"`
}

func (GetstorageSpaceRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetstorageSpaceRetrieveArgs)(nil)).Elem()
}


type GetstorageSpaceRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetstorageSpaceRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetstorageSpaceRetrieveResult)(nil)).Elem()
}

func (o GetstorageSpaceRetrieveResultOutput) ToGetstorageSpaceRetrieveResultOutput() GetstorageSpaceRetrieveResultOutput {
	return o
}

func (o GetstorageSpaceRetrieveResultOutput) ToGetstorageSpaceRetrieveResultOutputWithContext(ctx context.Context) GetstorageSpaceRetrieveResultOutput {
	return o
}

func (o GetstorageSpaceRetrieveResultOutput) ExtendedLocation() StorageSpacesResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) *StorageSpacesResponseExtendedLocation {
		return v.ExtendedLocation
	}).(StorageSpacesResponseExtendedLocationPtrOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Properties() StorageSpacesPropertiesResponseOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) StorageSpacesPropertiesResponse { return v.Properties }).(StorageSpacesPropertiesResponseOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetstorageSpaceRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetstorageSpaceRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetstorageSpaceRetrieveResultOutput{})
}
