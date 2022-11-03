


package v20190901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type PeeringService struct {
	pulumi.CustomResourceState

	Location               pulumi.StringOutput    `pulumi:"location"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	PeeringServiceLocation pulumi.StringPtrOutput `pulumi:"peeringServiceLocation"`
	PeeringServiceProvider pulumi.StringPtrOutput `pulumi:"peeringServiceProvider"`
	ProvisioningState      pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags                   pulumi.StringMapOutput `pulumi:"tags"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
}


func NewPeeringService(ctx *pulumi.Context,
	name string, args *PeeringServiceArgs, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20221001:PeeringService"),
		},
	})
	opts = append(opts, aliases)
	var resource PeeringService
	err := ctx.RegisterResource("azure-native:peering/v20190901preview:PeeringService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeeringService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringServiceState, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	var resource PeeringService
	err := ctx.ReadResource("azure-native:peering/v20190901preview:PeeringService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type peeringServiceState struct {
}

type PeeringServiceState struct {
}

func (PeeringServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceState)(nil)).Elem()
}

type peeringServiceArgs struct {
	Location               *string           `pulumi:"location"`
	PeeringServiceLocation *string           `pulumi:"peeringServiceLocation"`
	PeeringServiceName     *string           `pulumi:"peeringServiceName"`
	PeeringServiceProvider *string           `pulumi:"peeringServiceProvider"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type PeeringServiceArgs struct {
	Location               pulumi.StringPtrInput
	PeeringServiceLocation pulumi.StringPtrInput
	PeeringServiceName     pulumi.StringPtrInput
	PeeringServiceProvider pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (PeeringServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceArgs)(nil)).Elem()
}

type PeeringServiceInput interface {
	pulumi.Input

	ToPeeringServiceOutput() PeeringServiceOutput
	ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput
}

func (*PeeringService) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringService)(nil)).Elem()
}

func (i *PeeringService) ToPeeringServiceOutput() PeeringServiceOutput {
	return i.ToPeeringServiceOutputWithContext(context.Background())
}

func (i *PeeringService) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceOutput)
}

type PeeringServiceOutput struct{ *pulumi.OutputState }

func (PeeringServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringService)(nil)).Elem()
}

func (o PeeringServiceOutput) ToPeeringServiceOutput() PeeringServiceOutput {
	return o
}

func (o PeeringServiceOutput) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return o
}

func (o PeeringServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PeeringServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PeeringServiceOutput) PeeringServiceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.PeeringServiceLocation }).(pulumi.StringPtrOutput)
}

func (o PeeringServiceOutput) PeeringServiceProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringPtrOutput { return v.PeeringServiceProvider }).(pulumi.StringPtrOutput)
}

func (o PeeringServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PeeringServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PeeringServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeeringService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeeringServiceOutput{})
}
