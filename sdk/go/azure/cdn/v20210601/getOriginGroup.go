


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOriginGroup(ctx *pulumi.Context, args *LookupOriginGroupArgs, opts ...pulumi.InvokeOption) (*LookupOriginGroupResult, error) {
	var rv LookupOriginGroupResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getOriginGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOriginGroupArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	OriginGroupName   string `pulumi:"originGroupName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupOriginGroupResult struct {
	HealthProbeSettings                                   *HealthProbeParametersResponse                       `pulumi:"healthProbeSettings"`
	Id                                                    string                                               `pulumi:"id"`
	Name                                                  string                                               `pulumi:"name"`
	Origins                                               []ResourceReferenceResponse                          `pulumi:"origins"`
	ProvisioningState                                     string                                               `pulumi:"provisioningState"`
	ResourceState                                         string                                               `pulumi:"resourceState"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParametersResponse `pulumi:"responseBasedOriginErrorDetectionSettings"`
	SystemData                                            SystemDataResponse                                   `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                                 `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  string                                               `pulumi:"type"`
}

func LookupOriginGroupOutput(ctx *pulumi.Context, args LookupOriginGroupOutputArgs, opts ...pulumi.InvokeOption) LookupOriginGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOriginGroupResult, error) {
			args := v.(LookupOriginGroupArgs)
			r, err := LookupOriginGroup(ctx, &args, opts...)
			var s LookupOriginGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOriginGroupResultOutput)
}

type LookupOriginGroupOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	OriginGroupName   pulumi.StringInput `pulumi:"originGroupName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOriginGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginGroupArgs)(nil)).Elem()
}


type LookupOriginGroupResultOutput struct{ *pulumi.OutputState }

func (LookupOriginGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOriginGroupResult)(nil)).Elem()
}

func (o LookupOriginGroupResultOutput) ToLookupOriginGroupResultOutput() LookupOriginGroupResultOutput {
	return o
}

func (o LookupOriginGroupResultOutput) ToLookupOriginGroupResultOutputWithContext(ctx context.Context) LookupOriginGroupResultOutput {
	return o
}

func (o LookupOriginGroupResultOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *HealthProbeParametersResponse { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

func (o LookupOriginGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOriginGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOriginGroupResultOutput) Origins() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) []ResourceReferenceResponse { return v.Origins }).(ResourceReferenceResponseArrayOutput)
}

func (o LookupOriginGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOriginGroupResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupOriginGroupResultOutput) ResponseBasedOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *ResponseBasedOriginErrorDetectionParametersResponse {
		return v.ResponseBasedOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

func (o LookupOriginGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOriginGroupResultOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) *int { return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupOriginGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOriginGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOriginGroupResultOutput{})
}
