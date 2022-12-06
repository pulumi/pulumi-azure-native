


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvirtualNetworkRetrieve(ctx *pulumi.Context, args *GetvirtualNetworkRetrieveArgs, opts ...pulumi.InvokeOption) (*GetvirtualNetworkRetrieveResult, error) {
	var rv GetvirtualNetworkRetrieveResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20220501preview:getvirtualNetworkRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvirtualNetworkRetrieveArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VirtualNetworksName string `pulumi:"virtualNetworksName"`
}


type GetvirtualNetworkRetrieveResult struct {
	ExtendedLocation *VirtualNetworksResponseExtendedLocation `pulumi:"extendedLocation"`
	Id               string                                   `pulumi:"id"`
	Location         string                                   `pulumi:"location"`
	Name             string                                   `pulumi:"name"`
	Properties       VirtualNetworksPropertiesResponse        `pulumi:"properties"`
	SystemData       SystemDataResponse                       `pulumi:"systemData"`
	Tags             map[string]string                        `pulumi:"tags"`
	Type             string                                   `pulumi:"type"`
}

func GetvirtualNetworkRetrieveOutput(ctx *pulumi.Context, args GetvirtualNetworkRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetvirtualNetworkRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetvirtualNetworkRetrieveResult, error) {
			args := v.(GetvirtualNetworkRetrieveArgs)
			r, err := GetvirtualNetworkRetrieve(ctx, &args, opts...)
			var s GetvirtualNetworkRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetvirtualNetworkRetrieveResultOutput)
}

type GetvirtualNetworkRetrieveOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworksName pulumi.StringInput `pulumi:"virtualNetworksName"`
}

func (GetvirtualNetworkRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualNetworkRetrieveArgs)(nil)).Elem()
}


type GetvirtualNetworkRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetvirtualNetworkRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualNetworkRetrieveResult)(nil)).Elem()
}

func (o GetvirtualNetworkRetrieveResultOutput) ToGetvirtualNetworkRetrieveResultOutput() GetvirtualNetworkRetrieveResultOutput {
	return o
}

func (o GetvirtualNetworkRetrieveResultOutput) ToGetvirtualNetworkRetrieveResultOutputWithContext(ctx context.Context) GetvirtualNetworkRetrieveResultOutput {
	return o
}

func (o GetvirtualNetworkRetrieveResultOutput) ExtendedLocation() VirtualNetworksResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) *VirtualNetworksResponseExtendedLocation {
		return v.ExtendedLocation
	}).(VirtualNetworksResponseExtendedLocationPtrOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Properties() VirtualNetworksPropertiesResponseOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) VirtualNetworksPropertiesResponse { return v.Properties }).(VirtualNetworksPropertiesResponseOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetvirtualNetworkRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualNetworkRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetvirtualNetworkRetrieveResultOutput{})
}
