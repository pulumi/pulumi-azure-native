


package v20160402

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupOrigin(ctx *pulumi.Context, args *LookupOriginArgs, opts ...pulumi.InvokeOption) (*LookupOriginResult, error) {
	var rv LookupOriginResult
	err := ctx.Invoke("azure-native:cdn/v20160402:getOrigin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginName        string `pulumi:"originName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginResult struct {
	HostName          string `pulumi:"hostName"`
	HttpPort          *int   `pulumi:"httpPort"`
	HttpsPort         *int   `pulumi:"httpsPort"`
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ProvisioningState string `pulumi:"provisioningState"`
	ResourceState     string `pulumi:"resourceState"`
	Type              string `pulumi:"type"`
}

func LookupOriginOutput(ctx *pulumi.Context, args LookupOriginOutputArgs, opts ...pulumi.InvokeOption) LookupOriginResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOriginResult, error) {
			args := v.(LookupOriginArgs)
			r, err := LookupOrigin(ctx, &args, opts...)
			var s LookupOriginResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOriginResultOutput)
}

type LookupOriginOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	OriginName        pulumi.StringInput `pulumi:"originName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOriginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginArgs)(nil)).Elem()
}


type LookupOriginResultOutput struct{ *pulumi.OutputState }

func (LookupOriginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginResult)(nil)).Elem()
}

func (o LookupOriginResultOutput) ToLookupOriginResultOutput() LookupOriginResultOutput {
	return o
}

func (o LookupOriginResultOutput) ToLookupOriginResultOutputWithContext(ctx context.Context) LookupOriginResultOutput {
	return o
}

func (o LookupOriginResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o LookupOriginResultOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginResult) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o LookupOriginResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupOriginResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOriginResultOutput{})
}
