


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateRecordSet struct {
	pulumi.CustomResourceState

	ARecords         ARecordResponseArrayOutput    `pulumi:"aRecords"`
	AaaaRecords      AaaaRecordResponseArrayOutput `pulumi:"aaaaRecords"`
	CnameRecord      CnameRecordResponsePtrOutput  `pulumi:"cnameRecord"`
	Etag             pulumi.StringPtrOutput        `pulumi:"etag"`
	Fqdn             pulumi.StringOutput           `pulumi:"fqdn"`
	IsAutoRegistered pulumi.BoolOutput             `pulumi:"isAutoRegistered"`
	Metadata         pulumi.StringMapOutput        `pulumi:"metadata"`
	MxRecords        MxRecordResponseArrayOutput   `pulumi:"mxRecords"`
	Name             pulumi.StringOutput           `pulumi:"name"`
	PtrRecords       PtrRecordResponseArrayOutput  `pulumi:"ptrRecords"`
	SoaRecord        SoaRecordResponsePtrOutput    `pulumi:"soaRecord"`
	SrvRecords       SrvRecordResponseArrayOutput  `pulumi:"srvRecords"`
	Ttl              pulumi.Float64PtrOutput       `pulumi:"ttl"`
	TxtRecords       TxtRecordResponseArrayOutput  `pulumi:"txtRecords"`
	Type             pulumi.StringOutput           `pulumi:"type"`
}


func NewPrivateRecordSet(ctx *pulumi.Context,
	name string, args *PrivateRecordSetArgs, opts ...pulumi.ResourceOption) (*PrivateRecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateZoneName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateZoneName'")
	}
	if args.RecordType == nil {
		return nil, errors.New("invalid value for required argument 'RecordType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateRecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:PrivateRecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:PrivateRecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateRecordSet
	err := ctx.RegisterResource("azure-native:network/v20200601:PrivateRecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateRecordSetState, opts ...pulumi.ResourceOption) (*PrivateRecordSet, error) {
	var resource PrivateRecordSet
	err := ctx.ReadResource("azure-native:network/v20200601:PrivateRecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateRecordSetState struct {
}

type PrivateRecordSetState struct {
}

func (PrivateRecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateRecordSetState)(nil)).Elem()
}

type privateRecordSetArgs struct {
	ARecords              []ARecord         `pulumi:"aRecords"`
	AaaaRecords           []AaaaRecord      `pulumi:"aaaaRecords"`
	CnameRecord           *CnameRecord      `pulumi:"cnameRecord"`
	Metadata              map[string]string `pulumi:"metadata"`
	MxRecords             []MxRecord        `pulumi:"mxRecords"`
	PrivateZoneName       string            `pulumi:"privateZoneName"`
	PtrRecords            []PtrRecord       `pulumi:"ptrRecords"`
	RecordType            string            `pulumi:"recordType"`
	RelativeRecordSetName *string           `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SoaRecord             *SoaRecord        `pulumi:"soaRecord"`
	SrvRecords            []SrvRecord       `pulumi:"srvRecords"`
	Ttl                   *float64          `pulumi:"ttl"`
	TxtRecords            []TxtRecord       `pulumi:"txtRecords"`
}


type PrivateRecordSetArgs struct {
	ARecords              ARecordArrayInput
	AaaaRecords           AaaaRecordArrayInput
	CnameRecord           CnameRecordPtrInput
	Metadata              pulumi.StringMapInput
	MxRecords             MxRecordArrayInput
	PrivateZoneName       pulumi.StringInput
	PtrRecords            PtrRecordArrayInput
	RecordType            pulumi.StringInput
	RelativeRecordSetName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SoaRecord             SoaRecordPtrInput
	SrvRecords            SrvRecordArrayInput
	Ttl                   pulumi.Float64PtrInput
	TxtRecords            TxtRecordArrayInput
}

func (PrivateRecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateRecordSetArgs)(nil)).Elem()
}

type PrivateRecordSetInput interface {
	pulumi.Input

	ToPrivateRecordSetOutput() PrivateRecordSetOutput
	ToPrivateRecordSetOutputWithContext(ctx context.Context) PrivateRecordSetOutput
}

func (*PrivateRecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateRecordSet)(nil)).Elem()
}

func (i *PrivateRecordSet) ToPrivateRecordSetOutput() PrivateRecordSetOutput {
	return i.ToPrivateRecordSetOutputWithContext(context.Background())
}

func (i *PrivateRecordSet) ToPrivateRecordSetOutputWithContext(ctx context.Context) PrivateRecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateRecordSetOutput)
}

type PrivateRecordSetOutput struct{ *pulumi.OutputState }

func (PrivateRecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateRecordSet)(nil)).Elem()
}

func (o PrivateRecordSetOutput) ToPrivateRecordSetOutput() PrivateRecordSetOutput {
	return o
}

func (o PrivateRecordSetOutput) ToPrivateRecordSetOutputWithContext(ctx context.Context) PrivateRecordSetOutput {
	return o
}

func (o PrivateRecordSetOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) ARecordResponseArrayOutput { return v.ARecords }).(ARecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) AaaaRecordResponseArrayOutput { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v *PrivateRecordSet) CnameRecordResponsePtrOutput { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

func (o PrivateRecordSetOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateRecordSetOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o PrivateRecordSetOutput) IsAutoRegistered() pulumi.BoolOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.BoolOutput { return v.IsAutoRegistered }).(pulumi.BoolOutput)
}

func (o PrivateRecordSetOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o PrivateRecordSetOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) MxRecordResponseArrayOutput { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateRecordSetOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) PtrRecordResponseArrayOutput { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v *PrivateRecordSet) SoaRecordResponsePtrOutput { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

func (o PrivateRecordSetOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) SrvRecordResponseArrayOutput { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.Float64PtrOutput { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o PrivateRecordSetOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v *PrivateRecordSet) TxtRecordResponseArrayOutput { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

func (o PrivateRecordSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateRecordSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateRecordSetOutput{})
}
