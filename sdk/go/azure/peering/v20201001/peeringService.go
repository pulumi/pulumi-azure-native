


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PeeringService struct {
	pulumi.CustomResourceState

	Location               pulumi.StringOutput                `pulumi:"location"`
	Name                   pulumi.StringOutput                `pulumi:"name"`
	PeeringServiceLocation pulumi.StringPtrOutput             `pulumi:"peeringServiceLocation"`
	PeeringServiceProvider pulumi.StringPtrOutput             `pulumi:"peeringServiceProvider"`
	ProvisioningState      pulumi.StringOutput                `pulumi:"provisioningState"`
	Sku                    PeeringServiceSkuResponsePtrOutput `pulumi:"sku"`
	Tags                   pulumi.StringMapOutput             `pulumi:"tags"`
	Type                   pulumi.StringOutput                `pulumi:"type"`
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
			Type: pulumi.String("azure-native:peering/v20190901preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeeringService"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeeringService"),
		},
	})
	opts = append(opts, aliases)
	var resource PeeringService
	err := ctx.RegisterResource("azure-native:peering/v20201001:PeeringService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeeringService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringServiceState, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	var resource PeeringService
	err := ctx.ReadResource("azure-native:peering/v20201001:PeeringService", name, id, state, &resource, opts...)
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
	Location               *string            `pulumi:"location"`
	PeeringServiceLocation *string            `pulumi:"peeringServiceLocation"`
	PeeringServiceName     *string            `pulumi:"peeringServiceName"`
	PeeringServiceProvider *string            `pulumi:"peeringServiceProvider"`
	ResourceGroupName      string             `pulumi:"resourceGroupName"`
	Sku                    *PeeringServiceSku `pulumi:"sku"`
	Tags                   map[string]string  `pulumi:"tags"`
}


type PeeringServiceArgs struct {
	Location               pulumi.StringPtrInput
	PeeringServiceLocation pulumi.StringPtrInput
	PeeringServiceName     pulumi.StringPtrInput
	PeeringServiceProvider pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    PeeringServiceSkuPtrInput
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
	return reflect.TypeOf((*PeeringService)(nil))
}

func (i *PeeringService) ToPeeringServiceOutput() PeeringServiceOutput {
	return i.ToPeeringServiceOutputWithContext(context.Background())
}

func (i *PeeringService) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServiceOutput)
}

type PeeringServiceOutput struct{ *pulumi.OutputState }

func (PeeringServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeeringService)(nil))
}

func (o PeeringServiceOutput) ToPeeringServiceOutput() PeeringServiceOutput {
	return o
}

func (o PeeringServiceOutput) ToPeeringServiceOutputWithContext(ctx context.Context) PeeringServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PeeringServiceOutput{})
}
