


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NspLink struct {
	pulumi.CustomResourceState

	AutoApprovedRemotePerimeterResourceId pulumi.StringPtrOutput   `pulumi:"autoApprovedRemotePerimeterResourceId"`
	Description                           pulumi.StringPtrOutput   `pulumi:"description"`
	Etag                                  pulumi.StringOutput      `pulumi:"etag"`
	LocalInboundProfiles                  pulumi.StringArrayOutput `pulumi:"localInboundProfiles"`
	LocalOutboundProfiles                 pulumi.StringArrayOutput `pulumi:"localOutboundProfiles"`
	Name                                  pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState                     pulumi.StringOutput      `pulumi:"provisioningState"`
	RemoteInboundProfiles                 pulumi.StringArrayOutput `pulumi:"remoteInboundProfiles"`
	RemoteOutboundProfiles                pulumi.StringArrayOutput `pulumi:"remoteOutboundProfiles"`
	RemotePerimeterGuid                   pulumi.StringOutput      `pulumi:"remotePerimeterGuid"`
	Status                                pulumi.StringOutput      `pulumi:"status"`
	Type                                  pulumi.StringOutput      `pulumi:"type"`
}


func NewNspLink(ctx *pulumi.Context,
	name string, args *NspLinkArgs, opts ...pulumi.ResourceOption) (*NspLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NspLink
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NspLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNspLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspLinkState, opts ...pulumi.ResourceOption) (*NspLink, error) {
	var resource NspLink
	err := ctx.ReadResource("azure-native:network/v20210201preview:NspLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nspLinkState struct {
}

type NspLinkState struct {
}

func (NspLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspLinkState)(nil)).Elem()
}

type nspLinkArgs struct {
	AutoApprovedRemotePerimeterResourceId *string  `pulumi:"autoApprovedRemotePerimeterResourceId"`
	Description                           *string  `pulumi:"description"`
	LinkName                              *string  `pulumi:"linkName"`
	LocalInboundProfiles                  []string `pulumi:"localInboundProfiles"`
	NetworkSecurityPerimeterName          string   `pulumi:"networkSecurityPerimeterName"`
	RemoteInboundProfiles                 []string `pulumi:"remoteInboundProfiles"`
	ResourceGroupName                     string   `pulumi:"resourceGroupName"`
}


type NspLinkArgs struct {
	AutoApprovedRemotePerimeterResourceId pulumi.StringPtrInput
	Description                           pulumi.StringPtrInput
	LinkName                              pulumi.StringPtrInput
	LocalInboundProfiles                  pulumi.StringArrayInput
	NetworkSecurityPerimeterName          pulumi.StringInput
	RemoteInboundProfiles                 pulumi.StringArrayInput
	ResourceGroupName                     pulumi.StringInput
}

func (NspLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspLinkArgs)(nil)).Elem()
}

type NspLinkInput interface {
	pulumi.Input

	ToNspLinkOutput() NspLinkOutput
	ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput
}

func (*NspLink) ElementType() reflect.Type {
	return reflect.TypeOf((**NspLink)(nil)).Elem()
}

func (i *NspLink) ToNspLinkOutput() NspLinkOutput {
	return i.ToNspLinkOutputWithContext(context.Background())
}

func (i *NspLink) ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspLinkOutput)
}

type NspLinkOutput struct{ *pulumi.OutputState }

func (NspLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspLink)(nil)).Elem()
}

func (o NspLinkOutput) ToNspLinkOutput() NspLinkOutput {
	return o
}

func (o NspLinkOutput) ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput {
	return o
}

func (o NspLinkOutput) AutoApprovedRemotePerimeterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringPtrOutput { return v.AutoApprovedRemotePerimeterResourceId }).(pulumi.StringPtrOutput)
}

func (o NspLinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NspLinkOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NspLinkOutput) LocalInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.LocalInboundProfiles }).(pulumi.StringArrayOutput)
}

func (o NspLinkOutput) LocalOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.LocalOutboundProfiles }).(pulumi.StringArrayOutput)
}

func (o NspLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NspLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NspLinkOutput) RemoteInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.RemoteInboundProfiles }).(pulumi.StringArrayOutput)
}

func (o NspLinkOutput) RemoteOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.RemoteOutboundProfiles }).(pulumi.StringArrayOutput)
}

func (o NspLinkOutput) RemotePerimeterGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.RemotePerimeterGuid }).(pulumi.StringOutput)
}

func (o NspLinkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o NspLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspLinkOutput{})
}
