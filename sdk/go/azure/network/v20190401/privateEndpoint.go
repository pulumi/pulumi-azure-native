


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpoint struct {
	pulumi.CustomResourceState

	Etag                                pulumi.StringPtrOutput                          `pulumi:"etag"`
	Location                            pulumi.StringPtrOutput                          `pulumi:"location"`
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionResponseArrayOutput `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                pulumi.StringOutput                             `pulumi:"name"`
	NetworkInterfaces                   NetworkInterfaceResponseArrayOutput             `pulumi:"networkInterfaces"`
	PrivateLinkServiceConnections       PrivateLinkServiceConnectionResponseArrayOutput `pulumi:"privateLinkServiceConnections"`
	ProvisioningState                   pulumi.StringOutput                             `pulumi:"provisioningState"`
	Subnet                              SubnetResponsePtrOutput                         `pulumi:"subnet"`
	Tags                                pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                                pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpoint(ctx *pulumi.Context,
	name string, args *PrivateEndpointArgs, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:PrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpoint
	err := ctx.RegisterResource("azure-native:network/v20190401:PrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointState, opts ...pulumi.ResourceOption) (*PrivateEndpoint, error) {
	var resource PrivateEndpoint
	err := ctx.ReadResource("azure-native:network/v20190401:PrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointState struct {
}

type PrivateEndpointState struct {
}

func (PrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointState)(nil)).Elem()
}

type privateEndpointArgs struct {
	Etag                                *string                        `pulumi:"etag"`
	Id                                  *string                        `pulumi:"id"`
	Location                            *string                        `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection `pulumi:"manualPrivateLinkServiceConnections"`
	PrivateEndpointName                 *string                        `pulumi:"privateEndpointName"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnection `pulumi:"privateLinkServiceConnections"`
	ResourceGroupName                   string                         `pulumi:"resourceGroupName"`
	Subnet                              *SubnetType                    `pulumi:"subnet"`
	Tags                                map[string]string              `pulumi:"tags"`
}


type PrivateEndpointArgs struct {
	Etag                                pulumi.StringPtrInput
	Id                                  pulumi.StringPtrInput
	Location                            pulumi.StringPtrInput
	ManualPrivateLinkServiceConnections PrivateLinkServiceConnectionArrayInput
	PrivateEndpointName                 pulumi.StringPtrInput
	PrivateLinkServiceConnections       PrivateLinkServiceConnectionArrayInput
	ResourceGroupName                   pulumi.StringInput
	Subnet                              SubnetTypePtrInput
	Tags                                pulumi.StringMapInput
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointArgs)(nil)).Elem()
}

type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput
}

func (*PrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil))
}

func (i *PrivateEndpoint) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i *PrivateEndpoint) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil))
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
}
