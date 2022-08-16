


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFrontDoor(ctx *pulumi.Context, args *LookupFrontDoorArgs, opts ...pulumi.InvokeOption) (*LookupFrontDoorResult, error) {
	var rv LookupFrontDoorResult
	err := ctx.Invoke("azure-native:network/v20200401:getFrontDoor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFrontDoorArgs struct {
	FrontDoorName     string `pulumi:"frontDoorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFrontDoorResult struct {
	BackendPools          []BackendPoolResponse                `pulumi:"backendPools"`
	BackendPoolsSettings  *BackendPoolsSettingsResponse        `pulumi:"backendPoolsSettings"`
	Cname                 string                               `pulumi:"cname"`
	EnabledState          *string                              `pulumi:"enabledState"`
	FriendlyName          *string                              `pulumi:"friendlyName"`
	FrontdoorId           string                               `pulumi:"frontdoorId"`
	FrontendEndpoints     []FrontendEndpointResponse           `pulumi:"frontendEndpoints"`
	HealthProbeSettings   []HealthProbeSettingsModelResponse   `pulumi:"healthProbeSettings"`
	Id                    string                               `pulumi:"id"`
	LoadBalancingSettings []LoadBalancingSettingsModelResponse `pulumi:"loadBalancingSettings"`
	Location              *string                              `pulumi:"location"`
	Name                  string                               `pulumi:"name"`
	ProvisioningState     string                               `pulumi:"provisioningState"`
	ResourceState         string                               `pulumi:"resourceState"`
	RoutingRules          []RoutingRuleResponse                `pulumi:"routingRules"`
	RulesEngines          []RulesEngineResponse                `pulumi:"rulesEngines"`
	Tags                  map[string]string                    `pulumi:"tags"`
	Type                  string                               `pulumi:"type"`
}


func (val *LookupFrontDoorResult) Defaults() *LookupFrontDoorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackendPoolsSettings = tmp.BackendPoolsSettings.Defaults()

	return &tmp
}

func LookupFrontDoorOutput(ctx *pulumi.Context, args LookupFrontDoorOutputArgs, opts ...pulumi.InvokeOption) LookupFrontDoorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFrontDoorResult, error) {
			args := v.(LookupFrontDoorArgs)
			r, err := LookupFrontDoor(ctx, &args, opts...)
			var s LookupFrontDoorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFrontDoorResultOutput)
}

type LookupFrontDoorOutputArgs struct {
	FrontDoorName     pulumi.StringInput `pulumi:"frontDoorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFrontDoorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontDoorArgs)(nil)).Elem()
}


type LookupFrontDoorResultOutput struct{ *pulumi.OutputState }

func (LookupFrontDoorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontDoorResult)(nil)).Elem()
}

func (o LookupFrontDoorResultOutput) ToLookupFrontDoorResultOutput() LookupFrontDoorResultOutput {
	return o
}

func (o LookupFrontDoorResultOutput) ToLookupFrontDoorResultOutputWithContext(ctx context.Context) LookupFrontDoorResultOutput {
	return o
}

func (o LookupFrontDoorResultOutput) BackendPools() BackendPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []BackendPoolResponse { return v.BackendPools }).(BackendPoolResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) BackendPoolsSettings() BackendPoolsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) *BackendPoolsSettingsResponse { return v.BackendPoolsSettings }).(BackendPoolsSettingsResponsePtrOutput)
}

func (o LookupFrontDoorResultOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.Cname }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupFrontDoorResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupFrontDoorResultOutput) FrontdoorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.FrontdoorId }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) FrontendEndpoints() FrontendEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []FrontendEndpointResponse { return v.FrontendEndpoints }).(FrontendEndpointResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) HealthProbeSettings() HealthProbeSettingsModelResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []HealthProbeSettingsModelResponse { return v.HealthProbeSettings }).(HealthProbeSettingsModelResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) LoadBalancingSettings() LoadBalancingSettingsModelResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []LoadBalancingSettingsModelResponse { return v.LoadBalancingSettings }).(LoadBalancingSettingsModelResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFrontDoorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupFrontDoorResultOutput) RoutingRules() RoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []RoutingRuleResponse { return v.RoutingRules }).(RoutingRuleResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) RulesEngines() RulesEngineResponseArrayOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) []RulesEngineResponse { return v.RulesEngines }).(RulesEngineResponseArrayOutput)
}

func (o LookupFrontDoorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFrontDoorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontDoorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFrontDoorResultOutput{})
}
