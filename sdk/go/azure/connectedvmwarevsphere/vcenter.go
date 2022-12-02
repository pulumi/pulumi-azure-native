


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VCenter struct {
	pulumi.CustomResourceState

	ConnectionStatus   pulumi.StringOutput               `pulumi:"connectionStatus"`
	Credentials        VICredentialResponsePtrOutput     `pulumi:"credentials"`
	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	Fqdn               pulumi.StringOutput               `pulumi:"fqdn"`
	InstanceUuid       pulumi.StringOutput               `pulumi:"instanceUuid"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	Port               pulumi.IntPtrOutput               `pulumi:"port"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	Version            pulumi.StringOutput               `pulumi:"version"`
}


func NewVCenter(ctx *pulumi.Context,
	name string, args *VCenterArgs, opts ...pulumi.ResourceOption) (*VCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fqdn == nil {
		return nil, errors.New("invalid value for required argument 'Fqdn'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:VCenter"),
		},
	})
	opts = append(opts, aliases)
	var resource VCenter
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:VCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VCenterState, opts ...pulumi.ResourceOption) (*VCenter, error) {
	var resource VCenter
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:VCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vcenterState struct {
}

type VCenterState struct {
}

func (VCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*vcenterState)(nil)).Elem()
}

type vcenterArgs struct {
	Credentials       *VICredential     `pulumi:"credentials"`
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	Fqdn              string            `pulumi:"fqdn"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Port              *int              `pulumi:"port"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VcenterName       *string           `pulumi:"vcenterName"`
}


type VCenterArgs struct {
	Credentials       VICredentialPtrInput
	ExtendedLocation  ExtendedLocationPtrInput
	Fqdn              pulumi.StringInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Port              pulumi.IntPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VcenterName       pulumi.StringPtrInput
}

func (VCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vcenterArgs)(nil)).Elem()
}

type VCenterInput interface {
	pulumi.Input

	ToVCenterOutput() VCenterOutput
	ToVCenterOutputWithContext(ctx context.Context) VCenterOutput
}

func (*VCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenter)(nil)).Elem()
}

func (i *VCenter) ToVCenterOutput() VCenterOutput {
	return i.ToVCenterOutputWithContext(context.Background())
}

func (i *VCenter) ToVCenterOutputWithContext(ctx context.Context) VCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCenterOutput)
}

type VCenterOutput struct{ *pulumi.OutputState }

func (VCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenter)(nil)).Elem()
}

func (o VCenterOutput) ToVCenterOutput() VCenterOutput {
	return o
}

func (o VCenterOutput) ToVCenterOutputWithContext(ctx context.Context) VCenterOutput {
	return o
}

func (o VCenterOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VCenterOutput) Credentials() VICredentialResponsePtrOutput {
	return o.ApplyT(func(v *VCenter) VICredentialResponsePtrOutput { return v.Credentials }).(VICredentialResponsePtrOutput)
}

func (o VCenterOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o VCenterOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VCenter) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VCenterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o VCenterOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.InstanceUuid }).(pulumi.StringOutput)
}

func (o VCenterOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VCenterOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VCenter) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o VCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VCenterOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VCenter) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o VCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VCenterOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o VCenterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VCenterOutput{})
}
