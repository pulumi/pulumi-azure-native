


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSignalR(ctx *pulumi.Context, args *LookupSignalRArgs, opts ...pulumi.InvokeOption) (*LookupSignalRResult, error) {
	var rv LookupSignalRResult
	err := ctx.Invoke("azure-native:signalrservice/v20180301preview:getSignalR", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSignalRArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSignalRResult struct {
	ExternalIP        string               `pulumi:"externalIP"`
	HostName          string               `pulumi:"hostName"`
	HostNamePrefix    *string              `pulumi:"hostNamePrefix"`
	Id                string               `pulumi:"id"`
	Location          *string              `pulumi:"location"`
	Name              string               `pulumi:"name"`
	ProvisioningState string               `pulumi:"provisioningState"`
	PublicPort        int                  `pulumi:"publicPort"`
	ServerPort        int                  `pulumi:"serverPort"`
	Sku               *ResourceSkuResponse `pulumi:"sku"`
	Tags              map[string]string    `pulumi:"tags"`
	Type              string               `pulumi:"type"`
	Version           *string              `pulumi:"version"`
}

func LookupSignalROutput(ctx *pulumi.Context, args LookupSignalROutputArgs, opts ...pulumi.InvokeOption) LookupSignalRResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSignalRResult, error) {
			args := v.(LookupSignalRArgs)
			r, err := LookupSignalR(ctx, &args, opts...)
			var s LookupSignalRResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSignalRResultOutput)
}

type LookupSignalROutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupSignalROutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRArgs)(nil)).Elem()
}


type LookupSignalRResultOutput struct{ *pulumi.OutputState }

func (LookupSignalRResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSignalRResult)(nil)).Elem()
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutput() LookupSignalRResultOutput {
	return o
}

func (o LookupSignalRResultOutput) ToLookupSignalRResultOutputWithContext(ctx context.Context) LookupSignalRResultOutput {
	return o
}

func (o LookupSignalRResultOutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ExternalIP }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) HostNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.HostNamePrefix }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSignalRResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupSignalRResultOutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSignalRResult) int { return v.ServerPort }).(pulumi.IntOutput)
}

func (o LookupSignalRResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupSignalRResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSignalRResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSignalRResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSignalRResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSignalRResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSignalRResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSignalRResultOutput{})
}
