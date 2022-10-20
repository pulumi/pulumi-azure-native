


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AFDOriginGroup struct {
	pulumi.CustomResourceState

	DeploymentStatus                                      pulumi.StringOutput                                          `pulumi:"deploymentStatus"`
	HealthProbeSettings                                   HealthProbeParametersResponsePtrOutput                       `pulumi:"healthProbeSettings"`
	LoadBalancingSettings                                 LoadBalancingSettingsParametersResponsePtrOutput             `pulumi:"loadBalancingSettings"`
	Name                                                  pulumi.StringOutput                                          `pulumi:"name"`
	ProvisioningState                                     pulumi.StringOutput                                          `pulumi:"provisioningState"`
	ResponseBasedAfdOriginErrorDetectionSettings          ResponseBasedOriginErrorDetectionParametersResponsePtrOutput `pulumi:"responseBasedAfdOriginErrorDetectionSettings"`
	SessionAffinityState                                  pulumi.StringPtrOutput                                       `pulumi:"sessionAffinityState"`
	SystemData                                            SystemDataResponseOutput                                     `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrOutput                                          `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  pulumi.StringOutput                                          `pulumi:"type"`
}


func NewAFDOriginGroup(ctx *pulumi.Context,
	name string, args *AFDOriginGroupArgs, opts ...pulumi.ResourceOption) (*AFDOriginGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDOriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDOriginGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDOriginGroup
	err := ctx.RegisterResource("azure-native:cdn/v20200901:AFDOriginGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAFDOriginGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDOriginGroupState, opts ...pulumi.ResourceOption) (*AFDOriginGroup, error) {
	var resource AFDOriginGroup
	err := ctx.ReadResource("azure-native:cdn/v20200901:AFDOriginGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type afdoriginGroupState struct {
}

type AFDOriginGroupState struct {
}

func (AFDOriginGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginGroupState)(nil)).Elem()
}

type afdoriginGroupArgs struct {
	HealthProbeSettings                                   *HealthProbeParameters                       `pulumi:"healthProbeSettings"`
	LoadBalancingSettings                                 *LoadBalancingSettingsParameters             `pulumi:"loadBalancingSettings"`
	OriginGroupName                                       *string                                      `pulumi:"originGroupName"`
	ProfileName                                           string                                       `pulumi:"profileName"`
	ResourceGroupName                                     string                                       `pulumi:"resourceGroupName"`
	ResponseBasedAfdOriginErrorDetectionSettings          *ResponseBasedOriginErrorDetectionParameters `pulumi:"responseBasedAfdOriginErrorDetectionSettings"`
	SessionAffinityState                                  *string                                      `pulumi:"sessionAffinityState"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                         `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}


type AFDOriginGroupArgs struct {
	HealthProbeSettings                                   HealthProbeParametersPtrInput
	LoadBalancingSettings                                 LoadBalancingSettingsParametersPtrInput
	OriginGroupName                                       pulumi.StringPtrInput
	ProfileName                                           pulumi.StringInput
	ResourceGroupName                                     pulumi.StringInput
	ResponseBasedAfdOriginErrorDetectionSettings          ResponseBasedOriginErrorDetectionParametersPtrInput
	SessionAffinityState                                  pulumi.StringPtrInput
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput
}

func (AFDOriginGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginGroupArgs)(nil)).Elem()
}

type AFDOriginGroupInput interface {
	pulumi.Input

	ToAFDOriginGroupOutput() AFDOriginGroupOutput
	ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput
}

func (*AFDOriginGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOriginGroup)(nil)).Elem()
}

func (i *AFDOriginGroup) ToAFDOriginGroupOutput() AFDOriginGroupOutput {
	return i.ToAFDOriginGroupOutputWithContext(context.Background())
}

func (i *AFDOriginGroup) ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDOriginGroupOutput)
}

type AFDOriginGroupOutput struct{ *pulumi.OutputState }

func (AFDOriginGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOriginGroup)(nil)).Elem()
}

func (o AFDOriginGroupOutput) ToAFDOriginGroupOutput() AFDOriginGroupOutput {
	return o
}

func (o AFDOriginGroupOutput) ToAFDOriginGroupOutputWithContext(ctx context.Context) AFDOriginGroupOutput {
	return o
}

func (o AFDOriginGroupOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o AFDOriginGroupOutput) HealthProbeSettings() HealthProbeParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) HealthProbeParametersResponsePtrOutput { return v.HealthProbeSettings }).(HealthProbeParametersResponsePtrOutput)
}

func (o AFDOriginGroupOutput) LoadBalancingSettings() LoadBalancingSettingsParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) LoadBalancingSettingsParametersResponsePtrOutput {
		return v.LoadBalancingSettings
	}).(LoadBalancingSettingsParametersResponsePtrOutput)
}

func (o AFDOriginGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AFDOriginGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AFDOriginGroupOutput) ResponseBasedAfdOriginErrorDetectionSettings() ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) ResponseBasedOriginErrorDetectionParametersResponsePtrOutput {
		return v.ResponseBasedAfdOriginErrorDetectionSettings
	}).(ResponseBasedOriginErrorDetectionParametersResponsePtrOutput)
}

func (o AFDOriginGroupOutput) SessionAffinityState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringPtrOutput { return v.SessionAffinityState }).(pulumi.StringPtrOutput)
}

func (o AFDOriginGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDOriginGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AFDOriginGroupOutput) TrafficRestorationTimeToHealedOrNewEndpointsInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.IntPtrOutput {
		return v.TrafficRestorationTimeToHealedOrNewEndpointsInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o AFDOriginGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOriginGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDOriginGroupOutput{})
}
