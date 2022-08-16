


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkProfile(ctx *pulumi.Context, args *LookupNetworkProfileArgs, opts ...pulumi.InvokeOption) (*LookupNetworkProfileResult, error) {
	var rv LookupNetworkProfileResult
	err := ctx.Invoke("azure-native:network/v20190701:getNetworkProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkProfileArgs struct {
	Expand             *string `pulumi:"expand"`
	NetworkProfileName string  `pulumi:"networkProfileName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupNetworkProfileResult struct {
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfigurationResponse `pulumi:"containerNetworkInterfaceConfigurations"`
	ContainerNetworkInterfaces              []ContainerNetworkInterfaceResponse              `pulumi:"containerNetworkInterfaces"`
	Etag                                    *string                                          `pulumi:"etag"`
	Id                                      *string                                          `pulumi:"id"`
	Location                                *string                                          `pulumi:"location"`
	Name                                    string                                           `pulumi:"name"`
	ProvisioningState                       string                                           `pulumi:"provisioningState"`
	ResourceGuid                            string                                           `pulumi:"resourceGuid"`
	Tags                                    map[string]string                                `pulumi:"tags"`
	Type                                    string                                           `pulumi:"type"`
}

func LookupNetworkProfileOutput(ctx *pulumi.Context, args LookupNetworkProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkProfileResult, error) {
			args := v.(LookupNetworkProfileArgs)
			r, err := LookupNetworkProfile(ctx, &args, opts...)
			var s LookupNetworkProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkProfileResultOutput)
}

type LookupNetworkProfileOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	NetworkProfileName pulumi.StringInput    `pulumi:"networkProfileName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNetworkProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkProfileArgs)(nil)).Elem()
}


type LookupNetworkProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkProfileResult)(nil)).Elem()
}

func (o LookupNetworkProfileResultOutput) ToLookupNetworkProfileResultOutput() LookupNetworkProfileResultOutput {
	return o
}

func (o LookupNetworkProfileResultOutput) ToLookupNetworkProfileResultOutputWithContext(ctx context.Context) LookupNetworkProfileResultOutput {
	return o
}

func (o LookupNetworkProfileResultOutput) ContainerNetworkInterfaceConfigurations() ContainerNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) []ContainerNetworkInterfaceConfigurationResponse {
		return v.ContainerNetworkInterfaceConfigurations
	}).(ContainerNetworkInterfaceConfigurationResponseArrayOutput)
}

func (o LookupNetworkProfileResultOutput) ContainerNetworkInterfaces() ContainerNetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) []ContainerNetworkInterfaceResponse {
		return v.ContainerNetworkInterfaces
	}).(ContainerNetworkInterfaceResponseArrayOutput)
}

func (o LookupNetworkProfileResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkProfileResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkProfileResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupNetworkProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkProfileResultOutput{})
}
