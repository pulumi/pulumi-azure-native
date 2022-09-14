


package v20170801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ExpressRouteCircuit struct {
	pulumi.CustomResourceState

	AllowClassicOperations           pulumi.BoolPtrOutput                                          `pulumi:"allowClassicOperations"`
	Authorizations                   ExpressRouteCircuitAuthorizationResponseArrayOutput           `pulumi:"authorizations"`
	CircuitProvisioningState         pulumi.StringPtrOutput                                        `pulumi:"circuitProvisioningState"`
	Etag                             pulumi.StringOutput                                           `pulumi:"etag"`
	GatewayManagerEtag               pulumi.StringPtrOutput                                        `pulumi:"gatewayManagerEtag"`
	Location                         pulumi.StringPtrOutput                                        `pulumi:"location"`
	Name                             pulumi.StringOutput                                           `pulumi:"name"`
	Peerings                         ExpressRouteCircuitPeeringResponseArrayOutput                 `pulumi:"peerings"`
	ProvisioningState                pulumi.StringPtrOutput                                        `pulumi:"provisioningState"`
	ServiceKey                       pulumi.StringPtrOutput                                        `pulumi:"serviceKey"`
	ServiceProviderNotes             pulumi.StringPtrOutput                                        `pulumi:"serviceProviderNotes"`
	ServiceProviderProperties        ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput `pulumi:"serviceProviderProperties"`
	ServiceProviderProvisioningState pulumi.StringPtrOutput                                        `pulumi:"serviceProviderProvisioningState"`
	Sku                              ExpressRouteCircuitSkuResponsePtrOutput                       `pulumi:"sku"`
	Tags                             pulumi.StringMapOutput                                        `pulumi:"tags"`
	Type                             pulumi.StringOutput                                           `pulumi:"type"`
}


func NewExpressRouteCircuit(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCircuit"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCircuit"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuit
	err := ctx.RegisterResource("azure-native:network/v20170801:ExpressRouteCircuit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteCircuit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuit, error) {
	var resource ExpressRouteCircuit
	err := ctx.ReadResource("azure-native:network/v20170801:ExpressRouteCircuit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteCircuitState struct {
}

type ExpressRouteCircuitState struct {
}

func (ExpressRouteCircuitState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitState)(nil)).Elem()
}

type expressRouteCircuitArgs struct {
	AllowClassicOperations           *bool                                         `pulumi:"allowClassicOperations"`
	Authorizations                   []ExpressRouteCircuitAuthorizationType        `pulumi:"authorizations"`
	CircuitName                      *string                                       `pulumi:"circuitName"`
	CircuitProvisioningState         *string                                       `pulumi:"circuitProvisioningState"`
	GatewayManagerEtag               *string                                       `pulumi:"gatewayManagerEtag"`
	Id                               *string                                       `pulumi:"id"`
	Location                         *string                                       `pulumi:"location"`
	Peerings                         []ExpressRouteCircuitPeeringType              `pulumi:"peerings"`
	ProvisioningState                *string                                       `pulumi:"provisioningState"`
	ResourceGroupName                string                                        `pulumi:"resourceGroupName"`
	ServiceKey                       *string                                       `pulumi:"serviceKey"`
	ServiceProviderNotes             *string                                       `pulumi:"serviceProviderNotes"`
	ServiceProviderProperties        *ExpressRouteCircuitServiceProviderProperties `pulumi:"serviceProviderProperties"`
	ServiceProviderProvisioningState *string                                       `pulumi:"serviceProviderProvisioningState"`
	Sku                              *ExpressRouteCircuitSku                       `pulumi:"sku"`
	Tags                             map[string]string                             `pulumi:"tags"`
}


type ExpressRouteCircuitArgs struct {
	AllowClassicOperations           pulumi.BoolPtrInput
	Authorizations                   ExpressRouteCircuitAuthorizationTypeArrayInput
	CircuitName                      pulumi.StringPtrInput
	CircuitProvisioningState         pulumi.StringPtrInput
	GatewayManagerEtag               pulumi.StringPtrInput
	Id                               pulumi.StringPtrInput
	Location                         pulumi.StringPtrInput
	Peerings                         ExpressRouteCircuitPeeringTypeArrayInput
	ProvisioningState                pulumi.StringPtrInput
	ResourceGroupName                pulumi.StringInput
	ServiceKey                       pulumi.StringPtrInput
	ServiceProviderNotes             pulumi.StringPtrInput
	ServiceProviderProperties        ExpressRouteCircuitServiceProviderPropertiesPtrInput
	ServiceProviderProvisioningState pulumi.StringPtrInput
	Sku                              ExpressRouteCircuitSkuPtrInput
	Tags                             pulumi.StringMapInput
}

func (ExpressRouteCircuitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitArgs)(nil)).Elem()
}

type ExpressRouteCircuitInput interface {
	pulumi.Input

	ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput
	ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput
}

func (*ExpressRouteCircuit) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil)).Elem()
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return i.ToExpressRouteCircuitOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuit) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitOutput)
}

type ExpressRouteCircuitOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuit)(nil)).Elem()
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutput() ExpressRouteCircuitOutput {
	return o
}

func (o ExpressRouteCircuitOutput) ToExpressRouteCircuitOutputWithContext(ctx context.Context) ExpressRouteCircuitOutput {
	return o
}

func (o ExpressRouteCircuitOutput) AllowClassicOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.BoolPtrOutput { return v.AllowClassicOperations }).(pulumi.BoolPtrOutput)
}

func (o ExpressRouteCircuitOutput) Authorizations() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitAuthorizationResponseArrayOutput {
		return v.Authorizations
	}).(ExpressRouteCircuitAuthorizationResponseArrayOutput)
}

func (o ExpressRouteCircuitOutput) CircuitProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.CircuitProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitPeeringResponseArrayOutput { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o ExpressRouteCircuitOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) ServiceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) ServiceProviderNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceProviderNotes }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) ServiceProviderProperties() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
		return v.ServiceProviderProperties
	}).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

func (o ExpressRouteCircuitOutput) ServiceProviderProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringPtrOutput { return v.ServiceProviderProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitOutput) Sku() ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) ExpressRouteCircuitSkuResponsePtrOutput { return v.Sku }).(ExpressRouteCircuitSkuResponsePtrOutput)
}

func (o ExpressRouteCircuitOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExpressRouteCircuitOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitOutput{})
}
