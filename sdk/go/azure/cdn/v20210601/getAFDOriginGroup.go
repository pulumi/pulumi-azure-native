


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDOriginGroup(ctx *pulumi.Context, args *LookupAFDOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupAFDOriginGroupResult, error) {
	var rv LookupAFDOriginGroupResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getAFDOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDOriginGroupArgs struct {
	OriginGroupName   string `pulumi:"originGroupName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDOriginGroupResult struct {
	DeploymentStatus                                      string                                   `pulumi:"deploymentStatus"`
	HealthProbeSettings                                   *HealthProbeParametersResponse           `pulumi:"healthProbeSettings"`
	Id                                                    string                                   `pulumi:"id"`
	LoadBalancingSettings                                 *LoadBalancingSettingsParametersResponse `pulumi:"loadBalancingSettings"`
	Name                                                  string                                   `pulumi:"name"`
	ProfileName                                           string                                   `pulumi:"profileName"`
	ProvisioningState                                     string                                   `pulumi:"provisioningState"`
	SessionAffinityState                                  *string                                  `pulumi:"sessionAffinityState"`
	SystemData                                            SystemDataResponse                       `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                     `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  string                                   `pulumi:"type"`
}

func LookupAFDOriginGroupOutput(ctx *pulumi.Context, args LookupAFDOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAFDOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDOriginGroupResult, error) {
			args := v.(LookupAFDOriginGroupArgs)
			r, err := LookupAFDOriginGroup(ctx, &args, opts...)
			var s LookupAFDOriginGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDOriginGroupResultOutput)
}

type LookupAFDOriginGroupOutputArgs struct {
	OriginGroupName   pulumi.StringInput `pulumi:"originGroupName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupArgs)(nil)).Elem()
}


type LookupAFDOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAFDOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDOriginGroupResult)(nil)).Elem()
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutput() LookupAFDOriginGroupResultOutput {
	return o
}

func (o LookupAFDOriginGroupResultOutput) ToLookupAFDOriginGroupResultOutputWithContext(ctx context.Context) LookupAFDOriginGroupResultOutput {
	return o
}

func (o LookupAFDOriginGroupResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

func (o LookupAFDOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) LoadBalancingSettings() LoadBalancingSettingsParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *LoadBalancingSettingsParametersResponse {
		return v.LoadBalancingSettings
	}).(LoadBalancingSettingsParametersResponsePtrOutput)
}

func (o LookupAFDOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAFDOriginGroupResultOutput) SessionAffinityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *string { return v.SessionAffinityState }).(pulumi.StringPtrOutput)
}

func (o LookupAFDOriginGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAFDOriginGroupResultOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) *int {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o LookupAFDOriginGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDOriginGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDOriginGroupResultOutput{})
}
