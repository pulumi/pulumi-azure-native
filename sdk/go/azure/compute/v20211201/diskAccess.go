


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskAccess struct {
	pulumi.CustomResourceState

	ExtendedLocation           ExtendedLocationResponsePtrOutput            `pulumi:"extendedLocation"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	TimeCreated                pulumi.StringOutput                          `pulumi:"timeCreated"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewDiskAccess(ctx *pulumi.Context,
	name string, args *DiskAccessArgs, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:DiskAccess"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:DiskAccess"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskAccess
	err := ctx.RegisterResource("azure-native:compute/v20211201:DiskAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessState, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	var resource DiskAccess
	err := ctx.ReadResource("azure-native:compute/v20211201:DiskAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskAccessState struct {
}

type DiskAccessState struct {
}

func (DiskAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessState)(nil)).Elem()
}

type diskAccessArgs struct {
	DiskAccessName    *string           `pulumi:"diskAccessName"`
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type DiskAccessArgs struct {
	DiskAccessName    pulumi.StringPtrInput
	ExtendedLocation  ExtendedLocationPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (DiskAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessArgs)(nil)).Elem()
}

type DiskAccessInput interface {
	pulumi.Input

	ToDiskAccessOutput() DiskAccessOutput
	ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput
}

func (*DiskAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (i *DiskAccess) ToDiskAccessOutput() DiskAccessOutput {
	return i.ToDiskAccessOutputWithContext(context.Background())
}

func (i *DiskAccess) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessOutput)
}

type DiskAccessOutput struct{ *pulumi.OutputState }

func (DiskAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccess)(nil)).Elem()
}

func (o DiskAccessOutput) ToDiskAccessOutput() DiskAccessOutput {
	return o
}

func (o DiskAccessOutput) ToDiskAccessOutputWithContext(ctx context.Context) DiskAccessOutput {
	return o
}

func (o DiskAccessOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *DiskAccess) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o DiskAccessOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DiskAccessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskAccessOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DiskAccess) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o DiskAccessOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskAccessOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskAccessOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o DiskAccessOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskAccess) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskAccessOutput{})
}
