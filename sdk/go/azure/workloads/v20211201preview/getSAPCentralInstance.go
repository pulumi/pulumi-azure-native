


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSAPCentralInstance(ctx *pulumi.Context, args *LookupSAPCentralInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSAPCentralInstanceResult, error) {
	var rv LookupSAPCentralInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPCentralInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSAPCentralInstanceArgs struct {
	CentralInstanceName    string `pulumi:"centralInstanceName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
}


type LookupSAPCentralInstanceResult struct {
	EnqueueReplicationServerProperties *EnqueueReplicationServerPropertiesResponse `pulumi:"enqueueReplicationServerProperties"`
	EnqueueServerProperties            *EnqueueServerPropertiesResponse            `pulumi:"enqueueServerProperties"`
	Errors                             SAPVirtualInstanceErrorResponse             `pulumi:"errors"`
	GatewayServerProperties            *GatewayServerPropertiesResponse            `pulumi:"gatewayServerProperties"`
	Health                             string                                      `pulumi:"health"`
	Id                                 string                                      `pulumi:"id"`
	InstanceNo                         string                                      `pulumi:"instanceNo"`
	KernelPatch                        string                                      `pulumi:"kernelPatch"`
	KernelVersion                      string                                      `pulumi:"kernelVersion"`
	LoadBalancerDetails                LoadBalancerDetailsResponse                 `pulumi:"loadBalancerDetails"`
	Location                           string                                      `pulumi:"location"`
	MessageServerProperties            *MessageServerPropertiesResponse            `pulumi:"messageServerProperties"`
	Name                               string                                      `pulumi:"name"`
	ProvisioningState                  string                                      `pulumi:"provisioningState"`
	Status                             string                                      `pulumi:"status"`
	Subnet                             string                                      `pulumi:"subnet"`
	SystemData                         SystemDataResponse                          `pulumi:"systemData"`
	Tags                               map[string]string                           `pulumi:"tags"`
	Type                               string                                      `pulumi:"type"`
	VmDetails                          []CentralServerVmDetailsResponse            `pulumi:"vmDetails"`
}

func LookupSAPCentralInstanceOutput(ctx *pulumi.Context, args LookupSAPCentralInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSAPCentralInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSAPCentralInstanceResult, error) {
			args := v.(LookupSAPCentralInstanceArgs)
			r, err := LookupSAPCentralInstance(ctx, &args, opts...)
			var s LookupSAPCentralInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSAPCentralInstanceResultOutput)
}

type LookupSAPCentralInstanceOutputArgs struct {
	CentralInstanceName    pulumi.StringInput `pulumi:"centralInstanceName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SapVirtualInstanceName pulumi.StringInput `pulumi:"sapVirtualInstanceName"`
}

func (LookupSAPCentralInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPCentralInstanceArgs)(nil)).Elem()
}


type LookupSAPCentralInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSAPCentralInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPCentralInstanceResult)(nil)).Elem()
}

func (o LookupSAPCentralInstanceResultOutput) ToLookupSAPCentralInstanceResultOutput() LookupSAPCentralInstanceResultOutput {
	return o
}

func (o LookupSAPCentralInstanceResultOutput) ToLookupSAPCentralInstanceResultOutputWithContext(ctx context.Context) LookupSAPCentralInstanceResultOutput {
	return o
}

func (o LookupSAPCentralInstanceResultOutput) EnqueueReplicationServerProperties() EnqueueReplicationServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) *EnqueueReplicationServerPropertiesResponse {
		return v.EnqueueReplicationServerProperties
	}).(EnqueueReplicationServerPropertiesResponsePtrOutput)
}

func (o LookupSAPCentralInstanceResultOutput) EnqueueServerProperties() EnqueueServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) *EnqueueServerPropertiesResponse {
		return v.EnqueueServerProperties
	}).(EnqueueServerPropertiesResponsePtrOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) SAPVirtualInstanceErrorResponse { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o LookupSAPCentralInstanceResultOutput) GatewayServerProperties() GatewayServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) *GatewayServerPropertiesResponse {
		return v.GatewayServerProperties
	}).(GatewayServerPropertiesResponsePtrOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Health }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.InstanceNo }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.KernelPatch }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.KernelVersion }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) LoadBalancerDetailsResponse { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) MessageServerProperties() MessageServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) *MessageServerPropertiesResponse {
		return v.MessageServerProperties
	}).(MessageServerPropertiesResponsePtrOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSAPCentralInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSAPCentralInstanceResultOutput) VmDetails() CentralServerVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupSAPCentralInstanceResult) []CentralServerVmDetailsResponse { return v.VmDetails }).(CentralServerVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSAPCentralInstanceResultOutput{})
}
