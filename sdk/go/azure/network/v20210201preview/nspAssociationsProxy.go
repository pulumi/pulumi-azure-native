


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NspAssociationsProxy struct {
	pulumi.CustomResourceState

	AccessMode            pulumi.StringPtrOutput       `pulumi:"accessMode"`
	HasProvisioningIssues pulumi.StringOutput          `pulumi:"hasProvisioningIssues"`
	Location              pulumi.StringPtrOutput       `pulumi:"location"`
	Name                  pulumi.StringOutput          `pulumi:"name"`
	PrivateLinkResource   SubResourceResponsePtrOutput `pulumi:"privateLinkResource"`
	Profile               SubResourceResponsePtrOutput `pulumi:"profile"`
	ProvisioningState     pulumi.StringOutput          `pulumi:"provisioningState"`
	Tags                  pulumi.StringMapOutput       `pulumi:"tags"`
	Type                  pulumi.StringOutput          `pulumi:"type"`
}


func NewNspAssociationsProxy(ctx *pulumi.Context,
	name string, args *NspAssociationsProxyArgs, opts ...pulumi.ResourceOption) (*NspAssociationsProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NspAssociationsProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource NspAssociationsProxy
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NspAssociationsProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNspAssociationsProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspAssociationsProxyState, opts ...pulumi.ResourceOption) (*NspAssociationsProxy, error) {
	var resource NspAssociationsProxy
	err := ctx.ReadResource("azure-native:network/v20210201preview:NspAssociationsProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nspAssociationsProxyState struct {
}

type NspAssociationsProxyState struct {
}

func (NspAssociationsProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationsProxyState)(nil)).Elem()
}

type nspAssociationsProxyArgs struct {
	AccessMode                   *string           `pulumi:"accessMode"`
	AssociationName              *string           `pulumi:"associationName"`
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	Name                         *string           `pulumi:"name"`
	NetworkSecurityPerimeterName string            `pulumi:"networkSecurityPerimeterName"`
	PrivateLinkResource          *SubResource      `pulumi:"privateLinkResource"`
	Profile                      *SubResource      `pulumi:"profile"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type NspAssociationsProxyArgs struct {
	AccessMode                   pulumi.StringPtrInput
	AssociationName              pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	NetworkSecurityPerimeterName pulumi.StringInput
	PrivateLinkResource          SubResourcePtrInput
	Profile                      SubResourcePtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (NspAssociationsProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationsProxyArgs)(nil)).Elem()
}

type NspAssociationsProxyInput interface {
	pulumi.Input

	ToNspAssociationsProxyOutput() NspAssociationsProxyOutput
	ToNspAssociationsProxyOutputWithContext(ctx context.Context) NspAssociationsProxyOutput
}

func (*NspAssociationsProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociationsProxy)(nil)).Elem()
}

func (i *NspAssociationsProxy) ToNspAssociationsProxyOutput() NspAssociationsProxyOutput {
	return i.ToNspAssociationsProxyOutputWithContext(context.Background())
}

func (i *NspAssociationsProxy) ToNspAssociationsProxyOutputWithContext(ctx context.Context) NspAssociationsProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspAssociationsProxyOutput)
}

type NspAssociationsProxyOutput struct{ *pulumi.OutputState }

func (NspAssociationsProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociationsProxy)(nil)).Elem()
}

func (o NspAssociationsProxyOutput) ToNspAssociationsProxyOutput() NspAssociationsProxyOutput {
	return o
}

func (o NspAssociationsProxyOutput) ToNspAssociationsProxyOutputWithContext(ctx context.Context) NspAssociationsProxyOutput {
	return o
}

func (o NspAssociationsProxyOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringPtrOutput { return v.AccessMode }).(pulumi.StringPtrOutput)
}

func (o NspAssociationsProxyOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringOutput { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

func (o NspAssociationsProxyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NspAssociationsProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NspAssociationsProxyOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) SubResourceResponsePtrOutput { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

func (o NspAssociationsProxyOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) SubResourceResponsePtrOutput { return v.Profile }).(SubResourceResponsePtrOutput)
}

func (o NspAssociationsProxyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NspAssociationsProxyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NspAssociationsProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociationsProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspAssociationsProxyOutput{})
}
