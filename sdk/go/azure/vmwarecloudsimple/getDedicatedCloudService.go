


package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedCloudService(ctx *pulumi.Context, args *LookupDedicatedCloudServiceArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedCloudServiceResult, error) {
	var rv LookupDedicatedCloudServiceResult
	err := ctx.Invoke("azure-native:vmwarecloudsimple:getDedicatedCloudService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedCloudServiceArgs struct {
	DedicatedCloudServiceName string `pulumi:"dedicatedCloudServiceName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupDedicatedCloudServiceResult struct {
	GatewaySubnet      string            `pulumi:"gatewaySubnet"`
	Id                 string            `pulumi:"id"`
	IsAccountOnboarded string            `pulumi:"isAccountOnboarded"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	Nodes              int               `pulumi:"nodes"`
	ServiceURL         string            `pulumi:"serviceURL"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
}

func LookupDedicatedCloudServiceOutput(ctx *pulumi.Context, args LookupDedicatedCloudServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedCloudServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedCloudServiceResult, error) {
			args := v.(LookupDedicatedCloudServiceArgs)
			r, err := LookupDedicatedCloudService(ctx, &args, opts...)
			var s LookupDedicatedCloudServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedCloudServiceResultOutput)
}

type LookupDedicatedCloudServiceOutputArgs struct {
	DedicatedCloudServiceName pulumi.StringInput `pulumi:"dedicatedCloudServiceName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedCloudServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedCloudServiceArgs)(nil)).Elem()
}


type LookupDedicatedCloudServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedCloudServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedCloudServiceResult)(nil)).Elem()
}

func (o LookupDedicatedCloudServiceResultOutput) ToLookupDedicatedCloudServiceResultOutput() LookupDedicatedCloudServiceResultOutput {
	return o
}

func (o LookupDedicatedCloudServiceResultOutput) ToLookupDedicatedCloudServiceResultOutputWithContext(ctx context.Context) LookupDedicatedCloudServiceResultOutput {
	return o
}

func (o LookupDedicatedCloudServiceResultOutput) GatewaySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.GatewaySubnet }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) IsAccountOnboarded() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.IsAccountOnboarded }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Nodes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) int { return v.Nodes }).(pulumi.IntOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) ServiceURL() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.ServiceURL }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDedicatedCloudServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedCloudServiceResultOutput{})
}
