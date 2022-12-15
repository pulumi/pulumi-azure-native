


package v20220601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FluidRelayServer struct {
	pulumi.CustomResourceState

	Encryption          EncryptionPropertiesResponsePtrOutput `pulumi:"encryption"`
	FluidRelayEndpoints FluidRelayEndpointsResponseOutput     `pulumi:"fluidRelayEndpoints"`
	FrsTenantId         pulumi.StringOutput                   `pulumi:"frsTenantId"`
	Identity            IdentityResponsePtrOutput             `pulumi:"identity"`
	Location            pulumi.StringOutput                   `pulumi:"location"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState   pulumi.StringPtrOutput                `pulumi:"provisioningState"`
	Storagesku          pulumi.StringPtrOutput                `pulumi:"storagesku"`
	SystemData          SystemDataResponseOutput              `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                `pulumi:"tags"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
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
			Type: pulumi.String("azure-native:fluidrelay/v20210312preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210615preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210830preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210910preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220215:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220421:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220511:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220526:FluidRelayServer"),
		},
	})
	opts = append(opts, aliases)
	var resource FluidRelayServer
	err := ctx.RegisterResource("azure-native:fluidrelay/v20220601:FluidRelayServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFluidRelayServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FluidRelayServerState, opts ...pulumi.ResourceOption) (*FluidRelayServer, error) {
	var resource FluidRelayServer
	err := ctx.ReadResource("azure-native:fluidrelay/v20220601:FluidRelayServer", name, id, state, &resource, opts...)
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
	Encryption           *EncryptionProperties `pulumi:"encryption"`
	FluidRelayServerName *string               `pulumi:"fluidRelayServerName"`
	Identity             *Identity             `pulumi:"identity"`
	Location             *string               `pulumi:"location"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	ResourceGroup        string                `pulumi:"resourceGroup"`
	Storagesku           *string               `pulumi:"storagesku"`
	Tags                 map[string]string     `pulumi:"tags"`
}


type FluidRelayServerArgs struct {
	Encryption           EncryptionPropertiesPtrInput
	FluidRelayServerName pulumi.StringPtrInput
	Identity             IdentityPtrInput
	Location             pulumi.StringPtrInput
	ProvisioningState    pulumi.StringPtrInput
	ResourceGroup        pulumi.StringInput
	Storagesku           pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
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

func (o FluidRelayServerOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o FluidRelayServerOutput) FluidRelayEndpoints() FluidRelayEndpointsResponseOutput {
	return o.ApplyT(func(v *FluidRelayServer) FluidRelayEndpointsResponseOutput { return v.FluidRelayEndpoints }).(FluidRelayEndpointsResponseOutput)
}

func (o FluidRelayServerOutput) FrsTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.FrsTenantId }).(pulumi.StringOutput)
}

func (o FluidRelayServerOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o FluidRelayServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o FluidRelayServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FluidRelayServerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FluidRelayServerOutput) Storagesku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringPtrOutput { return v.Storagesku }).(pulumi.StringPtrOutput)
}

func (o FluidRelayServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FluidRelayServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FluidRelayServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FluidRelayServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FluidRelayServerOutput{})
}
