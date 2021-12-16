


package v20210312preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FluidRelayServer struct {
	pulumi.CustomResourceState

	FluidRelayEndpoints FluidRelayEndpointsResponseOutput `pulumi:"fluidRelayEndpoints"`
	FrsTenantId         pulumi.StringOutput               `pulumi:"frsTenantId"`
	Location            pulumi.StringOutput               `pulumi:"location"`
	Name                pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState   pulumi.StringPtrOutput            `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput            `pulumi:"tags"`
	Type                pulumi.StringOutput               `pulumi:"type"`
}


func NewFluidRelayServer(ctx *pulumi.Context,
	name string, args *FluidRelayServerArgs, opts ...pulumi.ResourceOption) (*FluidRelayServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:fluidrelay:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210615preview:FluidRelayServer"),
		},
	})
	opts = append(opts, aliases)
	var resource FluidRelayServer
	err := ctx.RegisterResource("azure-native:fluidrelay/v20210312preview:FluidRelayServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFluidRelayServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FluidRelayServerState, opts ...pulumi.ResourceOption) (*FluidRelayServer, error) {
	var resource FluidRelayServer
	err := ctx.ReadResource("azure-native:fluidrelay/v20210312preview:FluidRelayServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fluidRelayServerState struct {
}

type FluidRelayServerState struct {
}

func (FluidRelayServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fluidRelayServerState)(nil)).Elem()
}

type fluidRelayServerArgs struct {
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroup     string            `pulumi:"resourceGroup"`
	Tags              map[string]string `pulumi:"tags"`
}


type FluidRelayServerArgs struct {
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroup     pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (FluidRelayServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fluidRelayServerArgs)(nil)).Elem()
}

type FluidRelayServerInput interface {
	pulumi.Input

	ToFluidRelayServerOutput() FluidRelayServerOutput
	ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput
}

func (*FluidRelayServer) ElementType() reflect.Type {
	return reflect.TypeOf((**FluidRelayServer)(nil)).Elem()
}

func (i *FluidRelayServer) ToFluidRelayServerOutput() FluidRelayServerOutput {
	return i.ToFluidRelayServerOutputWithContext(context.Background())
}

func (i *FluidRelayServer) ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FluidRelayServerOutput)
}

type FluidRelayServerOutput struct{ *pulumi.OutputState }

func (FluidRelayServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FluidRelayServer)(nil)).Elem()
}

func (o FluidRelayServerOutput) ToFluidRelayServerOutput() FluidRelayServerOutput {
	return o
}

func (o FluidRelayServerOutput) ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FluidRelayServerOutput{})
}
