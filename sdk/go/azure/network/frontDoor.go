


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FrontDoor struct {
	pulumi.CustomResourceState

	BackendPools          BackendPoolResponseArrayOutput                `pulumi:"backendPools"`
	BackendPoolsSettings  BackendPoolsSettingsResponsePtrOutput         `pulumi:"backendPoolsSettings"`
	Cname                 pulumi.StringOutput                           `pulumi:"cname"`
	EnabledState          pulumi.StringPtrOutput                        `pulumi:"enabledState"`
	FriendlyName          pulumi.StringPtrOutput                        `pulumi:"friendlyName"`
	FrontdoorId           pulumi.StringOutput                           `pulumi:"frontdoorId"`
	FrontendEndpoints     FrontendEndpointResponseArrayOutput           `pulumi:"frontendEndpoints"`
	HealthProbeSettings   HealthProbeSettingsModelResponseArrayOutput   `pulumi:"healthProbeSettings"`
	LoadBalancingSettings LoadBalancingSettingsModelResponseArrayOutput `pulumi:"loadBalancingSettings"`
	Location              pulumi.StringPtrOutput                        `pulumi:"location"`
	Name                  pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                           `pulumi:"provisioningState"`
	ResourceState         pulumi.StringOutput                           `pulumi:"resourceState"`
	RoutingRules          RoutingRuleResponseArrayOutput                `pulumi:"routingRules"`
	RulesEngines          RulesEngineResponseArrayOutput                `pulumi:"rulesEngines"`
	Tags                  pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                  pulumi.StringOutput                           `pulumi:"type"`
}


func NewFrontDoor(ctx *pulumi.Context,
	name string, args *FrontDoorArgs, opts ...pulumi.ResourceOption) (*FrontDoor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.BackendPoolsSettings != nil {
		args.BackendPoolsSettings = args.BackendPoolsSettings.ToBackendPoolsSettingsPtrOutput().ApplyT(func(v *BackendPoolsSettings) *BackendPoolsSettings { return v.Defaults() }).(BackendPoolsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20190401:FrontDoor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190501:FrontDoor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:FrontDoor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:FrontDoor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:FrontDoor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210601:FrontDoor"),
		},
	})
	opts = append(opts, aliases)
	var resource FrontDoor
	err := ctx.RegisterResource("azure-native:network:FrontDoor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFrontDoor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontDoorState, opts ...pulumi.ResourceOption) (*FrontDoor, error) {
	var resource FrontDoor
	err := ctx.ReadResource("azure-native:network:FrontDoor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type frontDoorState struct {
}

type FrontDoorState struct {
}

func (FrontDoorState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorState)(nil)).Elem()
}

type frontDoorArgs struct {
	BackendPools          []BackendPool                `pulumi:"backendPools"`
	BackendPoolsSettings  *BackendPoolsSettings        `pulumi:"backendPoolsSettings"`
	EnabledState          *string                      `pulumi:"enabledState"`
	FriendlyName          *string                      `pulumi:"friendlyName"`
	FrontDoorName         *string                      `pulumi:"frontDoorName"`
	FrontendEndpoints     []FrontendEndpoint           `pulumi:"frontendEndpoints"`
	HealthProbeSettings   []HealthProbeSettingsModel   `pulumi:"healthProbeSettings"`
	LoadBalancingSettings []LoadBalancingSettingsModel `pulumi:"loadBalancingSettings"`
	Location              *string                      `pulumi:"location"`
	ResourceGroupName     string                       `pulumi:"resourceGroupName"`
	RoutingRules          []RoutingRule                `pulumi:"routingRules"`
	Tags                  map[string]string            `pulumi:"tags"`
}


type FrontDoorArgs struct {
	BackendPools          BackendPoolArrayInput
	BackendPoolsSettings  BackendPoolsSettingsPtrInput
	EnabledState          pulumi.StringPtrInput
	FriendlyName          pulumi.StringPtrInput
	FrontDoorName         pulumi.StringPtrInput
	FrontendEndpoints     FrontendEndpointArrayInput
	HealthProbeSettings   HealthProbeSettingsModelArrayInput
	LoadBalancingSettings LoadBalancingSettingsModelArrayInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	RoutingRules          RoutingRuleArrayInput
	Tags                  pulumi.StringMapInput
}

func (FrontDoorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontDoorArgs)(nil)).Elem()
}

type FrontDoorInput interface {
	pulumi.Input

	ToFrontDoorOutput() FrontDoorOutput
	ToFrontDoorOutputWithContext(ctx context.Context) FrontDoorOutput
}

func (*FrontDoor) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoor)(nil)).Elem()
}

func (i *FrontDoor) ToFrontDoorOutput() FrontDoorOutput {
	return i.ToFrontDoorOutputWithContext(context.Background())
}

func (i *FrontDoor) ToFrontDoorOutputWithContext(ctx context.Context) FrontDoorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorOutput)
}

type FrontDoorOutput struct{ *pulumi.OutputState }

func (FrontDoorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoor)(nil)).Elem()
}

func (o FrontDoorOutput) ToFrontDoorOutput() FrontDoorOutput {
	return o
}

func (o FrontDoorOutput) ToFrontDoorOutputWithContext(ctx context.Context) FrontDoorOutput {
	return o
}

func (o FrontDoorOutput) BackendPools() BackendPoolResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) BackendPoolResponseArrayOutput { return v.BackendPools }).(BackendPoolResponseArrayOutput)
}

func (o FrontDoorOutput) BackendPoolsSettings() BackendPoolsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *FrontDoor) BackendPoolsSettingsResponsePtrOutput { return v.BackendPoolsSettings }).(BackendPoolsSettingsResponsePtrOutput)
}

func (o FrontDoorOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

func (o FrontDoorOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o FrontDoorOutput) FrontdoorId() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.FrontdoorId }).(pulumi.StringOutput)
}

func (o FrontDoorOutput) FrontendEndpoints() FrontendEndpointResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) FrontendEndpointResponseArrayOutput { return v.FrontendEndpoints }).(FrontendEndpointResponseArrayOutput)
}

func (o FrontDoorOutput) HealthProbeSettings() HealthProbeSettingsModelResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) HealthProbeSettingsModelResponseArrayOutput { return v.HealthProbeSettings }).(HealthProbeSettingsModelResponseArrayOutput)
}

func (o FrontDoorOutput) LoadBalancingSettings() LoadBalancingSettingsModelResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) LoadBalancingSettingsModelResponseArrayOutput { return v.LoadBalancingSettings }).(LoadBalancingSettingsModelResponseArrayOutput)
}

func (o FrontDoorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o FrontDoorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FrontDoorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FrontDoorOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o FrontDoorOutput) RoutingRules() RoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) RoutingRuleResponseArrayOutput { return v.RoutingRules }).(RoutingRuleResponseArrayOutput)
}

func (o FrontDoorOutput) RulesEngines() RulesEngineResponseArrayOutput {
	return o.ApplyT(func(v *FrontDoor) RulesEngineResponseArrayOutput { return v.RulesEngines }).(RulesEngineResponseArrayOutput)
}

func (o FrontDoorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FrontDoorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontDoor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FrontDoorOutput{})
}
