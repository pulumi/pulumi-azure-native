


package managementpartner

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
			Type: pulumi.String("azure-nextgen:managementpartner:Partner"),
		},
		{
			Type: pulumi.String("azure-native:managementpartner/v20180201:Partner"),
		},
		{
			Type: pulumi.String("azure-nextgen:managementpartner/v20180201:Partner"),
		},
	})
	opts = append(opts, aliases)
	var resource Partner
	err := ctx.RegisterResource("azure-native:managementpartner:Partner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerState, opts ...pulumi.ResourceOption) (*Partner, error) {
	var resource Partner
	err := ctx.ReadResource("azure-native:managementpartner:Partner", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Partner)(nil))
}

func (i *Partner) ToPartnerOutput() PartnerOutput {
	return i.ToPartnerOutputWithContext(context.Background())
}

func (i *Partner) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerOutput)
}

type PartnerOutput struct{ *pulumi.OutputState }

func (PartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Partner)(nil))
}

func (o PartnerOutput) ToPartnerOutput() PartnerOutput {
	return o
}

func (o PartnerOutput) ToPartnerOutputWithContext(ctx context.Context) PartnerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PartnerOutput{})
}
