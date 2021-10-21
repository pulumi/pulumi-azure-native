


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalReachConnection struct {
	pulumi.CustomResourceState

	AddressPrefix           pulumi.StringOutput    `pulumi:"addressPrefix"`
	AuthorizationKey        pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	CircuitConnectionStatus pulumi.StringOutput    `pulumi:"circuitConnectionStatus"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	PeerExpressRouteCircuit pulumi.StringPtrOutput `pulumi:"peerExpressRouteCircuit"`
	ProvisioningState       pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                    pulumi.StringOutput    `pulumi:"type"`
}


func NewGlobalReachConnection(ctx *pulumi.Context,
	name string, args *GlobalReachConnectionArgs, opts ...pulumi.ResourceOption) (*GlobalReachConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:avs/v20200717preview:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210101preview:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:GlobalReachConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:avs/v20210601:GlobalReachConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource GlobalReachConnection
	err := ctx.RegisterResource("azure-native:avs/v20200717preview:GlobalReachConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGlobalReachConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalReachConnectionState, opts ...pulumi.ResourceOption) (*GlobalReachConnection, error) {
	var resource GlobalReachConnection
	err := ctx.ReadResource("azure-native:avs/v20200717preview:GlobalReachConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type globalReachConnectionState struct {
}

type GlobalReachConnectionState struct {
}

func (GlobalReachConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReachConnectionState)(nil)).Elem()
}

type globalReachConnectionArgs struct {
	AuthorizationKey          *string `pulumi:"authorizationKey"`
	GlobalReachConnectionName *string `pulumi:"globalReachConnectionName"`
	PeerExpressRouteCircuit   *string `pulumi:"peerExpressRouteCircuit"`
	PrivateCloudName          string  `pulumi:"privateCloudName"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
}


type GlobalReachConnectionArgs struct {
	AuthorizationKey          pulumi.StringPtrInput
	GlobalReachConnectionName pulumi.StringPtrInput
	PeerExpressRouteCircuit   pulumi.StringPtrInput
	PrivateCloudName          pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
}

func (GlobalReachConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalReachConnectionArgs)(nil)).Elem()
}

type GlobalReachConnectionInput interface {
	pulumi.Input

	ToGlobalReachConnectionOutput() GlobalReachConnectionOutput
	ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput
}

func (*GlobalReachConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalReachConnection)(nil))
}

func (i *GlobalReachConnection) ToGlobalReachConnectionOutput() GlobalReachConnectionOutput {
	return i.ToGlobalReachConnectionOutputWithContext(context.Background())
}

func (i *GlobalReachConnection) ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalReachConnectionOutput)
}

type GlobalReachConnectionOutput struct{ *pulumi.OutputState }

func (GlobalReachConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalReachConnection)(nil))
}

func (o GlobalReachConnectionOutput) ToGlobalReachConnectionOutput() GlobalReachConnectionOutput {
	return o
}

func (o GlobalReachConnectionOutput) ToGlobalReachConnectionOutputWithContext(ctx context.Context) GlobalReachConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlobalReachConnectionOutput{})
}
