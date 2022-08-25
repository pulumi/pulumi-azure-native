


package v20200415

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Origin struct {
	pulumi.CustomResourceState

	Enabled                    pulumi.BoolPtrOutput   `pulumi:"enabled"`
	HostName                   pulumi.StringOutput    `pulumi:"hostName"`
	HttpPort                   pulumi.IntPtrOutput    `pulumi:"httpPort"`
	HttpsPort                  pulumi.IntPtrOutput    `pulumi:"httpsPort"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	OriginHostHeader           pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	Priority                   pulumi.IntPtrOutput    `pulumi:"priority"`
	PrivateEndpointStatus      pulumi.StringOutput    `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           pulumi.StringPtrOutput `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage pulumi.StringPtrOutput `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        pulumi.StringPtrOutput `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      pulumi.StringPtrOutput `pulumi:"privateLinkResourceId"`
	ProvisioningState          pulumi.StringOutput    `pulumi:"provisioningState"`
	ResourceState              pulumi.StringOutput    `pulumi:"resourceState"`
	Type                       pulumi.StringOutput    `pulumi:"type"`
	Weight                     pulumi.IntPtrOutput    `pulumi:"weight"`
}


func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOption) (*Origin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Origin"),
		},
	})
	opts = append(opts, aliases)
	var resource Origin
	err := ctx.RegisterResource("azure-native:cdn/v20200415:Origin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginState, opts ...pulumi.ResourceOption) (*Origin, error) {
	var resource Origin
	err := ctx.ReadResource("azure-native:cdn/v20200415:Origin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type originState struct {
}

type OriginState struct {
}

func (OriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*originState)(nil)).Elem()
}

type originArgs struct {
	Enabled                    *bool   `pulumi:"enabled"`
	EndpointName               string  `pulumi:"endpointName"`
	HostName                   string  `pulumi:"hostName"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	OriginHostHeader           *string `pulumi:"originHostHeader"`
	OriginName                 *string `pulumi:"originName"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	PrivateLinkLocation        *string `pulumi:"privateLinkLocation"`
	PrivateLinkResourceId      *string `pulumi:"privateLinkResourceId"`
	ProfileName                string  `pulumi:"profileName"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	Weight                     *int    `pulumi:"weight"`
}


type OriginArgs struct {
	Enabled                    pulumi.BoolPtrInput
	EndpointName               pulumi.StringInput
	HostName                   pulumi.StringInput
	HttpPort                   pulumi.IntPtrInput
	HttpsPort                  pulumi.IntPtrInput
	OriginHostHeader           pulumi.StringPtrInput
	OriginName                 pulumi.StringPtrInput
	Priority                   pulumi.IntPtrInput
	PrivateLinkAlias           pulumi.StringPtrInput
	PrivateLinkApprovalMessage pulumi.StringPtrInput
	PrivateLinkLocation        pulumi.StringPtrInput
	PrivateLinkResourceId      pulumi.StringPtrInput
	ProfileName                pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	Weight                     pulumi.IntPtrInput
}

func (OriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originArgs)(nil)).Elem()
}

type OriginInput interface {
	pulumi.Input

	ToOriginOutput() OriginOutput
	ToOriginOutputWithContext(ctx context.Context) OriginOutput
}

func (*Origin) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (i *Origin) ToOriginOutput() OriginOutput {
	return i.ToOriginOutputWithContext(context.Background())
}

func (i *Origin) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginOutput)
}

type OriginOutput struct{ *pulumi.OutputState }

func (OriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (o OriginOutput) ToOriginOutput() OriginOutput {
	return o
}

func (o OriginOutput) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return o
}

func (o OriginOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o OriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o OriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o OriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o OriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OriginOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

func (o OriginOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o OriginOutput) PrivateEndpointStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.PrivateEndpointStatus }).(pulumi.StringOutput)
}

func (o OriginOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o OriginOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o OriginOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

func (o OriginOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o OriginOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OriginOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o OriginOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o OriginOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OriginOutput{})
}
