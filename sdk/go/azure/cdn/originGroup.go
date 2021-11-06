


package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OriginGroup struct {
	pulumi.CustomResourceState

	HealthProbeSettings                                   HealthProbeParametersResponsePtrOutput                       `pulumi:"healthProbeSettings"`
	Name                                                  pulumi.StringOutput                                          `pulumi:"name"`
	Origins                                               ResourceReferenceResponseArrayOutput                         `pulumi:"origins"`
	ProvisioningState                                     pulumi.StringOutput                                          `pulumi:"provisioningState"`
	ResourceState                                         pulumi.StringOutput                                          `pulumi:"resourceState"`
	ResponseBasedOriginErrorDetectionSettings             ResponseBasedOriginErrorDetectionParametersResponsePtrOutput `pulumi:"responseBasedOriginErrorDetectionSettings"`
	SystemData                                            SystemDataResponseOutput                                     `pulumi:"systemData"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrOutput                                          `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
	Type                                                  pulumi.StringOutput                                          `pulumi:"type"`
}


func NewOriginGroup(ctx *pulumi.Context,
	name string, args *OriginGroupArgs, opts ...pulumi.ResourceOption) (*OriginGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.Origins == nil {
		return nil, errors.New("invalid value for required argument 'Origins'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:cdn:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20191231:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200331:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200415:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:OriginGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:cdn/v20200901:OriginGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource OriginGroup
	err := ctx.RegisterResource("azure-native:cdn:OriginGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOriginGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginGroupState, opts ...pulumi.ResourceOption) (*OriginGroup, error) {
	var resource OriginGroup
	err := ctx.ReadResource("azure-native:cdn:OriginGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type originGroupState struct {
}

type OriginGroupState struct {
}

func (OriginGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*originGroupState)(nil)).Elem()
}

type originGroupArgs struct {
	EndpointName                                          string                                       `pulumi:"endpointName"`
	HealthProbeSettings                                   *HealthProbeParameters                       `pulumi:"healthProbeSettings"`
	OriginGroupName                                       *string                                      `pulumi:"originGroupName"`
	Origins                                               []ResourceReference                          `pulumi:"origins"`
	ProfileName                                           string                                       `pulumi:"profileName"`
	ResourceGroupName                                     string                                       `pulumi:"resourceGroupName"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParameters `pulumi:"responseBasedOriginErrorDetectionSettings"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                         `pulumi:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes"`
}


type OriginGroupArgs struct {
	EndpointName                                          pulumi.StringInput
	HealthProbeSettings                                   HealthProbeParametersPtrInput
	OriginGroupName                                       pulumi.StringPtrInput
	Origins                                               ResourceReferenceArrayInput
	ProfileName                                           pulumi.StringInput
	ResourceGroupName                                     pulumi.StringInput
	ResponseBasedOriginErrorDetectionSettings             ResponseBasedOriginErrorDetectionParametersPtrInput
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes pulumi.IntPtrInput
}

func (OriginGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originGroupArgs)(nil)).Elem()
}

type OriginGroupInput interface {
	pulumi.Input

	ToOriginGroupOutput() OriginGroupOutput
	ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput
}

func (*OriginGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginGroup)(nil))
}

func (i *OriginGroup) ToOriginGroupOutput() OriginGroupOutput {
	return i.ToOriginGroupOutputWithContext(context.Background())
}

func (i *OriginGroup) ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginGroupOutput)
}

type OriginGroupOutput struct{ *pulumi.OutputState }

func (OriginGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginGroup)(nil))
}

func (o OriginGroupOutput) ToOriginGroupOutput() OriginGroupOutput {
	return o
}

func (o OriginGroupOutput) ToOriginGroupOutputWithContext(ctx context.Context) OriginGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OriginGroupOutput{})
}
