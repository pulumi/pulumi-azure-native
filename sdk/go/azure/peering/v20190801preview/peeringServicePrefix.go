


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PeeringServicePrefix struct {
	pulumi.CustomResourceState

	LearnedType           pulumi.StringPtrOutput `pulumi:"learnedType"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Prefix                pulumi.StringPtrOutput `pulumi:"prefix"`
	PrefixValidationState pulumi.StringPtrOutput `pulumi:"prefixValidationState"`
	ProvisioningState     pulumi.StringOutput    `pulumi:"provisioningState"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewPeeringServicePrefix(ctx *pulumi.Context,
	name string, args *PeeringServicePrefixArgs, opts ...pulumi.ResourceOption) (*PeeringServicePrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringServiceName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeeringServicePrefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeeringServicePrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource PeeringServicePrefix
	err := ctx.RegisterResource("azure-native:peering/v20190801preview:PeeringServicePrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeeringServicePrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringServicePrefixState, opts ...pulumi.ResourceOption) (*PeeringServicePrefix, error) {
	var resource PeeringServicePrefix
	err := ctx.ReadResource("azure-native:peering/v20190801preview:PeeringServicePrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type peeringServicePrefixState struct {
}

type PeeringServicePrefixState struct {
}

func (PeeringServicePrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServicePrefixState)(nil)).Elem()
}

type peeringServicePrefixArgs struct {
	LearnedType           *string `pulumi:"learnedType"`
	PeeringServiceName    string  `pulumi:"peeringServiceName"`
	Prefix                *string `pulumi:"prefix"`
	PrefixName            *string `pulumi:"prefixName"`
	PrefixValidationState *string `pulumi:"prefixValidationState"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type PeeringServicePrefixArgs struct {
	LearnedType           pulumi.StringPtrInput
	PeeringServiceName    pulumi.StringInput
	Prefix                pulumi.StringPtrInput
	PrefixName            pulumi.StringPtrInput
	PrefixValidationState pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
}

func (PeeringServicePrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServicePrefixArgs)(nil)).Elem()
}

type PeeringServicePrefixInput interface {
	pulumi.Input

	ToPeeringServicePrefixOutput() PeeringServicePrefixOutput
	ToPeeringServicePrefixOutputWithContext(ctx context.Context) PeeringServicePrefixOutput
}

func (*PeeringServicePrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServicePrefix)(nil)).Elem()
}

func (i *PeeringServicePrefix) ToPeeringServicePrefixOutput() PeeringServicePrefixOutput {
	return i.ToPeeringServicePrefixOutputWithContext(context.Background())
}

func (i *PeeringServicePrefix) ToPeeringServicePrefixOutputWithContext(ctx context.Context) PeeringServicePrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringServicePrefixOutput)
}

type PeeringServicePrefixOutput struct{ *pulumi.OutputState }

func (PeeringServicePrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeeringServicePrefix)(nil)).Elem()
}

func (o PeeringServicePrefixOutput) ToPeeringServicePrefixOutput() PeeringServicePrefixOutput {
	return o
}

func (o PeeringServicePrefixOutput) ToPeeringServicePrefixOutputWithContext(ctx context.Context) PeeringServicePrefixOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PeeringServicePrefixOutput{})
}
