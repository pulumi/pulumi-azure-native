


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Identity                          IdentityResponsePtrOutput                       `pulumi:"identity"`
	Location                          pulumi.StringPtrOutput                          `pulumi:"location"`
	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	Sku                               SkuResponsePtrOutput                            `pulumi:"sku"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:machinelearningservices:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:machinelearningservices:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionState struct {
}

type PrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	Identity                          *Identity                         `pulumi:"identity"`
	Location                          *string                           `pulumi:"location"`
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	Sku                               *Sku                              `pulumi:"sku"`
	Tags                              map[string]string                 `pulumi:"tags"`
	WorkspaceName                     string                            `pulumi:"workspaceName"`
}


type PrivateEndpointConnectionArgs struct {
	Identity                          IdentityPtrInput
	Location                          pulumi.StringPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	Sku                               SkuPtrInput
	Tags                              pulumi.StringMapInput
	WorkspaceName                     pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o PrivateEndpointConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o PrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
