


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAvailabilityGroupListener(ctx *pulumi.Context, args *LookupAvailabilityGroupListenerArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilityGroupListenerResult, error) {
	var rv LookupAvailabilityGroupListenerResult
	err := ctx.Invoke("azure-native:sqlvirtualmachine/v20220801preview:getAvailabilityGroupListener", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilityGroupListenerArgs struct {
	AvailabilityGroupListenerName string  `pulumi:"availabilityGroupListenerName"`
	Expand                        *string `pulumi:"expand"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName    string  `pulumi:"sqlVirtualMachineGroupName"`
}


type LookupAvailabilityGroupListenerResult struct {
	AvailabilityGroupConfiguration           *AgConfigurationResponse             `pulumi:"availabilityGroupConfiguration"`
	AvailabilityGroupName                    *string                              `pulumi:"availabilityGroupName"`
	CreateDefaultAvailabilityGroupIfNotExist *bool                                `pulumi:"createDefaultAvailabilityGroupIfNotExist"`
	Id                                       string                               `pulumi:"id"`
	LoadBalancerConfigurations               []LoadBalancerConfigurationResponse  `pulumi:"loadBalancerConfigurations"`
	MultiSubnetIpConfigurations              []MultiSubnetIpConfigurationResponse `pulumi:"multiSubnetIpConfigurations"`
	Name                                     string                               `pulumi:"name"`
	Port                                     *int                                 `pulumi:"port"`
	ProvisioningState                        string                               `pulumi:"provisioningState"`
	SystemData                               SystemDataResponse                   `pulumi:"systemData"`
	Type                                     string                               `pulumi:"type"`
}

func LookupAvailabilityGroupListenerOutput(ctx *pulumi.Context, args LookupAvailabilityGroupListenerOutputArgs, opts ...pulumi.InvokeOption) LookupAvailabilityGroupListenerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAvailabilityGroupListenerResult, error) {
			args := v.(LookupAvailabilityGroupListenerArgs)
			r, err := LookupAvailabilityGroupListener(ctx, &args, opts...)
			var s LookupAvailabilityGroupListenerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAvailabilityGroupListenerResultOutput)
}

type LookupAvailabilityGroupListenerOutputArgs struct {
	AvailabilityGroupListenerName pulumi.StringInput    `pulumi:"availabilityGroupListenerName"`
	Expand                        pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName             pulumi.StringInput    `pulumi:"resourceGroupName"`
	SqlVirtualMachineGroupName    pulumi.StringInput    `pulumi:"sqlVirtualMachineGroupName"`
}

func (LookupAvailabilityGroupListenerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilityGroupListenerArgs)(nil)).Elem()
}


type LookupAvailabilityGroupListenerResultOutput struct{ *pulumi.OutputState }

func (LookupAvailabilityGroupListenerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAvailabilityGroupListenerResult)(nil)).Elem()
}

func (o LookupAvailabilityGroupListenerResultOutput) ToLookupAvailabilityGroupListenerResultOutput() LookupAvailabilityGroupListenerResultOutput {
	return o
}

func (o LookupAvailabilityGroupListenerResultOutput) ToLookupAvailabilityGroupListenerResultOutputWithContext(ctx context.Context) LookupAvailabilityGroupListenerResultOutput {
	return o
}

func (o LookupAvailabilityGroupListenerResultOutput) AvailabilityGroupConfiguration() AgConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) *AgConfigurationResponse {
		return v.AvailabilityGroupConfiguration
	}).(AgConfigurationResponsePtrOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) AvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) *string { return v.AvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) CreateDefaultAvailabilityGroupIfNotExist() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) *bool { return v.CreateDefaultAvailabilityGroupIfNotExist }).(pulumi.BoolPtrOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) LoadBalancerConfigurations() LoadBalancerConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) []LoadBalancerConfigurationResponse {
		return v.LoadBalancerConfigurations
	}).(LoadBalancerConfigurationResponseArrayOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) MultiSubnetIpConfigurations() MultiSubnetIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) []MultiSubnetIpConfigurationResponse {
		return v.MultiSubnetIpConfigurations
	}).(MultiSubnetIpConfigurationResponseArrayOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAvailabilityGroupListenerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAvailabilityGroupListenerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAvailabilityGroupListenerResultOutput{})
}
