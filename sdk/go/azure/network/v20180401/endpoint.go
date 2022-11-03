


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Endpoint struct {
	pulumi.CustomResourceState

	CustomHeaders         EndpointPropertiesResponseCustomHeadersArrayOutput `pulumi:"customHeaders"`
	EndpointLocation      pulumi.StringPtrOutput                             `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrOutput                             `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrOutput                             `pulumi:"endpointStatus"`
	GeoMapping            pulumi.StringArrayOutput                           `pulumi:"geoMapping"`
	MinChildEndpoints     pulumi.Float64PtrOutput                            `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrOutput                             `pulumi:"name"`
	Priority              pulumi.Float64PtrOutput                            `pulumi:"priority"`
	Subnets               EndpointPropertiesResponseSubnetsArrayOutput       `pulumi:"subnets"`
	Target                pulumi.StringPtrOutput                             `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrOutput                             `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrOutput                             `pulumi:"type"`
	Weight                pulumi.Float64PtrOutput                            `pulumi:"weight"`
}


func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20151101:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170501:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:network/v20180401:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:network/v20180401:Endpoint", name, id, state, &resource, opts...)
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
	CustomHeaders         []EndpointPropertiesCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                           `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                           `pulumi:"endpointMonitorStatus"`
	EndpointName          *string                           `pulumi:"endpointName"`
	EndpointStatus        *string                           `pulumi:"endpointStatus"`
	EndpointType          string                            `pulumi:"endpointType"`
	GeoMapping            []string                          `pulumi:"geoMapping"`
	Id                    *string                           `pulumi:"id"`
	MinChildEndpoints     *float64                          `pulumi:"minChildEndpoints"`
	Name                  *string                           `pulumi:"name"`
	Priority              *float64                          `pulumi:"priority"`
	ProfileName           string                            `pulumi:"profileName"`
	ResourceGroupName     string                            `pulumi:"resourceGroupName"`
	Subnets               []EndpointPropertiesSubnets       `pulumi:"subnets"`
	Target                *string                           `pulumi:"target"`
	TargetResourceId      *string                           `pulumi:"targetResourceId"`
	Type                  *string                           `pulumi:"type"`
	Weight                *float64                          `pulumi:"weight"`
}


type EndpointArgs struct {
	CustomHeaders         EndpointPropertiesCustomHeadersArrayInput
	EndpointLocation      pulumi.StringPtrInput
	EndpointMonitorStatus pulumi.StringPtrInput
	EndpointName          pulumi.StringPtrInput
	EndpointStatus        pulumi.StringPtrInput
	EndpointType          pulumi.StringInput
	GeoMapping            pulumi.StringArrayInput
	Id                    pulumi.StringPtrInput
	MinChildEndpoints     pulumi.Float64PtrInput
	Name                  pulumi.StringPtrInput
	Priority              pulumi.Float64PtrInput
	ProfileName           pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Subnets               EndpointPropertiesSubnetsArrayInput
	Target                pulumi.StringPtrInput
	TargetResourceId      pulumi.StringPtrInput
	Type                  pulumi.StringPtrInput
	Weight                pulumi.Float64PtrInput
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

func (o EndpointOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesResponseCustomHeadersArrayOutput { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

func (o EndpointOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o EndpointOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointOutput) Subnets() EndpointPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesResponseSubnetsArrayOutput { return v.Subnets }).(EndpointPropertiesResponseSubnetsArrayOutput)
}

func (o EndpointOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.Weight }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
