


package deviceupdate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionProxy struct {
	pulumi.CustomResourceState

	ETag                  pulumi.StringOutput                    `pulumi:"eTag"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                    `pulumi:"provisioningState"`
	RemotePrivateEndpoint RemotePrivateEndpointResponsePtrOutput `pulumi:"remotePrivateEndpoint"`
	SystemData            SystemDataResponseOutput               `pulumi:"systemData"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
}


func NewPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionProxyArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deviceupdate:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-nextgen:deviceupdate/v20200301preview:PrivateEndpointConnectionProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionProxy
	err := ctx.RegisterResource("azure-native:deviceupdate:PrivateEndpointConnectionProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionProxyState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	var resource PrivateEndpointConnectionProxy
	err := ctx.ReadResource("azure-native:deviceupdate:PrivateEndpointConnectionProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionProxyState struct {
}

type PrivateEndpointConnectionProxyState struct {
}

func (PrivateEndpointConnectionProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyState)(nil)).Elem()
}

type privateEndpointConnectionProxyArgs struct {
	AccountName                      string                 `pulumi:"accountName"`
	PrivateEndpointConnectionProxyId *string                `pulumi:"privateEndpointConnectionProxyId"`
	RemotePrivateEndpoint            *RemotePrivateEndpoint `pulumi:"remotePrivateEndpoint"`
	ResourceGroupName                string                 `pulumi:"resourceGroupName"`
}


type PrivateEndpointConnectionProxyArgs struct {
	AccountName                      pulumi.StringInput
	PrivateEndpointConnectionProxyId pulumi.StringPtrInput
	RemotePrivateEndpoint            RemotePrivateEndpointPtrInput
	ResourceGroupName                pulumi.StringInput
}

func (PrivateEndpointConnectionProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyArgs)(nil)).Elem()
}

type PrivateEndpointConnectionProxyInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput
	ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput
}

func (*PrivateEndpointConnectionProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProxy)(nil))
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return i.ToPrivateEndpointConnectionProxyOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionProxyOutput)
}

type PrivateEndpointConnectionProxyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProxy)(nil))
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return o
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionProxyInput)(nil)).Elem(), &PrivateEndpointConnectionProxy{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionProxyOutput{})
}
