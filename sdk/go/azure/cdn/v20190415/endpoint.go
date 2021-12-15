


package v20190415

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Endpoint struct {
	pulumi.CustomResourceState

	ContentTypesToCompress     pulumi.StringArrayOutput                                          `pulumi:"contentTypesToCompress"`
	DeliveryPolicy             EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput `pulumi:"deliveryPolicy"`
	GeoFilters                 GeoFilterResponseArrayOutput                                      `pulumi:"geoFilters"`
	HostName                   pulumi.StringOutput                                               `pulumi:"hostName"`
	IsCompressionEnabled       pulumi.BoolPtrOutput                                              `pulumi:"isCompressionEnabled"`
	IsHttpAllowed              pulumi.BoolPtrOutput                                              `pulumi:"isHttpAllowed"`
	IsHttpsAllowed             pulumi.BoolPtrOutput                                              `pulumi:"isHttpsAllowed"`
	Location                   pulumi.StringOutput                                               `pulumi:"location"`
	Name                       pulumi.StringOutput                                               `pulumi:"name"`
	OptimizationType           pulumi.StringPtrOutput                                            `pulumi:"optimizationType"`
	OriginHostHeader           pulumi.StringPtrOutput                                            `pulumi:"originHostHeader"`
	OriginPath                 pulumi.StringPtrOutput                                            `pulumi:"originPath"`
	Origins                    DeepCreatedOriginResponseArrayOutput                              `pulumi:"origins"`
	ProbePath                  pulumi.StringPtrOutput                                            `pulumi:"probePath"`
	ProvisioningState          pulumi.StringOutput                                               `pulumi:"provisioningState"`
	QueryStringCachingBehavior pulumi.StringPtrOutput                                            `pulumi:"queryStringCachingBehavior"`
	ResourceState              pulumi.StringOutput                                               `pulumi:"resourceState"`
	Tags                       pulumi.StringMapOutput                                            `pulumi:"tags"`
	Type                       pulumi.StringOutput                                               `pulumi:"type"`
}


func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:cdn:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:cdn/v20190415:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:cdn/v20190415:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	ContentTypesToCompress     []string                                          `pulumi:"contentTypesToCompress"`
	DeliveryPolicy             *EndpointPropertiesUpdateParametersDeliveryPolicy `pulumi:"deliveryPolicy"`
	EndpointName               *string                                           `pulumi:"endpointName"`
	GeoFilters                 []GeoFilter                                       `pulumi:"geoFilters"`
	IsCompressionEnabled       *bool                                             `pulumi:"isCompressionEnabled"`
	IsHttpAllowed              *bool                                             `pulumi:"isHttpAllowed"`
	IsHttpsAllowed             *bool                                             `pulumi:"isHttpsAllowed"`
	Location                   *string                                           `pulumi:"location"`
	OptimizationType           *string                                           `pulumi:"optimizationType"`
	OriginHostHeader           *string                                           `pulumi:"originHostHeader"`
	OriginPath                 *string                                           `pulumi:"originPath"`
	Origins                    []DeepCreatedOrigin                               `pulumi:"origins"`
	ProbePath                  *string                                           `pulumi:"probePath"`
	ProfileName                string                                            `pulumi:"profileName"`
	QueryStringCachingBehavior *QueryStringCachingBehavior                       `pulumi:"queryStringCachingBehavior"`
	ResourceGroupName          string                                            `pulumi:"resourceGroupName"`
	Tags                       map[string]string                                 `pulumi:"tags"`
}


type EndpointArgs struct {
	ContentTypesToCompress     pulumi.StringArrayInput
	DeliveryPolicy             EndpointPropertiesUpdateParametersDeliveryPolicyPtrInput
	EndpointName               pulumi.StringPtrInput
	GeoFilters                 GeoFilterArrayInput
	IsCompressionEnabled       pulumi.BoolPtrInput
	IsHttpAllowed              pulumi.BoolPtrInput
	IsHttpsAllowed             pulumi.BoolPtrInput
	Location                   pulumi.StringPtrInput
	OptimizationType           pulumi.StringPtrInput
	OriginHostHeader           pulumi.StringPtrInput
	OriginPath                 pulumi.StringPtrInput
	Origins                    DeepCreatedOriginArrayInput
	ProbePath                  pulumi.StringPtrInput
	ProfileName                pulumi.StringInput
	QueryStringCachingBehavior QueryStringCachingBehaviorPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
