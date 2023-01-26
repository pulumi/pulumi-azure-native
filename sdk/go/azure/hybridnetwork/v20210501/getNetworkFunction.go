


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkFunction(ctx *pulumi.Context, args *LookupNetworkFunctionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkFunctionResult, error) {
	var rv LookupNetworkFunctionResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20210501:getNetworkFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkFunctionArgs struct {
	NetworkFunctionName string `pulumi:"networkFunctionName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNetworkFunctionResult struct {
	Device                                 *SubResourceResponse                       `pulumi:"device"`
	Etag                                   *string                                    `pulumi:"etag"`
	Id                                     string                                     `pulumi:"id"`
	Location                               string                                     `pulumi:"location"`
	ManagedApplication                     SubResourceResponse                        `pulumi:"managedApplication"`
	ManagedApplicationParameters           interface{}                                `pulumi:"managedApplicationParameters"`
	Name                                   string                                     `pulumi:"name"`
	NetworkFunctionContainerConfigurations interface{}                                `pulumi:"networkFunctionContainerConfigurations"`
	NetworkFunctionUserConfigurations      []NetworkFunctionUserConfigurationResponse `pulumi:"networkFunctionUserConfigurations"`
	ProvisioningState                      string                                     `pulumi:"provisioningState"`
	ServiceKey                             string                                     `pulumi:"serviceKey"`
	SkuName                                *string                                    `pulumi:"skuName"`
	SkuType                                string                                     `pulumi:"skuType"`
	SystemData                             SystemDataResponse                         `pulumi:"systemData"`
	Tags                                   map[string]string                          `pulumi:"tags"`
	Type                                   string                                     `pulumi:"type"`
	VendorName                             *string                                    `pulumi:"vendorName"`
	VendorProvisioningState                string                                     `pulumi:"vendorProvisioningState"`
}

func LookupNetworkFunctionOutput(ctx *pulumi.Context, args LookupNetworkFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkFunctionResult, error) {
			args := v.(LookupNetworkFunctionArgs)
			r, err := LookupNetworkFunction(ctx, &args, opts...)
			var s LookupNetworkFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkFunctionResultOutput)
}

type LookupNetworkFunctionOutputArgs struct {
	NetworkFunctionName pulumi.StringInput `pulumi:"networkFunctionName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionArgs)(nil)).Elem()
}


type LookupNetworkFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionResult)(nil)).Elem()
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutput() LookupNetworkFunctionResultOutput {
	return o
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutputWithContext(ctx context.Context) LookupNetworkFunctionResultOutput {
	return o
}

func (o LookupNetworkFunctionResultOutput) Device() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *SubResourceResponse { return v.Device }).(SubResourceResponsePtrOutput)
}

func (o LookupNetworkFunctionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) ManagedApplication() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) SubResourceResponse { return v.ManagedApplication }).(SubResourceResponseOutput)
}

func (o LookupNetworkFunctionResultOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) interface{} { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

func (o LookupNetworkFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) NetworkFunctionContainerConfigurations() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) interface{} { return v.NetworkFunctionContainerConfigurations }).(pulumi.AnyOutput)
}

func (o LookupNetworkFunctionResultOutput) NetworkFunctionUserConfigurations() NetworkFunctionUserConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) []NetworkFunctionUserConfigurationResponse {
		return v.NetworkFunctionUserConfigurations
	}).(NetworkFunctionUserConfigurationResponseArrayOutput)
}

func (o LookupNetworkFunctionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.ServiceKey }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) SkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.SkuName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkFunctionResultOutput) SkuType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.SkuType }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNetworkFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNetworkFunctionResultOutput) VendorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.VendorName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkFunctionResultOutput) VendorProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.VendorProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkFunctionResultOutput{})
}
