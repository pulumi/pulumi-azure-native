


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Host struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput            `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	MoName             pulumi.StringOutput               `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput            `pulumi:"moRefId"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput            `pulumi:"vCenterId"`
}


func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:Host"),
		},
	})
	opts = append(opts, aliases)
	var resource Host
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hostState struct {
}

type HostState struct {
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	HostName          *string           `pulumi:"hostName"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	MoRefId           *string           `pulumi:"moRefId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VCenterId         *string           `pulumi:"vCenterId"`
}


type HostArgs struct {
	ExtendedLocation  ExtendedLocationPtrInput
	HostName          pulumi.StringPtrInput
	InventoryItemId   pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VCenterId         pulumi.StringPtrInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((*Host)(nil))
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Host)(nil))
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HostOutput{})
}
