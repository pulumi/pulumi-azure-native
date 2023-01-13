


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFrontendsInterface(ctx *pulumi.Context, args *LookupFrontendsInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupFrontendsInterfaceResult, error) {
	var rv LookupFrontendsInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking/v20221001preview:getFrontendsInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFrontendsInterfaceArgs struct {
	FrontendName          string `pulumi:"frontendName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TrafficControllerName string `pulumi:"trafficControllerName"`
}


type LookupFrontendsInterfaceResult struct {
	Id                string                               `pulumi:"id"`
	IpAddressVersion  *string                              `pulumi:"ipAddressVersion"`
	Location          string                               `pulumi:"location"`
	Mode              *string                              `pulumi:"mode"`
	Name              string                               `pulumi:"name"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	PublicIPAddress   *FrontendPropertiesIPAddressResponse `pulumi:"publicIPAddress"`
	SystemData        SystemDataResponse                   `pulumi:"systemData"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}

func LookupFrontendsInterfaceOutput(ctx *pulumi.Context, args LookupFrontendsInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupFrontendsInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFrontendsInterfaceResult, error) {
			args := v.(LookupFrontendsInterfaceArgs)
			r, err := LookupFrontendsInterface(ctx, &args, opts...)
			var s LookupFrontendsInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFrontendsInterfaceResultOutput)
}

type LookupFrontendsInterfaceOutputArgs struct {
	FrontendName          pulumi.StringInput `pulumi:"frontendName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupFrontendsInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendsInterfaceArgs)(nil)).Elem()
}


type LookupFrontendsInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupFrontendsInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendsInterfaceResult)(nil)).Elem()
}

func (o LookupFrontendsInterfaceResultOutput) ToLookupFrontendsInterfaceResultOutput() LookupFrontendsInterfaceResultOutput {
	return o
}

func (o LookupFrontendsInterfaceResultOutput) ToLookupFrontendsInterfaceResultOutputWithContext(ctx context.Context) LookupFrontendsInterfaceResultOutput {
	return o
}

func (o LookupFrontendsInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFrontendsInterfaceResultOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *string { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

func (o LookupFrontendsInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupFrontendsInterfaceResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupFrontendsInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFrontendsInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFrontendsInterfaceResultOutput) PublicIPAddress() FrontendPropertiesIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *FrontendPropertiesIPAddressResponse { return v.PublicIPAddress }).(FrontendPropertiesIPAddressResponsePtrOutput)
}

func (o LookupFrontendsInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupFrontendsInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFrontendsInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFrontendsInterfaceResultOutput{})
}
