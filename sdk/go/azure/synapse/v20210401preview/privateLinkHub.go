


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkHub struct {
	pulumi.CustomResourceState

	Location                   pulumi.StringOutput                                                `pulumi:"location"`
	Name                       pulumi.StringOutput                                                `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringPtrOutput                                             `pulumi:"provisioningState"`
	Tags                       pulumi.StringMapOutput                                             `pulumi:"tags"`
	Type                       pulumi.StringOutput                                                `pulumi:"type"`
}


func NewPrivateLinkHub(ctx *pulumi.Context,
	name string, args *PrivateLinkHubArgs, opts ...pulumi.ResourceOption) (*PrivateLinkHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:PrivateLinkHub"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkHub
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:PrivateLinkHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkHubState, opts ...pulumi.ResourceOption) (*PrivateLinkHub, error) {
	var resource PrivateLinkHub
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:PrivateLinkHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkHubState struct {
}

type PrivateLinkHubState struct {
}

func (PrivateLinkHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkHubState)(nil)).Elem()
}

type privateLinkHubArgs struct {
	Location           *string           `pulumi:"location"`
	PrivateLinkHubName *string           `pulumi:"privateLinkHubName"`
	ProvisioningState  *string           `pulumi:"provisioningState"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
}


type PrivateLinkHubArgs struct {
	Location           pulumi.StringPtrInput
	PrivateLinkHubName pulumi.StringPtrInput
	ProvisioningState  pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (PrivateLinkHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkHubArgs)(nil)).Elem()
}

type PrivateLinkHubInput interface {
	pulumi.Input

	ToPrivateLinkHubOutput() PrivateLinkHubOutput
	ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput
}

func (*PrivateLinkHub) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkHub)(nil)).Elem()
}

func (i *PrivateLinkHub) ToPrivateLinkHubOutput() PrivateLinkHubOutput {
	return i.ToPrivateLinkHubOutputWithContext(context.Background())
}

func (i *PrivateLinkHub) ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkHubOutput)
}

type PrivateLinkHubOutput struct{ *pulumi.OutputState }

func (PrivateLinkHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkHub)(nil)).Elem()
}

func (o PrivateLinkHubOutput) ToPrivateLinkHubOutput() PrivateLinkHubOutput {
	return o
}

func (o PrivateLinkHubOutput) ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput {
	return o
}

func (o PrivateLinkHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkHubOutput) PrivateEndpointConnections() PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkHub) PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput)
}

func (o PrivateLinkHubOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkHubOutput{})
}
