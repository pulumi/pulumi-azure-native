


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateZone struct {
	pulumi.CustomResourceState

	Etag                                           pulumi.StringPtrOutput `pulumi:"etag"`
	InternalId                                     pulumi.StringOutput    `pulumi:"internalId"`
	Location                                       pulumi.StringPtrOutput `pulumi:"location"`
	MaxNumberOfRecordSets                          pulumi.Float64Output   `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfVirtualNetworkLinks                 pulumi.Float64Output   `pulumi:"maxNumberOfVirtualNetworkLinks"`
	MaxNumberOfVirtualNetworkLinksWithRegistration pulumi.Float64Output   `pulumi:"maxNumberOfVirtualNetworkLinksWithRegistration"`
	Name                                           pulumi.StringOutput    `pulumi:"name"`
	NumberOfRecordSets                             pulumi.Float64Output   `pulumi:"numberOfRecordSets"`
	NumberOfVirtualNetworkLinks                    pulumi.Float64Output   `pulumi:"numberOfVirtualNetworkLinks"`
	NumberOfVirtualNetworkLinksWithRegistration    pulumi.Float64Output   `pulumi:"numberOfVirtualNetworkLinksWithRegistration"`
	ProvisioningState                              pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags                                           pulumi.StringMapOutput `pulumi:"tags"`
	Type                                           pulumi.StringOutput    `pulumi:"type"`
}


func NewPrivateZone(ctx *pulumi.Context,
	name string, args *PrivateZoneArgs, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:PrivateZone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:PrivateZone"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateZone
	err := ctx.RegisterResource("azure-native:network/v20200601:PrivateZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateZoneState, opts ...pulumi.ResourceOption) (*PrivateZone, error) {
	var resource PrivateZone
	err := ctx.ReadResource("azure-native:network/v20200601:PrivateZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateZoneState struct {
}

type PrivateZoneState struct {
}

func (PrivateZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneState)(nil)).Elem()
}

type privateZoneArgs struct {
	Etag              *string           `pulumi:"etag"`
	Location          *string           `pulumi:"location"`
	PrivateZoneName   *string           `pulumi:"privateZoneName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type PrivateZoneArgs struct {
	Etag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	PrivateZoneName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (PrivateZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateZoneArgs)(nil)).Elem()
}

type PrivateZoneInput interface {
	pulumi.Input

	ToPrivateZoneOutput() PrivateZoneOutput
	ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput
}

func (*PrivateZone) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateZone)(nil))
}

func (i *PrivateZone) ToPrivateZoneOutput() PrivateZoneOutput {
	return i.ToPrivateZoneOutputWithContext(context.Background())
}

func (i *PrivateZone) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateZoneOutput)
}

type PrivateZoneOutput struct{ *pulumi.OutputState }

func (PrivateZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateZone)(nil))
}

func (o PrivateZoneOutput) ToPrivateZoneOutput() PrivateZoneOutput {
	return o
}

func (o PrivateZoneOutput) ToPrivateZoneOutputWithContext(ctx context.Context) PrivateZoneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateZoneOutput{})
}
