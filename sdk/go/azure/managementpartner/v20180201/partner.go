


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Partner struct {
	pulumi.CustomResourceState

	CreatedTime pulumi.StringPtrOutput `pulumi:"createdTime"`
	Etag        pulumi.IntPtrOutput    `pulumi:"etag"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	ObjectId    pulumi.StringPtrOutput `pulumi:"objectId"`
	PartnerId   pulumi.StringPtrOutput `pulumi:"partnerId"`
	PartnerName pulumi.StringPtrOutput `pulumi:"partnerName"`
	TenantId    pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	UpdatedTime pulumi.StringPtrOutput `pulumi:"updatedTime"`
	Version     pulumi.IntPtrOutput    `pulumi:"version"`
}


func NewPartner(ctx *pulumi.Context,
	name string, args *PartnerArgs, opts ...pulumi.ResourceOption) (*Partner, error) {
	if args == nil {
		args = &PartnerArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managementpartner:Partner"),
		},
	})
	opts = append(opts, aliases)
	var resource Partner
	err := ctx.RegisterResource("azure-native:managementpartner/v20180201:Partner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerState, opts ...pulumi.ResourceOption) (*Partner, error) {
	var resource Partner
	err := ctx.ReadResource("azure-native:managementpartner/v20180201:Partner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerState struct {
}

type PartnerState struct {
}

func (PartnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerState)(nil)).Elem()
}

type partnerArgs struct {
	PartnerId *string `pulumi:"partnerId"`
}


type PartnerArgs struct {
	PartnerId pulumi.StringPtrInput
}

func (PartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerArgs)(nil)).Elem()
}

type PartnerInput interface {
	pulumi.Input

	ToPartnerOutput() PartnerOutput
	ToPartnerOutputWithContext(ctx context.Context) PartnerOutput
}

func (*Partner) ElementType() reflect.Type {
	return reflect.TypeOf((**Partner)(nil)).Elem()
}

func (i *Partner) ToPartnerOutput() PartnerOutput {
	return i.ToPartnerOutputWithContext(context.Background())
}

func (i *Partner) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerOutput)
}

type PartnerOutput struct{ *pulumi.OutputState }

func (PartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Partner)(nil)).Elem()
}

func (o PartnerOutput) ToPartnerOutput() PartnerOutput {
	return o
}

func (o PartnerOutput) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return o
}

func (o PartnerOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) Etag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.IntPtrOutput { return v.Etag }).(pulumi.IntPtrOutput)
}

func (o PartnerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) PartnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.PartnerId }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.PartnerName }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PartnerOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.StringPtrOutput { return v.UpdatedTime }).(pulumi.StringPtrOutput)
}

func (o PartnerOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Partner) pulumi.IntPtrOutput { return v.Version }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerOutput{})
}
