


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvirtualnetworkRetrieve(ctx *pulumi.Context, args *GetvirtualnetworkRetrieveArgs, opts ...pulumi.InvokeOption) (*GetvirtualnetworkRetrieveResult, error) {
	var rv GetvirtualnetworkRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210901preview:getvirtualnetworkRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvirtualnetworkRetrieveArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VirtualnetworksName string `pulumi:"virtualnetworksName"`
}


type GetvirtualnetworkRetrieveResult struct {
	ExtendedLocation  *ExtendedLocationResponse                  `pulumi:"extendedLocation"`
	Id                string                                     `pulumi:"id"`
	Location          string                                     `pulumi:"location"`
	Name              string                                     `pulumi:"name"`
	NetworkType       *string                                    `pulumi:"networkType"`
	ProvisioningState string                                     `pulumi:"provisioningState"`
	ResourceName      *string                                    `pulumi:"resourceName"`
	Status            VirtualNetworkStatusResponse               `pulumi:"status"`
	Subnets           []VirtualnetworksPropertiesResponseSubnets `pulumi:"subnets"`
	SystemData        SystemDataResponse                         `pulumi:"systemData"`
	Tags              map[string]string                          `pulumi:"tags"`
	Type              string                                     `pulumi:"type"`
}

func GetvirtualnetworkRetrieveOutput(ctx *pulumi.Context, args GetvirtualnetworkRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetvirtualnetworkRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetvirtualnetworkRetrieveResult, error) {
			args := v.(GetvirtualnetworkRetrieveArgs)
			r, err := GetvirtualnetworkRetrieve(ctx, &args, opts...)
			var s GetvirtualnetworkRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetvirtualnetworkRetrieveResultOutput)
}

type GetvirtualnetworkRetrieveOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualnetworksName pulumi.StringInput `pulumi:"virtualnetworksName"`
}

func (GetvirtualnetworkRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualnetworkRetrieveArgs)(nil)).Elem()
}


type GetvirtualnetworkRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetvirtualnetworkRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualnetworkRetrieveResult)(nil)).Elem()
}

func (o GetvirtualnetworkRetrieveResultOutput) ToGetvirtualnetworkRetrieveResultOutput() GetvirtualnetworkRetrieveResultOutput {
	return o
}

func (o GetvirtualnetworkRetrieveResultOutput) ToGetvirtualnetworkRetrieveResultOutputWithContext(ctx context.Context) GetvirtualnetworkRetrieveResultOutput {
	return o
}

func (o GetvirtualnetworkRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Status() VirtualNetworkStatusResponseOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) VirtualNetworkStatusResponse { return v.Status }).(VirtualNetworkStatusResponseOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Subnets() VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) []VirtualnetworksPropertiesResponseSubnets { return v.Subnets }).(VirtualnetworksPropertiesResponseSubnetsArrayOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetvirtualnetworkRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualnetworkRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetvirtualnetworkRetrieveResultOutput{})
}
