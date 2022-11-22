


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayCustomDomain struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                         `pulumi:"name"`
	Properties GatewayCustomDomainPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                    `pulumi:"systemData"`
	Type       pulumi.StringOutput                         `pulumi:"type"`
}


func NewGatewayCustomDomain(ctx *pulumi.Context,
	name string, args *GatewayCustomDomainArgs, opts ...pulumi.ResourceOption) (*GatewayCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:GatewayCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:GatewayCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayCustomDomain
	err := ctx.RegisterResource("azure-native:appplatform/v20220501preview:GatewayCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCustomDomainState, opts ...pulumi.ResourceOption) (*GatewayCustomDomain, error) {
	var resource GatewayCustomDomain
	err := ctx.ReadResource("azure-native:appplatform/v20220501preview:GatewayCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayCustomDomainState struct {
}

type GatewayCustomDomainState struct {
}

func (GatewayCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCustomDomainState)(nil)).Elem()
}

type gatewayCustomDomainArgs struct {
	DomainName        *string                        `pulumi:"domainName"`
	GatewayName       string                         `pulumi:"gatewayName"`
	Properties        *GatewayCustomDomainProperties `pulumi:"properties"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	ServiceName       string                         `pulumi:"serviceName"`
}


type GatewayCustomDomainArgs struct {
	DomainName        pulumi.StringPtrInput
	GatewayName       pulumi.StringInput
	Properties        GatewayCustomDomainPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (GatewayCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCustomDomainArgs)(nil)).Elem()
}

type GatewayCustomDomainInput interface {
	pulumi.Input

	ToGatewayCustomDomainOutput() GatewayCustomDomainOutput
	ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput
}

func (*GatewayCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomain)(nil)).Elem()
}

func (i *GatewayCustomDomain) ToGatewayCustomDomainOutput() GatewayCustomDomainOutput {
	return i.ToGatewayCustomDomainOutputWithContext(context.Background())
}

func (i *GatewayCustomDomain) ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCustomDomainOutput)
}

type GatewayCustomDomainOutput struct{ *pulumi.OutputState }

func (GatewayCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCustomDomain)(nil)).Elem()
}

func (o GatewayCustomDomainOutput) ToGatewayCustomDomainOutput() GatewayCustomDomainOutput {
	return o
}

func (o GatewayCustomDomainOutput) ToGatewayCustomDomainOutputWithContext(ctx context.Context) GatewayCustomDomainOutput {
	return o
}

func (o GatewayCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayCustomDomainOutput) Properties() GatewayCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) GatewayCustomDomainPropertiesResponseOutput { return v.Properties }).(GatewayCustomDomainPropertiesResponseOutput)
}

func (o GatewayCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GatewayCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayCustomDomainOutput{})
}
